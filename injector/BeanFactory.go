package injector

import (
	"fmt"
	"github.com/dengpju/higo-express/express"
	"reflect"
	"regexp"
	"strings"
)

var BeanFactory *BeanFactoryImpl

func init() {
	BeanFactory = NewBeanFactory()
}

type BeanFactoryImpl struct {
	beanMapper BeanMapper
	exprMap    map[string]interface{}
}

func NewBeanFactory() *BeanFactoryImpl {
	return &BeanFactoryImpl{beanMapper: make(BeanMapper), exprMap: make(map[string]interface{})}
}

func (this *BeanFactoryImpl) SetExprMap(key string, val interface{}) {
	this.exprMap[key] = val
	express.SetFuncMap(key, val)
}

func (this *BeanFactoryImpl) GetExprMap() map[string]interface{} {
	return this.exprMap
}

func (this *BeanFactoryImpl) Set(values ...interface{}) {
	if values == nil || len(values) == 0 {
		return
	}
	for _, v := range values {
		this.beanMapper.add(v)
	}
}

func (this *BeanFactoryImpl) Get(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	value := this.beanMapper.get(v)
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
			if field.Tag.Get("inject") == "-" { // 单例
				if value := this.Get(field.Type); value != nil { // 容器中如果存在
					v.Field(i).Set(reflect.ValueOf(value))
					this.Apply(value)
				}
			} else { // 多例
				ret := express.Run(field.Tag.Get("inject"))
				if ret != nil && !ret.IsEmpty() {
					retValue := ret[0]
					if retValue != nil {
						v.Field(i).Set(reflect.ValueOf(retValue))
						this.Apply(retValue)
					}
				}
			}
		}
	}
}

func (this *BeanFactoryImpl) Config(cfgs ...interface{}) {
	for _, cfg := range cfgs {
		t := reflect.TypeOf(cfg)
		if t.Kind() != reflect.Ptr {
			panic("required ptr object")
		}
		this.Set(cfg)
		this.SetExprMap(t.Elem().Name(), cfg) // 自动构建
		v := reflect.ValueOf(cfg)
		for i := 0; i < t.NumMethod(); i++ {
			method := v.Method(i)
			typeRegexp := regexp.MustCompile(`func\((.*)\)`)
			regParams := typeRegexp.FindStringSubmatch(fmt.Sprintf("%s", method.Type()))
			if "" != regParams[1] {
				params := make([]reflect.Value, 0)
				args := strings.Split(regParams[1], ",")
				for _, a := range args {
					trimArgType := strings.Trim(a, " ")
					if "string" == trimArgType {
						params = append(params, reflect.ValueOf(""))
					} else if "int" == trimArgType {
						params = append(params, reflect.ValueOf(0))
					} else if "int64" == trimArgType {
						params = append(params, reflect.ValueOf(int64(0)))
					}
				}
				callRet := method.Call(params)
				if callRet != nil && len(callRet) == 1 {
					this.Set(callRet[0].Interface())
				}
			} else {
				callRet := method.Call(nil)
				if callRet != nil && len(callRet) == 1 {
					this.Set(callRet[0].Interface())
				}
			}
		}
	}
}
