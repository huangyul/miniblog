package token

import (
	"errors"
	"sync"
)

// Config 包括 token 包的配置选项
type Config struct {
	key         string
	identityKey string
}

// ErrMissingHeader 表示请求头为空
var ErrMissingHeader = errors.New("the length of the `Authorization header is zeor`")

var (
	once   sync.Once
	config = Config{"Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5", "identityKey"}
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

func Parse(tokenString string, key string) (string, error) {
	//// 解析 token
	//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//	if _, ok := token.Method.(*jwt.SigningMethod)
	//})
	return "", nil
}
