package confluent

import (
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.opentelemetry.io/otel/propagation"
)

type KafkaMessageTextMapCarrier struct {
	msg *kafka.Message
}

var _ propagation.TextMapCarrier = &KafkaMessageTextMapCarrier{}

// Get returns the value associated with the passed key.
func (carrier *KafkaMessageTextMapCarrier) Get(key string) string {
	for i := 0; i < len(carrier.msg.Headers); i++ {
		header := carrier.msg.Headers[i]
		if strings.Compare("md-"+key, header.Key) == 0 {
			return string(header.Value)
		}
	}
	return ""
}

// Set stores the key-value pair.
func (carrier *KafkaMessageTextMapCarrier) Set(key string, value string) {
	carrier.msg.Headers = append(carrier.msg.Headers, kafka.Header{
		Key:   "md-" + key,
		Value: []byte(value),
	})
}

// Keys lists the keys stored in this carrier.
func (carrier *KafkaMessageTextMapCarrier) Keys() []string {
	return nil
}
