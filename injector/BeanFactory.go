package injector

import (
	"log"
	"reflect"
)

var BeanFactory *BeanFactoryImpl

func init()  {
	BeanFactory=NewBeanFactory()
}

type BeanFactoryImpl struct {
	beanMapper BeanMapper
	ExprMap map[string]interface{}
}

func NewBeanFactory() *BeanFactoryImpl {
	return &BeanFactoryImpl{beanMapper:make(BeanMapper),ExprMap: make(map[string]interface{})}
}

func (this *BeanFactoryImpl)Set(values ...interface{})  {
	if values==nil || len(values)==0 {
		return
	}
	for _,v:=range values {
		this.beanMapper.add(v)
	}
}

func (this *BeanFactoryImpl)Get(v interface{}) interface{} {
	if v==nil {
		return nil
	}
	value:=this.beanMapper.get(v)
	if value.IsValid() {
		return value.Interface()
	}
	return nil
}

func (this *BeanFactoryImpl) Apply(bean interface{}) {
	if bean == nil {
		return
	}
	v := reflect.ValueOf(bean)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if v.Field(i).CanSet() && field.Tag.Get("inject") != "" {
			if field.Tag.Get("inject") == "-" {
				if value := this.Get(field.Type); value != nil {
					v.Field(i).Set(reflect.ValueOf(value))
				}
			}else{
				log.Panicln("表达式")
				//ret:=expr.BeanExpr(field.Tag.Get("inject"),this.ExprMap)
				//if ret!=nil && !ret.IsEmpty() {
				//	v.Field(i).Set(reflect.ValueOf(ret[0]))
				//}
			}
		}
	}
}