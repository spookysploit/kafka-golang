package handler

import (
	"fmt"
	"log/slog"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleMessage(message []byte, topic kafka.TopicPartition, cn int) error {
	logMsg := fmt.Sprintf(
		"CONSUMER #%d, Message from kafka with offset %v, partition=%d", cn, topic.Offset, topic.Partition,
	)
	slog.Info(logMsg)
	return nil
}
