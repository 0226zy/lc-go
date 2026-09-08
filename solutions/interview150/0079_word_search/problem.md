# 79. 单词搜索 (Word Search)

## 题目描述

给定一个 `m x n` 二维字符网格 `board` 和一个字符串单词 `word`。如果 `word` 存在于网格中，返回 `true`；否则，返回 `false`。

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

这是 **网格回溯（ Flood Fill 式 DFS）** 的标准模板：把单词匹配看成在网格上“走路”——从某个格子出发，每一步走到上下左右相邻且字符匹配的格子，看能不能恰好走出整个单词。

模板要点有三：

1. **入口枚举**：单词可以从任意格子开始，所以对每个格子都尝试做一次 DFS（首字符不匹配会立刻返回）。
2. **访问标记**：同一个格子不能重复用。最省事的办法是把走过的格子临时改成哨兵值（如 `0`，题目保证只会出现英文字母），递归返回后再 **恢复原值**——这就是“标记后还原”。
3. **找到即返回**：任意一条路径匹配成功就算成功，所以每层递归一旦发现下层返回 true，就立刻向上返回 true，不再尝试其他分支。

回溯三要素：

- **路径**：已匹配到的字符数 `k`（坐标 i, j 随递归参数传递）。
- **选择列表**：上下左右四个相邻格子中字符等于 `word[k+1]` 的格子。
- **结束条件**：`k == len(word)-1` 且字符匹配，返回 true。

### 算法步骤

1. 遍历网格每个格子 `(i, j)`：
   - 从 `(i, j, 0)` 开始 DFS。
2. DFS `(i, j, k)`：
   - 若 `board[i][j] != word[k]`，返回 false；
   - 若 `k` 已是单词末位，返回 true；
   - 把 `board[i][j]` 暂存并置为哨兵值 `0`（标记访问）；
   - 依次尝试四个方向，越界的跳过，递归 `(ni, nj, k+1)`；
   - 若任一方向返回 true，先还原当前格再返回 true；
   - 四个方向都不行，还原当前格，返回 false。
3. 所有起点都失败则返回 false。

### 复杂度分析

- **时间复杂度**: O(m × n × 4ᴸ)，L 为单词长度。每个起点最多展开一棵深度为 L、分支为 4 的搜索树（剪枝后实际远小于此）。
- **空间复杂度**: O(L)，递归栈深度最大为 L。

## 代码实现

```go
// dirs 四个移动方向：上、下、左、右
var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func Exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])

    var backtrack func(i, j, k int) bool
    backtrack = func(i, j, k int) bool {
        if board[i][j] != word[k] {
            return false
        }
        if k == len(word)-1 {
            return true // 最后一个字符也匹配，找到单词
        }
        // 标记为已访问（0 不会出现在棋盘里），避免同一路径重复使用
        temp := board[i][j]
        board[i][j] = 0
        for _, d := range dirs {
            ni, nj := i+d[0], j+d[1]
            if ni >= 0 && ni < m && nj >= 0 && nj < n && backtrack(ni, nj, k+1) {
                board[i][j] = temp // 还原现场
                return true
            }
        }
        board[i][j] = temp // 回溯：恢复原字符
        return false
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if backtrack(i, j, 0) {
                return true
            }
        }
    }
    return false
}
```

**执行过程示例**（`board` 如上，`word = "ABCCED"`）：

```
起点(0,0) 'A' 匹配 word[0] → 标记访问:
  (0,1) 'B' 匹配 word[1] → 标记:
    (0,2) 'C' 匹配 word[2] → 标记:
      (1,2) 'C' 匹配 word[3] → 标记:
        (2,2) 'E' 匹配 word[4] → 标记:
          (2,1) 'D' 匹配 word[5]（末位）→ 返回 true
一路向上返回 true
```
