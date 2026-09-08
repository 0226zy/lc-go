package ipo

import (
	"container/heap"
	"sort"
)

// project 项目：capital 为启动所需资本，profit 为完成后的纯利润
type project struct {
	capital int
	profit  int
}

// profitHeap 大顶堆（按利润），堆顶为当前可做项目中利润最高的
type profitHeap []int

func (h profitHeap) Len() int            { return len(h) }
func (h profitHeap) Less(i, j int) bool  { return h[i] > h[j] } // 大顶堆
func (h profitHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *profitHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *profitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// FindMaximizedCapital IPO
// 最多完成 k 个项目，从初始资本 w 出发，每个项目需要 capital[i] 的启动资本，
// 完成后获得 profits[i] 的纯利润。返回最终能积累的最大资本。
// 贪心策略：每次从“当前资本能启动”的项目中选利润最大的一个去做，
// 用排序 + 指针解锁项目，用大顶堆快速取最大利润。
// 时间复杂度: O(n log n) 排序主导  空间复杂度: O(n)
func FindMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	projects := make([]project, n)
	for i := 0; i < n; i++ {
		projects[i] = project{capital: capital[i], profit: profits[i]}
	}
	// 按所需资本从小到大排序，配合指针逐步解锁
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	h := &profitHeap{}
	i := 0
	for step := 0; step < k; step++ {
		// 把所有资本不超过当前资本的项目加入大顶堆（每个项目只进堆一次）
		for i < n && projects[i].capital <= w {
			heap.Push(h, projects[i].profit)
			i++
		}
		if h.Len() == 0 {
			break // 没有可做的项目，提前结束
		}
		// 贪心：选利润最大的项目，资本增加
		w += heap.Pop(h).(int)
	}
	return w
}
