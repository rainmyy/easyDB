package bind

import (
	. "github.com/rainmyy/easyDB/library/strategy"
)

type Array struct {
	length int
	value  []map[string]interface{}
}

func ArrayInterface() *Array {
	return &Array{value: make([]map[string]interface{}, 0)}
}

/**
* 获取数据树的map和slice
 */
func (a *Array) Bind(treeList []*TreeStruct) {
	var treeMapList = make([]map[string]interface{}, 0)
	var getBindMap func(tree []*TreeStruct) []map[string]interface{}
	/***
	* 递归方式获取
	 */
	getBindMap = func(tree []*TreeStruct) []map[string]interface{} {
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
	a.value = treeMapList
}

func (a *Array) GetValue() interface{} {
	return a.value
}
