package cache

import "errors"

// ConfigCache 外层统一入参结构体
type ConfigCache struct {
	Driver string            // 驱动名: redis / memcache
	Config map[string]string // 自定义配置
}

// Cache 统一缓存接口
type Cache interface {
	Set(key string, value interface{}, expireSeconds int) error
	Get(key string) (interface{}, error)
	Del(key string) error
	Close() error
}

// Driver 驱动工厂接口
type Driver interface {
	Open(cfg map[string]string) (Cache, error)
}

var drivers = make(map[string]Driver)

// Register 注册驱动
func Register(name string, d Driver) {
	drivers[name] = d
}

// NewCache 工厂方法创建缓存实例
func NewCache(conf ConfigCache) (Cache, error) {
	d, ok := drivers[conf.Driver]
	if !ok {
		return nil, errors.New("不支持的缓存驱动: " + conf.Driver)
	}
	return d.Open(conf.Config)
}
