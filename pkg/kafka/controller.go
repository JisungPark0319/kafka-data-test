package kafka

import (
	"net"
	"strconv"
	"strings"

	"github.com/segmentio/kafka-go"
)

type Controller struct {
	conn *kafka.Conn
}

func NewController(brokers []string) (*Controller, error) {
	kafkaController := &Controller{}
	conn, err := kafka.Dial("tcp", strings.Join(brokers, ","))
	if err != nil {
		return kafkaController, err
	}
	controller, err := conn.Controller()
	if err != nil {
		return kafkaController, err
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return kafkaController, err
	}

	kafkaController.conn = controllerConn
	return kafkaController, nil
}
func (c *Controller) Close() error {
	return c.conn.Close()
}

func (c *Controller) CreateTopic(topic string, partitions int, factor int) error {
	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     partitions,
			ReplicationFactor: factor,
		},
	}

	err := c.conn.CreateTopics(topicConfigs...)
	return err
}

func (c *Controller) GetPartitionCount(topic string) (int, error) {
	partitions, err := c.conn.ReadPartitions(topic)
	if err != nil {
		return 0, nil
	}

	return len(partitions), nil
}
