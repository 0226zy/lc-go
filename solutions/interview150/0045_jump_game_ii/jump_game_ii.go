package jumpgameii

import "math"

// Jump 跳跃游戏 II（标准 DP 数组版）
// 给定一个非负整数数组 nums，你最初位于数组的第一个下标，每个元素表示在该位置
// 可以跳跃的最大长度。返回到达最后一个下标所需的最少跳跃次数（题目保证可达）。
// dp[i] 表示从下标 0 跳到下标 i 的最少跳跃次数。
// dp[i] = min(dp[j] + 1)，其中 j < i 且 j + nums[j] >= i。
// 时间复杂度: O(n²)  空间复杂度: O(n)
func Jump(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32 // 初始化为极大值，表示尚未到达
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if j+nums[j] >= i { // 前驱 j 能一步跳到 i
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

// JumpAlternative 跳跃游戏 II（贪心版，BFS 分层思想）
// 维护当前跳能到达的边界 end 和下一跳的最远位置 farthest，
// 每走到边界就必须再跳一次，把边界推进到 farthest。
// 时间复杂度: O(n)  空间复杂度: O(1)
func JumpAlternative(nums []int) int {
	jumps, end, farthest := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i]) // 更新下一跳的最远可达位置
		if i == end {                       // 走完当前跳覆盖的边界，必须再跳一次
			jumps++
			end = farthest
		}
	}
	return jumps
}
