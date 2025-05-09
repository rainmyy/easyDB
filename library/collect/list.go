package collect

import (
	"sync"
)

/***
*双向链表
 */

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type ListObj struct {
	head   *Node
	tail   *Node
	length uint
	mutex  *sync.RWMutex
}

func ListInstance() *ListObj {
	return &ListObj{mutex: new(sync.RWMutex)}
}

// 尾部压入数据
func (this *ListObj) Append(data interface{}) bool {
	if data == nil {
		return false
	}
	this.mutex.Lock()
	defer this.mutex.Unlock()
	var node = new(Node)
	node.data = data
	if this.length == 0 {
		this.head = node
		this.tail = node
		this.length = 1
		return true
	}
	tail := this.tail
	node.prev = tail
	tail.next = node
	this.tail = node
	this.length += 1
	return true
}

func (this *ListObj) Insert(index uint, data interface{}) bool {
	if data == nil {
		return false
	}
	if index > this.length {
		return false
	}
	this.mutex.Lock()
	defer this.mutex.Unlock()
	var node = new(Node)
	if index == 0 {
		node.next = this.head
		this.head.prev = node
		this.head = node
		this.length += 1
		return true
	}
	var i uint
	ptr := this.head
	for i = 1; i < index; i++ {
		ptr = ptr.next
	}
	next := ptr.next
	ptr.next = node
	node.prev = ptr
	next.prev = node
	node.next = next
	this.length += 1
	return true
}

func (this *ListObj) Delete(index uint) bool {
	if this == nil || index > this.length-1 {
		return false
	}

	this.mutex.Lock()
	defer this.mutex.Unlock()
	if index == 0 {
		head := this.head.next
		this.head = head
		this.head.prev = nil
		if this.length == 1 {
			this.tail = nil
		}
		this.length -= 1
		return true
	}
	ptr := this.head
	var i uint
	for i = 1; i < index; i++ {
		ptr = ptr.next
	}
	next := ptr.next
	ptr.next = next.next
	next.next.prev = ptr
	if index == this.length-1 {
		this.tail = ptr
	}
	this.length -= 1
	return true
}

func (this *ListObj) Get(index uint) *Node {
	if this == nil || index > this.length-1 {
		return nil
	}
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	node := this.head
	for i := 0; i < int(index); i++ {
		node = node.next
	}
	return node
}

/**
* 查找链表中的元素
 */
func (this *ListObj) Find(data interface{}) *Node {
	if this == nil {
		return nil
	}
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	head := this.head
	tail := this.tail
	var start uint = 0
	var end uint = this.length - 1
	for start != end {
		if head.data == data {
			return head
		} else if tail.data == data {
			return tail
		}
		head = head.next
		tail = tail.prev
		start++
		end--
	}
	return nil
}
