package cache

import (
	"fmt"
	"strconv"
	"time"

	redispool "github.com/gomodule/redigo/redis"
)

type redisPoolDriver struct{}
type redisPoolCache struct {
	client *redispool.Pool
}

func init() {
	Register("redispool", new(redisPoolDriver))
}

// CacheRedisPool 缓存连接池
var CacheRedisPool *redispool.Pool

func (rp *redisPoolDriver) Open(cfg map[string]string) (Cache, error) {

	addr := cfg["addr"]
	pwd := cfg["password"]
	db, _ := strconv.Atoi(cfg["db"])
	MaxIdle, _ := strconv.Atoi(cfg["MaxIdle"])
	MaxActive, _ := strconv.Atoi(cfg["MaxActive"])
	IdleTimeout, _ := strconv.ParseInt(cfg["IdleTimeout"], 10, 64)

	if len(pwd) == 0 {
		CacheRedisPool = &redispool.Pool{
			MaxIdle:     MaxIdle,
			MaxActive:   MaxActive,
			IdleTimeout: time.Duration(IdleTimeout) * time.Second,
			Dial: func() (redispool.Conn, error) {
				return redispool.Dial("tcp", addr, redispool.DialDatabase(db))
			},
		}
	} else {
		pwdOption := redispool.DialPassword(pwd)
		CacheRedisPool = &redispool.Pool{
			MaxIdle:     MaxIdle,
			MaxActive:   MaxActive,
			IdleTimeout: time.Duration(IdleTimeout) * time.Second,
			Dial: func() (redispool.Conn, error) {
				return redispool.Dial("tcp", addr, pwdOption, redispool.DialDatabase(db))
			},
		}
	}

	return &redisPoolCache{client: CacheRedisPool}, nil

}

func (rp *redisPoolCache) Set(key string, value interface{}, expire int) error {
	cacheConn := rp.client.Get()
	defer cacheConn.Close()
	_, err := redispool.String(cacheConn.Do("SETEX", key, expire, string(value.(string))))
	return err
}

func (rp *redisPoolCache) Get(key string) (interface{}, error) {
	cacheConn := rp.client.Get()
	defer cacheConn.Close()
	result, err := redispool.String(cacheConn.Do("GET", key))
	if err != nil {
		if err == redispool.ErrNil {
			result = ""
		} else {
			fmt.Println("cacheConn.Do(\"GET\", key)", err)
			return nil, err
		}
	}
	return result, err
}

func (rp *redisPoolCache) Del(key string) error {
	cacheConn := rp.client.Get()
	defer cacheConn.Close()
	_, err := redispool.String(cacheConn.Do("DEL", key))
	return err
}

func (rp *redisPoolCache) Close() error {
	return rp.client.Close()
}
