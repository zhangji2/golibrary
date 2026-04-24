package encode

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 密码相关函数

// EncryptPassword 加密用户密码
func EncryptPassword(password string, salt string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	return string(hash)
}

// CheckPassword 检查密码是否正确
func CheckPassword(oldPassword, newPassword, salt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(newPassword+salt))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
