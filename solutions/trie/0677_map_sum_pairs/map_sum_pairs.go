package mapsumpairs

// MapSum 键值映射
type MapSum struct {
	// TODO: 定义前缀树结构
}

// Constructor 初始化对象
func Constructor() MapSum {
	// TODO: 实现
	return MapSum{}
}

// Insert 插入或更新键值对
func (m *MapSum) Insert(key string, val int) {
	// TODO: 实现
}

// Sum 返回所有以 prefix 为前缀的键的值总和
// 时间复杂度: O(L + 子树大小)  空间复杂度: O(L) 递归栈
func (m *MapSum) Sum(prefix string) int {
	// TODO: 实现
	return 0
}
