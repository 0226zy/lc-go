# 51. N 皇后 (N-Queens)

## 题目描述

按照国际象棋的规则，皇后可以攻击与之处在**同一行、同一列或同一斜线**上的棋子。

**n 皇后问题**研究的是如何将 `n` 个皇后放置在 `n×n` 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 `n`，返回所有不同的 n 皇后问题的解决方案。

每一种解法包含一个不同的 n 皇后问题的棋子放置方案，该方案中 `'Q'` 和 `'.'` 分别代表了皇后和空位。

### 示例 1

```
输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：4 皇后问题存在两个不同的解法。
```

### 示例 2

```
输入：n = 1
输出：[["Q"]]
```

## 提示

- `1 <= n <= 9`

## 题目解析

### 核心思路

经典的**回溯**问题。关键观察：每行恰好放一个皇后（放两个必然同行冲突，放不满 n 行则凑不齐 n 个皇后），因此可以**逐行决策**——第 `row` 层递归只决定第 `row` 行的皇后放在哪一列，天然避免了同行冲突，只需检测列与两条对角线。

判断对角线冲突的技巧：

- 同一条**主对角线**（左上—右下）上的格子，`行 - 列` 的值恒定；
- 同一条**副对角线**（右上—左下）上的格子，`行 + 列` 的值恒定。

用三个集合（`cols`、`diag1`、`diag2`）分别记录已占用的列、`r-c`、`r+c`，放置前 O(1) 即可判断冲突，无需像朴素做法那样逐格扫描。

回溯三要素：

- **路径**：`queens[r]` 记录第 `r` 行皇后所在列（已做出的选择）。
- **选择列表**：当前行 `row` 的 n 个列中，不与已有皇后冲突（不在 `cols`/`diag1`/`diag2` 中）的列。
- **结束条件**：`row == n`，即 n 行全部放置完毕，把 `queens` 转成棋盘字符串收集进结果。

### 算法步骤

1. 初始化 `queens` 数组（长度 n，初值 -1）和三个占用集合。
2. 从第 0 行开始回溯 `backtrack(row)`：
   - 若 `row == n`，根据 `queens` 生成棋盘字符串快照，加入结果，返回；
   - 否则遍历当前行的每一列 `col`：
     - 若 `col`、`row-col`、`row+col` 任一已被占用，跳过（剪枝）；
     - 做选择：记录 `queens[row] = col`，三个集合打标记；
     - 递归 `backtrack(row + 1)`；
     - 撤销选择：恢复 `queens[row]`，删除三个集合中的标记。
3. 返回结果。

### 复杂度分析

- **时间复杂度**：O(n!)。第一行有 n 种选择，第二行至多 n-1 种……最坏情况上界为 n!，对角线剪枝会进一步减少实际搜索量。
- **空间复杂度**：O(n)。递归栈深度为 n，`queens` 数组与三个集合均为 O(n)（不计输出占用）。

## 代码实现

```go
package nqueens

// SolveNQueens N 皇后
// 将 n 个皇后放置在 n×n 的棋盘上，使皇后彼此不能互相攻击，返回所有不同的放置方案。
// 时间复杂度: O(n!)  空间复杂度: O(n)
func SolveNQueens(n int) [][]string {
	queens := make([]int, n) // queens[r] 表示第 r 行皇后所在的列
	for i := range queens {
		queens[i] = -1
	}
	cols := make(map[int]bool)  // 已被占用的列
	diag1 := make(map[int]bool) // 主对角线方向：r - c 恒定
	diag2 := make(map[int]bool) // 副对角线方向：r + c 恒定
	result := make([][]string, 0)

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			board := make([]string, n)
			for r := 0; r < n; r++ {
				line := make([]byte, n)
				for c := 0; c < n; c++ {
					line[c] = '.'
				}
				line[queens[r]] = 'Q'
				board[r] = string(line)
			}
			result = append(result, board)
			return
		}
		for col := 0; col < n; col++ {
			d1, d2 := row-col, row+col
			if cols[col] || diag1[d1] || diag2[d2] {
				continue // 列或对角线冲突，剪枝
			}
			queens[row] = col
			cols[col], diag1[d1], diag2[d2] = true, true, true
			backtrack(row + 1)
			queens[row] = -1
			delete(cols, col)
			delete(diag1, d1)
			delete(diag2, d2)
		}
	}
	backtrack(0)
	return result
}
```

**执行过程示例**（`n = 4`）：

```
backtrack(0)：第 0 行
  试 col=0 → queens=[0,-,-,-]
    backtrack(1)：col=0 列冲突，col=1 对角线冲突
      试 col=2 → queens=[0,2,-,-]
        backtrack(2)：col=0 对角线冲突，col=1 对角线冲突，
                      col=2 列冲突，col=3 对角线冲突 → 全灭，回溯
      试 col=3 → queens=[0,3,-,-]
        backtrack(2)：仅 col=1 可放 → queens=[0,3,1,-]
          backtrack(3)：col=0 列冲突，col=1 列冲突，
                        col=2 对角线冲突，col=3 列冲突 → 全灭，回溯
      第 0 行放 col=0 无解，撤销
  试 col=1 → queens=[1,-,-,-]
    backtrack(1)：col=0、1、2 均冲突，仅 col=3 可放 → queens=[1,3,-,-]
      backtrack(2)：仅 col=0 可放 → queens=[1,3,0,-]
        backtrack(3)：仅 col=2 可放 → queens=[1,3,0,2]
          row==4，收集解 [".Q..","...Q","Q...","..Q."]
  试 col=2（与 col=1 对称）→ queens=[2,0,3,1]
          收集解 ["..Q.","Q...","...Q",".Q.."]
  试 col=3（与 col=0 对称）→ 无解
结果：共 2 个解
```
