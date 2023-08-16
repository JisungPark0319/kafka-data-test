package handler

import (
	"context"
	"encoding/json"
	"math/rand"
	"random-data/pkg/generator"
	"random-data/pkg/kafka"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func ProducerHandler(wg *sync.WaitGroup, brokers []string, topic string, count int, dataType string) {
	defer wg.Done()

	cfg := kafka.ProducerConfig{
		Brokers: brokers,
		Topic:   topic,
	}
	producer, err := kafka.NewProducer(cfg)
	if err != nil {
		logrus.Error("producer access fail")
		return
	}

	loopCount := 1
	for {
		latency := rand.Float64() * 100 * float64(time.Millisecond)

		var genData any
		if dataType == "image" {
			genData = generator.NewImage(1000, 500)
		} else {
			genData = generator.NewUser()
		}

		data, err := json.Marshal(genData)
		if err != nil {
			logrus.Error("json marshal fail")
		}
		message := kafka.Message{
			Key:   uuid.NewString(),
			Value: string(data),
		}

		sendStartTime := time.Now()
		if err := producer.SendMessage(context.Background(), message); err != nil {
			logrus.Error("send fail")
		}
		logrus.Infof("%d send data - latency: %v\n", loopCount, time.Since(sendStartTime))

		if count != 0 {
			if loopCount == count {
				break
			}
		}
		loopCount++

		time.Sleep(time.Duration(latency))
	}
}
