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
*实现树状结构
 */
func initTreeFunc(bytesList [][]byte) []*strategy.TreeStruct {
	currentTree := strategy.TreeInstance()
	//分隔符，91:'[' 46:'.' 58:':'
	var segment = []int{int(common.LeftBracket), int(common.Colon)}
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
