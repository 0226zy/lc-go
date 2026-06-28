package implementtrieprefixtree

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

// Trie 前缀树（字典树）数据结构
type Trie struct {
	root *TrieNode
}

// Constructor 初始化前缀树对象
func Constructor() Trie {
	return Trie{root: &TrieNode{}}
}

// Insert 向前缀树中插入字符串 word
func (t *Trie) Insert(word string) {
	curr := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &TrieNode{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
}

// Search 如果字符串 word 在前缀树中，返回 true；否则，返回 false
func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return curr.isEnd
}

// StartsWith 如果之前已经插入的字符串 word 的前缀之一为 prefix，返回 true；否则，返回 false
func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for _, ch := range prefix {
		idx := ch - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return true
}
