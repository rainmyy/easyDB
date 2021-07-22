package bind

import (
	"bytes"
	"reflect"

	"github.com/rainmyy/easyDB/library/strategy"
)

type Binder interface {
	Bind(treeList []*strategy.TreeStruct, obj interface{})
	GetValue() interface{}
}

/***
*一套绑定参数的方法，默认将数据转化成字符串
 */
/**
* 绑定实体和参数
 */
func DefaultBind(tree []*strategy.TreeStruct, obj Binder) (interface{}, error) {

	value := reflect.ValueOf(obj)
	var buffer interface{}
	obj.Bind(tree, obj)
	switch value.Kind() {
	case reflect.String:
		buffer = bytes.NewBuffer([]byte{})
		buffer = obj.GetValue()
		//buffer = DefaultBindString(tree)
	case reflect.Map:
		buffer = make([]map[string]interface{}, 0)
		//buffer = DefaultBindMap(tree)
	case reflect.Struct:
		buffer = obj
		//buffer = DefaultBindStruct(tree, obj)
	}

	return buffer, nil
}
