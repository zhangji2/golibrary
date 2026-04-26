package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDriver struct{}
type postgresDB struct {
	db *gorm.DB
}

func init() {
	Register("postgres", &postgresDriver{})
}

func (d *postgresDriver) Open(config map[string]string) (DB, error) {
	host := config["host"]
	port := config["port"]
	user := config["user"]
	password := config["password"]
	dbname := config["dbname"]
	sslmode := config["sslmode"]

	// 1. 目标库DSN
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	// 2. 尝试连接
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		return &postgresDB{db: gormDB}, nil
	}

	// 3. 数据库不存在 → 创建
	errMsg := strings.ToLower(err.Error())
	if strings.Contains(errMsg, strings.ToLower("database \""+dbname+"\" does not exist")) {
		fmt.Printf("数据库 %s 不存在，正在创建...\n", dbname)
		dsnAdmin := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=postgres sslmode=%s",
			host, port, user, password, sslmode,
		)
		dbAdmin, errAdmin := sql.Open("postgres", dsnAdmin)
		if errAdmin != nil {
			return nil, errAdmin
		}
		defer dbAdmin.Close()

		// 创建
		_, execErr := dbAdmin.Exec(fmt.Sprintf(`CREATE DATABASE "%s"`, dbname))
		if execErr != nil {
			return nil, execErr
		}

		// 重连
		gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}

	return &postgresDB{db: gormDB}, err
}

func (p *postgresDB) DB() *gorm.DB {
	return p.db
}

func (p *postgresDB) Close() error {
	sdb, _ := p.db.DB()
	return sdb.Close()
}
