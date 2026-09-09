# 0072. 编辑距离 (Edit Distance)

## 题目描述

给你两个单词 `word1` 和 `word2`，请返回将 `word1` 转换成 `word2` 所使用的最少操作数。

你可以对一个单词进行如下三种操作：

- 插入一个字符
- 删除一个字符
- 替换一个字符

### 示例

```
输入: word1 = "horse", word2 = "ros"
输出: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
```

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

### 提示

- `0 <= word1.length, word2.length <= 500`
- `word1` 和 `word2` 由小写英文字母组成

## 四步拆解

本题是最经典的**二维动态规划（双序列 DP）**问题，也是「两个字符串匹配类」DP 模板（编辑距离、最长公共子序列、不同的子序列等）的母题。

### 1. 状态定义

- `dp[i][j]` 表示**把 `word1` 的前 `i` 个字符转换成 `word2` 的前 `j` 个字符所需的最少操作数**（下标表示前缀长度，不是字符下标）。
- 最终答案：`dp[m][n]`（m、n 分别是两个单词的长度），即两个完整单词之间的编辑距离。

为什么用「前缀长度」定义？因为三种操作（增、删、换）都作用在单词末尾时，问题恰好可以拆成规模更小的前缀问题，二维数组就能描述全部状态。

### 2. 边界条件 + 遍历顺序

- base case：
  - `dp[i][0] = i`：`word2` 为空，把 `word1` 前 i 个字符全部删掉，需要 i 次删除；
  - `dp[0][j] = j`：`word1` 为空，要变出 `word2` 前 j 个字符，需要 j 次插入。
- 遍历顺序：**i 从 1 到 m、j 从 1 到 n 正序双重循环**。`dp[i][j]` 依赖 `dp[i-1][j]`（正上方）、`dp[i][j-1]`（左侧）、`dp[i-1][j-1]`（左上），正序遍历时这三个子问题都已求好。

### 3. 状态转移方程

比较两个单词的末尾字符 `word1[i-1]` 和 `word2[j-1]`：

```
如果 word1[i-1] == word2[j-1]:
    dp[i][j] = dp[i-1][j-1]                        // 末尾相同，不需要任何操作
否则:
    dp[i][j] = 1 + min(
        dp[i-1][j],     // 删除：删掉 word1[i-1]，再转化前 i-1 个字符
        dp[i][j-1],     // 插入：在 word1 末尾插入 word2[j-1]，等价于先转化出前 j-1 个
        dp[i-1][j-1],   // 替换：把 word1[i-1] 改成 word2[j-1]，再转化前 i-1 / j-1 个
    )
```

推导：考虑最后一对字符。若相同，这对接力「免费」，答案等于两个前缀的编辑距离。若不同，最后一步必然是三种操作之一——删除、插入或替换，各花费 1 次操作，剩下的就是规模更小的子问题。三种情况穷尽所有可能，取 min 再加 1。

边界特殊情况：某一单词为空时只剩纯删除 / 纯插入，已由 base case 覆盖。

### 4. 样例验证 + 代码实现

手动推导 `word1 = "horse"`, `word2 = "ros"` 的 dp 表（行是 word1 前缀，列是 word2 前缀）：

| dp | "" | r | o | s |
| --- | --- | --- | --- | --- |
| **""** | 0 | 1 | 2 | 3 |
| **h** | 1 | 1 | 2 | 3 |
| **o** | 2 | 2 | 1 | 2 |
| **r** | 3 | 2 | 2 | 2 |
| **s** | 4 | 3 | 3 | 2 |
| **e** | 5 | 4 | 4 | **3** |

- `dp[2][2]`（"ho"→"ro"）：末尾 'o' == 'o'，取 `dp[1][1] = 1`
- `dp[5][3]`（"horse"→"ros"）：'e' ≠ 's'，`1 + min(dp[4][3]=2, dp[5][2]=4, dp[4][2]=3) = 3` ✅（与示例一致）

代码与上面的推导一一对应：

```go
// MinDistance 编辑距离（标准 DP 数组版）
func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i // base case：word2 为空，i 次删除
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // base case：word1 为空，j 次插入
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
			}
		}
	}
	return dp[m][n]
}
```

## 空间优化版（滚动一维数组）

`dp[i][j]` 只依赖上一行（`dp[i-1][j]`、`dp[i-1][j-1]`）和本行左侧（`dp[i][j-1]`）三个值。用一维数组按行滚动：更新 `dp[j]` 之前它是「正上方」的值，再额外用一个变量 `prev` 记住「左上」的值（即更新前的 `dp[j-1]`），就能只开 O(min(m,n)) 的数组。让较短的单词做列，空间更省：

```go
// MinDistanceOptimized 编辑距离（滚动一维数组空间优化版）
func MinDistanceOptimized(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m < n { // 让 word2 是较短的那个，压缩空间
		word1, word2 = word2, word1
		m, n = n, m
	}
	dp := make([]int, n+1)
	for j := 0; j <= n; j++ {
		dp[j] = j // 第 0 行：空串转化需 j 次插入
	}
	for i := 1; i <= m; i++ {
		prev := dp[0] // prev 记录左上 dp[i-1][j-1]
		dp[0] = i
		for j := 1; j <= n; j++ {
			tmp := dp[j] // 暂存正上方 dp[i-1][j]，作为下一轮的左上
			if word1[i-1] == word2[j-1] {
				dp[j] = prev
			} else {
				dp[j] = 1 + min(dp[j], min(dp[j-1], prev))
			}
			prev = tmp
		}
	}
	return dp[n]
}
```

## 复杂度分析

| 版本 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- |
| DP 数组版 | O(m*n) | O(m*n) |
| 滚动一维数组版 | O(m*n) | O(min(m,n)) |
