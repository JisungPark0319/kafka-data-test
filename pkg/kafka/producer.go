package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Writer *kafka.Writer
}

type ProducerConfig struct {
	Brokers []string
	Topic   string
}

func NewProducer(cfg ProducerConfig) (*Producer, error) {
	p := &Producer{}
	config := kafka.WriterConfig{
		Brokers:      cfg.Brokers,
		Topic:        cfg.Topic,
		BatchTimeout: time.Millisecond * 1,
	}

	p.Writer = kafka.NewWriter(config)

	return p, nil
}

func (p *Producer) SendMessage(ctx context.Context, message Message) error {
	err := p.Writer.WriteMessages(
		ctx,
		kafka.Message{
			Key:   []byte(message.Key),
			Value: []byte(message.Value),
		},
	)

	if err != nil {
		return fmt.Errorf("failed to wrtie message: %v", err)
	}

	return nil
}
