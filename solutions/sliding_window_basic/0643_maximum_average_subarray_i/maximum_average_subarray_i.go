package maximumaveragesubarrayi

// FindMaxAverage 子数组最大平均数 I
// 给定 n 个整数的数组 nums 和整数 k，找出长度恰好为 k 的连续子数组中平均数最大的，
// 返回该最大平均数。采用固定长度 k 的滑动窗口：先计算前 k 个元素之和，
// 之后每次右移一格，减去左端移出元素、加上右端新进入元素，维护最大窗口和。
// 时间复杂度: O(n) 数组只扫描一遍  空间复杂度: O(1) 只使用常数个变量
func FindMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	// 非法输入：空数组或非正窗口长度或窗口比数组还长
	if n == 0 || k <= 0 || k > n {
		return 0.0
	}
	// 先计算第一个窗口（前 k 个元素）的和
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	// 滑动窗口右移：每次加入右端新元素、减去左端移出元素
	for i := k; i < n; i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}
