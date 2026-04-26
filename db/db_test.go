package db

import (
	"fmt"
	"log"
	"testing"
)

// User 测试模型
type TestModelUser struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size:32"`
	Age  int
}

func dbHandler(conf ConfigDB) error {
	dbClient, err := NewDB(conf)
	if err != nil {
		log.Println("数据库("+conf.Driver+") 初始化失败:", err)
	} else {

		fmt.Println("数据库(" + conf.Driver + ")  连接成功")

		// 自动建表
		_ = dbClient.DB().AutoMigrate(&TestModelUser{})

		// 1. 增
		fmt.Println("\n--- 新增")
		u1 := &TestModelUser{Name: "张三", Age: 20}
		dbClient.DB().Create(u1)
		fmt.Println("新增成功 ID:", u1.ID)

		u2 := &TestModelUser{Name: "王五", Age: 31}
		dbClient.DB().Create(u2)
		fmt.Println("新增成功 ID:", u2.ID)

		// 2. 查
		fmt.Println("\n--- 查询")
		var u TestModelUser
		dbClient.DB().First(&u, u1.ID)
		fmt.Println("查询结果:", u)

		// 3. 改
		fmt.Println("\n--- 修改")
		dbClient.DB().Model(&u).Update("age", 22)
		fmt.Println("修改后年龄:", u.Age)

		// 4. 删
		fmt.Println("\n--- 删除")
		dbClient.DB().Delete(&u)
		fmt.Println("删除成功")

		defer dbClient.Close()
	}
	return nil
}

func TestMysql(t *testing.T) {
	mysqlConf := ConfigDB{
		Driver: "mysql",
		Config: map[string]string{
			"host":     "127.0.0.1",
			"port":     "8306",
			"user":     "root",
			"password": "admin",
			"dbname":   "test",
		},
	}
	dbHandler(mysqlConf)
}

func TestPostgreSQL(t *testing.T) {
	pgConf := ConfigDB{
		Driver: "postgres",
		Config: map[string]string{
			"host":     "127.0.0.1",
			"port":     "5432",
			"user":     "zero",
			"password": "ServBay.dev",
			"dbname":   "test",
			"sslmode":  "disable",
		},
	}
	dbHandler(pgConf)
}

func TestSQLite(t *testing.T) {
	sqliteConf := ConfigDB{
		Driver: "sqlite",
		Config: map[string]string{
			"path": "test.db",
		},
	}
	dbHandler(sqliteConf)
}
