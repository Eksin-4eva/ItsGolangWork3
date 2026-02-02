package utils

import (
	"errors"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 定义密钥
var jwtSecret = []byte("TodolistKey")

type Claims struct {
	ID                   uint   `json:"id"`
	UserName             string `json:"user_name"`
	jwt.RegisteredClaims        // 包含 ExpiredAt, Issuer
}

// 签发Token
func GenerateToken(id uint, userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour) // token有效期24小时

	claims := Claims{
		ID:       id,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "Tarisu",
		},
	}

	//签名
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 解析Token函数
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 检测算法是不是 HMAC (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		//验证token是否有效，转为claims结构
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
