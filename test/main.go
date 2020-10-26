package main

import (
	"fmt"
	"github.com/dengpju/higo-ioc/config"
	"github.com/dengpju/higo-ioc/injector"
	"github.com/dengpju/higo-ioc/test/services"
)

func main()  {
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

	serviceConfig:=config.NewServiceConfig()
	injector.BeanFactory.ExprMap = map[string]interface{}{
		"ServiceConfig":serviceConfig,
	}

	injector.BeanFactory.Set(serviceConfig)
	userService:=services.NewUserService()
	injector.BeanFactory.Apply(userService)
	fmt.Println(userService.Order)
}
