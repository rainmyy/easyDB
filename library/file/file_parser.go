package file

import (
	"fmt"

	"github.com/rainmyy/easyDB/library/common"
	"github.com/rainmyy/easyDB/library/strategy"
)

/**
* 解析数据，将数据解析成树形结构进行存储
 */
func (this *File) Parser(objType int) ([]*strategy.TreeStruct, error) {
	var tree []*strategy.TreeStruct
	switch objType {
	case common.IniType:
		tree = this.ParserIniContent()
	case common.YamlType:
		tree = this.ParserYamlContent()
	case common.JsonType:
		tree = this.ParserjSONContent()
	case common.DataType:
		tree = this.ParserContent()
	default:
		tree = this.ParserContent()
	}

	if tree == nil {
		return nil, fmt.Errorf("data is none")
	}
	return tree, nil
}

func (this *File) ParserContent() []*strategy.TreeStruct {
	return nil
}

func (this *File) ParserjSONContent() []*strategy.TreeStruct {
	return nil
}
func (this *File) ParserYamlContent() []*strategy.TreeStruct {
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
func (this *File) ParserIniContent() []*strategy.TreeStruct {
	if this.content == nil {
		return nil
	}
	bytesList := [][]byte{}
	hasSlash := false
	bytes := []byte{}
	if this.content[len(this.content)-1] != 10 {
		this.content = append(this.content, 10)
	}
	for i := 0; i < len(this.content); i++ {
		value := this.content[i]
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
	byteTreeList := initTreeFunc(bytesList)
	return byteTreeList
}
