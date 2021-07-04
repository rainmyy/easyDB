package bind

import "github.com/rainmyy/easyDB/library/strategy"

/**
* 序列化tree:{key:name, key2:{name:1, name2:2}, key3:[{name:1, name2:3}, {name:3, name:4}]}
 */
func BindString(tree *strategy.TreeStruct, obj *string) {
	currentTree := tree

	for currentTree != nil {

	}
}

func UnBindString(str string, tree *strategy.TreeStruct) {

}
