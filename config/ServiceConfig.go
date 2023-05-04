package config

import (
	"github.com/dunpju/higo-ioc/injector"
	"github.com/dunpju/higo-ioc/test/services"
)

type ServiceConfig struct {
	injector.BeanConfig
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (this *ServiceConfig) OrderService() *services.OrderService {
	return services.NewOrderService()
}

func (this *ServiceConfig) DbService() *services.DbService {
	return services.NewDbService()
}
