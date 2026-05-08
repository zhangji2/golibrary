package logger

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	// 初始化log配置
	var LogConfig LogConfigStruct
	LogConfig.Path = "../tmp/service/logs"
	LogConfig.Path = "/Users/zhangji/open_source/golibrary/tmp/service/logs"
	// 2. 转换为绝对路径
	absPath, err := filepath.Abs(LogConfig.Path)
	if err != nil {
		log.Fatalf("路径转换失败：%v", err)
	}
	LogConfig.Path = absPath
	LogConfig.FileName = "test00"
	LogConfig.Debug = true
	LogConfig.UdpIp = "127.0.0.1:9001"
	Log = NewLogger(LogConfig)
}

func TestLogger(t *testing.T) {

	// 测试不同级别的日志
	Log.WithFields(logrus.Fields{
		"task_id": "123456",
		"module":  "auth",
	}).Info("用户登录成功")

	Log.WithFields(logrus.Fields{
		"task_id": "123457",
		"module":  "payment",
	}).Warn("支付金额异常，请检查")

	Log.WithFields(logrus.Fields{
		"task_id": "123458",
		"module":  "database",
	}).Error("数据库连接失败")

	Log.WithFields(logrus.Fields{
		"task_id": "123459",
		"module":  "api",
	}).Debug("API请求详情: method=GET, path=/users")

	time.Sleep(5 * time.Second) // 等待异步发送完成

	Log.Info("服务启动成功")
	Log.Warn("系统负载较高")
	Log.Error("数据库连接异常")

	Log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}

func TestUdp(t *testing.T) {
	// 1. 解析服务端地址
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9001")
	if err != nil {
		fmt.Printf("解析服务端地址失败：%v\n", err)
		os.Exit(1)
	}

	// 2. 建立 UDP 连接
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Printf("连接服务端失败：%v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 3. 发送消息
	msg := "Hello UDP Server from Go!"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Printf("发送消息失败：%v\n", err)
		os.Exit(1)
	}
	fmt.Println("已发送：", msg)

	// 4. 接收服务端的回复
	buf := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Printf("接收回复失败：%v\n", err)
		return
	}
	fmt.Println("服务端回复：", string(buf[:n]))
}

// 在项目中使用：先引入包，再调函数。
// zlogger "github.com/zhangji2/golibrary/logger"
