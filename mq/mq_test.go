package mq

import (
	"fmt"
	"testing"
	"time"
)

var msgHandler = func(msg []byte) error {
	fmt.Printf("✅ 消费成功：%s\n", string(msg))
	return nil
}

func mqHandler(conf ConfigMQ) error {
	mqClient, err := NewMQ(conf)
	if err != nil {
		fmt.Println("MQ 连接失败:", err)
		return err
	}
	defer mqClient.Close()

	// 订阅
	_ = mqClient.Subscribe("test_topic", msgHandler)

	// 发送
	_ = mqClient.Send("test_topic", []byte("hello go mq driver"))
	_ = mqClient.Send("test_topic", []byte(conf.Driver+" 普通消息"))
	_ = mqClient.SendDelay("test_topic", []byte(conf.Driver+" 延时消息（3秒）"), 3)

	time.Sleep(3 * time.Second)
	fmt.Println("运行完成")

	return nil

}

func TestRedis(t *testing.T) {
	mqConf := ConfigMQ{
		Driver: "redis",
		Config: map[string]string{
			"addr":     "127.0.0.1:6379",
			"password": "foobared",
			"db":       "16",
		},
	}
	mqHandler(mqConf)
}

func TestRabbitMQ(t *testing.T) {
	rabbitConf := ConfigMQ{
		Driver: "rabbitmq",
		Config: map[string]string{
			"uri": "amqp://admin:123456@127.0.0.1:5672/default_vhost",
		},
	}
	mqHandler(rabbitConf)
}

func TestKafka(t *testing.T) {
	kafkaConf := ConfigMQ{
		Driver: "kafka",
		Config: map[string]string{
			"brokers": "127.0.0.1:9092",
		},
	}
	mqHandler(kafkaConf)
}

func TestRocketMQ(t *testing.T) {
	rocketConf := ConfigMQ{
		Driver: "rocketmq",
		Config: map[string]string{
			"endpoint": "127.0.0.1:9876",
			"group":    "mq_consumer_group",
		},
	}
	mqHandler(rocketConf)
}
