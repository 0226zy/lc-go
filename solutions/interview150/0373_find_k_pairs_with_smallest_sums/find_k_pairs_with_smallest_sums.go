package findkpairswithsmallestsums

import "container/heap"

// pair 数对下标组合：nums1[i] 与 nums2[j]
type pair struct {
	i, j int
}

// pairHeap 最小堆，按 nums1[i]+nums2[j] 的和比较
type pairHeap struct {
	pairs      []pair
	nums1      []int
	nums2      []int
}

func (h pairHeap) Len() int { return len(h.pairs) }
func (h pairHeap) Less(a, b int) bool {
	s1 := h.nums1[h.pairs[a].i] + h.nums2[h.pairs[a].j]
	s2 := h.nums1[h.pairs[b].i] + h.nums2[h.pairs[b].j]
	return s1 < s2
}
func (h pairHeap) Swap(a, b int) { h.pairs[a], h.pairs[b] = h.pairs[b], h.pairs[a] }
func (h *pairHeap) Push(x interface{}) {
	h.pairs = append(h.pairs, x.(pair))
}
func (h *pairHeap) Pop() interface{} {
	old := h.pairs
	n := len(old)
	x := old[n-1]
	h.pairs = old[:n-1]
	return x
}

// KSmallestPairs 查找和最小的 K 对数字
// 给定两个升序数组 nums1、nums2，找出和最小的 k 个数对 (u,v)。
// 把 nums1 的每个元素看作一条“队列头”，用最小堆做多路归并：
// 堆里始终保存每行的当前最小候选，每次弹出全局最小，再把同行的下一个推入。
// 时间复杂度: O(k log k)  空间复杂度: O(k)
func KSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k <= 0 {
		return nil
	}
	// 每行最多只需要前 k 个元素（第 k 个以后的元素不可能进入前 k 小）
	n := len(nums1)
	if n > k {
		n = k
	}
	h := &pairHeap{nums1: nums1, nums2: nums2}
	for i := 0; i < n; i++ {
		heap.Push(h, pair{i, 0}) // 初始：每行以 j=0 的数对入堆
	}

	result := make([][]int, 0, k)
	for h.Len() > 0 && len(result) < k {
		p := heap.Pop(h).(pair)
		result = append(result, []int{nums1[p.i], nums2[p.j]})
		if p.j+1 < len(nums2) {
			heap.Push(h, pair{p.i, p.j + 1}) // 同 nums1[i] 的下一个数对成为新候选
		}
	}
	return result
}
