package jumpgame

// CanJump 跳跃游戏（标准 DP 数组版）
// dp[i] 表示下标 i 是否可达。
// dp[i] = 存在某个 j < i，使得 dp[j] 为 true 且 j + nums[j] >= i。
// 时间复杂度: O(n²)  空间复杂度: O(n)
func CanJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[0] = true // base case：起点可达
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && j+nums[j] >= i { // 前驱 j 可达且能一步跳到 i
				dp[i] = true
				break // 找到一个前驱即可
			}
		}
	}
	return dp[n-1]
}

// CanJumpAlternative 跳跃游戏（贪心版）
// 维护最远可达下标 farthest，遍历时不断更新；
// 若某位置超出 farthest 则不可达，farthest 覆盖终点则可达。
// 时间复杂度: O(n)  空间复杂度: O(1)
func CanJumpAlternative(nums []int) bool {
	farthest := 0
	for i := 0; i < len(nums); i++ {
		if i > farthest { // 当前位置本身不可达，后面更不可能
			return false
		}
		farthest = max(farthest, i+nums[i])
		if farthest >= len(nums)-1 { // 已能覆盖终点
			return true
		}
	}
	return true
}
