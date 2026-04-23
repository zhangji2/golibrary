package utils

import (
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

// IP工具函数

// LocalIP 获取本地ip
func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// ServerIP 获取服务器的外网ip
func ServerIP(s ...string) string {
	// 默认值
	defaultS := "https://api.ipify.org?format=text" //国外的
	if len(s) > 0 && s[0] != "" {
		// https://api.test.com/ip.php //自定义一个国内的接口url
		defaultS = s[0]
	}

	// 创建带超时的 HTTP 客户端
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	// 访问外部服务获取 IP
	resp, err := client.Get(defaultS)

	if err != nil {
		// fmt.Errorf("请求公网 IP 失败: %v", err)
		return ""
	}
	defer resp.Body.Close()
	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// fmt.Errorf("读取响应失败: %v", err)
		return ""
	}
	ip := strings.TrimSpace(string(body))
	// 校验是否为合法 IP
	if net.ParseIP(ip) == nil {
		// fmt.Errorf("返回的 IP 无效: %s", ip)
		return ""
	}
	return ip
}

// ClientIP 获取客户端ip
// func ClientIP(c *gin.Context) string {
// 	// 1. 优先从 X-Forwarded-For 获取
// 	ip := c.GetHeader("X-Forwarded-For")
// 	if ip != "" {
// 		return strings.Split(ip, ",")[0]
// 	}
// 	// 2. 从 X-Real-IP 获取
// 	ip = c.GetHeader("X-Real-IP")
// 	if ip != "" {
// 		return ip
// 	}
// 	// 3. 兜底用 Gin 内置方法
// 	return c.ClientIP()
// }
