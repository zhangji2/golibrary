package cache

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type memcacheDriver struct{}
type memcacheCache struct {
	client *memcache.Client
}

func init() {
	Register("memcache", new(memcacheDriver))
}

func (m *memcacheDriver) Open(cfg map[string]string) (Cache, error) {
	addr := cfg["addr"]
	cli := memcache.New(addr)
	if err := cli.Ping(); err != nil {
		return nil, fmt.Errorf("memcache 连接失败: %w", err)
	}
	return &memcacheCache{client: cli}, nil
}

func (m *memcacheCache) Set(key string, value interface{}, expire int) error {
	s, ok := value.(string)
	if !ok {
		return fmt.Errorf("memcache 仅支持 string 类型值")
	}
	return m.client.Set(&memcache.Item{
		Key:        key,
		Value:      []byte(s),
		Expiration: int32(expire),
	})
}

func (m *memcacheCache) Get(key string) (interface{}, error) {
	item, err := m.client.Get(key)
	if err != nil {
		return nil, err
	}
	return string(item.Value), nil
}

func (m *memcacheCache) Del(key string) error {
	return m.client.Delete(key)
}

func (m *memcacheCache) Close() error {
	return nil
}
