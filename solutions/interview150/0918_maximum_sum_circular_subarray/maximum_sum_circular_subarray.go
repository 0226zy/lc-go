package maximumsumcircularsubarray

// MaxSubarraySumCircular 环形子数组的最大和
// 环形数组的最大子数组和 = max(普通最大子数组和, 总和 - 最小子数组和)。
// 一次遍历同时维护最大子数组和与最小子数组和；
// 全为负数时子数组非空，不能取“总和 - 最小和 = 0”，需特判直接返回最大子数组和。
// 时间复杂度: O(n) 一次遍历  空间复杂度: O(1) 常数变量
func MaxSubarraySumCircular(nums []int) int {
	total := nums[0]
	maxCur, maxSum := nums[0], nums[0] // 最大子数组和（Kadane）
	minCur, minSum := nums[0], nums[0] // 最小子数组和（反向 Kadane）
	for i := 1; i < len(nums); i++ {
		x := nums[i]
		total += x
		// 最大子数组和：前面的和为负资产就抛弃
		if maxCur+x > x {
			maxCur += x
		} else {
			maxCur = x
		}
		if maxCur > maxSum {
			maxSum = maxCur
		}
		// 最小子数组和：前面的和为正资产就抛弃
		if minCur+x < x {
			minCur += x
		} else {
			minCur = x
		}
		if minCur < minSum {
			minSum = minCur
		}
	}
	// 全是负数时，子数组非空，不能取“总和 - 最小和 = 0”
	if maxSum < 0 {
		return maxSum
	}
	// 跨环情形 = 总和 - 中间最小子数组和
	if rest := total - minSum; rest > maxSum {
		maxSum = rest
	}
	return maxSum
}
