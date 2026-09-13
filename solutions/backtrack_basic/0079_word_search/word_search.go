package wordsearch

// directions 上下左右四个方向的坐标偏移
var directions = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// Exist 单词搜索
// 给定 m x n 的字符网格 board 和单词 word，判断 word 是否能由网格中横纵相邻的格子按顺序组成（同一格子不可复用）。
// 时间复杂度: O(m * n * 4^L)  L 为单词长度  空间复杂度: O(L) 递归栈深度
func Exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	// 预剪枝：统计网格字符频次，若某字符数量不足以覆盖单词需求则直接失败
	if !enoughChars(board, word) {
		return false
	}

	// dfs 表示当前站在 (i, j)，正准备匹配单词的第 k 个字符
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		// 把当前格子改写成哨兵字符，相当于“已访问”标记
		ch := board[i][j]
		board[i][j] = '#'
		// 朝四个方向继续匹配下一个字符，任一方向成功即成功
		for _, d := range directions {
			if dfs(i+d[0], j+d[1], k+1) {
				board[i][j] = ch
				return true
			}
		}
		// 四个方向都失败，回溯：把格子还原，让别的路径还能使用
		board[i][j] = ch
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// enoughChars 检查网格中各字符的出现次数是否不少于单词所需
func enoughChars(board [][]byte, word string) bool {
	var boardCnt, wordCnt [128]int
	for _, row := range board {
		for _, c := range row {
			boardCnt[c]++
		}
	}
	for i := 0; i < len(word); i++ {
		wordCnt[word[i]]++
	}
	for c := 'A'; c <= 'z'; c++ {
		if wordCnt[c] > boardCnt[c] {
			return false
		}
	}
	return true
}
