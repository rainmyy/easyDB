package bind

import (
	"bytes"

	"github.com/rainmyy/easyDB/library/common"
	"github.com/rainmyy/easyDB/library/strategy"
)

/**
* 序列化tree:{key:name, key2:{name:1, name2:2}, key3:[{name:1, name2:3}, {name:3, name:4}]}
 */
func BindString(treeList []*strategy.TreeStruct, ret string) (buffer bytes.Buffer) {
	if len(treeList) == 0 {
		return
	}
	buffer.WriteString("")
	for _, val := range treeList {
		nodeList := val.GetNode()
		if len(nodeList) == 0 {
			continue
		}
		node := nodeList[0]
		if val.IsLeaf() {
			buffer.WriteString(string(node.GetName()))
			buffer.WriteRune(common.Colon)
			buffer.WriteString(string(node.GetData()))
		} else {
			buffer.WriteRune(common.LeftRrance)
		}
	}
	return
	//return json.Marshal(currentTree)
}

func UnBindString(str string, tree *strategy.TreeStruct) {

}
