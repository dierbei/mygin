package models

import "fmt"

type UserModel struct {
	UserId   int
	UserName string
}

func (m *UserModel) String() string {
	return fmt.Sprintf("user_id:%d,user_name=%s", m.UserId, m.UserName)
}
