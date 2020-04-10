package tree

type Trie struct {
	next   map[rune]*Trie
	isWord bool
}

func NewTrie() *Trie {
	root := new(Trie)
	root.next = make(map[rune]*Trie, 26)
	root.isWord = false
	return root
}

func (t *Trie) Insert(word string) {
	node := t
	for _, v := range word {
		if node.next[v] == nil {
			node.next[v] = NewTrie()
		}
		node = node.next[v]
	}
	node.isWord = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, v := range word {
		if node.next[v] == nil {
			return false
		}
		node = node.next[v]
	}
	return node.isWord
}

func (t *Trie) StartWith(prefix string) bool {
	node := t
	for _, v := range prefix {
		if node.next[v] == nil {
			return false
		}
		node = node.next[v]
	}
	return true
}
