package models

import "fmt"

type UserModel struct {
	UserId   int	`uri:"id" binding:"required"`
	UserName string
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (m *UserModel) String() string {
	return fmt.Sprintf("user_id:%d,user_name=%s", m.UserId, m.UserName)
}
