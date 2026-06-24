package jumpgame

// CanJump 跳跃游戏
// 判断是否能够到达数组的最后一个下标。
// 时间复杂度: O(n)  空间复杂度: O(1)
func CanJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	jump := 0
	for i := 0; i <= jump; i++ {
		jump = max(nums[i]+i, jump)
		if jump >= len(nums)-1 {
			return true
		}
	}
	return false
}
