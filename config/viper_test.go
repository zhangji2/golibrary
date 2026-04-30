package config

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewViper(t *testing.T) {

	// 初始化Viper配置
	var ViperConfigInfo ViperConfigStruct
	ViperConfigInfo.Path = "./"
	ViperConfigInfo.File = "config_test_app"

	// 初始化Nacos配置
	var appConfig AppConfigStruct

	err := NewViper(ViperConfigInfo, &appConfig)
	if err != nil {
		fmt.Println("\n初始化Nacos失败:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		panic(err)
	}

	// 打印结果
	fmt.Printf("\n获取App配置成功：%v", appConfig)

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
