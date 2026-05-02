package config

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ViperConfigStruct struct {
	Path     string
	FileName string
}

func NewViper(viperConfig ViperConfigStruct, configInfo interface{}) error {
	err := fmt.Errorf("config file path is empty")
	if viperConfig.Path == "" || viperConfig.FileName == "" {
		fmt.Println("\nViper配置无效:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		return err
	}

	fmt.Println("\nviperConfig", viperConfig)
	fmt.Println("\nconfigInfo", configInfo)

	v := viper.New()
	v.AddConfigPath(viperConfig.Path)
	v.SetConfigName(viperConfig.FileName)

	err = v.ReadInConfig()
	if err != nil {
		fmt.Println("\nFailed to load local configuration file", err)
		return err
	}

	// 兼容如 config_test_nacos.yaml 这类以顶层 key 包裹的配置结构
	if v.IsSet("nacos") {
		err = v.UnmarshalKey("nacos", configInfo)
	} else {
		err = v.Unmarshal(configInfo)
	}
	if err != nil {
		fmt.Println("\nParsing local configuration file failed", err)
		return err
	}

	// 监听配置是否有变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		err = v.Unmarshal(configInfo)
		if err != nil {
			fmt.Println("\nConfiguration changes, parsing local configuration file failed", err)
			return
		}
		fmt.Println("\nViper config file changed:", configInfo)
	})
	return nil
}
