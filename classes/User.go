package classes

import (
	"github.com/gin-gonic/gin"
	"mygin/goft"
	"net/http"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (u *UserClass) UserList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"result": "ok",
		})
	}
}

func (u *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/userlist", u.UserList())
}
