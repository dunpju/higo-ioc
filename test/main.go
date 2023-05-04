package main

import (
	"fmt"
	"github.com/dunpju/higo-annotation/anno"
	"github.com/dunpju/higo-ioc/config"
	"github.com/dunpju/higo-ioc/injector"
	"github.com/dunpju/higo-ioc/test/services"
)

func main() {
	//uid:=123
	//userService:=services.NewUserService(services.NewOrderService())
	//userService.GetUserInfo(uid)
	//userService.GetOrderInfo(uid)

	//injector.BeanFactory.Set(services.NewOrderService())
	//order:=injector.BeanFactory.Get((*services.UserService)(nil))
	//fmt.Println(order)

	//injector.BeanFactory.Set(services.NewOrderService())
	//
	//userService:=services.NewUserService()
	//injector.BeanFactory.Apply(userService)
	//fmt.Println(userService.Order)

	serviceConfig := config.NewServiceConfig()
	//injector.BeanFactory.SetExprMap("ServiceConfig", serviceConfig)
	injector.BeanFactory.Config(serviceConfig)

	//order := services.NewOrderService()
	//order.Version = "2.0"
	//injector.BeanFactory.Set(order)

	anno.Config.Set("user.age", "520")
	userService := services.NewUserService()
	injector.BeanFactory.Apply(userService)
	fmt.Println(userService.Age)
	fmt.Println(userService.Order)
	adminService := services.NewAdminService()
	injector.BeanFactory.Apply(adminService)
	fmt.Println(adminService.Order)
	fmt.Println(adminService.Order.Db)
}
