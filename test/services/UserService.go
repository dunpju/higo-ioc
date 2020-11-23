package services

import "fmt"

type UserService struct {
	Order IOrder `inject:"ServiceConfig.OrderService()"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService)GetUserInfo(uid int)  {
	fmt.Println("uid",uid)
}
