package services

import "fmt"

type OrderService struct {
	Version string
}

func NewOrderService() *OrderService {
	fmt.Println("初始化 OrderService")
	return &OrderService{Version:"3.0"}
}

func (this *OrderService)GetOrderInfo(uid int)  {
	fmt.Println("uid",uid,"订单信息")
}
