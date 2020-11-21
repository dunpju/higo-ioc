package injector

import (
	"github.com/dengpju/higo-express/express"
	"log"
	"reflect"
)

var BeanFactory *BeanFactoryImpl

func init()  {
	BeanFactory=NewBeanFactory()
}

type BeanFactoryImpl struct {
	beanMapper BeanMapper
	exprMap map[string]interface{}
}

func NewBeanFactory() *BeanFactoryImpl {
	return &BeanFactoryImpl{beanMapper:make(BeanMapper),exprMap: make(map[string]interface{})}
}

func (this *BeanFactoryImpl)SetExprMap(key string, val interface{})  {
	this.exprMap[key] = val
	express.SetFuncMap(key, val)
}

func (this *BeanFactoryImpl)GetExprMap() map[string]interface{} {
	return this.exprMap
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
				log.Println("表达式")
				ret := express.Run(field.Tag.Get("inject"))
				if ret != nil && !ret.IsEmpty() {
					v.Field(i).Set(reflect.ValueOf(ret[0]))
				}
			}
		}
	}
}