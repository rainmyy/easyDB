
type Tree struct {
	node     *Node
	children *Tree
}

type Node struct {
	data []byte
	len  int
}

