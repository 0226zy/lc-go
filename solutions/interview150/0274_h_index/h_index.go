package hindex

import "sort"

// HIndex H 指数
// 给定一个整数数组 citations，citations[i] 表示研究者的第 i 篇论文被引用的次数。
// h 指数的定义：一名研究者的 h 指数是指其发表了 h 篇论文，每篇至少被引用 h 次，
// 且其余论文每篇被引用次数不超过 h 次。计算并返回该研究者的 h 指数。
// 时间复杂度: O(n log n) 排序  空间复杂度: O(log n) 排序递归栈
func HIndex(citations []int) int {
	n := len(citations)
	// 按引用次数从大到小排序
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	// 排序后，citations[i] 是第 (i+1) 高的引用次数。
	// 如果 citations[i] >= i+1，说明至少有 i+1 篇论文引用数 >= i+1；
	// 从左往右找最后一个满足条件的 i，h = i+1。
	for i := 0; i < n; i++ {
		if citations[i] < i+1 {
			return i
		}
	}
	return n
}

// HIndexBucket H 指数（计数桶解法）
// 时间复杂度: O(n)  空间复杂度: O(n)
func HIndexBucket(citations []int) int {
	n := len(citations)
	// 引用次数 >= n 的论文统一放入桶 n（因为 h 最大不会超过 n）
	buckets := make([]int, n+1)
	for _, c := range citations {
		if c >= n {
			buckets[n]++
		} else {
			buckets[c]++
		}
	}
	// 从桶 n 往桶 0 累加论文数，首个满足"累计篇数 >= 引用次数"的就是 h
	papers := 0
	for h := n; h >= 0; h-- {
		papers += buckets[h]
		if papers >= h {
			return h
		}
	}
	return 0
}
