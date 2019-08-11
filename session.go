package go_redis_session

import (
	"context"

	"crypto/rand"

	"github.com/gin-gonic/gin"
	"github.com/xiangrui2019/redis"
)

type Session struct {
	ctx       context.Context
	redisconn redis.Client
}

func NewSession(redisoptions redis.Options) *Session {
	return &Session{
		ctx:       context.Background(),
		redisconn: redis.New(redisoptions),
	}
}

func (session *Session) Set(context *gin.Context, key string, value string) error {

}

func (session *Session) Get(context *gin.Context, key string) (string, error) {

}

func (session *Session) randomToken(bits int) string {
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	result := make([]byte, bits)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
