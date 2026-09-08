package wordsearchii

// TrieNode 单词搜索 II 用的字典树节点
// children[i] 对应字母 'a'+i 的子节点；word 非空表示从根到该节点的路径构成一个完整单词
type TrieNode struct {
	children [26]*TrieNode
	word     string
}

// insert 向字典树插入单词
func (root *TrieNode) insert(word string) {
	node := root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.word = word
}

// FindWords 单词搜索 II
// 在棋盘上找出所有出现在 words 中的单词：单词由上下左右相邻格子依次构成，同一格不能重复使用。
// 算法：先把 words 建字典树，再从每个格子出发做 DFS，走到字典树不存在的分支就剪枝。
// 时间复杂度: O(m*n*4^L) L 为最长单词长度  空间复杂度: O(单词总字符数 + m*n) 字典树 + 访问标记
func FindWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return nil
	}

	root := &TrieNode{}
	for _, w := range words {
		root.insert(w)
	}

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var result []string
	var dfs func(r, c int, node *TrieNode)
	dfs = func(r, c int, node *TrieNode) {
		// 越界、已访问、字典树中无对应分支，均剪枝
		if r < 0 || r >= m || c < 0 || c >= n || visited[r][c] {
			return
		}
		child := node.children[board[r][c]-'a']
		if child == nil {
			return
		}
		// 命中完整单词：记录下来并清空，避免重复收集同一单词
		if child.word != "" {
			result = append(result, child.word)
			child.word = ""
		}

		visited[r][c] = true
		dfs(r+1, c, child)
		dfs(r-1, c, child)
		dfs(r, c+1, child)
		dfs(r, c-1, child)
		visited[r][c] = false // 回溯：恢复访问标记，供其他路径使用
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			dfs(r, c, root)
		}
	}
	return result
}
