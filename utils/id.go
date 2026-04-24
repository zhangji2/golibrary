package utils

import "github.com/google/uuid"

// ID工具函数

//GetGuid 唯一识别码
func GetGuid() string {
	return uuid.New().String()
}
