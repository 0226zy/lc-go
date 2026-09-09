package bestsightseeingpair

// MaxScoreSightseeingPair 最佳观光组合（标准 DP 数组版）
// 得分 values[i] + values[j] + i - j 可拆为 (values[i]+i) + (values[j]-j)。
// dp[j] 表示下标不超过 j 的位置中 values[i]+i 的最大值；
// 枚举每个 j 作为右端点，得分 = dp[j-1] + values[j] - j，取全局最大。
// 时间复杂度: O(n)  空间复杂度: O(n)
func MaxScoreSightseeingPair(values []int) int {
	n := len(values)
	dp := make([]int, n)
	dp[0] = values[0] // base case：候选左端点只有下标 0
	ans := 0
	for j := 1; j < n; j++ {
		// j 作为右端点：左端点取 0..j-1 中 values[i]+i 的最大值
		if score := dp[j-1] + values[j] - j; score > ans {
			ans = score
		}
		// 更新前缀最优值，供后续右端点使用
		if v := values[j] + j; v > dp[j-1] {
			dp[j] = v
		} else {
			dp[j] = dp[j-1]
		}
	}
	return ans
}

// MaxScoreSightseeingPairOptimized 最佳观光组合（滚动变量空间优化版）
// dp[j] 只依赖 dp[j-1]，用一个滚动变量 best 代替 dp 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxScoreSightseeingPairOptimized(values []int) int {
	ans := 0
	best := values[0] // 当前 j 左边 values[i]+i 的最大值，即 dp[j-1]
	for j := 1; j < len(values); j++ {
		if score := best + values[j] - j; score > ans {
			ans = score
		}
		if v := values[j] + j; v > best {
			best = v
		}
	}
	return ans
}
