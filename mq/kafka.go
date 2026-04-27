package mq

import (
	"strings"
	"time"

	"github.com/IBM/sarama"
)

type kafkaDriver struct{}
type kafkaMQ struct {
	producer sarama.SyncProducer
	consumer sarama.Consumer
	brokers  []string
}

func init() {
	Register("kafka", &kafkaDriver{})
}

func (d *kafkaDriver) Open(cfg map[string]string) (MQ, error) {
	brokers := strings.Split(cfg["brokers"], ",")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Consumer.Return.Errors = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		_ = producer.Close()
		return nil, err
	}
	return &kafkaMQ{producer: producer, consumer: consumer, brokers: brokers}, nil
}

func (k *kafkaMQ) Send(topic string, msg []byte) error {
	_, _, err := k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	})
	return err
}

// 延时发送（本地时间轮实现）
func (k *kafkaMQ) SendDelay(topic string, msg []byte, delaySeconds int) error {
	go func() {
		time.Sleep(time.Duration(delaySeconds) * time.Second)
		_ = k.Send(topic, msg)
	}()
	return nil
}

func (k *kafkaMQ) Subscribe(topic string, handler func([]byte) error) error {
	parts, err := k.consumer.Partitions(topic)
	if err != nil {
		return err
	}
	for _, p := range parts {
		pc, err := k.consumer.ConsumePartition(topic, p, sarama.OffsetNewest)
		if err != nil {
			continue
		}
		go func(pc sarama.PartitionConsumer) {
			defer pc.Close()
			for m := range pc.Messages() {
				_ = handler(m.Value)
			}
		}(pc)
	}
	return nil
}

func (k *kafkaMQ) Close() error {
	_ = k.producer.Close()
	_ = k.consumer.Close()
	return nil
}
