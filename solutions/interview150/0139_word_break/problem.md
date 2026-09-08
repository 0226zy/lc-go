# 139. 单词拆分 (Word Break)

## 题目描述

给你一个字符串 `s` 和一个字符串列表 `wordDict` 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 `s` 则返回 `true` 。

**注意**：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。

### 示例 1

```
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
```

### 示例 2

```
输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
     注意，你可以重复使用字典中的单词。
```

### 示例 3

```
输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
```

## 提示

- `1 <= s.length <= 300`
- `1 <= wordDict.length <= 1000`
- `1 <= wordDict[i].length <= 20`
- `s` 和 `wordDict[i]` 仅由小写英文字母组成
- `wordDict` 中的所有字符串互不相同

## 题目解析

### 核心思路

这是一道典型的**一维动态规划**（完全背包思想的判定版）题目。

关键观察：字符串 `s` 能否被拆分，取决于**它的某个前缀能否被拆分**。这满足最优子结构和无后效性：

- 定义 `dp[i]` 表示 `s` 的前 `i` 个字符（即 `s[:i]`）能否用字典单词拼出。
- 边界：`dp[0] = true`，空串不需要任何单词，视为可以拆分。
- 转移：对 `i >= 1`，枚举最后一个单词的起点 `j`（`0 <= j < i`），只要 `dp[j] == true` 且 `s[j:i]` 在字典中，则 `dp[i] = true`。即：

  ```
  dp[i] = OR{ dp[j] && dict.contains(s[j:i]) } ，其中 0 <= j < i
  ```

- 答案：`dp[n]`（`n = len(s)`）。

为了加速子串查询，把 `wordDict` 放进哈希集合 `map[string]struct{}`，查询一次为 O(k)（k 为子串长度，用于哈希计算）。

一个小优化：字典单词最长为 `maxLen`，所以 `s[j:i]` 长度超过 `maxLen` 时必然不在字典里，`j` 只需从 `i - maxLen` 开始枚举（不小于 0），把内层枚举范围从 O(n) 降到 O(maxLen)。

本题属于**「一维线性 DP / 完全背包判定」**模板，核心是「枚举最后一段，查表验证」。

### 算法步骤

1. 将 `wordDict` 存入哈希集合 `dict`，并统计最长单词长度 `maxLen`。
2. 令 `n = len(s)`，创建 `dp := make([]bool, n+1)`，置 `dp[0] = true`。
3. 外层循环 `i` 从 `1` 到 `n`：
   - 计算 `start = max(i - maxLen, 0)`。
   - 内层循环 `j` 从 `start` 到 `i-1`：若 `dp[j]` 为真且 `s[j:i]` 在 `dict` 中，则 `dp[i] = true`，跳出内层循环。
4. 返回 `dp[n]`。

### 复杂度分析

- **时间复杂度**: O(n · m · k)，其中 `n = len(s)`，`m = maxLen`（最长字典单词长度，≤20），`k` 为子串哈希/比较的开销（每次截取 `s[j:i]` 并查询集合需要 O(k)）。若不按 `maxLen` 剪枝则为 O(n² · k)。
- **空间复杂度**: O(n + d)，`dp` 数组占 O(n)，字典哈希集合占 O(d)（d 为字典总字符数）。

## 代码实现

```go
func WordBreak(s string, wordDict []string) bool {
    // 字典放入哈希集合加速查询，同时记录最长单词长度用于剪枝
    dict := make(map[string]struct{}, len(wordDict))
    maxLen := 0
    for _, w := range wordDict {
        dict[w] = struct{}{}
        if len(w) > maxLen {
            maxLen = len(w)
        }
    }

    n := len(s)
    // dp[i] 表示 s[:i] 能否用字典单词拼出，空串可以
    dp := make([]bool, n+1)
    dp[0] = true

    for i := 1; i <= n; i++ {
        // 只枚举最后一个单词的起点 j，长度超过 maxLen 的子串一定不在字典中
        start := i - maxLen
        if start < 0 {
            start = 0
        }
        for j := start; j < i; j++ {
            if dp[j] {
                if _, ok := dict[s[j:i]]; ok {
                    dp[i] = true
                    break
                }
            }
        }
    }
    return dp[n]
}
```

**执行过程示例**（`s = "leetcode"`, `wordDict = ["leet","code"]`）：

```
dp[0] = true
i=4:  s[0:4]="leet" 在字典中 → dp[4] = true
i=8:  dp[4] 为 true 且 s[4:8]="code" 在字典中 → dp[8] = true
返回 dp[8] = true
```
