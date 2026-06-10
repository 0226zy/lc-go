package heap

// MinHeap 最小堆
// 基于切片实现的二叉最小堆，支持 Push、Pop、Peek、Size 操作。
// 时间复杂度：Push O(log n)，Pop O(log n)，Peek O(1)
type MinHeap struct {
	data []int
}

// NewMinHeap 创建一个空的最小堆
func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: make([]int, 0),
	}
}

// NewMinHeapFromSlice 从切片构建最小堆（堆化）
// 时间复杂度: O(n)
func NewMinHeapFromSlice(nums []int) *MinHeap {
	ret := NewMinHeap()
	for _, num := range nums {
		ret.Push(num)
	}
	return ret
}

// Push 向堆中插入元素
// 时间复杂度: O(log n)
func (h *MinHeap) Push(val int) {
	h.data = append(h.data, val)
	h.up(len(h.data) - 1)
}

// Pop 移除并返回堆顶（最小值）
// 时间复杂度: O(log n)
func (h *MinHeap) Pop() int {
	ret := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down(0)
	return ret
}

// Peek 返回堆顶元素（最小值），不移除
// 时间复杂度: O(1)
func (h *MinHeap) Peek() int {
	return h.data[0]
}

// Size 返回堆中元素个数
func (h *MinHeap) Size() int {
	return len(h.data)
}

// IsEmpty 判断堆是否为空
func (h *MinHeap) IsEmpty() bool {
	return len(h.data) == 0
}

// Data 返回堆内数据的副本
func (h *MinHeap) Data() []int {
	ret := make([]int, len(h.data))
	copy(ret, h.data)
	return ret
}
func (h *MinHeap) up(idx int) {
	for idx >= 0 {
		pattern := (idx - 1) / 2
		if h.data[idx] >= h.data[pattern] {
			break
		}
		h.data[idx], h.data[pattern] = h.data[pattern], h.data[idx]
		idx = pattern
	}
}
func (h *MinHeap) down(idx int) {
	currIdx := idx
	n := h.Size()
	for {
		minIdx := currIdx
		left, right := 2*currIdx+1, 2*currIdx+2

		if left < n && h.data[left] < h.data[currIdx] {
			minIdx = left
		}
		if right < n && h.data[right] < h.data[currIdx] {
			minIdx = right
		}
		// 没有找到可以交换的点，左右节点都比当前节点小
		if minIdx == idx {
			break
		}
		h.data[currIdx], h.data[minIdx] = h.data[minIdx], h.data[currIdx]

		// 处理被交换后的节点
		currIdx = minIdx

	}
}
