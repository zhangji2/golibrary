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

type nacosConfig struct {
	Addr        string
	Port        uint64
	Scheme      string
	ContextPath string
	NameSpace   string
	DataID      string
	DataGroup   string
	User        string
	Password    string
	RuntimeDir  string
	LogLevel    string
}

// 初始化Nacos配置
var NacosConfig nacosConfig

// 初始化配置中心客户端
var NaocsConfigClient config_client.IConfigClient

// 初始化服务中心客户端
var NaocsNamingClient naming_client.INamingClient

func NewNacos(config nacosConfig, configInfo interface{}) error {
	err := fmt.Errorf("nacos address is empty")
	if config.Addr == "" {
		fmt.Println("连接Nacos配置中心失败1:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		return err
	}

	sc := []constant.ServerConfig{
		{
			IpAddr: config.Addr,
			Port:   config.Port,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         config.NameSpace, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              config.RuntimeDir + "/logs",
		CacheDir:            config.RuntimeDir + "/cache",
		LogRollingConfig: &constant.ClientLogRollingConfig{
			MaxSize: 100,
			MaxAge:  3,
		},
		LogLevel: config.LogLevel,
		Username: config.User,
		Password: config.Password,
	}

	NaocsConfigClient, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		fmt.Println("连接Nacos配置中心失败:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		panic(err)
	}

	NaocsNamingClient, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		fmt.Println("连接Nacos服务中心失败:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		panic(err)
	}

	data, err := NaocsConfigClient.GetConfig(vo.ConfigParam{
		DataId: config.DataID,
		Group:  config.DataGroup,
	})
	if err != nil {
		fmt.Println("获取Nacos远程配置失败:", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		panic(err)
	}

	// fmt.Println("Nacos原始内容：\n", data, " ", time.Now().Format("2006-01-02 15:04:05"))
	err = yaml.Unmarshal([]byte(data), configInfo)
	if err != nil {
		fmt.Println("监听远程配置中心失败", err, " ", time.Now().Format("2006-01-02 15:04:05"))
		return err
	}

	// 监听配置变化
	err = NaocsConfigClient.ListenConfig(vo.ConfigParam{
		DataId: config.DataID,
		Group:  config.DataGroup,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("Nacos配置有变化 group:" + group + ", dataId:" + dataId)
			err = yaml.Unmarshal([]byte(data), configInfo)
			if err != nil {
				fmt.Println("监听远程配置中心失败", err, " ", time.Now().Format("2006-01-02 15:04:05"))
				return
			}
		},
	})

	if err != nil {
		fmt.Println("监听远程配置中心初始化失败", err, " ", time.Now().Format("2006-01-02 15:04:05"))
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
		fmt.Printf("连接服务中心失败:%s|%s time:%s \n", name, err.Error(), time.Now().Format("2006-01-02 15:04:05"))
		return nil, err
	}
	instance, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: name,
		Clusters:    clusters,
	})
	if err != nil {
		fmt.Printf("获取健康服务实例错误:%s|%s time:%s \n", name, err.Error(), time.Now().Format("2006-01-02 15:04:05"))
		return nil, err
	}
	return instance, nil
}

// 获取多个服务
func (s *nacosNonfigCenter) GetAllService(name string, clusters []string) ([]model.Instance, error) {
	namingClient, err := NacosConfigCenter.GetNamingClient()
	if err != nil {
		fmt.Printf("连接服务中心失败:%s|%s time:%s \n", name, err.Error(), time.Now().Format("2006-01-02 15:04:05"))
		return nil, err
	}
	instance, err := namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: name,
		Clusters:    clusters,
	})
	if err != nil {
		fmt.Printf("获取所有服务实例错误:%s|%s time:%s \n", name, err.Error(), time.Now().Format("2006-01-02 15:04:05"))
		return nil, err
	}
	return instance, nil
}

// 注册服务
func (s *nacosNonfigCenter) RegisterServer(naming nacosNaming) {
	fmt.Println("开始Nacos注册服务", "|time:", time.Now().Format("2006-01-02 15:04:05"))

	client, err := s.GetNamingClient()
	if err != nil {
		fmt.Println("连接Nacos注册中心失败:", err, "|time:", time.Now().Format("2006-01-02 15:04:05"))
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
			fmt.Println("Nacos注册服务失败-Grpg:", err, "|time:", time.Now().Format("2006-01-02 15:04:05"))
			panic(err)
		}
		fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
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
			fmt.Println("Nacos注册服务失败-Http:", err, "|time:", time.Now().Format("2006-01-02 15:04:05"))
			panic(err)
		}
		fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
	}
}
