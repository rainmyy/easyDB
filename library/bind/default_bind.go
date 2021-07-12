package bind

import (
	"bytes"
	"reflect"

	"github.com/rainmyy/easyDB/library/strategy"
)

/***
*一套绑定参数的方法，默认将
 */
/**
* 绑定实体和参数
 */
func DefaultBind(tree []*strategy.TreeStruct, obj interface{}) (interface{}, error) {
	value := reflect.ValueOf(obj)
	var buffer interface{}
	switch value.Kind() {
	case reflect.String:
		buffer = bytes.NewBuffer([]byte{})
		buffer = DefaultBindString(tree)
	case reflect.Map:
		buffer = make([]map[string]interface{}, 0)
		buffer = DefaultBindMap(tree)
	case reflect.Struct:
		buffer = obj
		buffer = DefaultBindStruct(tree, obj)
	}
	return buffer, nil
}
