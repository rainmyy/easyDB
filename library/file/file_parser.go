package file

import (
	"github.com/rainmyy/easydb/library/strategy"
)

/**
* 解析
 */
func (this *File) Parser(obj *interface{}) {
	tree := ParserIniContent(this.content)
	bind(tree, obj)
}

/**
* 绑定实体和参数
 */
func bind(tree *strategy.TreeStruct, obj *interface{}) {

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
