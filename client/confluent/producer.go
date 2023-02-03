package confluent

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-bamboo/pkg/log"
	otelext "github.com/go-bamboo/pkg/otel"
	"github.com/go-bamboo/pkg/queue"
	"github.com/go-kratos/kratos/v2/metrics"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// KafkaProducer 生产者
type KafkaProducer struct {
	pub        *kafka.Producer
	tracer     trace.Tracer
	propagator propagation.TextMapPropagator
	pubCounter metrics.Counter // 发送次数
	// topic      string
}

func MustNewPusher(c *ProducerConf) queue.Pusher {
	pub, err := NewPusher(c)
	if err != nil {
		log.Fatal(err)
	}
	return pub
}

func NewPusher(c *ProducerConf) (queue.Pusher, error) {
	var config = make(kafka.ConfigMap)
	config["bootstrap.servers"] = c.BootstrapServers
	config["api.version.request"] = true
	config["security.protocol"] = c.SecurityProtocol
	config["sasl.mechanisms"] = "PLAIN"
	config["sasl.username"] = "kafkauser"
	config["sasl.password"] = "Kafkauser-music-bee"
	// config["enable.ssl.certificate.verification"] = "true"
	config["ssl.ca.location"] = c.Ssl.CaLocation

	pub, err := kafka.NewProducer(&config)
	if err != nil {
		err = WrapError(err)
		return nil, err
	}
	tracingPub := &KafkaProducer{
		pub:        pub,
		tracer:     otel.Tracer("kafka"),
		propagator: propagation.NewCompositeTextMapPropagator(otelext.Metadata{}, propagation.Baggage{}, otelext.TraceContext{}),
		// topic:      c.Topic,
	}
	return tracingPub, nil
}

func (p *KafkaProducer) Name() string {
	return "confluent"
}

func (p *KafkaProducer) Push(ctx context.Context, topic string, key, value []byte) error {
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   key,
		Value: value,
	}
	operation := "pub:" + topic
	ctx, span := p.tracer.Start(ctx, operation, trace.WithSpanKind(trace.SpanKindConsumer))
	p.propagator.Inject(ctx, &KafkaMessageTextMapCarrier{msg: msg})
	span.SetAttributes(
		attribute.String("kafka.topic", *msg.TopicPartition.Topic),
		attribute.String("kafka.key", string(msg.Key)),
	)
	if err := p.pub.Produce(msg, deliveryChan); err != nil {
		span.RecordError(err)
		return WrapError(err)
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		err := WrapError(m.TopicPartition.Error)
		return err
	}
	log.Debugf("Delivered message to topic %s [%d] at offset %v", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	return nil
}

func (p *KafkaProducer) Close() error {
	p.pub.Close()
	return nil
}
