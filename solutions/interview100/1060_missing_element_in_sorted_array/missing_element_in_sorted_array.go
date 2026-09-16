package missingelementinsortedarray

// MissingElement 有序数组中的缺失元素
// nums 严格递增，求在其最小值与最大值之间（不足时可延伸到最大值之外）
// 按顺序排列的第 k 个缺失整数（k 从 1 开始计数）。
// 时间复杂度: O(log n)  空间复杂度: O(1)
func MissingElement(nums []int, k int) int {
	n := len(nums)
	// missing(i) 表示 [nums[0], nums[i]] 区间内缺失的整数个数
	missing := func(i int) int {
		return nums[i] - nums[0] - i
	}

	// 二分找最大的下标 idx，使得 missing(idx) < k
	// missing(0) = 0 < k 恒成立，因此 idx 一定存在
	lo, hi := 0, n-1
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if missing(mid) < k {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	// 第 k 个缺失数落在 nums[idx] 之后
	return nums[lo] + (k - missing(lo))
}
