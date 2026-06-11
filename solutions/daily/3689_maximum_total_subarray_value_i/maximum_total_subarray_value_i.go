package maxtotalsubarray

import "math"

// MaxTotalSubarrayValue 最大子数组总价值 I
// 给定整数数组 nums 和整数 k，选择恰好 k 个非空子数组（可重叠、可重复选择），
// 使得所有子数组的 (max - min) 之和最大。
// 时间复杂度: O(?)  空间复杂度: O(?)
func MaxTotalSubarrayValue(nums []int, k int) int64 {
	m := math.MinInt64
	n := math.MaxInt64
	for _, v := range nums {
		m = max(m, v)
		n = min(n, v)
	}
	return int64((m - n) * k)
}
