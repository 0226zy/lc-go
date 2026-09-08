package kthlargestelementinanarray

import (
	"container/heap"
	"sort"
)

// minHeap 小顶堆，堆顶为当前堆中的最小值
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

// FindKthLargest 数组中的第K个最大元素
// 返回数组中第 k 个最大的元素（将数组按从大到小排序后的第 k 个元素）。
// 维护一个大小为 k 的小顶堆：堆顶始终是这 k 个数里的最小值，
// 也就是整个数组第 k 大的元素。
// 时间复杂度: O(n log k)  空间复杂度: O(k)
func FindKthLargest(nums []int, k int) int {
	h := minHeap{}
	for _, num := range nums {
		heap.Push(&h, num)
		if h.Len() > k {
			heap.Pop(&h) // 弹出堆中最小值，堆中始终保留最大的 k 个数
		}
	}
	return h[0]
}

// FindKthLargestBySort 直接排序对照解法
// 时间复杂度: O(n log n)  空间复杂度: O(log n)
func FindKthLargestBySort(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
