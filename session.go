package go_redis_session

import (
	"context"
	"errors"
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
	token, err := context.Cookie(key)

	if err != nil {
		return "", err
	}

	item, err := session.redisconn.Get(session.ctx, token)

	if err != nil {
		return "", err
	}

	if string(item.Value) == "" {
		return "", errors.New("not found.")
	}

	return string(item.Value), nil
}

func (session *Session) Delete(context *gin.Context, key string) error {
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
