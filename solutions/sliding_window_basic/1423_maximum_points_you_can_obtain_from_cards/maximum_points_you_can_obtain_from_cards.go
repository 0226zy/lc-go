package maximumpointsyoucanobtainfromcards

// MaxScore 可获得的最大点数
// 给定整数数组 cardPoints 和整数 k，每次只能从数组开头或末尾拿一张牌，
// 恰好拿走 k 张牌，返回能获得的最大点数之和。
// 等价转化：拿走 k 张牌后，剩下的必为长度为 n-k 的连续子数组，
// 因此答案 = 总和 - 长度为 n-k 的连续子数组的最小和（固定长度滑动窗口）。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	if n == 0 || k <= 0 {
		return 0
	}

	// 数组总点数
	total := 0
	for _, p := range cardPoints {
		total += p
	}

	windowLen := n - k
	// k == n 时必须拿走所有牌，直接返回总和
	if windowLen == 0 {
		return total
	}

	// 初始窗口：前 windowLen 张牌的点数和
	sum := 0
	for i := 0; i < windowLen; i++ {
		sum += cardPoints[i]
	}
	minSum := sum

	// 窗口向右滑动：移出左端元素、移入右端元素，维护最小窗口和
	for right := windowLen; right < n; right++ {
		sum += cardPoints[right] - cardPoints[right-windowLen]
		if sum < minSum {
			minSum = sum
		}
	}

	// 最大拿牌点数 = 总和 - 最小剩余窗口和
	return total - minSum
}
