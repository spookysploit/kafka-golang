package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/spookysploit/kafka-golang/internal/kafka"
)

const (
	topic        = "my-topic"
	numberOfKeys = 20
)

var address = []string{"localhost:9091", "localhost:9092", "localhost:9093"}

func main() {
	log := setupLogger()
	p, err := kafka.NewProducer(address)
	if err != nil {
		log.Error("failed to create new producer", slog.String("error", err.Error()))
	}

	keys := generateUUIDString()
	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("Kafka message %d", i)
		key := keys[i%numberOfKeys]
		if err := p.Produce(msg, topic, key, time.Now()); err != nil {
			log.Error("failed to send message to topic", slog.String("error", err.Error()))
		}
	}

}

func setupLogger() *slog.Logger {
	var log *slog.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return log
}

func generateUUIDString() [numberOfKeys]string {
	var uuids [numberOfKeys]string
	for i := 0; i < numberOfKeys; i++ {
		uuids[i] = uuid.NewString()
	}
	return uuids
}
