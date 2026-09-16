package maximumaveragesubarrayii

// FindMaxAverage 子数组最大平均数 II
// 给定一个整数数组 nums 和一个整数 k，找出长度大于等于 k 的连续子数组中平均数的最大值。
// 时间复杂度: O(n log(R/ε))，R 为值域范围，ε 为判定精度  空间复杂度: O(1)
func FindMaxAverage(nums []int, k int) float64 {
	// 最大平均数必然落在 [min(nums), max(nums)] 区间内
	lo, hi := float64(nums[0]), float64(nums[0])
	for _, v := range nums {
		f := float64(v)
		if f < lo {
			lo = f
		}
		if f > hi {
			hi = f
		}
	}
	// 二分答案：平均数 >= mid 是否可行具有单调性
	for hi-lo > 1e-6 {
		mid := (lo + hi) / 2
		if canAchieve(nums, k, mid) {
			lo = mid
		} else {
			hi = mid
		}
	}
	return lo
}

// canAchieve 判定是否存在长度 >= k 的子数组，其平均数 >= avg
// 等价于：把每个元素减去 avg 后，是否存在长度 >= k 且元素和 >= 0 的子数组
func canAchieve(nums []int, k int, avg float64) bool {
	pre := 0.0    // 当前位置的前缀和（sum(nums[0..j]) - avg 累计）
	minPre := 0.0 // 滞后 k 位之前（含）的最小前缀和
	lag := 0.0    // 滞后 k 位的前缀和，用于更新 minPre
	for j := 0; j < len(nums); j++ {
		pre += float64(nums[j]) - avg
		// 当 j >= k 时，把 P[j-k+1] 纳入 minPre 的候选
		if j >= k {
			lag += float64(nums[j-k]) - avg
			if lag < minPre {
				minPre = lag
			}
		}
		// 子数组长度至少为 k：pre - minPre 即某个长度 >= k 的子数组和
		if j >= k-1 && pre-minPre >= 0 {
			return true
		}
	}
	return false
}
