package handshakesthatdontcross

// NumberOfWays 不相交的握手
// numPeople 个人围成一圈，两两握手且握手连线互不交叉，返回方案总数对 10^9+7 取余的结果。
// 答案即第 numPeople/2 个卡特兰数，用动态规划递推。
// 时间复杂度: O(n^2)  空间复杂度: O(n)
func NumberOfWays(numPeople int) int {
	const mod = 1_000_000_007

	// 奇数个人无法两两配对
	if numPeople%2 == 1 {
		return 0
	}

	// dp[i] 表示 i 个人围成圈时的不相交握手方案数
	dp := make([]int, numPeople+1)
	dp[0] = 1
	for i := 2; i <= numPeople; i += 2 {
		// 固定 1 号人与第 j+2 号人握手，把圆分成 j 人和 i-2-j 人两个独立区域
		for j := 0; j <= i-2; j += 2 {
			dp[i] = (dp[i] + dp[j]*dp[i-2-j]) % mod
		}
	}
	return dp[numPeople]
}
