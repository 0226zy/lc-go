package painthouseii

import "math"

// MinCostII 粉刷房子 II
// 一排 n 个房子，每个房子可刷 k 种颜色之一，costs[i][j] 表示第 i 个房子
// 刷成颜色 j 的花费；要求相邻房子颜色不同，返回刷完所有房子的最小花费。
// 时间复杂度: O(n*k) 每行只需最小值/次小值完成转移  空间复杂度: O(k) 滚动数组
func MinCostII(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	k := len(costs[0])
	if k == 1 {
		// 只有一种颜色时，仅当只有一栋房子才存在合法方案
		if n == 1 {
			return costs[0][0]
		}
		return -1
	}

	// dp[j] 表示刷到当前房子、且当前房子刷颜色 j 时的最小总花费
	dp := make([]int, k)
	copy(dp, costs[0])
	min1, min2, minIdx := twoMin(dp)

	for i := 1; i < n; i++ {
		nMin1, nMin2, nMinIdx := math.MaxInt, math.MaxInt, -1
		for j := 0; j < k; j++ {
			// 不能与上一栋同色：上一栋取最小值颜色时退而求其次
			prev := min1
			if j == minIdx {
				prev = min2
			}
			dp[j] = costs[i][j] + prev
			// 顺带维护本行的最小值与次小值
			if dp[j] < nMin1 {
				nMin1, nMin2, nMinIdx = dp[j], nMin1, j
			} else if dp[j] < nMin2 {
				nMin2 = dp[j]
			}
		}
		min1, min2, minIdx = nMin1, nMin2, nMinIdx
	}
	return min1
}

// twoMin 返回切片中的最小值、次小值以及最小值的下标
func twoMin(nums []int) (min1, min2, minIdx int) {
	min1, min2, minIdx = math.MaxInt, math.MaxInt, -1
	for i, v := range nums {
		if v < min1 {
			min1, min2, minIdx = v, min1, i
		} else if v < min2 {
			min2 = v
		}
	}
	return
}
