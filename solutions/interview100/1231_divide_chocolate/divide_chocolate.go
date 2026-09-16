package dividechocolate

// MaximizeSweetness 分享巧克力
// 将 sweetness 切成 k+1 个连续的非空块，自己拿走甜度之和最小的那块，
// 返回最佳切割策略下这块甜度的最大值。
// 时间复杂度: O(n log S)，S 为数组元素总和  空间复杂度: O(1)
func MaximizeSweetness(sweetness []int, k int) int {
	sum := 0
	for _, s := range sweetness {
		sum += s
	}

	// check(x): 能否切出至少 k+1 块、每块甜度之和都 >= x
	check := func(x int) bool {
		cnt, cur := 0, 0
		for _, s := range sweetness {
			cur += s
			// 当前块达到目标甜度，贪心切一刀，重新开始累加
			if cur >= x {
				cnt++
				cur = 0
			}
		}
		return cnt >= k+1
	}

	// 在 [lo, hi] 上二分满足 check 的最大甜度
	lo, hi := 1, sum
	for lo < hi {
		mid := (lo + hi + 1) / 2 // 向上取整，避免死循环
		if check(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}
