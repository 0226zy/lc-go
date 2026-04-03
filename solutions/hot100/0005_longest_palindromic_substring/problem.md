# 5. 最长回文子串 (Longest Palindromic Substring)

## 题目描述

给你一个字符串 `s`，找到 `s` 中最长的回文子串。

### 示例 1

```
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
```

### 示例 2

```
输入：s = "cbbd"
输出："bb"
```

### 示例 3

```
输入：s = "a"
输出："a"
```

### 示例 4

```
输入：s = "ac"
输出："a"
```

## 提示

- `1 <= s.length <= 1000`
- `s` 仅由数字和英文字母组成

## 题目解析

### 核心思路（中心扩展）

回文串关于「中心」对称。枚举每个可能的中心，向两侧扩展，只要两端字符相同就继续扩。

- **奇数长度回文**：中心为下标 `i` 的一个字符，初始 `left = right = i`。
- **偶数长度回文**：中心为 `i` 与 `i+1` 之间，初始 `left = i`、`right = i+1`，若 `s[left] != s[right]` 则长度为 0，扩展函数会立即返回。

对每个中心得到当前最长回文区间，与全局最优比较并更新起止下标。

### 复杂度

- **时间复杂度**：O(n²)，共约 2n 个中心，每次扩展最多 O(n)。
- **空间复杂度**：O(1)，只使用若干下标变量。

若需 **O(n)** 可研究 Manacher 算法，本题数据规模下中心扩展足够。

## 代码实现

```go
func longestPalindrome(s string) string {
    if len(s) == 0 {
        return ""
    }
    start, end := 0, 0
    for i := range s {
        if l, r := expand(s, i, i); r-l > end-start {
            start, end = l, r
        }
        if l, r := expand(s, i, i+1); r-l > end-start {
            start, end = l, r
        }
    }
    return s[start : end+1]
}

func expand(s string, left, right int) (int, int) {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return left + 1, right - 1
}
```
