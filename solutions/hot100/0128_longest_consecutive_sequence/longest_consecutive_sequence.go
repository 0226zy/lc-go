package longestconsecutivesequence

// LongestConsecutive 最长连续序列
// 给定一个未排序的整数数组 nums，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 时间复杂度: O(n)  空间复杂度: O(n)
func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	maxLen := 0
	for num := range numSet {
		// 只有当 num-1 不在集合中时，num 才是一个连续序列的起点
		if _, ok := numSet[num-1]; ok {
			continue
		}

		currentNum := num
		currentLen := 1
		for {
			if _, ok := numSet[currentNum+1]; ok {
				currentNum++
				currentLen++
			} else {
				break
			}
		}

		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}
