type trideNode struct {
	char     string
	isEnding bool
	children map[rune]*trideNode
}

func NewTrieNode(char string) *trideNode {
	return &trideNode{
		char:     char,
		isEnding: false,
		children: make(map[rune]*trideNode),
	}
}

type Trie struct {
	root *trideNode
}

func NewTrie() *Trie {
	trideNode := NewTrieNode("/")
	return &Trie{trideNode}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, code := range word {
		value, ok := node.children[code]
		fmt.Println(code, value)
		if !ok {
			value = NewTrieNode(string(code))
			node.children[code] = value
		}
		node = value
	}
	node.isEnding = true
}

func (t *Trie) Find(word string) bool {
	node := t.root
	for _, code := range word {
		value, ok := node.children[code]
		if !ok {
			return false
		}
		node = value
	}
	if node.isEnding == false {
		return false
	}
	return true
}