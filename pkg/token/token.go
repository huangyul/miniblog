package token

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

// Config 包括 token 包的配置
type Config struct {
	key         string
	identityKey string
}

var ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")

var (
	config = Config{
		key:         "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
		identityKey: "identityKey",
	}
	once sync.Once
)

// Init 设置包级别的配置
func Init(key string, identityKey string) {
	once.Do(func() {
		if key != "" {
			config.key = key
		}
		if identityKey != "" {
			config.identityKey = identityKey
		}
	})
}

// Parse 使用密钥 key 来解析 token
func Parse(tokenString string, key string) (string, error) {
	// 解析 token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// 判断 token 使用的加密算法是正确
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(key), nil
	})
	// 解析失败
	if err != nil {
		return "", err
	}

	// 获取 token 的主体
	var identityKey string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		identityKey = claims[config.identityKey].(string)
	}
	return identityKey, nil
}

// ParseRequest 从请求头中获取 token，并传递给 Parse 函数解析
func ParseRequest(ctx *gin.Context) (string, error) {
	header := ctx.Request.Header.Get("Authorization")

	if len(header) == 0 {
		return "", ErrMissingHeader
	}

	var t string
	fmt.Sscanf(header, "Bearer %s", &t)

	return Parse(t, config.key)
}

// Sign 生成 token 签名
func Sign(identityKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.identityKey: identityKey,
		"nbf":              time.Now().Unix(),
		"iat":              time.Now().Unix(),
		"exp":              time.Now().Add(10000 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.key))

	return tokenString, err
}
