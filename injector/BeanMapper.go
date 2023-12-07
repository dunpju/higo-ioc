package injector

import (
	"reflect"
	"sync"
)

type BeanMapper struct {
	m sync.Map
}

func NewBeanMapper() *BeanMapper {
	return &BeanMapper{}
}

func (this *BeanMapper) add(bean interface{}) {
	t := reflect.TypeOf(bean)
	if t.Kind() != reflect.Ptr {
		panic("require ptr object")
	}
	this.m.Store(t, reflect.ValueOf(bean))
}

func (this *BeanMapper) get(bean interface{}) reflect.Value {
	var t reflect.Type
	if bt, ok := bean.(reflect.Type); ok {
		t = bt
	} else {
		t = reflect.TypeOf(bean)
	}
	if v, ok := this.m.Load(t); ok {
		return v.(reflect.Value)
	}
	// 处理接口 继承
	v := reflect.Value{}
	this.m.Range(func(key, value any) bool {
		if key.(reflect.Type).Implements(t) {
			v = value.(reflect.Value)
			return false
		}
		return true
	})
	return v
}
