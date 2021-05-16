package injector

type IBean interface {
	Provider()
}

type BeanConfig struct {
}

func NewBeanConfig() *BeanConfig {
	return &BeanConfig{}
}

func (this *BeanConfig) Provider() {}