package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

// JWT 验证结构体
type JWT struct {
	SecretKey     string `yaml:"secret_key"`     // 密钥
	ExpireMinutes int    `yaml:"expire_minutes"` // 过期时间
}

var j JWT
var jwtKeySlice = []byte(j.SecretKey)

// Claims 生成token
type Claims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

// SignToken 生成token并返回
func SignToken(id int64, name string) (string, error) {
	expirationTime := time.Now().Add(1024 * time.Second)
	claims := &Claims{
		Username: name,
		ID:       id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		}}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKeySlice)
	return token, err
}

// ParseToken 解析token并返回claimsid
func ParseToken(token string) (int64, error) {
	claims := &Claims{}

	tk, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKeySlice, nil
	})

	if err == nil && tk.Valid {
		return claims.ID, nil
	}

	return -1, err
}
