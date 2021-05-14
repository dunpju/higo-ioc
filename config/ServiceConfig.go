package config

import "github.com/dengpju/higo-ioc/test/services"

type ServiceConfig struct {
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (this *ServiceConfig) Provider() {
}

func (this *ServiceConfig) OrderService() *services.OrderService {
	return services.NewOrderService()
}

func (this *ServiceConfig) DbService() *services.DbService {
	return services.NewDbService()
}
