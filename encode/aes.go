package encode

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

var ErrInvalidPadding = errors.New("invalid PKCS padding")

func AesEncryptCBC(a string, k string) string {
	origData := []byte(a) // 待加密的数据
	key := []byte(k)      // 加密的密钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted := make([]byte, len(origData))                    // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return hex.EncodeToString(encrypted)
}

// AesDecryptCBC 解密 Base64 编码的 AES-CBC 密文（与 AesEncryptCBC 配对）。
func AesDecryptCBC(a string, k string) (string, error) {
	key := []byte(k)
	raw, err := hex.DecodeString(a)
	if err != nil {
		return "hex.DecodeString", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "aes.NewCipher", err
	}
	blockSize := block.BlockSize()
	if len(raw) == 0 || len(raw)%blockSize != 0 {
		return "block.BlockSize", fmt.Errorf("invalid ciphertext length: %d blockSize: %d", len(raw), blockSize)
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	decrypted := make([]byte, len(raw))
	blockMode.CryptBlocks(decrypted, raw)
	aaa, err := pkcs5UnPadding(decrypted, blockSize)
	return string(aaa), err
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte, blockSize int) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, ErrInvalidPadding
	}
	if length%blockSize != 0 {
		return nil, ErrInvalidPadding
	}
	unpadding := int(origData[length-1])
	if unpadding <= 0 || unpadding > blockSize || unpadding > length {
		return nil, ErrInvalidPadding
	}
	for i := 0; i < unpadding; i++ {
		if origData[length-1-i] != byte(unpadding) {
			return nil, ErrInvalidPadding
		}
	}
	return origData[:length-unpadding], nil
}

func AesEncryptECB(a string, k string) string {
	origData := []byte(a)
	key := []byte(k)
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return hex.EncodeToString(encrypted)
}

func AesDecryptECB(a string, k string) (string, error) {
	key := []byte(k)
	raw, err := hex.DecodeString(a)
	if err != nil {
		return "hex.DecodeString", err
	}
	cipher, err := aes.NewCipher(generateKey(key))
	if err != nil {
		return "aes.NewCipher", err
	}
	decrypted := make([]byte, len(raw))
	for bs, be := 0, cipher.BlockSize(); bs < len(raw); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], raw[bs:be])
	}
	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return string(decrypted[:trim]), nil
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func AesEncryptCFB(a string, k string) (string, error) {
	origData := []byte(a)
	key := []byte(k)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "aes.NewCipher", err
	}
	encrypted := make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "io.ReadFull", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return hex.EncodeToString(encrypted), nil
}

func AesDecryptCFB(a string, k string) (string, error) {
	key := []byte(k) // 加密的密钥
	block, err := aes.NewCipher(key)
	if err != nil {
		return "aes.NewCipher", err
	}
	raw, err := hex.DecodeString(a)
	if err != nil {
		return "hex.DecodeString", err
	}
	if len(raw) < aes.BlockSize {
		return "aes.BlockSize", fmt.Errorf("ciphertext too short length: %d blockSize: %d", len(raw), aes.BlockSize)
	}
	iv := raw[:aes.BlockSize]
	encrypted := raw[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return string(encrypted), nil
}
