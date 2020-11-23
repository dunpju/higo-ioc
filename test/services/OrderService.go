package services

import "fmt"

type IOrder interface {
	Name() string
}

type OrderService struct {
	Version string
	Db *DbService `inject:"ServiceConfig.DbService()"`
}

func NewOrderService() *OrderService {
	fmt.Println("初始化 OrderService")
	return &OrderService{Version:"3.0"}
}

func (this *OrderService) GetOrderInfo(uid int) {
	fmt.Println("uid", uid, "订单信息")
}

func (this *OrderService) Name() string {
	return "OrderService"
}