package file

import (
	"reflect"

	"github.com/rainmyy/easyDB/library/strategy"
)

/**
* 解析
 */
func (this *File) Parser(obj interface{}) interface{} {
	tree := ParserIniContent(this.content)
	if tree == nil {
		return nil
	}
	result := bindObj(tree, obj)
	return result
}

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
		list = ""
		objFunc = func(args []*strategy.NodeStruct, list interface{}) interface{} {
			for _, val := range args {
				if len(val.GetName()) == 0 {
					continue
				}
			}
			return ""
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

/**
*解析ini格式配置文件
*desc:
*[test]
*    [..params]
*        name:name1
*        key:value
*    [...params]
*        name:name2
*        key:value
 */
func ParserIniContent(content []byte) *strategy.TreeStruct {
	if content == nil {
		return nil
	}
	bytesList := [][]byte{}
	hasSlash := false
	bytes := []byte{}
	if content[len(content)-1] != 10 {
		content = append(content, 10)
	}
	for i := 0; i < len(content); i++ {
		value := content[i]
		//出现斜杠过滤
		if value == 47 {
			hasSlash = true
			continue
		}
		if hasSlash {
			if value == 10 {
				hasSlash = false
			}
			continue
		}
		// 通过\n截取长度
		if value != 10 && value != 32 {
			bytes = append(bytes, value)
		} else if len(bytes) > 0 {
			bytesList = append(bytesList, bytes)
			bytes = []byte{}
		}
	}
	if len(bytesList) == 0 {
		return nil
	}
	//数据以树型结构存储
	//byteTreeList := []*strategy.TreeStruct{}
	byteTreeList := initTreeFunc(bytesList)
	return byteTreeList
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
