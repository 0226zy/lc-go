package jumpgameii

// Jump 跳跃游戏 II
// 给定一个非负整数数组 nums，你最初位于数组的第一个下标。数组中的每个元素
// 表示你在该位置可以跳跃的最大长度。返回到达最后一个下标所需的最少跳跃次数。
// 题目保证你总是可以到达最后一个下标。
// 时间复杂度: O(n) 单次遍历  空间复杂度: O(1) 常数空间
func Jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0 // 已在终点，无需跳跃
	}
	jumps := 0    // 已使用的跳跃次数
	currEnd := 0  // 当前这一跳能覆盖的最远下标
	maxReach := 0 // 下一跳能覆盖的最远下标
	for i := 0; i < n-1; i++ {
		if reach := i + nums[i]; reach > maxReach {
			maxReach = reach // 更新下一跳的最远可达位置
		}
		if i == currEnd {
			// 走到了当前这一跳覆盖范围的边界，必须使用下一跳
			jumps++
			currEnd = maxReach
		}
	}
	return jumps
}
