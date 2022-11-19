package confluent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-bamboo/pkg/rescue"
	"github.com/go-bamboo/pkg/tracing"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metrics"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type (
	kafkaQueue struct {
		c       *ConsumerConf
		handler queue.ConsumeHandler

		sub        *kafka.Consumer
		tracer     trace.Tracer
		propagator propagation.TextMapPropagator
		subCounter metrics.Counter // 发送次数

		wg  sync.WaitGroup
		ctx context.Context
		cf  context.CancelFunc
	}

	kafkaQueues struct {
		queues []*kafkaQueue
	}
)

func MustNewQueue(c *ConsumerConf, handler queue.ConsumeHandler) queue.MessageQueue {
	q, err := NewQueue(c, handler)
	if err != nil {
		log.Fatal(err)
	}
	return q
}

func NewQueue(c *ConsumerConf, handler queue.ConsumeHandler) (queue.MessageQueue, error) {
	q := kafkaQueues{}
	cc, err := newKafkaQueue(c, handler)
	if err != nil {
		return nil, err
	}
	q.queues = append(q.queues, cc)
	return &q, nil
}

func (q kafkaQueues) Name() string {
	return "confluent"
}

func (q kafkaQueues) Start(ctx context.Context) error {
	for _, queue := range q.queues {
		queue.Start(ctx)
	}
	return nil
}

func (q kafkaQueues) Stop(ctx context.Context) error {
	for _, queue := range q.queues {
		queue.Stop(ctx)
	}
	return nil
}

func newKafkaQueue(c *ConsumerConf, handler queue.ConsumeHandler) (k *kafkaQueue, err error) {
	// sub
	var config = make(kafka.ConfigMap)
	config["bootstrap.servers"] = c.BootstrapServers
	config["api.version.request"] = true
	config["security.protocol"] = c.SecurityProtocol
	config["sasl.mechanisms"] = "PLAIN"
	config["sasl.username"] = c.Sasl.User
	config["sasl.password"] = c.Sasl.Password
	// config["enable.ssl.certificate.verification"] = "true"
	config["ssl.ca.location"] = c.Ssl.CaLocation

	config["group.id"] = c.Group
	config["session.timeout.ms"] = 6000
	config["default.topic.config"] = kafka.ConfigMap{"auto.offset.reset": "earliest"}

	sub, err := kafka.NewConsumer(&config)
	if err != nil {
		err = WrapError(err)
		return nil, err
	}

	ctx, cf := context.WithCancel(context.Background())
	k = &kafkaQueue{
		c:       c,
		handler: handler,

		sub:        sub,
		tracer:     otel.Tracer("kafka"),
		propagator: propagation.NewCompositeTextMapPropagator(tracing.Metadata{}, propagation.Baggage{}, tracing.TraceContext{}),

		ctx: ctx,
		cf:  cf,
	}
	return
}

func (c *kafkaQueue) Start(context.Context) error {
	log.Infof("start cunsumer topic:%v", c.c.Topics)
	c.wg.Add(1)
	go c.consumGroupTopic(c.c.Topics)
	log.Infof("start kafka consumer.")
	return nil
}

func (c *kafkaQueue) Stop(context.Context) error {
	c.cf()
	c.wg.Wait()
	if err := c.sub.Close(); err != nil {
		return err
	}
	log.Infof("stop kafka consumer.")
	return nil
}

func (c *kafkaQueue) consumGroupTopic(topics []string) {
	defer rescue.Recover(func() {
		c.wg.Done()
		log.Warnf("kafka consume group topic done")
	})
	if err := c.sub.SubscribeTopics(topics, nil); err != nil {
		log.Errorf("Failed to get %v the list of partition: %v", topics, err)
		return
	}
	log.Infof("topics => %v", topics)
	ctx := context.TODO()
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			// ms
			ev := c.sub.Poll(1000)
			if ev == nil {
				time.Sleep(time.Second)
				continue
			}
			switch msg := ev.(type) {
			case *kafka.Message:
				log.Debugw("handle kafka msg", "headers", msg.Headers, "topic", msg.TopicPartition.Topic, "Partition", msg.TopicPartition.Partition, "Offset", msg.TopicPartition.Offset)
				cCtx, cf := context.WithTimeout(ctx, 60*time.Second)
				cctx, span := c.tracer.Start(cCtx, "sub:"+*msg.TopicPartition.Topic, trace.WithSpanKind(trace.SpanKindConsumer))
				c.propagator.Inject(cctx, &KafkaMessageTextMapCarrier{msg: msg})
				span.SetAttributes(
					attribute.String("kafka.topic", *msg.TopicPartition.Topic),
					attribute.String("kafka.key", string(msg.Key)),
				)
				if err := c.handler.Consume(cctx, *msg.TopicPartition.Topic, msg.Key, msg.Value); err != nil {
					// 直接放弃的消息
					se := errors.FromError(err)
					log.Errorw(fmt.Sprintf("%+v", err), "code", se.Code, "reason", se.Reason)
					span.RecordError(err)
				}
				// 确认消费
				_, err := c.sub.CommitMessage(msg)
				if err != nil {
					span.RecordError(err)
					log.Errorf("err: %v", err)
				}
				cf()
			case kafka.OffsetsCommitted:
				log.Infof("kafka offsets committed", "topic", msg.Offsets[0].Topic, "Partition", msg.Offsets[0].Partition, "Offset", msg.Offsets[0].Offset)
			case kafka.PartitionEOF:
				log.Errorw(fmt.Sprintf("%+v", msg.Error), "topic", msg.Topic, "Partition", msg.Partition, "Offset", msg.Offset)
			case kafka.Error:
				log.Errorw(fmt.Sprintf("%+v", msg.Error()), "code", msg.Code())
			default:
				log.Warnf("Ignored %v", msg)
			}
		}
	}
}
