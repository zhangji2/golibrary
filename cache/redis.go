package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/zhangji2/golibrary/logger"
)

type redisDriver struct{}
type redisCache struct {
	client *redis.Client
	ctx    context.Context
}

func init() {
	Register("redis", new(redisDriver))
}

func (r *redisDriver) Open(cfg map[string]string) (Cache, error) {
	addr := cfg["addr"]
	pwd := cfg["password"]
	dbNum, _ := strconv.Atoi(cfg["db"])

	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       dbNum,
	})

	ctx := context.Background()
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis 连接失败: %w", err)
	}

	return &redisCache{
		client: cli,
		ctx:    ctx,
	}, nil
}

func (r *redisCache) Set(key string, value interface{}, expire int) error {
	return r.client.Set(r.ctx, key, value, time.Duration(expire)*time.Second).Err()
}

func (r *redisCache) Get(key string) (interface{}, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *redisCache) Lock(key string, value interface{}, expire int) (bool, error) {
	err := r.client.SetNX(r.ctx, key, value, time.Duration(expire)*time.Second).Err()
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"mark": "isLockTask"}).Error(err)
		return false, err
	}
	return true, nil
}

func (r *redisCache) Del(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

func (r *redisCache) Close() error {
	return r.client.Close()
}
