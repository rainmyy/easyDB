package bind

import (
	"fmt"
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
func bindObj(tree *strategy.TreeStruct, obj interface{}) func(args ...interface{}) interface{} {
	value := reflect.ValueOf(obj)
	var list interface{}
	var objFunc func(args []*strategy.NodeStruct, list interface{}) interface{}
	var recurNode func(tree *strategy.TreeStruct, list interface{})
	switch value.Kind() {
	case reflect.String:
		list = "{%s}"
		objFunc = func(args []*strategy.NodeStruct, list interface{}) interface{} {
			for _, val := range args {
				nodeName := common.Bytes2str(val.GetName())
				nodeData := common.Bytes2str(val.GetData())
				if len(nodeName) == 0 {
					continue
				}
				temp := "{%s}"
				if len(nodeData) > 0 {
					if len(args) > 1 {
						temp = fmt.Sprintf("[{%s:%s}]", nodeName, nodeData)
						list = fmt.Sprintf("["+list.(string)+",%s]", temp)
					} else {
						list = fmt.Sprintf("["+list.(string)+"]", temp)
					}
				} else {
					temp = fmt.Sprintf(list.(string), nodeName, "")
				}
			}
			fmt.Print(list)
			//fmt.Print(list)
			return list
		}
	case reflect.Map:
		list = make(map[string]interface{})
	case reflect.Struct:
		list = obj
	}
	recurNode = func(tree *strategy.TreeStruct, list interface{}) {
		if tree == nil {
			return
		}
		if tree.GetNode() != nil {
			dataList := tree.GetNode()
			list = objFunc(dataList, list)
		}
		for _, val := range tree.GetChildren() {
			recurNode(val, list)
		}
	}
	recurNode(tree, list)
	return nil
}
