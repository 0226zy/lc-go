package implementtrieprefixtree

// Trie 前缀树（字典树）节点
// children[i] 对应字母 'a'+i 的子节点，isEnd 标记该节点是否是一个单词的结尾
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

// Constructor 初始化前缀树
func Constructor() Trie {
	return Trie{}
}

// Insert 向前缀树中插入字符串 word
// 时间复杂度: O(L) L 为 word 长度  空间复杂度: O(L) 最坏情况新增 L 个节点
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Trie{}
		}
		node = node.children[idx]
	}
	node.isEnd = true

}

// Search 查找字符串 word 是否在前缀树中（必须是完整单词）
// 时间复杂度: O(L)  空间复杂度: O(1)
func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd

}

// StartsWith 判断前缀 prefix 是否匹配任意已插入单词
// 时间复杂度: O(L)  空间复杂度: O(1)
func (t *Trie) StartsWith(prefix string) bool {
	node := t.searchPrefix(prefix)
	return node != nil
}

// searchPrefix 沿前缀逐字符下钻，返回最后一个字符对应的节点；前缀不存在返回 nil
func (t *Trie) searchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return nil
		}
		node = node.children[idx]
	}
	return node
}
