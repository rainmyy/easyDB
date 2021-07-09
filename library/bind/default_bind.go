package bind

import (
	"bytes"
	"reflect"

	"github.com/rainmyy/easyDB/library/common"
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
	var buffer *bytes.Buffer
	switch value.Kind() {
	case reflect.String:
		buffer = new(bytes.Buffer)
		buffer.WriteRune(common.LeftRrance)
		BindString(tree, buffer)
		buffer.WriteRune(common.RightRrance)
	case reflect.Map:
	case reflect.Struct:

	}
	return buffer, nil
}
