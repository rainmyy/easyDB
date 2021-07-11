package bind

import (
	"reflect"

	"github.com/rainmyy/easyDB/library/strategy"
)

func DefaultBindStruct(treeList []*strategy.TreeStruct, obj interface{}) []interface{} {
	st := reflect.TypeOf(obj)
}
