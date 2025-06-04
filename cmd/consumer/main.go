package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spookysploit/kafka-golang/internal/handler"
	"github.com/spookysploit/kafka-golang/internal/kafka"
)

const (
	topic         = "my-topic"
	consumerGroup = "my-consumer-group"
)

var address = []string{"localhost:9091", "localhost:9092", "localhost:9093"}

func main() {
	h := handler.NewHandler()

	c1, err := kafka.NewConsumer(h, address, topic, consumerGroup, 1)
	if err != nil {
		log.Fatal(err)
	}

	c2, err := kafka.NewConsumer(h, address, topic, consumerGroup, 2)
	if err != nil {
		log.Fatal(err)
	}

	c3, err := kafka.NewConsumer(h, address, topic, consumerGroup, 3)
	if err != nil {
		log.Fatal(err)
	}

	go c1.Start()
	go c2.Start()
	go c3.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	if err := c1.Stop(); err != nil {
		log.Println("error stopping c1:", err)
	}
	if err := c2.Stop(); err != nil {
		log.Println("error stopping c2:", err)
	}
	if err := c3.Stop(); err != nil {
		log.Println("error stopping c3:", err)
	}
}
