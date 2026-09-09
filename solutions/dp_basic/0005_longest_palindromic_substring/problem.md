# 0005. 最长回文子串 (Longest Palindromic Substring)

## 题目描述

给你一个字符串 `s`，找到 `s` 中最长的**回文子串**。

回文串指正着读和反着读都一样的字符串，例如 `"aba"`、`"abba"`。

### 示例

```
输入: s = "babad"
输出: "bab"
解释: "aba" 同样是符合题意的答案。
```

```
输入: s = "cbbd"
输出: "bb"
```

### 提示

- `1 <= s.length <= 1000`
- `s` 仅由数字和英文字母组成

## 四步拆解

### 1. 状态定义

- `dp[i][j]` 表示**子串 `s[i..j]`（含两端）是否为回文串**，值为布尔类型。
- 最终答案：遍历中所有 `dp[i][j] == true` 的子串里，长度 `j - i + 1` 最大的那个。

为什么用二维状态？因为「`s[i..j]` 是否回文」取决于两端字符 `s[i]`、`s[j]` 是否相等，以及**去掉两端后的内部子串** `s[i+1..j-1]` 是否回文——这正是子问题，用两个下标才能描述。

### 2. 边界条件 + 遍历顺序

- base case：
  - 单个字符一定是回文：`dp[i][i] = true`；
  - 两个相邻字符相等即回文：`s[i] == s[j]` 时 `dp[i][j] = true`（长度 2 时内部子串为空，视为回文）。
- 遍历顺序：转移 `dp[i][j]` 依赖 `dp[i+1][j-1]`，即**更短的内部子串**，所以必须保证短子串先算。最稳妥的写法是**按子串长度 `length` 从 1 到 n 递增枚举**，外层是长度，内层是起点 `i`，终点 `j = i + length - 1`。

### 3. 状态转移方程

```
dp[i][j] = (s[i] == s[j]) && (j - i < 2 || dp[i+1][j-1])
```

推导：考察子串 `s[i..j]`，要么两端字符不相等，直接不是回文；要么两端相等，此时再看内部 `s[i+1..j-1]` 是否回文。当子串长度小于 3（`j - i < 2`）时没有内部子串，两端相等就是回文。每填一个 `true` 就顺手比较长度，更新答案的起点和最大长度。

### 4. 样例验证 + 代码实现

手动推导 `s = "babad"`（只列出为 `true` 的 `dp[i][j]`）：

| 长度 | 子串 | dp 值 |
| --- | --- | --- |
| 1 | b, a, b, a, d | 全为 true |
| 2 | "ba" false, "ab" false, "ba" false, "ad" false | 无 |
| 3 | "bab" true（两端 b 相等，内部 "a" 回文）, "aba" true, "bad" false | 最长更新为 3 |
| 4 | "baba" false, "abad" false | 无 |
| 5 | "babad" false | 无 |

最长回文长度为 3，对应 `"bab"`（或 `"aba"`）✅（与示例一致）

代码与上面的推导一一对应：

```go
// LongestPalindrome 最长回文子串（标准 DP 数组版）
func LongestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	start, maxLen := 0, 1
	for length := 1; length <= n; length++ { // 按子串长度枚举
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				dp[i][j] = length < 3 || dp[i+1][j-1]
			}
			if dp[i][j] && length > maxLen {
				start, maxLen = i, length
			}
		}
	}
	return s[start : start+maxLen]
}
```

## 中心扩展法（不使用 DP 数组）

换一个角度：回文串一定有一个「对称中心」。枚举每个字符（奇数长度回文）和每对相邻字符之间（偶数长度回文）作为中心，向两边扩展直到两端字符不相等，扩展出的最长串就是以该中心的最长回文。中心一共 `2n - 1` 个，每个中心最多扩展 O(n) 次，全程不需要额外数组，空间 O(1)：

```go
// LongestPalindromeAlternative 最长回文子串（中心扩展法）
func LongestPalindromeAlternative(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	start, maxLen := 0, 1
	expand := func(l, r int) { // 从中心 l,r 向两边扩展
		for l >= 0 && r < n && s[l] == s[r] {
			if r-l+1 > maxLen {
				start, maxLen = l, r-l+1
			}
			l--
			r++
		}
	}
	for i := 0; i < n; i++ {
		expand(i, i)   // 奇数长度，中心是单个字符
		expand(i, i+1) // 偶数长度，中心是两个字符之间
	}
	return s[start : start+maxLen]
}
```

## 复杂度分析

| 版本 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- |
| DP 数组版 | O(n²) | O(n²) |
| 中心扩展法 | O(n²) | O(1) |

**进阶**：本题还有 O(n) 的 Manacher 算法，通过改造字符串消除奇偶差异并复用已算出的回文半径，面试中可作加分项了解。
