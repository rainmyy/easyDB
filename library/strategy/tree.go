package strategy

import "time"

type TreeStruct struct {
	node         []*NodeStruct
	children     []*TreeStruct
	parent       *TreeStruct
	high         int //计算层高
	leaf         bool
	childLeafNum int
}

type NodeStruct struct {
	data   []byte
	name   []byte
	length int
	/**
	* 数据的创建时间，隐藏数据不展示，通过该key值进行树的检索和排序
	* 第一次创建数据时初始化该值，直到NodeStruct被回收前该值保持不变
	 */
	createtime time.Time
	/**
	* 数据的更新时间，隐藏数据不展示，通过该key值进行树的检索和排序
	* 第一次创建数据时初始化该值，元素修改时修改该值
	 */
	updatetime time.Time

	/**
	* 每次修改数据
	 */
	backup []byte
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
	if len(node.data) > 0 {
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
	for _, val := range this.children {
		if val.IsLeaf() == true {
			this.childLeafNum++
		}
	}
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

func (this *TreeStruct) IsLeaf() bool {
	return this.leaf
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

func (this *NodeStruct) UpdateData(value []byte) *NodeStruct {
	nodeData := this.data
	hasDiff := false
	if len(this.data) == len(value) {
		for i := 0; i < len(nodeData); i++ {
			if this.data[i] != value[i] {
				hasDiff = true
			}
		}
	}
	if !hasDiff {
		return this
	}
	this.backup = nodeData
	this.data = value
	this.updatetime = time.Now()
	return this
}

func NodeInstance(key []byte, value []byte) *NodeStruct {
	return &NodeStruct{
		name:       key,
		data:       value,
		createtime: time.Now(),
		updatetime: time.Now(),
	}
}

func (this *NodeStruct) GetData() []byte {
	return this.data
}
func (this *NodeStruct) GetName() []byte {
	return this.name
}
