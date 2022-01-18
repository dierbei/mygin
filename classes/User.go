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

func (u *UserClass) UserList(ctx *gin.Context) goft.Models {
	users := []*models.UserModel{
		{
			UserId:   5207101145,
			UserName: "Dierbei",
		},
		{
			UserId:   5200000,
			UserName: "Xiaolatiao",
		},
	}

	return goft.MakeModels(users)
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
