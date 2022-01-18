package middleware

import "log"

type UserMiddle struct {
}

func NewUserMiddle() *UserMiddle {
	return &UserMiddle{}
}

func (mid *UserMiddle) OnRequest() error {
	log.Println("这是用户中间件")
	return nil
}
