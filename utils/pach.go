package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 路径工具函数

// RecurisonListPath 路径分词
func RecurisonListPath(path string, slice *[]string) {
	if path == "/" {
		return
	}
	path2 := filepath.Dir(path)
	RecurisonListPath(path2, slice)
	*slice = append(*slice, path)
}

// GetCurrPath 获取当前运行路径
func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}
