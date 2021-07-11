package bind

import (
	"github.com/rainmyy/easyDB/library/strategy"
)

/**
* 获取数据树的map和slice
 */
func DefaultBindMap(treeList []*strategy.TreeStruct) []map[string]interface{} {
	var treeMapList = make([]map[string]interface{}, 0)
	var getBindMap func(tree []*strategy.TreeStruct) []map[string]interface{}
	getBindMap = func(tree []*strategy.TreeStruct) []map[string]interface{} {
		if len(tree) == 0 {
			return nil
		}
		var treeMapList = make([]map[string]interface{}, 0)
		for _, val := range tree {
			nodeList := val.GetNode()
			node := nodeList[0]
			var treeMap = make(map[string]interface{})
			if val.IsLeaf() {
				treeMap[string(node.GetName())] = string(node.GetData())
			} else {
				childrenNum := len(val.GetChildren())
				nodeName := string(node.GetName())
				if childrenNum > 1 {
					treeSlice := getBindMap(val.GetChildren())
					if len(treeSlice) == 0 {
						continue
					}
					treeMap[nodeName] = treeSlice
				} else {
					res := getBindMap(val.GetChildren())
					treeMap[nodeName] = make(map[string]interface{})
					if len(res) == 0 {
						treeMap[nodeName] = nil
					} else {
						treeMap[nodeName] = res[0]
					}
				}
				if len(treeMap) == 0 {
					continue
				}
			}
			treeMapList = append(treeMapList, treeMap)
		}
		return treeMapList
	}
	treeMapList = getBindMap(treeList)
	return treeMapList
}
