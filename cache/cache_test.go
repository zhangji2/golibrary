package cache

import (
	"fmt"
	"log"
	"testing"
)

func cacheHandler(conf ConfigCache) error {
	cacheClient, err := NewCache(conf)
	if err != nil {
		log.Println("redis 初始化失败:", err)
	} else {
		_ = cacheClient.Set("demo_key", "hello_"+conf.Driver, 300)
		val, _ := cacheClient.Get("demo_key")
		fmt.Println("Redis 取值:", val)
		defer cacheClient.Close()
	}
	return nil
}

func TestRedis(t *testing.T) {
	redisConf := ConfigCache{
		Driver: "redis",
		Config: map[string]string{
			"addr":     "127.0.0.1:6379",
			"password": "foobared",
			"db":       "0",
		},
	}
	cacheHandler(redisConf)
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
	cacheHandler(redisPoolConf)
}

func TestMemcache(t *testing.T) {
	memConf := ConfigCache{
		Driver: "memcache",
		Config: map[string]string{
			"addr": "127.0.0.1:11211",
		},
	}
	cacheHandler(memConf)
}

// 在项目中使用：先引入包，再调函数。
// zcache "github.com/zhangji2/golibrary/cache"
