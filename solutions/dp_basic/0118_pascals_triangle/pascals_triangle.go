package pascalstriangle

// Generate 杨辉三角（标准 DP 数组版）
// 生成杨辉三角的前 numRows 行，每个数等于它左上方和右上方的数之和。
// dp[i][j] 表示第 i 行第 j 个数，dp[i][j] = dp[i-1][j-1] + dp[i-1][j]，每行首尾固定为 1。
// 时间复杂度: O(n²)  空间复杂度: O(n²)（输出本身的大小）
func Generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)
		dp[i][0], dp[i][i] = 1, 1 // base case：每行首尾固定为 1
		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j] // 左上方 + 右上方
		}
	}
	return dp
}
