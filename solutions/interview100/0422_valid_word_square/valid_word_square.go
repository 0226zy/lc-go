package validwordsquare

// ValidWordSquare 有效的单词方块
// 把单词数组按行排成字符矩阵，判断对每个 k，第 k 行与第 k 列是否完全相同。
// 各行长度可以不同，因此访问 words[j][i] 前必须检查下标是否越界。
// 时间复杂度: O(n*m) n 为单词数，m 为最长单词长度  空间复杂度: O(1)
func ValidWordSquare(words []string) bool {
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			// 第 j 行不存在，或第 j 行没有第 i 个字符，行列不对称
			if j >= len(words) || i >= len(words[j]) {
				return false
			}
			if words[i][j] != words[j][i] {
				return false
			}
		}
	}
	return true
}
