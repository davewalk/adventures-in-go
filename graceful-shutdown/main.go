package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
)

var (
	kafkaHost string
	inTopic   string
)

func init() {
	flag.StringVar(&kafkaHost, "kafkaHost", "localhost:9092", "Kafka host, including port")
	flag.StringVar(&inTopic, "inTopic", "product", "The Kafka topic to consume from")

	flag.Parse()
}

func main() {
	checkForMessages := true

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	master, err := sarama.NewConsumer([]string{kafkaHost}, config)
	if err != nil {
		panic(err)
	}

	// Consume from the beginning so that the messages stream through
	consumer, err := master.ConsumePartition(inTopic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)
		sig := <-sigchan
		fmt.Printf("Got a signal: %v\n", sig)

		fmt.Println("Closing Kafka consumer...")
		consumer.Close()
		master.Close()

		checkForMessages = false
	}()

	for checkForMessages {
		select {
		case err = <-consumer.Errors():
			fmt.Printf("got a consumer error: %v", err)
		case _ = <-consumer.Messages():
			fmt.Printf("Got a message...")
			time.Sleep(5 * time.Second)
			fmt.Printf("done with message\n")
		}
	}

	os.Exit(0)
}
