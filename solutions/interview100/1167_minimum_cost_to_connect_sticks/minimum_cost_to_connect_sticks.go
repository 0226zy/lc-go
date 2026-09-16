package minimumcosttoconnectsticks

import "container/heap"

// intHeap 实现 heap.Interface 的最小堆
type intHeap []int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	top := old[n-1]
	*h = old[:n-1]
	return top
}

// ConnectSticks 连接木棍的最低费用
// 每次连接两根长度为 x、y 的木棍费用为 x + y，返回连接成一根木棍的最低总费用。
// 贪心策略：每步用最小堆选取当前最短的两根木棍合并（哈夫曼合并）。
// 时间复杂度: O(n log n)  空间复杂度: O(n)
func ConnectSticks(sticks []int) int {
	// 少于两根木棍无需连接
	if len(sticks) < 2 {
		return 0
	}

	h := intHeap(append([]int(nil), sticks...))
	heap.Init(&h)

	total := 0
	// 每次弹出最短的两根合并，费用累加后把新木棍压回堆
	for h.Len() > 1 {
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)
		total += x + y
		heap.Push(&h, x+y)
	}
	return total
}
