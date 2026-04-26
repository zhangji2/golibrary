package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDriver struct{}
type mysqlDB struct {
	db *gorm.DB
}

func init() {
	Register("mysql", &mysqlDriver{})
}

func (d *mysqlDriver) Open(config map[string]string) (DB, error) {
	host := config["host"]
	port := config["port"]
	user := config["user"]
	password := config["password"]
	dbname := config["dbname"]

	// 1. 拼接目标库 DSN
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname,
	)

	// 2. 尝试直接连接目标库
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		// 连接成功，直接返回
		return &mysqlDB{db: gormDB}, nil
	}

	// 3. 只有报错【不存在】才自动创建
	errMsg := strings.ToLower(err.Error())
	if strings.Contains(errMsg, "unknown database") || strings.Contains(errMsg, "doesn't exist") {
		fmt.Printf("数据库 %s 不存在，正在创建...\n", dbname)
		// 连接管理库
		dsnAdmin := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/mysql?charset=utf8mb4&parseTime=True&loc=Local",
			user, password, host, port,
		)
		dbAdmin, errAdmin := sql.Open("mysql", dsnAdmin)
		if errAdmin != nil {
			return nil, fmt.Errorf("连接管理库失败: %w", errAdmin)
		}
		defer dbAdmin.Close()

		// 创建数据库
		_, execErr := dbAdmin.Exec(fmt.Sprintf(
			"CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci",
			dbname,
		))
		if execErr != nil {
			return nil, fmt.Errorf("创建数据库失败: %w", execErr)
		}

		// 重新连接
		gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}

	return &mysqlDB{db: gormDB}, nil
}

func (m *mysqlDB) DB() *gorm.DB {
	return m.db
}

func (m *mysqlDB) Close() error {
	sdb, _ := m.db.DB()
	return sdb.Close()
}
