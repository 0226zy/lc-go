# 212. 单词搜索 II (Word Search II)

## 题目描述

给定一个 `m x n` 二维字符网格 `board` 和一个单词列表 `words`，**返回所有在网格中出现的单词**。

单词必须按照字母顺序，通过相邻的单元格构成，其中"相邻"单元格是指水平相邻或垂直相邻的单元格。**同一个单元格内的字母在一个单词中不允许被重复使用**。

### 示例 1

```
输入: board = [["o","a","a","n"],
               ["e","t","a","e"],
               ["i","h","k","r"],
               ["i","f","l","v"]],
      words = ["oath","pea","eat","rain"]
输出: ["eat","oath"]
```

### 示例 2

```
输入: board = [["a","b"],["c","d"]], words = ["abcb"]
输出: []
```

## 提示

- `m == board.length`
- `n == board[i].length`
- `1 <= m, n <= 12`
- `1 <= words.length <= 3 * 10^4`
- `1 <= words[i].length <= 10`
- `board[i][j]` 和 `words[i]` 由小写英文字母组成
- `words` 中的所有字符串**互不相同**

## 题目解析

### 核心思路

这是 79 题「单词搜索」的加强版：单词从 1 个变成最多 3 万个。如果对每个单词各跑一次棋盘 DFS，复杂度是 O(单词数 × m × n × 4^L)，必然超时。

突破口：**把单词列表建一棵前缀树（Trie）**，然后从棋盘的每个格子出发做一次 DFS，但 DFS 的走向由 Trie 引导：

- 走到某个格子，如果 Trie 当前节点**没有**对应字母的子节点，说明任何单词都不会以当前路径为前缀，**立即剪枝**；
- 走到 Trie 中 `word` 非空的节点，说明找到了一个完整单词，收入答案。

这样所有单词在同一次棋盘遍历中被一次性找出，重叠的前缀也只遍历一次。本质上是**回溯（棋盘 DFS）+ 前缀树剪枝**的组合模板。

三个必须处理的细节：

1. **去重**：同一个单词可能在棋盘上找到多条路径，Trie 节点的 `word` 字段在首次命中后清空，天然去重；
2. **访问标记回溯**：进入格子标 `visited`，递归返回后撤销标记，保证"同一格不重复用"只对当前路径生效；
3. 提前特判空输入。

### 算法步骤

1. 把 `words` 全部插入 Trie（节点存 `children` 和 `word`，`word` 只在单词结尾处非空）。
2. 从棋盘每个格子 `(r, c)` 出发做 DFS，参数为当前 Trie 节点：
   - 越界或已访问，返回；
   - 取 `board[r][c]` 对应子节点，为空则剪枝返回；
   - 子节点 `word` 非空：答案加入该单词，并清空 `word`；
   - 标记 `visited[r][c] = true`，向四个方向递归；
   - 撤销标记（回溯）。
3. 返回收集到的单词列表。

### 复杂度分析

设 m、n 为棋盘尺寸，L 为最长单词长度：

- **时间复杂度**: O(m × n × 4^L)。每个起点最坏走 4^L 条路径，但 Trie 剪枝使实际复杂度远低于上界；建 Trie 为 O(单词总字符数)
- **空间复杂度**: O(单词总字符数 + m × n)，字典树与访问标记数组

## 代码实现

```go
// TrieNode 字典树节点：children[i] 对应字母 'a'+i，word 非空表示路径构成完整单词
type TrieNode struct {
	children [26]*TrieNode
	word     string
}

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

func FindWords(board [][]byte, words []string) []string {
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
		// 命中完整单词：记录并清空，避免重复收集
		if child.word != "" {
			result = append(result, child.word)
			child.word = ""
		}

		visited[r][c] = true
		dfs(r+1, c, child)
		dfs(r-1, c, child)
		dfs(r, c+1, child)
		dfs(r, c-1, child)
		visited[r][c] = false // 回溯：恢复访问标记
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			dfs(r, c, root)
		}
	}
	return result
}
```

**执行过程示例**（示例 1 中找 `"oath"`，路径为 o(0,0) → a(0,1) → t(1,1) → h(2,1)）：

```
从 (0,0) 的 'o' 出发，Trie 沿 o 分支:
  -> 邻居 (0,1)='a'，oa 分支存在，继续
       -> 邻居 (1,1)='t'，oat 分支存在，继续
            -> 邻居 (2,1)='h'，oath 分支存在且 word 非空，命中 "oath"，加入答案
visited 标记保证路径上 'o'-'a'-'t' 不会被同一路径二次使用；
每个方向的递归返回后逐格撤销标记，其他起点/路径仍可复用这些格子。
```
