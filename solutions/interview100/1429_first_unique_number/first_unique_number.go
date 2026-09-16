package firstuniquenumber

// FirstUnique 第一个唯一数字
// 维护一个整数序列，支持高效查询第一个唯一数字以及追加新数字。
// Add 与 ShowFirstUnique 均摊时间复杂度: O(1)  空间复杂度: O(n)
type FirstUnique struct {
	count map[int]int // 每个数字的出现次数
	queue []int       // 当前仍唯一的候选数字队列，保持加入顺序
}

// Constructor 用 nums 初始化 FirstUnique
func Constructor(nums []int) FirstUnique {
	fu := FirstUnique{
		count: make(map[int]int, len(nums)),
	}
	for _, v := range nums {
		fu.Add(v)
	}
	return fu
}

// ShowFirstUnique 返回序列中第一个唯一数字，不存在时返回 -1
func (fu *FirstUnique) ShowFirstUnique() int {
	// 惰性删除：队首元素若已不再唯一则弹出
	for len(fu.queue) > 0 && fu.count[fu.queue[0]] > 1 {
		fu.queue = fu.queue[1:]
	}
	if len(fu.queue) == 0 {
		return -1
	}
	return fu.queue[0]
}

// Add 将 value 追加到序列中
func (fu *FirstUnique) Add(value int) {
	// 首次出现的数字才有资格成为唯一数字，加入候选队列
	if fu.count[value] == 0 {
		fu.queue = append(fu.queue, value)
	}
	fu.count[value]++
}
