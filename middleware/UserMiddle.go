package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

type UserMiddle struct {
}

func NewUserMiddle() *UserMiddle {
	return &UserMiddle{}
}

func (mid *UserMiddle) OnRequest(ctx *gin.Context) error {
	log.Println("这是用户中间件, name:", ctx.Query("name"))
	return nil
}
