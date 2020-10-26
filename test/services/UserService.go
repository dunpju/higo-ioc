package services

import "fmt"

type UserService struct {
	Order *OrderService `inject:"ServiceConfig.OrderService()"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService)GetUserInfo(uid int)  {
	fmt.Println("uid",uid)
}

func (this *UserService)GetOrderInfo(uid int)  {
	this.Order.GetOrderInfo(uid)
}