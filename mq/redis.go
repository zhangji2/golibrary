package mq

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisDriver struct{}
type redisMQ struct {
	client *redis.Client
	ctx    context.Context
}

func init() {
	Register("redis", &redisDriver{})
}

func (d *redisDriver) Open(cfg map[string]string) (MQ, error) {
	dbNum, _ := strconv.Atoi(cfg["db"])
	client := redis.NewClient(&redis.Options{
		Addr:     cfg["addr"],
		Password: cfg["password"],
		DB:       dbNum,
	})
	return &redisMQ{client: client, ctx: context.Background()}, nil
}

// 普通发送
func (r *redisMQ) Send(topic string, msg []byte) error {
	return r.client.LPush(r.ctx, topic, msg).Err()
}

// 延时发送（zset 实现）
func (r *redisMQ) SendDelay(topic string, msg []byte, delaySeconds int) error {
	delayKey := "delay_queue:" + topic
	score := float64(time.Now().Unix() + int64(delaySeconds))
	if err := r.client.ZAdd(r.ctx, delayKey, &redis.Z{Score: score, Member: msg}).Err(); err != nil {
		return err
	}
	go r.startDelayWorker(topic)
	return nil
}

// 延时消息后台转移
func (r *redisMQ) startDelayWorker(topic string) {
	delayKey := "delay_queue:" + topic
	for {
		time.Sleep(200 * time.Millisecond)
		now := time.Now().Unix()
		res, err := r.client.ZRangeByScore(r.ctx, delayKey, &redis.ZRangeBy{
			Min: "0",
			Max: strconv.FormatInt(now, 10),
		}).Result()
		if err != nil || len(res) == 0 {
			continue
		}
		for _, msg := range res {
			if err := r.client.ZRem(r.ctx, delayKey, msg).Err(); err != nil {
				continue
			}
			_ = r.Send(topic, []byte(msg))
		}
	}
}

func (r *redisMQ) Subscribe(topic string, handler func([]byte) error) error {
	go func() {
		for {
			msg, err := r.client.BRPop(r.ctx, 0, topic).Result()
			if err != nil || len(msg) < 2 {
				time.Sleep(1 * time.Second)
				continue
			}
			_ = handler([]byte(msg[1]))
		}
	}()
	return nil
}

func (r *redisMQ) Close() error {
	return r.client.Close()
}
