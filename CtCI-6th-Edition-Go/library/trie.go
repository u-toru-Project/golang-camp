package library

type TrieNode struct {
	Children   map[rune]*TrieNode
	Terminates bool
	Character  rune
}

func NewTrieNode() *TrieNode {
	return &TrieNode{Children: make(map[rune]*TrieNode)}
}

func (n *TrieNode) AddWord(word string) {
	if word == "" {
		return
	}
	first := []rune(word)[0]
	child, ok := n.Children[first]
	if !ok {
		child = NewTrieNode()
		child.Character = first
		n.Children[first] = child
	}
	rest := string([]rune(word)[1:])
	if rest != "" {
		child.AddWord(rest)
		return
	}
	child.Terminates = true
}

func (n *TrieNode) GetChild(character rune) *TrieNode {
	return n.Children[character]
}

type Trie struct {
	Root *TrieNode
}

func NewTrie(words []string) *Trie {
	trie := &Trie{Root: NewTrieNode()}
	for _, word := range words {
		trie.Root.AddWord(word)
	}
	return trie
}

func (t *Trie) Contains(prefix string, exact bool) bool {
	node := t.Root
	for _, character := range prefix {
		node = node.GetChild(character)
		if node == nil {
			return false
		}
	}
	return !exact || node.Terminates
}
