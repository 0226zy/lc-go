# 79. 单词搜索 (Word Search)

## 题目描述

给定一个 `m x n` 的二维字符网格 `board` 和一个字符串单词 `word`。如果 `word` 存在于网格中，返回 `true`；否则，返回 `false`。

单词必须按照字母顺序，通过相邻的单元格构成，其中 **相邻** 单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母 **不允许被重复使用**。

### 示例 1

```
输入: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出: true
```

### 示例 2

```
输入: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出: true
```

### 示例 3

```
输入: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出: false
```

## 提示

- `m == board.length`
- `n = board[i].length`
- `1 <= m, n <= 6`
- `1 <= word.length <= 15`
- `board` 和 `word` 仅由大小写英文字母组成

## 题目解析

### 核心思路

这是一道经典的 **网格回溯** 题。把单词匹配想象成在网格上“走格子”：选一个起点，依次走到上下左右相邻且字符匹配的格子，看能否恰好走完整个单词。走不通就退回上一步换个方向，这就是回溯。

回溯三要素：

- **路径**：已经匹配的字符个数 `k`（当前坐标 `i, j` 随递归参数传递，隐含在路径中）。
- **选择列表**：当前格子上下左右四个相邻格子中，字符等于 `word[k]` 的格子。
- **结束条件**：`k == len(word)-1` 且当前格字符匹配，说明单词全部匹配成功，返回 true。

DFS 模板要点与标记/还原策略：

1. **入口枚举**：单词可以从任意格子开始，因此对每个格子都发起一次 DFS。
2. **越界与失配剪枝**：把“越界或字符不匹配则返回 false”放在递归函数开头，调用方就可以放心遍历四个方向。
3. **标记与还原**：进入格子后先把它改写成哨兵字符（如 `'#'`，题目保证只会出现英文字母），防止同一条路径重复使用；离开前无论成败都恢复原值。也可以额外开一个 `visited` 数组，代价是 O(m×n) 额外空间，原位标记更省。
4. **预剪枝（可选优化）**：先统计网格与单词的字符频次，若单词里某字符的需求数超过网格中的总量，直接返回 false，能省掉大量无效搜索。

### 算法步骤

1. 统计 `board` 与 `word` 的字符频次，做预剪枝。
2. 枚举每个格子 `(i, j)` 作为起点，调用 `dfs(i, j, 0)`。
3. `dfs(i, j, k)` 的逻辑：
   - 越界或 `board[i][j] != word[k]`，返回 false；
   - `k` 到达单词末位，返回 true；
   - 暂存当前字符并把格子置为哨兵值；
   - 依次向上下左右递归 `dfs(ni, nj, k+1)`，任一方向返回 true 则还原当前格并返回 true；
   - 四个方向都失败，还原当前格，返回 false。
4. 所有起点均失败则返回 false。

### 复杂度分析

- **时间复杂度**：O(m × n × 4ᴸ)，L 为单词长度。每个起点最坏展开一棵深度 L、分支因子 4 的搜索树，实际因剪枝远小于该上界。
- **空间复杂度**：O(L)，即递归栈的最大深度。

## 代码实现

```go
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
```

**执行过程示例**（`board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]`，`word = "SEE"`）：

```
预剪枝通过（S、E 数量均足够）。
起点 (0,0) 'A' != 'S'，直接失败；(0,1) 'B'、 (0,2) 'C' 同理。
起点 (0,3) 'E' != 'S'，失败。
起点 (1,0) 'S' 匹配 word[0] → 标记为 '#':
  向上 (0,0) 'A' != word[1]('E')，失败
  向下 (2,0) 'A' != 'E'，失败
  向右 (1,1) 'F' != 'E'，失败
  四个方向全失败 → 还原 (1,0) 为 'S'，返回 false
起点 (1,3) 'S' 匹配 word[0] → 标记为 '#':
  向上 (0,3) 'E' 匹配 word[1] → 标记为 '#':
    向下回到 (1,3) 是 '#'，不匹配
    向右越界
    向左 (0,2) 'C' != word[2]('E')，失败
    → 还原 (0,3) 为 'E'，返回 false
  向下 (2,3) 'E' 匹配 word[1] → 标记为 '#':
    向左 (2,2) 'E' 匹配 word[2]，且已是单词末位 → 返回 true
  → 还原 (2,3)，向上返回 true
找到单词，Exist 返回 true
```
