package injector

import "fmt"

type IBean interface {
	Provider() IBean
}

type BeanConfig struct {
}

func NewBeanConfig() *BeanConfig {
	return &BeanConfig{}
}

func (this *BeanConfig) Provider() IBean {
	fmt.Printf("%T\n", this)
	return this
}