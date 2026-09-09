package uglynumberii

// NthUglyNumber 丑数 II（三指针 DP）
// 丑数是质因子只含 2、3、5 的正整数（1 是第 1 个），求第 n 个丑数。
// dp[i] 表示第 i 个丑数，dp[i] = min(dp[p2]*2, dp[p3]*3, dp[p5]*5)，
// 三个指针记录下一个该被 2/3/5 乘的丑数下标，命中即前进（去重）。
// 时间复杂度: O(n)  空间复杂度: O(n)
func NthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1             // base case：1 是第 1 个丑数
	p2, p3, p5 := 1, 1, 1 // 分别指向下一个该乘 2、3、5 的丑数
	for i := 2; i <= n; i++ {
		n2, n3, n5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(n2, n3, n5)
		// 不用 else if：同一个值可能由多条链产生（如 6=2*3=3*2），都要跳过去重
		if dp[i] == n2 {
			p2++
		}
		if dp[i] == n3 {
			p3++
		}
		if dp[i] == n5 {
			p5++
		}
	}
	return dp[n]
}
