package utils

import "crypto/sha256"

// Encrypt 随机加密盐对字符串加密，返回哈希值
func Encrypt(s string, salt []byte) []byte {
	e := sha256.New()
	e.Write([]byte(s))
	e.Write(salt)
	return e.Sum(nil)
}
