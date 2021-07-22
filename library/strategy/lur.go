package strategy

/**
* 数据读写最近最少使用算法
 */
type linkedNode struct {
	key   interface{}
	value interface{}
	pre   *linkedNode //上一个节点
	post  *linkedNode //下一个节点
}
type lRUCache struct {
	count      int
	capacity   int
	cache      map[interface{}]*linkedNode
	head, tail *linkedNode
}

func NewLRUCache(capacity int) *lRUCache {
	l := &lRUCache{}
	l.capacity = capacity
	l.cache = make(map[interface{}]*linkedNode)
	l.head = newLinkeNode()
	l.head.pre = nil
	l.tail = newLinkeNode()
	l.tail.post = nil
	l.head.post = l.tail
	l.tail.pre = l.head
	return l
}
func newLinkeNode() *linkedNode {
	return &linkedNode{}
}
func (l *lRUCache) Get(key interface{}) interface{} {
	node, ok := l.cache[key]
	if !ok {
		return -1
	}

	l.moveTohead(node)
	return node.value
}

func (l *lRUCache) Put(key interface{}, value interface{}) {
	node, ok := l.cache[key]
	if ok {
		node.value = value
		l.moveTohead(node)
		return
	}
	linkeNode := newLinkeNode()
	linkeNode.key = key
	linkeNode.value = value
	l.cache[key] = linkeNode
	l.addNode(linkeNode)
	l.count++
	if l.count > l.capacity {
		tail := l.popTail()
		if tail == nil {
			return
		}
		if _, ok := l.cache[tail.key]; ok {
			delete(l.cache, tail.key)
		}
		l.count--
	}
}

func (l *lRUCache) addNode(node *linkedNode) {
	node.pre = l.head
	node.post = l.head.post
	l.head.post.pre = node
	l.head.post = node
}

func (l *lRUCache) removeNode(node *linkedNode) {
	node.pre = newLinkeNode()
	pre := node.pre
	post := node.post
	pre.post = post
	post.pre = pre
}

func (l *lRUCache) moveTohead(node *linkedNode) {
	l.removeNode(node)
	l.addNode(node)
}

func (l *lRUCache) popTail() *linkedNode {
	res := l.tail.pre

	if res == nil {
		return nil
	}
	l.removeNode(res)
	return res
}
