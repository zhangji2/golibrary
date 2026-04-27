package mq

import "errors"

// ConfigMQ 外部统一配置
type ConfigMQ struct {
	Driver string            // redis/rabbitmq/kafka/rocketmq
	Config map[string]string // 连接参数
}

// MQ 统一消息队列接口（支持普通+延时消息）
type MQ interface {
	// 普通发送
	Send(topic string, message []byte) error
	// 延时发送（秒）
	SendDelay(topic string, message []byte, delaySeconds int) error
	// 订阅
	Subscribe(topic string, handler func(msg []byte) error) error
	// 关闭
	Close() error
}

// Driver 驱动接口
type Driver interface {
	Open(config map[string]string) (MQ, error)
}

// 驱动注册中心
var drivers = make(map[string]Driver)

func Register(name string, driver Driver) {
	drivers[name] = driver
}

func NewMQ(conf ConfigMQ) (MQ, error) {
	driver, ok := drivers[conf.Driver]
	if !ok {
		return nil, errors.New("unsupported mq driver: " + conf.Driver)
	}
	return driver.Open(conf.Config)
}
