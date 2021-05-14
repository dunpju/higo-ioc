package services

import (
	"fmt"
	"github.com/dengpju/higo-annotation/anno"
)

type UserService struct {
	Order IOrder      `inject:"ServiceConfig.OrderService()"`
	Age   *anno.Value `prefix:"user.age" json:"age"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) GetUserInfo(uid int) {
	fmt.Println("uid", uid)
}
