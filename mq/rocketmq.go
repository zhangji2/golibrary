package mq

import (
	"context"
	"os"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

// 关闭 RocketMQ 疯狂刷屏日志！！
func init() {
	_ = os.Setenv("ROCKETMQ_GO_LOG_LEVEL", "error")
}

type rocketDriver struct{}
type rocketMQ struct {
	producer rocketmq.Producer
	consumer rocketmq.PushConsumer
}

func init() {
	Register("rocketmq", &rocketDriver{})
}

func (d *rocketDriver) Open(cfg map[string]string) (MQ, error) {
	endpoint := cfg["endpoint"]
	group := cfg["group"]

	// 生产者
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{endpoint}),
		producer.WithRetry(1),
		producer.WithSendMsgTimeout(3000),
	)
	if err != nil {
		return nil, err
	}
	if err = p.Start(); err != nil {
		return nil, err
	}

	// 消费者（不启动也可以，只发消息）
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{endpoint}),
		consumer.WithGroupName(group),
		consumer.WithConsumerModel(consumer.Clustering),
	)
	if err != nil {
		_ = p.Shutdown()
		return nil, err
	}

	return &rocketMQ{producer: p, consumer: c}, nil
}

func (r *rocketMQ) Send(topic string, msg []byte) error {
	_, err := r.producer.SendSync(context.Background(), primitive.NewMessage(topic, msg))
	return err
}

func (r *rocketMQ) SendDelay(topic string, msg []byte, delaySeconds int) error {
	m := primitive.NewMessage(topic, msg)
	level := getDelayLevel(delaySeconds)
	m.WithDelayTimeLevel(int(level))
	_, err := r.producer.SendSync(context.Background(), m)
	return err
}

func getDelayLevel(sec int) int32 {
	switch {
	case sec <= 1:
		return 1
	case sec <= 5:
		return 2
	case sec <= 10:
		return 3
	case sec <= 30:
		return 4
	case sec <= 60:
		return 5
	default:
		return 6
	}
}

// 订阅
func (r *rocketMQ) Subscribe(topic string, handler func(msg []byte) error) error {
	selector := consumer.MessageSelector{}
	err := r.consumer.Subscribe(topic, selector, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			_ = handler(msg.Body)
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		return err
	}
	return r.consumer.Start()
}

func (r *rocketMQ) Close() error {
	_ = r.producer.Shutdown()
	_ = r.consumer.Shutdown()
	return nil
}
