package utils

import (
	"testing"
)

func TestLocalIP(t *testing.T) {
	t.Log("LocalIP:", LocalIP())
}

func TestServerIP(t *testing.T) {
	t.Log("ServerIP:", ServerIP())
	t.Log("ServerIP:", ServerIP("https://api.test.net/ip.php"))
}

// 在项目中使用：先引入包，再调函数。
// zutils "github.com/zhangji2/golibrary/utils"
// fmt.Printf("InArray: %v\n", zutils.InArray(arr_a, "apple"))
// fmt.Printf("ClientIP: %v\n", zutils.ClientIP())
// fmt.Printf("LocalIP: %v\n", zutils.LocalIP())
// fmt.Printf("ServerIP: %v\n", zutils.ServerIP())
// fmt.Printf("RemoteIP: %v\n", zutils.RemoteIP())
