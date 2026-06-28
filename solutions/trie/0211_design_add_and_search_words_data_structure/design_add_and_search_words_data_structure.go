package designaddandsearchwordsdatastructure

// WordDictionary 添加与搜索单词的数据结构
type WordDictionary struct {
	// TODO: 定义前缀树结构
}

// Constructor 初始化词典对象
func Constructor() WordDictionary {
	// TODO: 实现
	return WordDictionary{}
}

// AddWord 将 word 添加到数据结构中
func (d *WordDictionary) AddWord(word string) {
	// TODO: 实现
}

// Search 如果存在与 word 匹配的字符串返回 true，'.' 可匹配任意字母
// 时间复杂度: O(26^L) 最坏（全通配符）  空间复杂度: O(L) 递归栈
func (d *WordDictionary) Search(word string) bool {
	// TODO: 实现
	return false
}
