package handler

import (
	"context"
	"random-data/pkg/kafka"
	"sync"

	"github.com/sirupsen/logrus"
)

func ConsumerHandler(wg *sync.WaitGroup, brokers []string, topic string, groupID string) {
	defer wg.Done()

	cfg := kafka.ConsumerConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	}
	consumer := kafka.NewConsumer(cfg)

	for {
		message, err := consumer.ReadMessage(context.Background())
		if err != nil {
			logrus.Errorf("failed read message in %s topic\n", topic)
			continue
		}
		logrus.Infof("%s topic read message - {%s: %s}\n", topic, message.Key, message.Value)
	}
}
