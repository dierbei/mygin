package classes

import (
	"github.com/gin-gonic/gin"
	"mygin/goft"
	"mygin/models"
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

func (u *UserClass) UserDetail(ctx *gin.Context) goft.Model {
	return &models.UserModel{
		UserId:   5207101145,
		UserName: "Dierbei",
	}
}

func (u *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/userlist", u.UserList)
	goft.Handle("GET", "/userdetail", u.UserDetail)
}
