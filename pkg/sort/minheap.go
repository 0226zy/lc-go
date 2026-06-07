package sort

// MinHeap 通用最小堆实现
// 通过 LessFunc 比较函数支持任意类型
type MinHeap[T any] struct {
	data    []T
	lessFunc func(a, b T) bool
}

// NewMinHeap 创建并初始化最小堆
func NewMinHeap[T any](lessFunc func(a, b T) bool) *MinHeap[T] {
	return &MinHeap[T]{
		data:    make([]T, 0),
		lessFunc: lessFunc,
	}
}

// Len 返回堆的大小
func (h *MinHeap[T]) Len() int {
	return len(h.data)
}

// Less 比较两个元素
func (h *MinHeap[T]) Less(i, j int) bool {
	return h.lessFunc(h.data[i], h.data[j])
}

// Swap 交换两个元素的位置
func (h *MinHeap[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Push 向堆中插入一个元素，O(log n)
func (h *MinHeap[T]) Push(x T) {
	h.data = append(h.data, x)
	h.heapifyUp(h.Len() - 1)
}

// Pop 从堆中弹出最小值，O(log n)
func (h *MinHeap[T]) Pop() T {
	if h.Len() == 0 {
		var zero T
		return zero
	}
	// 将堆顶与最后一个元素交换
	minVal := h.data[0]
	lastIdx := h.Len() - 1
	h.data[0] = h.data[lastIdx]
	h.data = h.data[:lastIdx]
	// 下沉调整
	h.heapifyDown(0)
	return minVal
}

// Peek 返回堆顶元素但不弹出
func (h *MinHeap[T]) Peek() T {
	if h.Len() == 0 {
		var zero T
		return zero
	}
	return h.data[0]
}

// heapifyUp 自下而上调整堆
func (h *MinHeap[T]) heapifyUp(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2
		if !h.Less(idx, parent) {
			break
		}
		h.Swap(idx, parent)
		idx = parent
	}
}

// heapifyDown 自上而下调整堆
func (h *MinHeap[T]) heapifyDown(idx int) {
	n := h.Len()
	for {
		left := 2*idx + 1
		right := 2*idx + 2
		smallest := idx

		if left < n && h.Less(left, smallest) {
			smallest = left
		}
		if right < n && h.Less(right, smallest) {
			smallest = right
		}
		if smallest == idx {
			break
		}
		h.Swap(idx, smallest)
		idx = smallest
	}
}
