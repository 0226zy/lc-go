package minstack

// MinStack 最小栈
// 支持 push、pop、top 操作，并能在常数时间内检索到最小元素。
// 使用辅助栈同步保存当前最小值，辅助栈栈顶即主栈当前最小元素。
type MinStack struct {
	data    []int // 主栈，存放所有元素
	minData []int // 辅助栈，栈顶为当前主栈的最小值
}

// Constructor 初始化最小栈
// 时间复杂度: O(1)  空间复杂度: O(1)
func Constructor() MinStack {
	return MinStack{}
}

// Push 将元素 val 推入堆栈
// 时间复杂度: O(1)  空间复杂度: O(1)
func (s *MinStack) Push(val int) {
	s.data = append(s.data, val)
	// 注意用 <=，保证重复的最小值会重复入辅助栈，弹出时一一对应
	if len(s.minData) == 0 || val <= s.minData[len(s.minData)-1] {
		s.minData = append(s.minData, val)
	}
}

// Pop 删除堆栈顶部的元素
// 时间复杂度: O(1)  空间复杂度: O(1)
func (s *MinStack) Pop() {
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	if top == s.minData[len(s.minData)-1] {
		s.minData = s.minData[:len(s.minData)-1]
	}
}

// Top 获取堆栈顶部的元素
// 时间复杂度: O(1)  空间复杂度: O(1)
func (s *MinStack) Top() int {
	return s.data[len(s.data)-1]
}

// GetMin 获取堆栈中的最小元素
// 时间复杂度: O(1)  空间复杂度: O(1)
func (s *MinStack) GetMin() int {
	return s.minData[len(s.minData)-1]
}
