package bind

import (
	. "github.com/rainmyy/easyDB/library/strategy"
)

type Binder interface {
	Bind(treeList []*TreeStruct)
	UnBind(obj interface{})
	GetValue() interface{}
}

/**
* 绑定实体和参数
 */
func DefaultBind(tree []*TreeStruct, obj Binder) interface{} {
	obj.Bind(tree)
	bindData := obj.GetValue()
	return bindData
}

func DefaultUnBind(obj interface{}) {

}
