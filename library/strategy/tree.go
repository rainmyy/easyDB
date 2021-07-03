package strategy

type TreeStruct struct {
	node     []*NodeStruct
	children []*TreeStruct
	parent   *TreeStruct
	high     int //计算层高
	leaf     bool
}

type NodeStruct struct {
	data   interface{}
	name   []byte
	length int
}

func (this *TreeStruct) GetNode() []*NodeStruct {
	if this.node == nil {
		return nil
	}
	return this.node
}

func (this *TreeStruct) SetNode(node *NodeStruct) *TreeStruct {
	if node == nil {
		return this
	}
	this.node = append(this.node, node)
	if node.data != nil {
		this.leaf = true
	}
	return this
}
func (this *TreeStruct) GetChildren() []*TreeStruct {
	if this.children == nil {
		return nil
	}
	return this.children
}

func (this *TreeStruct) SetChildren(children *TreeStruct) *TreeStruct {
	if children == nil {
		return this
	}
	children.SetParent(this)
	children.SetHight(this.high + 1)
	this.children = append(this.children, children)
	return this
}

/**
*上一个节点
 */
func (this *TreeStruct) GetParent() *TreeStruct {
	if this.parent == nil {
		return this
	}
	return this.parent
}

func (this *TreeStruct) SetParent(tree *TreeStruct) *TreeStruct {
	if tree == nil {
		return this
	}
	this.parent = tree
	return this
}

func (this *TreeStruct) GetRoot() *TreeStruct {
	if this.IsRoot() == true {
		return this
	}
	for this.parent != nil {
		this = this.parent
	}
	return this
}

func (this *TreeStruct) GetHight() int {
	return this.high
}

func (this *TreeStruct) SetHight(hight int) *TreeStruct {
	this.high = hight
	return this
}

func (this *TreeStruct) IsRoot() bool {
	if this.parent != nil {
		return false
	}
	return true
}

func TreeInstance() *TreeStruct {
	return &TreeStruct{
		node:     make([]*NodeStruct, 0),
		children: make([]*TreeStruct, 0),
		parent:   new(TreeStruct),
	}
}

func (this *NodeStruct) GetData() interface{} {
	return this.data
}

func NodeInstance(key []byte, value interface{}) *NodeStruct {
	lenth := 0
	switch value := value.(type) {
	case string:
	case []byte:
	case map[interface{}]interface{}:
	case []interface{}:
		lenth = len(value)
	}
	return &NodeStruct{name: key, data: value, length: lenth}
}
