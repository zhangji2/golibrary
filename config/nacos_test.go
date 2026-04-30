package config

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type AppConfigInfoStruct struct {
	App struct {
		Name  string `yaml:"name"`
		Debug bool   `yaml:"debug"`
	} `yaml:"app"`
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		Db       int    `yaml:"db"`
	} `yaml:"redis"`
	Mysql struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"mysql"`
	RabbitMQ struct {
		Uri string `yaml:"uri"`
	} `yaml:"rabbitmq"`
}

// test-app.yaml
// app:
//   name: test
//   debug: true
// redis:
//   addr: host.docker.internal:6379
//   password: foobared
//   db: 20
// mysql:
//   dsn: root:123@tcp(127.0.0.1:3306)/test
// rabbitmq:
//   uri: amqp://admin:123456@127.0.0.1:5672/default_vhost

var AppConfInfo AppConfigInfoStruct

func TestNacosConfigClient(t *testing.T) {

	// 初始化Nacos配置
	var nacosConfig nacosConfig

	nacosConfig.Addr = "127.0.0.1"
	nacosConfig.Port = uint64(18848)
	nacosConfig.User = "nacos"
	nacosConfig.Password = "nacos"
	nacosConfig.ContextPath = "/naocs"
	nacosConfig.NameSpace = "f1ec9270-ef54-4991-ad7a-bd02b741f375"
	nacosConfig.DataID = "test-app.yaml"
	nacosConfig.DataGroup = "DEFAULT_GROUP"
	nacosConfig.RuntimeDir = "../tmp/nacos"
	nacosConfig.LogLevel = "info"

	cnfFromEnv := os.Getenv("CONFIG_FROM_ENV")
	if len(cnfFromEnv) > 0 {
		nacosConfig.Addr = os.Getenv("NACOS_ADDR")
		nacosPort, _ := strconv.Atoi(os.Getenv("NACOS_PORT"))
		nacosConfig.Port = uint64(nacosPort)
		nacosConfig.User = os.Getenv("NACOS_USER")
		nacosConfig.Password = os.Getenv("NACOS_PASSWORD")
		nacosConfig.ContextPath = os.Getenv("NACOS_CONTEXT_PATH")
		nacosConfig.NameSpace = os.Getenv("NACOS_NAME_SPACE")
		nacosConfig.DataID = os.Getenv("NACOS_DATA_ID")
		nacosConfig.DataGroup = os.Getenv("NACOS_DATA_GROUP")
		nacosConfig.RuntimeDir = os.Getenv("NACOS_Runtime_Dir")
		nacosConfig.LogLevel = os.Getenv("NACOS_LOG_LEVEL")
	}

	err := NewNacos(nacosConfig, &AppConfInfo)
	if err != nil {
		fmt.Println("初始化Nacos失败:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		panic(err)
	}

	// 打印结果
	fmt.Printf("获取Nacos配置成功：%v", AppConfInfo)

	// 断言结果
	require.True(t, AppConfInfo.App.Debug)                  // require 失败直接终止，后面的代码不会执行
	assert.Nil(t, err)                                      // 必须是 nil
	assert.NoError(t, err)                                  // 必须没有错误
	assert.NotEmpty(t, AppConfInfo.Redis.Addr)              // 不为空
	assert.Equal(t, AppConfInfo.Redis.Password, "foobared") // 字符串相等
	assert.Greater(t, AppConfInfo.Redis.Db, 1)              // 大于1
	assert.Equal(t, AppConfInfo.Redis.Db, 20)               // 等于20
	assert.GreaterOrEqual(t, AppConfInfo.Redis.Db, 1)       // 大于等于1
	assert.Less(t, AppConfInfo.Redis.Db, 26)                // 小于16
	assert.Contains(t, AppConfInfo.Mysql.Dsn, "3306")       // 字符串包含
}
