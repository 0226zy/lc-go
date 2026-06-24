package longestcommonsubsequence

// LongestCommonSubsequence 最长公共子序列
// 返回两个字符串的最长公共子序列的长度。
// 时间复杂度: O(m*n)  空间复杂度: O(n)
func LongestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	f := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		f[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				f[i][j] = f[i-1][j-1] + 1
			} else {
				f[i][j] = max(f[i-1][j], f[i][j-1])
			}
		}
	}
	return f[m][n]
}
