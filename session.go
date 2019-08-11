package go_redis_session

import (
	"context"
	"time"

	"math/rand"

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

func (session *Session) Set(context *gin.Context, key string, value string, expire time.Duration) error {
	token := session.randomToken(128)

	err := session.redisconn.Set(session.ctx, &redis.Item{
		Key:   token,
		Value: []byte(value),
		TTL:   int32(expire),
	})

	if err != nil {
		return err
	}

	context.SetCookie(key, value, int(expire), "/", "", false, true)

	return nil
}

func (session *Session) Get(context *gin.Context, key string) (string, error) {

}

func (session *Session) randomToken(bits int) string {
	seeder := rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	result := make([]byte, bits)
	for i := range result {
		result[i] = chars[seeder.Intn(len(chars))]
	}
	return string(result)
}
