package config

import (
	"fmt"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"gopkg.in/yaml.v2"
)

type NacosConfigStruct struct {
	Addr        string `mapstructure:"addr"`
	Port        uint64 `mapstructure:"port"`
	Scheme      string `mapstructure:"scheme"`
	ContextPath string `mapstructure:"context_path"`
	NameSpace   string `mapstructure:"name_space"`
	DataID      string `mapstructure:"data_id"`
	DataGroup   string `mapstructure:"data_group"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	RuntimeDir  string `mapstructure:"runtime_dir"`
	LogLevel    string `mapstructure:"log_level"`
}

// 初始化配置中心客户端
var NaocsConfigClient config_client.IConfigClient

// 初始化服务中心客户端
var NaocsNamingClient naming_client.INamingClient

func NewNacos(nacosConfig NacosConfigStruct, configInfo interface{}) error {
	err := fmt.Errorf("nacos address is empty")
	if nacosConfig.Addr == "" {
		fmt.Println("\n连接Nacos配置中心失败1:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		return err
	}

	sc := []constant.ServerConfig{
		{
			IpAddr: nacosConfig.Addr,
			Port:   nacosConfig.Port,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         nacosConfig.NameSpace, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              nacosConfig.RuntimeDir + "/logs",
		CacheDir:            nacosConfig.RuntimeDir + "/cache",
		LogRollingConfig: &constant.ClientLogRollingConfig{
			MaxSize: 100,
			MaxAge:  3,
		},
		LogLevel: nacosConfig.LogLevel,
		Username: nacosConfig.User,
		Password: nacosConfig.Password,
	}

	NaocsConfigClient, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		fmt.Println("\n连接Nacos配置中心失败:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		panic(err)
	}

	NaocsNamingClient, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		fmt.Println("\n连接Nacos服务中心失败:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		panic(err)
	}

	data, err := NaocsConfigClient.GetConfig(vo.ConfigParam{
		DataId: nacosConfig.DataID,
		Group:  nacosConfig.DataGroup,
	})
	if err != nil {
		fmt.Println("\n获取Nacos远程配置失败:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		panic(err)
	}

	// fmt.Println("Nacos原始内容：\n", data, " ", time.Now().Format("2006-01-02 15:04:05"))
	err = yaml.Unmarshal([]byte(data), configInfo)
	if err != nil {
		fmt.Println("\n监听远程配置中心失败", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		return err
	}

	// 监听配置是否有变化
	err = NaocsConfigClient.ListenConfig(vo.ConfigParam{
		DataId: nacosConfig.DataID,
		Group:  nacosConfig.DataGroup,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("Nacos配置有变化 group:" + group + ", dataId:" + dataId)
			err = yaml.Unmarshal([]byte(data), configInfo)
			if err != nil {
				fmt.Println("\n监听远程配置中心失败", err, " ", time.Now().Format("2006-01-02 15:04:05"))
				return
			}
			fmt.Println("\nNacos config file changed:", configInfo)
		},
	})

	if err != nil {
		fmt.Println("\n监听远程配置中心初始化失败", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		panic(err)
	}

	return nil
}

type nacosNaming struct {
	Ip          string
	PortOutGrpc uint64
	PortOutHttp uint64
	ServiceName string
}

// 初始化Nacos服务注册
var NacosNaming nacosNaming

// 服务注册客户端
type nacosNonfigCenter struct {
}

// 初始化注册中心
var NacosConfigCenter nacosNonfigCenter

// 获取配置注册中心
func (s *nacosNonfigCenter) GetConfigClient() (config_client.IConfigClient, error) {
	return NaocsConfigClient, nil
}

// 获取服务注册中心
func (s *nacosNonfigCenter) GetNamingClient() (naming_client.INamingClient, error) {
	return NaocsNamingClient, nil
}

// 获取一个服务
func (s *nacosNonfigCenter) GetOneService(name string, clusters []string) (*model.Instance, error) {
	namingClient, err := NacosConfigCenter.GetNamingClient()
	if err != nil {
		fmt.Printf("\n连接服务中心失败:%s|%s time:%s \n", name, err.Error(), time.Now().Format("2006-01-02 15:04:05"))
		return nil, err
	}
	instance, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
		Clusters:    clusters,
	})
	if err != nil {
		fmt.Printf("\n获取健康服务实例错误:%s|%s time:%s \n", name, err.Error(), time.Now().Format("2006-01-02 15:04:05"))
		return nil, err
	}
	return instance, nil
}

// 获取多个服务
func (s *nacosNonfigCenter) GetAllService(name string, clusters []string) ([]model.Instance, error) {
	namingClient, err := NacosConfigCenter.GetNamingClient()
	if err != nil {
		fmt.Printf("\n连接服务中心失败:%s|%s time:%s \n", name, err.Error(), time.Now().Format("2006-01-02 15:04:05"))
		return nil, err
	}
	instance, err := namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: name,
		Clusters:    clusters,
	})
	if err != nil {
		fmt.Printf("\n获取所有服务实例错误:%s|%s time:%s \n", name, err.Error(), time.Now().Format("2006-01-02 15:04:05"))
		return nil, err
	}
	return instance, nil
}

// 注册服务
func (s *nacosNonfigCenter) RegisterServer(naming nacosNaming) {
	fmt.Println("\n开始Nacos注册服务", "|time:", time.Now().Format("2006-01-02 15:04:05"))

	client, err := s.GetNamingClient()
	if err != nil {
		fmt.Println("\n连接Nacos注册中心失败:", err, "|time:", time.Now().Format("2006-01-02 15:04:05"))
		panic(err)
	}

	//注册grpc服务
	if naming.PortOutGrpc > 0 {
		param := vo.RegisterInstanceParam{
			Ip:          naming.Ip,
			Port:        naming.PortOutGrpc,
			ServiceName: fmt.Sprintf("%s.%s", naming.ServiceName, "grpc"),
			Weight:      10,
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
			Metadata:    map[string]string{},
		}
		success, err := client.RegisterInstance(param)
		if err != nil {
			fmt.Println("\nNacos注册服务失败-Grpg:", err, "|time:", time.Now().Format("2006-01-02 15:04:05"))
			panic(err)
		}
		fmt.Printf("\nRegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
	}

	//注册http服务
	if naming.PortOutHttp > 0 {
		param := vo.RegisterInstanceParam{
			Ip:          naming.Ip,
			Port:        naming.PortOutHttp,
			ServiceName: fmt.Sprintf("%s.%s", naming.ServiceName, "http"),
			Weight:      10,
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
			Metadata:    map[string]string{},
		}
		success, err := client.RegisterInstance(param)
		if err != nil {
			fmt.Println("\nNacos注册服务失败-Http:", err, "|time:", time.Now().Format("2006-01-02 15:04:05"))
			panic(err)
		}
		fmt.Printf("\nRegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
	}
}
