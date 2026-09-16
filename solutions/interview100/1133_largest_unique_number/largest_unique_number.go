package largestuniquenumber

// LargestUniqueNumber 最大唯一数
// 返回 nums 中只出现一次的整数里的最大值，若不存在则返回 -1。
// 时间复杂度: O(n + U)，U = 1001 为值域大小  空间复杂度: O(U)
func LargestUniqueNumber(nums []int) int {
	// 值域为 [0, 1000]，用计数数组统计出现次数
	var cnt [1001]int
	for _, v := range nums {
		cnt[v]++
	}
	// 从大到小扫描，第一个恰好出现一次的值即为答案
	for v := 1000; v >= 0; v-- {
		if cnt[v] == 1 {
			return v
		}
	}
	return -1
}
