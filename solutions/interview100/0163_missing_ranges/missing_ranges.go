package missingranges

// FindMissingRanges 缺失的区间
// 给定排序整数数组 nums（元素均在 [lower, upper] 内且互不相同），
// 返回恰好覆盖 [lower, upper] 中所有未出现整数的最大缺失区间列表。
// 时间复杂度: O(n) 遍历数组一遍  空间复杂度: O(1) 不计返回值
func FindMissingRanges(nums []int, lower int, upper int) [][]int {
	ans := make([][]int, 0)
	// 收集区间 [start, end]，仅在非空时加入结果
	addRange := func(start, end int) {
		if start <= end {
			ans = append(ans, []int{start, end})
		}
	}

	if len(nums) == 0 {
		addRange(lower, upper)
		return ans
	}

	// 首段：[lower, nums[0]-1]
	addRange(lower, nums[0]-1)
	// 中间段：相邻元素之间的空隙
	for i := 1; i < len(nums); i++ {
		addRange(nums[i-1]+1, nums[i]-1)
	}
	// 尾段：[nums[n-1]+1, upper]
	addRange(nums[len(nums)-1]+1, upper)
	return ans
}
