package file

import (
	"fmt"
	"sort"

	"github.com/rainmyy/easyDB/library/common"
	"github.com/rainmyy/easyDB/library/strategy"
)

func InIntSliceSortedFunc(stack []int) func(int) bool {
	sort.Ints(stack)
	return func(needle int) bool {
		index := sort.SearchInts(stack, needle)
		return index < len(stack) && stack[index] == needle
	}
}

/**
*实现树状结构
 */
func initTreeFunc(bytesList [][]byte) []*strategy.TreeStruct {
	currentTree := strategy.TreeInstance()
	//分隔符，91:'[' 46:'.' 58:'.'
	var segment = []int{int(common.LeftBracket), int(common.Period)}
	infunc := InIntSliceSortedFunc(segment)
	var rootTree = currentTree
	//根节点设置为1
	currentTree.SetHight(1)
	for i := 0; i < len(bytesList); i++ {
		bytes := bytesList[i]
		bytesLen := len(bytes)
		if bytesLen == 0 {
			continue
		}
		tempNum := 0
		for j := 0; j < bytesLen; j++ {
			if infunc(int(bytes[j])) {
				tempNum++
			}
		}
		treeStruct := strategy.TreeInstance()
		currentHigh := currentTree.GetHight()
		var nodeStruct *strategy.NodeStruct
		if tempNum > 0 && len(bytes) > tempNum {
			bytes = bytes[tempNum : bytesLen-1]

			nodeStruct = strategy.NodeInstance(bytes, []byte{})
			for tempNum < currentHigh {
				currentTree = currentTree.GetParent()
				currentHigh = currentTree.GetHight()
			}
			treeStruct.SetNode(nodeStruct)
			treeStruct.SetParent(currentTree)
			currentTree.SetChildren(treeStruct)
			currentTree = treeStruct
		} else if tempNum == 0 {
			//key:vaule类型的值
			separatorPlace := common.SlicePlace(byte(common.Colon), bytes)
			if separatorPlace <= 0 {
				continue
			}
			key := bytes[0:separatorPlace]
			value := bytes[separatorPlace+1 : bytesLen]
			nodeStruct = strategy.NodeInstance(key, value)
			if currentTree == nil {
				continue
			}
			treeStruct.SetNode(nodeStruct)
			treeStruct.SetParent(currentTree)
			currentTree.SetChildren(treeStruct)
		}
	}
	return rootTree.GetChildren()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parserDataFunc(file *File, objType int, data []byte) ([]*strategy.TreeStruct, error) {
	switch objType {
	case common.IniType:
		return ParserIniContent(data)
	case common.YamlType:
		return ParserYamlContent(data)
	case common.JsonType:
		return ParserjSONContent(data)
	case common.DataType:
		return ParserContent(data)
	default:
		return ParserContent(data)
	}
}

func ParserContent(data []byte) ([]*strategy.TreeStruct, error) {
	return nil, nil
}

func ParserjSONContent(data []byte) ([]*strategy.TreeStruct, error) {
	return nil, nil
}
func ParserYamlContent(data []byte) ([]*strategy.TreeStruct, error) {
	return nil, nil
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
func ParserIniContent(data []byte) ([]*strategy.TreeStruct, error) {
	if data == nil {
		return nil, fmt.Errorf("content is nil")
	}
	bytesList := [][]byte{}
	hasSlash := false
	bytes := []byte{}
	if data[len(data)-1] != 10 {
		data = append(data, byte(common.LineBreak))
	}
	for i := 0; i < len(data); i++ {
		value := data[i]
		//出现斜杠过滤
		if value == byte(common.Slash) || value == byte(common.Hash) || value == byte(common.Asterisk) {
			hasSlash = true
			continue
		}
		if hasSlash {
			if value == byte(common.LineBreak) {
				hasSlash = false
			}
			continue
		}
		// 通过\n截取长度
		if value != byte(common.LineBreak) && value != byte(common.Blank) {
			bytes = append(bytes, value)
		} else if len(bytes) > 0 {
			bytesList = append(bytesList, bytes)
			bytes = []byte{}
		}
	}
	if len(bytesList) == 0 {
		return nil, fmt.Errorf("bytes is empty")
	}
	//数据以树型结构存储
	byteTreeList := initTreeFunc(bytesList)
	return byteTreeList, nil
}
