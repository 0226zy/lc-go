package jumpgame

// CanJump 跳跃游戏
// 给定一个非负整数数组 nums，你最初位于数组的第一个下标。数组中的每个元素
// 表示你在该位置可以跳跃的最大长度。判断你是否能够到达最后一个下标。
// 时间复杂度: O(n) 单次遍历  空间复杂度: O(1) 常数空间
func CanJump(nums []int) bool {
	maxReach := 0 // 当前能到达的最远下标
	for i := 0; i < len(nums); i++ {
		// 如果当前下标已经超过了能到达的最远位置，说明走不到这里，直接失败
		if i > maxReach {
			return false
		}
		// 从位置 i 出发最远可以到达 i + nums[i]，更新最远可达下标
		if reach := i + nums[i]; reach > maxReach {
			maxReach = reach
		}
		// 已经能到达终点（或更远），可以提前返回
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return true
}
