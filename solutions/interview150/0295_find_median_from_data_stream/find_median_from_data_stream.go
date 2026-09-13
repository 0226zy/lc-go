package findmedianfromdatastream

import "container/heap"

// maxHeap 大顶堆：存放数据流中较小的一半，堆顶为这一半的最大值
type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// minHeap 小顶堆：存放数据流中较大的一半，堆顶为这一半的最小值
type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MedianFinder 数据流的中位数
// 维护两个堆：low（大顶堆）存较小一半，high（小顶堆）存较大一半。
// 不变量：len(low) == len(high) 或 len(low) == len(high)+1，且 low 堆顶 <= high 堆顶。
type MedianFinder struct {
	low  maxHeap // 较小一半
	high minHeap // 较大一半
}

// Constructor 初始化 MedianFinder
func Constructor() MedianFinder {
	return MedianFinder{}
}

// AddNum 将整数 num 加入数据流
// 先按大小放入对应堆，再通过一次 pop/push 恢复两边长度差不超过 1。
// 时间复杂度: O(log n)  空间复杂度: O(1) 额外（整体存储 O(n)）
func (mf *MedianFinder) AddNum(num int) {
	if mf.low.Len() == 0 || num <= mf.low[0] {
		heap.Push(&mf.low, num)
	} else {
		heap.Push(&mf.high, num)
	}

	// 平衡：较小半至多比更大半多 1 个元素
	if mf.low.Len() > mf.high.Len()+1 {
		heap.Push(&mf.high, heap.Pop(&mf.low).(int))
	} else if mf.high.Len() > mf.low.Len() {
		heap.Push(&mf.low, heap.Pop(&mf.high).(int))
	}
}

// FindMedian 返回当前所有元素的中位数
// 奇数个元素时中位数为 low 堆顶；偶数个为两个堆顶的平均值。
// 时间复杂度: O(1)  空间复杂度: O(1)
func (mf *MedianFinder) FindMedian() float64 {
	if mf.low.Len() > mf.high.Len() {
		return float64(mf.low[0])
	}
	return float64(mf.low[0]+mf.high[0]) / 2.0
}
