package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	Reader *kafka.Reader
}

type ConsumerConfig struct {
	Brokers   []string
	GroupID   string
	Partition int
	Topic     string
}

func NewConsumer(cfg ConsumerConfig) *Consumer {
	rCfg := kafka.ReaderConfig{
		Brokers:  cfg.Brokers,
		Topic:    cfg.Topic,
		MaxBytes: 10e6, // 10MB
	}
	if cfg.GroupID != "" {
		rCfg.GroupID = cfg.GroupID
	}
	if cfg.Partition != 0 {
		rCfg.Partition = cfg.Partition
	}

	r := kafka.NewReader(rCfg)

	return &Consumer{
		Reader: r,
	}
}

func (c *Consumer) SetPartition(partition int) {
	cfg := c.Reader.Config()
	cfg.Partition = partition
	// r.Reader = kafka.NewReader(cfg)
}

func (c *Consumer) ReadMessage(ctx context.Context) (Message, error) {
	message := Message{}
	m, err := c.Reader.ReadMessage(ctx)
	if err != nil {
		return message, err
	}
	message = Message{
		Key:   string(m.Key),
		Value: string(m.Value),
	}

	return message, nil
}
