package classes

import (
	"github.com/gin-gonic/gin"
	"mygin/goft"
)

type UserClass struct {
}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (u *UserClass) UserList(ctx *gin.Context) string {
	//return func(ctx *gin.Context) {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"result": "ok",
	//	})
	//}
	return "Abc"
}

func (u *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/userlist", u.UserList)
}
