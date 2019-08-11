package go_redis_session

import "github.com/gin-gonic/gin"

type Session struct{}

func NewSession() *Session {
	return &Session{}
}

func (session *Session) Set(context *gin.Context, key string, value string) error {

}

func (session *Session) Get(context *gin.Context, key string) (string, error) {

}
