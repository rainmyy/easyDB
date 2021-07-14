package strategy

import "sync"

const treVal = 0.75

type Stack struct {
	top    *node
	length int
	lock   *sync.RWMutex
}

type node struct {
	value interface{}
	prev  *node
}

func StackInstance() *Stack {
	return &Stack{nil, 0, &sync.RWMutex{}}
}
func (this *Stack) Len() int {
	return this.length
}
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}

	return this.top.value
}

func (this *Stack) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

func (this *Stack) Push(value interface{}) {
	this.lock.Lock()
	this.lock.Unlock()
	n := &node{value, this.top}
	this.top = n
	this.length++
}
