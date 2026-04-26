package cache

import (
	"fmt"
	"log"
	"testing"
)

func TestRedis(t *testing.T) {
	redisConf := ConfigCache{
		Driver: "redis",
		Config: map[string]string{
			"addr":     "127.0.0.1:6379",
			"password": "foobared",
			"db":       "0",
		},
	}
	cacheCli, err := NewCache(redisConf)
	if err != nil {
		log.Println("redis 初始化失败:", err)
	} else {
		_ = cacheCli.Set("demo_key", "hello_redis", 300)
		val, _ := cacheCli.Get("demo_key")
		fmt.Println("Redis 取值:", val)
		defer cacheCli.Close()
	}
}

func TestRedisPool(t *testing.T) {
	redisPoolConf := ConfigCache{
		Driver: "redispool",
		Config: map[string]string{
			"addr":        "127.0.0.1:6379",
			"password":    "foobared",
			"db":          "0",
			"MaxIdle":     "10",
			"MaxActive":   "100",
			"IdleTimeout": "300",
		},
	}
	cacheCli, err := NewCache(redisPoolConf)
	if err != nil {
		log.Println("redis 初始化失败:", err)
	} else {
		_ = cacheCli.Set("demo_key", "hello_redis_pool", 300)
		val, _ := cacheCli.Get("demo_key")
		fmt.Println("Redis 取值:", val)
		defer cacheCli.Close()
	}
}

func TestMemcache(t *testing.T) {
	memConf := ConfigCache{
		Driver: "memcache",
		Config: map[string]string{
			"addr": "127.0.0.1:11211",
		},
	}
	cacheCli, err := NewCache(memConf)
	if err != nil {
		log.Println("memcache 初始化失败:", err)
	} else {
		_ = cacheCli.Set("demo_key", "hello_memcache", 300)
		val, _ := cacheCli.Get("demo_key")
		fmt.Println("Memcache 取值:", val)
		defer cacheCli.Close()
	}
}
