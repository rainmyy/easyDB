package file

import (
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
func ParserIniContent(content []byte) []*strategy.TreeStruct {
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
	byteTreeList := initTreeFunc(bytesList)
	return byteTreeList
}

/**
*实现树状结构
 */
func initTreeFunc(bytesList [][]byte) []*strategy.TreeStruct {
	currentTree := strategy.TreeInstance()
	//分隔符，91:'[' 46:'.' 58:':'
	var segment = []int{91, 46}
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
			separatorPlace := common.SlicePlace(58, bytes)
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
