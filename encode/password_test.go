package encode

import (
	"testing"
)

var str_password_encode = "apple,banana,apple,cherry,banana"
var str_password_salt = "your_salt_here"

func TestEncryptPassword(t *testing.T) {
	t.Log("EncryptPassword:", EncryptPassword(str_password_encode, str_password_salt))
}

func TestCheckPassword(t *testing.T) {
	t.Log("CheckPassword:", CheckPassword(EncryptPassword(str_password_encode, str_password_salt), str_password_encode, str_password_salt))
}

// 在项目中使用：先引入包，再调函数。
// zencode "github.com/zhangji2/golibrary/encode"
// fmt.Printf("EncryptPassword: %v\n", zencode.EncryptPassword(str_password_encode, str_password_salt))
// fmt.Printf("CheckPassword: %v\n", zencode.CheckPassword(zencode.EncryptPassword(str_password_encode, str_password_salt), str_password_encode, str_password_salt))
