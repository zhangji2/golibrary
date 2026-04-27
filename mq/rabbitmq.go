package mq

import (
	"strconv"

	"github.com/streadway/amqp"
)

type rabbitDriver struct{}
type rabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func init() {
	Register("rabbitmq", &rabbitDriver{})
}

func (d *rabbitDriver) Open(cfg map[string]string) (MQ, error) {
	conn, err := amqp.Dial(cfg["uri"])
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, err
	}
	return &rabbitMQ{conn: conn, channel: ch}, nil
}

func (r *rabbitMQ) Send(topic string, msg []byte) error {
	_, _ = r.channel.QueueDeclare(topic, true, false, false, false, nil)
	return r.channel.Publish("", topic, false, false, amqp.Publishing{Body: msg})
}

// 延时发送
func (r *rabbitMQ) SendDelay(topic string, msg []byte, delaySeconds int) error {
	delayQueue := "delay." + topic
	args := amqp.Table{
		"x-dead-letter-exchange":    "",
		"x-dead-letter-routing-key": topic,
	}
	_, _ = r.channel.QueueDeclare(delayQueue, true, false, false, false, args)
	return r.channel.Publish("", delayQueue, false, false, amqp.Publishing{
		Body:       msg,
		Expiration: strconv.Itoa(delaySeconds * 1000),
	})
}

func (r *rabbitMQ) Subscribe(topic string, handler func([]byte) error) error {
	_, _ = r.channel.QueueDeclare(topic, true, false, false, false, nil)
	msgs, _ := r.channel.Consume(topic, "", true, false, false, false, nil)
	go func() {
		for d := range msgs {
			_ = handler(d.Body)
		}
	}()
	return nil
}

func (r *rabbitMQ) Close() error {
	_ = r.channel.Close()
	return r.conn.Close()
}
