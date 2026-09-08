package designaddandsearchwordsdatastructure

// WordDictionary 支持添加单词与按模式搜索的词典
// Search 支持 '.' 通配符，匹配任意一个字母
// 底层是字典树：children[i] 对应字母 'a'+i 的子节点，isEnd 标记单词结尾
type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

// Constructor 初始化词典
func Constructor() WordDictionary {
	return WordDictionary{}
}

// AddWord 将 word 加入词典
// 时间复杂度: O(L) L 为 word 长度  空间复杂度: O(L) 最坏情况新增 L 个节点
func (d *WordDictionary) AddWord(word string) {
	node := d
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &WordDictionary{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

// Search 查找 word 是否在词典中，'.' 可匹配任意一个字母
// 时间复杂度: 最坏 O(26^L) 全是通配符时退化为多路回溯，一般 O(L)  空间复杂度: O(L) 递归深度
func (d *WordDictionary) Search(word string) bool {
	return d.searchFrom(word, 0)
}

// searchFrom 从第 index 个字符开始在当前子树下回溯搜索
func (d *WordDictionary) searchFrom(word string, index int) bool {
	if index == len(word) {
		return d.isEnd
	}
	ch := word[index]
	if ch == '.' {
		// 通配符：尝试所有存在的子分支
		for _, child := range d.children {
			if child != nil && child.searchFrom(word, index+1) {
				return true
			}
		}
		return false
	}
	child := d.children[ch-'a']
	if child == nil {
		return false
	}
	return child.searchFrom(word, index+1)
}
