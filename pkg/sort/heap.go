package sort

// Heap 通用堆实现（支持最小堆和最大堆）
// 通过 lessFunc 比较函数决定堆的类型：
//   - 若 lessFunc = a < b，则为最小堆（父节点小于子节点）
//   - 若 lessFunc = a > b，则为最大堆（父节点大于子节点）
type Heap[T any] struct {
	data    []T
	lessFunc func(a, b T) bool
}

// NewHeap 创建并初始化通用堆
// lessFunc 参数决定堆的类型：
//   - lessFunc(a, b) = a < b → 最小堆
//   - lessFunc(a, b) = a > b → 最大堆
func NewHeap[T any](lessFunc func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data:    make([]T, 0),
		lessFunc: lessFunc,
	}
}

// NewMinHeap 创建最小堆（便捷构造函数）
// 对于可比较的类型，使用 func(a, b T) bool { return a < b }
func NewMinHeap[T any](less func(a, b T) bool) *Heap[T] {
	return NewHeap(less)
}

// NewMaxHeap 创建最大堆（便捷构造函数）
// 对于可比较的类型，使用 func(a, b T) bool { return a > b }
func NewMaxHeap[T any](less func(a, b T) bool) *Heap[T] {
	return NewHeap(less)
}

// Len 返回堆的大小
func (h *Heap[T]) Len() int {
	return len(h.data)
}

// Less 比较两个元素（供 heapifyUp/heapifyDown 使用）
func (h *Heap[T]) Less(i, j int) bool {
	return h.lessFunc(h.data[i], h.data[j])
}

// Swap 交换两个元素的位置
func (h *Heap[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Push 向堆中插入一个元素，O(log n)
func (h *Heap[T]) Push(x T) {
	h.data = append(h.data, x)
	h.heapifyUp(h.Len() - 1)
}

// Pop 从堆中弹出堆顶元素，O(log n)
// 对于最小堆：弹出最小值；对于最大堆：弹出最大值
func (h *Heap[T]) Pop() T {
	if h.Len() == 0 {
		var zero T
		return zero
	}
	// 将堆顶与最后一个元素交换
	top := h.data[0]
	lastIdx := h.Len() - 1
	h.data[0] = h.data[lastIdx]
	h.data = h.data[:lastIdx]
	// 下沉调整
	h.heapifyDown(0)
	return top
}

// Peek 返回堆顶元素但不弹出
// 对于最小堆：返回最小值；对于最大堆：返回最大值
func (h *Heap[T]) Peek() T {
	if h.Len() == 0 {
		var zero T
		return zero
	}
	return h.data[0]
}

// heapifyUp 自下而上调整堆（插入后调用）
func (h *Heap[T]) heapifyUp(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2
		if !h.Less(idx, parent) {
			break
		}
		h.Swap(idx, parent)
		idx = parent
	}
}

// heapifyDown 自上而下调整堆（弹出后调用）
func (h *Heap[T]) heapifyDown(idx int) {
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
