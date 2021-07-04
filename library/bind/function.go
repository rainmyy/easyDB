package bind

import "github.com/rainmyy/easyDB/library/strategy"

func recurGetTreeNode(tree *strategy.TreeStruct) {
	if tree == nil || tree.IsLeaf() {
		return
	}

}
