package decodeways

// NumDecodings 解码方法（标准 DP 数组版）
// 数字字符串按 '1'->'A' ... '26'->'Z' 映射，求解码方法总数。
// dp[i] 表示前 i 个字符的解码方法数：末尾数字单独解码加 dp[i-1]，
// 末两位在 10~26 之间可合并解码再加 dp[i-2]。
// 时间复杂度: O(n)  空间复杂度: O(n)
func NumDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1 // base case：空串视为一种解码，作为转移基石
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1] // 最后一个数字单独解码
		}
		if i >= 2 {
			two := int(s[i-2]-'0')*10 + int(s[i-1]-'0')
			if two >= 10 && two <= 26 {
				dp[i] += dp[i-2] // 最后两个数字合并解码
			}
		}
	}
	return dp[n]
}

// NumDecodingsOptimized 解码方法（滚动变量空间优化版）
// dp[i] 只依赖 dp[i-1]、dp[i-2]，用两个滚动变量代替 dp 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func NumDecodingsOptimized(s string) int {
	n := len(s)
	if s[0] == '0' {
		return 0 // 前导零无法解码
	}
	prev2, prev1 := 1, 1 // 分别代表 dp[0]、dp[1]（首字符非 '0' 时 dp[1]=1）
	for i := 2; i <= n; i++ {
		cur := 0
		if s[i-1] != '0' {
			cur += prev1 // 最后一个数字单独解码
		}
		two := int(s[i-2]-'0')*10 + int(s[i-1]-'0')
		if two >= 10 && two <= 26 {
			cur += prev2 // 最后两个数字合并解码
		}
		prev2, prev1 = prev1, cur
	}
	return prev1
}
