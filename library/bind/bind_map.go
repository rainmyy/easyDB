package bind

import (
	"github.com/rainmyy/easyDB/library/strategy"
)

func DefaultBindMap(treeList []*strategy.TreeStruct) map[string]interface{} {
	var treeMap = make(map[string]interface{})
	var getBindMap func(tree []*strategy.TreeStruct, treetreeMap map[string]interface{})
	getBindMap = func(tree []*strategy.TreeStruct, treetreeMap map[string]interface{}) {
		for _, val := range treeList {
			nodeList := val.GetNode()
			node := nodeList[0]
			tempMap := make(map[string]interface{})
			if val.IsLeaf() {
				tempMap[string(node.GetName())] = string(node.GetData())
			} else {
				childrenNum := len(val.GetChildren())
				nodeName := string(node.GetName())
				if childrenNum > 1 {
					tempMap[nodeName] = make([]map[string]interface{}, childrenNum)
					getBindMap(val.GetChildren(), tempMap)
				}
			}
		}
	}
	getBindMap(treeList, treeMap)
	return treeMap
}
