package maximumproductsubarray

// MaxProduct 乘积最大子数组（动态规划）
// 同时维护以当前元素结尾的最大乘积与最小乘积，因为负数可能让最小翻为最大。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	curMax, curMin := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		n := nums[i]
		// 必须用旧值计算，故先暂存
		a := curMax * n
		b := curMin * n
		curMax = max(n, a, b)
		curMin = min(n, a, b)
		if curMax > result {
			result = curMax
		}
	}
	return result
}

// MaxProductPrefixSuffix 乘积最大子数组（前缀/后缀扫描）
// 从左到右累乘一遍、从右到左累乘一遍，取最大值。遇 0 重置为 1。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxProductPrefixSuffix(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	// 从左到右
	prefix := 1
	for _, n := range nums {
		prefix *= n
		if prefix > result {
			result = prefix
		}
		if n == 0 {
			prefix = 1 // 遇 0 重置，相当于分段
		}
	}
	// 从右到左
	suffix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		suffix *= nums[i]
		if suffix > result {
			result = suffix
		}
		if nums[i] == 0 {
			suffix = 1
		}
	}
	return result
}

func max(a ...int) int {
	m := a[0]
	for _, v := range a[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func min(a ...int) int {
	m := a[0]
	for _, v := range a[1:] {
		if v < m {
			m = v
		}
	}
	return m
}
