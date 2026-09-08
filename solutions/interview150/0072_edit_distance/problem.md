# 72. 编辑距离 (Edit Distance)

## 题目描述

给你两个单词 `word1` 和 `word2`，请返回将 `word1` 转换成 `word2` 所使用的最少操作数。

你可以对一个单词进行如下三种操作：

- 插入一个字符
- 删除一个字符
- 替换一个字符

### 示例 1

```
输入: word1 = "horse", word2 = "ros"
输出: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
```

### 示例 2

```
输入: word1 = "intention", word2 = "execution"
输出: 5
解释:
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
```

## 提示

- `0 <= word1.length, word2.length <= 500`
- `word1` 和 `word2` 仅由小写英文字母组成

## 题目解析

### 核心思路

本题是最经典的**二维动态规划（双序列 DP）**问题。定义状态：

- `dp[i][j]` 表示把 `word1` 的前 `i` 个字符转换成 `word2` 的前 `j` 个字符所需的最少操作数。

**边界（初始化）**：

- `dp[0][j] = j`：`word1` 为空串时，只能通过 `j` 次**插入**得到 `word2[:j]`。
- `dp[i][0] = i`：`word2` 为空串时，只能通过 `i` 次**删除**把 `word1[:i]` 变空。

**状态转移**，考察两个单词的末尾字符 `word1[i-1]` 与 `word2[j-1]`：

- **末尾字符相同**（`word1[i-1] == word2[j-1]`）：这一对字符不需要任何操作，问题规模直接缩小一格：
  `dp[i][j] = dp[i-1][j-1]`
- **末尾字符不同**：必须做一次操作，取三种操作中代价最小的再加 1：
  - **删除** `word1[i-1]`：剩下的 `word1[:i-1]` 去匹配 `word2[:j]` → `dp[i-1][j]`
  - **插入** `word2[j-1]` 到 `word1` 末尾：插入后末尾已经匹配，等价于 `word1[:i]` 匹配 `word2[:j-1]` → `dp[i][j-1]`
  - **替换** `word1[i-1]` 为 `word2[j-1]`：替换后末尾匹配，问题缩小一格 → `dp[i-1][j-1]`

  即 `dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])`。

**空间优化（滚动数组）**：转移只依赖上一行（`dp[i-1][j]`、`dp[i-1][j-1]`）和本行左侧（`dp[i][j-1]`），因此用一维数组 `dp[j]` 滚动更新即可。遍历时用一个变量 `prev` 暂存左上角的 `dp[i-1][j-1]`（它会在更新 `dp[j]` 时被覆盖）。

最终答案为 `dp[m][n]`。

### 算法步骤

1. 令 `m = len(word1)`，`n = len(word2)`，创建长度为 `n+1` 的一维数组 `dp`。
2. 初始化 `dp[j] = j`（对应 `word1` 为空串，全插入）。
3. 外层循环 `i` 从 1 到 `m`：
   - `prev = dp[0]`（即 `dp[i-1][0]`），然后把 `dp[0]` 更新为 `i`（即 `dp[i][0]`，全删除）。
   - 内层循环 `j` 从 1 到 `n`：
     - 暂存 `old = dp[j]`（即 `dp[i-1][j]`）。
     - 若 `word1[i-1] == word2[j-1]`，则 `dp[j] = prev`；
     - 否则 `dp[j] = 1 + min(old, dp[j-1], prev)`（分别对应删、插、换）。
     - `prev = old`，为下一轮提供左上角值。
4. 返回 `dp[n]`。

### 复杂度分析

- **时间复杂度**: O(m×n)，两层循环遍历整个 DP 表，每次转移是 O(1)。
- **空间复杂度**: O(n)，使用滚动数组只保留一行状态（原始二维 DP 为 O(m×n)）。

## 代码实现

```go
func MinDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)

    // dp[j] 表示当前已处理的 word1 前缀到 word2[:j] 的编辑距离
    // 初始状态对应 word1 为空串：把空串变成 word2[:j] 需要插入 j 次
    dp := make([]int, n+1)
    for j := 0; j <= n; j++ {
        dp[j] = j
    }

    for i := 1; i <= m; i++ {
        prev := dp[0] // prev 记录 dp[i-1][j-1]
        dp[0] = i     // dp[i][0]：word1[:i] 变成空串需要删除 i 次
        for j := 1; j <= n; j++ {
            old := dp[j] // 暂存 dp[i-1][j]，供下一轮作为 dp[i-1][j-1]
            if word1[i-1] == word2[j-1] {
                dp[j] = prev
            } else {
                // 1 + min(删除 dp[i-1][j], 插入 dp[i][j-1], 替换 dp[i-1][j-1])
                dp[j] = 1 + min(old, min(dp[j-1], prev))
            }
            prev = old
        }
    }

    return dp[n]
}
```

**执行过程示例**（`word1 = "horse"`, `word2 = "ros"`）的二维 DP 表：

```
        ""  r  o  s
   ""    0  1  2  3
   h     1  1  2  3
   o     2  2  1  2
   r     3  2  2  2
   s     4  3  3  2
   e     5  4  4  3
```

右下角 `dp[5][3] = 3`，即 "horse" → "ros" 的最少操作数为 3。
