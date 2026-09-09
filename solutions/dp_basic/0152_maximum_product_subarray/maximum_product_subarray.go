package maximumproductsubarray

// MaxProduct 乘积最大子数组（标准 DP 数组版）
// 求乘积最大的非空连续子数组的乘积。
// 负数会让最大变最小、最小变最大，因此同时维护两个状态：
// dpMax[i] 表示以 nums[i] 结尾的最大乘积，dpMin[i] 表示以 nums[i] 结尾的最小乘积。
// dpMax[i] = max(nums[i], dpMax[i-1]*nums[i], dpMin[i-1]*nums[i])
// dpMin[i] = min(nums[i], dpMin[i-1]*nums[i], dpMax[i-1]*nums[i])
// 时间复杂度: O(n)  空间复杂度: O(n)
func MaxProduct(nums []int) int {
	n := len(nums)
	dpMax := make([]int, n)
	dpMin := make([]int, n)
	dpMax[0], dpMin[0] = nums[0], nums[0] // base case：单元素
	ans := nums[0]
	for i := 1; i < n; i++ {
		// 三个候选：单独成段、接在前面的最大乘积后、接在前面的最小乘积后
		dpMax[i] = max(nums[i], max(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]))
		dpMin[i] = min(nums[i], min(dpMin[i-1]*nums[i], dpMax[i-1]*nums[i]))
		ans = max(ans, dpMax[i]) // 最大子数组可以在任何位置结尾
	}
	return ans
}

// MaxProductOptimized 乘积最大子数组（滚动变量空间优化版）
// dpMax/dpMin 只依赖前一个位置，用两个滚动变量代替 dp 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxProductOptimized(nums []int) int {
	curMax, curMin, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		prevMax := curMax // 先保存旧值，避免更新 curMin 时读到新 curMax
		curMax = max(nums[i], max(curMax*nums[i], curMin*nums[i]))
		curMin = min(nums[i], min(curMin*nums[i], prevMax*nums[i]))
		ans = max(ans, curMax)
	}
	return ans
}
