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

func TestConfig(t *testing.T) {

	// 初始化Viper配置
	var ViperConfig ViperConfigStruct
	ViperConfig.Path = "./"
	ViperConfig.File = "config_test_app"

	// 初始化Nacos配置
	var appConfig AppConfigStruct

	err := NewViper(ViperConfig, &appConfig)
	if err != nil {
		fmt.Println("\n初始化Nacos失败1:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		panic(err)
	}

	// 打印结果
	fmt.Printf("\n获取App配置成功1：%v", appConfig)

	// 初始化配置
	var nacosConfig NacosConfigStruct

	ViperConfig.File = "config_test_nacos"

	err = NewViper(ViperConfig, &nacosConfig)
	if err != nil {
		fmt.Println("\n初始化Nacos失败2:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		panic(err)
	}
	// 打印结果
	fmt.Printf("\n获取Nacos配置成功1：%v", nacosConfig)

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

	// 打印结果
	fmt.Printf("\n获取Nacos配置成功2：%v", nacosConfig)

	err = NewNacos(nacosConfig, &appConfig)
	if err != nil {
		fmt.Println("\n初始化Nacos失败:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		panic(err)
	}

	// 打印结果
	fmt.Printf("\n获取App配置成功2：%v", appConfig)

	// 断言结果
	require.True(t, appConfig.App.Debug)                  // require 失败直接终止，后面的代码不会执行
	assert.Nil(t, err)                                    // 必须是 nil
	assert.NoError(t, err)                                // 必须没有错误
	assert.NotEmpty(t, appConfig.Redis.Addr)              // 不为空
	assert.Equal(t, appConfig.Redis.Password, "foobared") // 字符串相等
	assert.Greater(t, appConfig.Redis.Db, 1)              // 大于1
	assert.Equal(t, appConfig.Redis.Db, 20)               // 等于20
	assert.GreaterOrEqual(t, appConfig.Redis.Db, 1)       // 大于等于1
	assert.Less(t, appConfig.Redis.Db, 26)                // 小于16
	assert.Contains(t, appConfig.Mysql.Dsn, "3306")       // 字符串包含
}
