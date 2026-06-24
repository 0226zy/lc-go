package editdistance

// MinDistance 编辑距离
// 返回将 word1 转换成 word2 所使用的最少操作数。
// 时间复杂度: O(m*n)  空间复杂度: O(n)
func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	f := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		f[i] = make([]int, n+1)
	}
	for i := 0; i < m+1; i++ {
		f[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		f[0][j] = j
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				f[i][j] = f[i-1][j-1]
				continue
			}
			f[i][j] = min(f[i-1][j-1], f[i-1][j], f[i][j-1]) + 1
		}
	}
	return f[m][n]
}
