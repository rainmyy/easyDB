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
func bindObj(tree []*strategy.TreeStruct, obj interface{}) (*bytes.Buffer, error) {
	value := reflect.ValueOf(obj)
	buffer := bytes.NewBuffer([]byte{})
	switch value.Kind() {
	case reflect.String:
		buffer = DefaultBindString(tree)
	case reflect.Map:
		_ = DefaultBindMap(tree)
	case reflect.Struct:

	}
	return buffer, nil
}
