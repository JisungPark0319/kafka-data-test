package handler

import (
	"random-data/pkg/kafka"
	"sync"

	"github.com/sirupsen/logrus"
)

func CreateTopic(wg *sync.WaitGroup, brokers []string, topic string, partition int) {
	defer wg.Done()
	controller, err := kafka.NewController(brokers)
	if err != nil {
		logrus.Error("fail create kafka controller")
		return
	}
	if err := controller.CreateTopic(topic, partition, 1); err != nil {
		logrus.Errorf("fali create topci - %s\n", topic)
		return
	}
	logrus.Info("success create topic")
}

func GetPartition(brokers []string, topic string) (int, error) {
	controller, err := kafka.NewController(brokers)
	if err != nil {
		return 0, err
	}
	partition, err := controller.GetPartitionCount(topic)

	return partition, err
}
