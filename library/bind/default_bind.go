package bind

import (
	"github.com/rainmyy/easyDB/library/strategy"
)

type Binder interface {
	Bind(treeList []*strategy.TreeStruct)
	GetValue() interface{}
}

/***
*一套绑定参数的方法，默认将数据转化成字符串
 */
/**
* 绑定实体和参数
 */
func DefaultBind(tree []*strategy.TreeStruct, obj Binder) interface{} {
	obj.Bind(tree)
	bindData := obj.GetValue()
	return bindData
}
