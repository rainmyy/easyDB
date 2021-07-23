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
	if len(treeList) > 1 {
		buffer.WriteRune(common.LeftBracket)
	}
	BindString(treeList, buffer)
	if len(treeList) > 1 {
		buffer.WriteRune(common.RightBracket)
	}
	s.value = buffer
}

func (s *String) GetValue() interface{} {
	return s.value
}
func StrigInstance() *String {
	return new(String)
}

/**
* 序列化tree:[{"test":[{"params":[{"name":"name1"},{"key":"value"},{"count":{"value":"www"}}]},{"params":[{"name":"name2"},{"key":"value"}]}]}]
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
	//currentNum := leafNum
	for key, val := range treeList {
		nodeList := val.GetNode()
		if len(nodeList) <= 0 {
			continue
		}
		node := nodeList[0]
		buffer.WriteRune(common.LeftRrance)
		if val.IsLeaf() {
			buffer.WriteString(formatBytes(node.GetName()))
			buffer.WriteRune(common.Colon)
			buffer.WriteString(formatBytes(node.GetData()))
		} else {
			childrenNum = len(val.GetChildren())
			buffer.WriteString(formatBytes(node.GetName()))
			buffer.WriteRune(common.Colon)
			if childrenNum > 1 {
				buffer.WriteRune(common.LeftBracket)
			}
			BindString(val.GetChildren(), buffer)
			if childrenNum > 1 {
				buffer.WriteRune(common.RightBracket)
			}
		}
		buffer.WriteRune(common.RightRrance)
		if key != len(treeList)-1 {
			buffer.WriteRune(common.Comma)
		}
	}
	return childrenNum
}

func formatBytes(bytes []byte) string {
	str := common.Bytes2string(bytes)
	if str == "" {
		return str
	}
	return "\"" + str + "\""
}

func UnBindString(str string, tree *strategy.TreeStruct) {

}
