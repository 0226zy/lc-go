package minstack

// MinStack 最小栈（辅助栈实现）
// 数据栈存储实际值，辅助栈存储当前最小值，所有操作 O(1)。
type MinStack struct {
	data []int
	mins []int
}

// Constructor 初始化 MinStack
func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		mins: make([]int, 0),
	}
}

// Push 将元素 val 推入堆栈
func (s *MinStack) Push(val int) {
	s.data = append(s.data, val)
	if len(s.mins) == 0 || val <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, val)
	} else {
		s.mins = append(s.mins, s.mins[len(s.mins)-1])
	}
}

// Pop 删除堆栈顶部的元素
func (s *MinStack) Pop() {
	s.data = s.data[:len(s.data)-1]
	s.mins = s.mins[:len(s.mins)-1]
}

// Top 获取堆栈顶部的元素
func (s *MinStack) Top() int {
	return s.data[len(s.data)-1]
}

// GetMin 获取堆栈中的最小元素
func (s *MinStack) GetMin() int {
	return s.mins[len(s.mins)-1]
}

// MinStackDiff 最小栈（差值法实现）
// 使用一个栈存储当前值与最小值的差值，节省辅助栈空间。
type MinStackDiff struct {
	stack  []int
	minVal int
}

// ConstructorDiff 初始化 MinStackDiff
func ConstructorDiff() MinStackDiff {
	return MinStackDiff{
		stack:  make([]int, 0),
		minVal: 0,
	}
}

// Push 将元素 val 推入堆栈
func (s *MinStackDiff) Push(val int) {
	if len(s.stack) == 0 {
		s.minVal = val
		s.stack = append(s.stack, 0)
		return
	}
	diff := val - s.minVal
	s.stack = append(s.stack, diff)
	if diff < 0 {
		s.minVal = val
	}
}

// Pop 删除堆栈顶部的元素
func (s *MinStackDiff) Pop() {
	diff := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if diff < 0 {
		s.minVal = s.minVal - diff
	}
}

// Top 获取堆栈顶部的元素
func (s *MinStackDiff) Top() int {
	diff := s.stack[len(s.stack)-1]
	if diff < 0 {
		return s.minVal
	}
	return s.minVal + diff
}

// GetMin 获取堆栈中的最小元素
func (s *MinStackDiff) GetMin() int {
	return s.minVal
}
