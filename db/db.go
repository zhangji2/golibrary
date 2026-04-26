package db

import (
	"errors"

	"gorm.io/gorm"
)

// ConfigDB 外部统一传入的配置
type ConfigDB struct {
	Driver string            // mysql / postgres / sqlite
	Config map[string]string // 连接参数
}

// DB 统一数据库接口
type DB interface {
	DB() *gorm.DB
	Close() error
}

// Driver 驱动接口
type Driver interface {
	Open(config map[string]string) (DB, error)
}

// 驱动注册中心
var drivers = make(map[string]Driver)

// Register 注册数据库驱动
func Register(name string, driver Driver) {
	drivers[name] = driver
}

// NewDB 工厂方法：根据配置创建数据库实例
func NewDB(conf ConfigDB) (DB, error) {
	driver, ok := drivers[conf.Driver]
	if !ok {
		return nil, errors.New("unsupported database driver: " + conf.Driver)
	}
	return driver.Open(conf.Config)
}
