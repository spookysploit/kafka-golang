package kafka

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer struct {
	producer *kafka.Producer
}

const (
	flushTimeout = 5000
)

var errUnknownType = errors.New("unknow type event")

func NewProducer(address []string) (*Producer, error) {
	conf := &kafka.ConfigMap{
		"bootstrap.servers": strings.Join(address, ","),
	}
	p, err := kafka.NewProducer(conf)
	if err != nil {
		return nil, fmt.Errorf("error when creating producer: %w", err)
	}
	return &Producer{producer: p}, nil
}

func (p *Producer) Produce(message, topic, key string, t time.Time) error {
	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:     []byte(message),
		Key:       []byte(key),
		Timestamp: t,
	}

	kafkaChan := make(chan kafka.Event)

	if err := p.producer.Produce(kafkaMsg, kafkaChan); err != nil {
		return err
	}

	e := <-kafkaChan
	switch eType := e.(type) {
	case *kafka.Message:
		return nil
	case kafka.Error:
		return eType
	default:
		return errUnknownType
	}
}

func (p *Producer) Close() {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
}
