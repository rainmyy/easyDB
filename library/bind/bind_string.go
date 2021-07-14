package bind

import (
	"bytes"

	"github.com/rainmyy/easyDB/library/common"
	"github.com/rainmyy/easyDB/library/strategy"
)

type String struct {
	value *bytes.Buffer
}

func (s *String) Bind(treeList []*strategy.TreeStruct) {
	var buffer = bytes.NewBuffer([]byte{})
	buffer = new(bytes.Buffer)
	buffer.WriteRune(common.LeftRrance)
	BindString(treeList, buffer)
	buffer.WriteRune(common.RightRrance)
	s.value = buffer
}

func (s *String) GetValue() interface{} {
	return s.value
}
func StrigInstance() *String {
	return new(String)
}

/**
* 序列化tree:{key:name, key2:{name:1, name2:2}, key3:[{name:1, name2:3}, {name:3, name:4}]}
 */
func BindString(treeList []*strategy.TreeStruct, buffer *bytes.Buffer) int {
	if len(treeList) == 0 {
		return 0
	}
	childrenNum := 0
	leafNum := 0
	for _, val := range treeList {
		if val.IsLeaf() {
			leafNum++
		}
	}
	currentNum := leafNum
	for key, val := range treeList {
		nodeList := val.GetNode()
		node := nodeList[0]
		if val.IsLeaf() {
			if currentNum == leafNum {
				buffer.WriteRune(common.LeftRrance)
			}
			buffer.WriteString(string(node.GetName()))
			buffer.WriteRune(common.Colon)
			buffer.WriteString(string(node.GetData()))
			leafNum--
			if leafNum == 0 {
				buffer.WriteRune(common.RightRrance)
			}
		} else {
			childrenNum = len(val.GetChildren())
			buffer.WriteString(string(node.GetName()))
			buffer.WriteRune(common.Colon)
			if childrenNum > 1 {
				buffer.WriteRune(common.LeftBracket)
			}
			BindString(val.GetChildren(), buffer)
			if childrenNum > 1 {
				buffer.WriteRune(common.RightBracket)
			}
		}
		if key != len(treeList)-1 {
			buffer.WriteRune(common.Comma)
		}
	}
	return childrenNum
}

func UnBindString(str string, tree *strategy.TreeStruct) {

}
