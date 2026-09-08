# 5. 最长回文子串 (Longest Palindromic Substring)

## 题目描述

给你一个字符串 `s`，找到 `s` 中最长的回文子串。

回文子串是指正序（从左向右）和倒序（从右向左）读都是一样的子串。如果存在多个长度相同的最长回文子串，返回其中任意一个即可。

### 示例 1

```
输入: s = "babad"
输出: "bab"
解释: "aba" 同样是符合题意的答案。
```

### 示例 2

```
输入: s = "cbbd"
输出: "bb"
```

## 提示

- `1 <= s.length <= 1000`
- `s` 仅由数字和英文字母组成

## 题目解析

### 核心思路

本题采用**中心扩展法**：回文串一定关于某个中心对称，因此可以枚举每一个可能的“对称中心”，从这个中心向两侧同时扩展，直到两侧字符不相等或越界，得到以该中心对称的最长回文串；所有中心中长度最大的那个即为答案。

关键点在于对称中心有两种形态：

- **奇数长度回文**：中心是某个字符 `s[i]`，例如 `"bab"` 以 `'a'` 为中心，对应调用 `expand(s, i, i)`。
- **偶数长度回文**：中心位于两个字符 `s[i]` 和 `s[i+1]` 之间，例如 `"bb"`，对应调用 `expand(s, i, i+1)`。

对每个位置 `i`，两种形态都要尝试一次，所以一共枚举 `2n - 1` 个中心。枚举过程中记录当前最长回文的起始下标 `start` 和长度 `maxLen`，最后用切片 `s[start : start+maxLen]` 取出结果，无需在扩展过程中反复构造字符串。

本题属于**“回文串中心扩展”**模板，相比动态规划写法（O(n^2) 时间 + O(n^2) 空间），中心扩展把空间压到了 O(1)。

### 算法步骤

1. 若 `len(s) < 2`，单字符（或空串）本身就是回文，直接返回 `s`。
2. 初始化 `start = 0`、`maxLen = 1`（由约束 `s.length >= 1` 保证至少有一个字符的回文）。
3. 遍历每个下标 `i`，分别尝试两种中心：
   - `expand(s, i, i)`：奇数长度中心，若返回长度 `l > maxLen`，则更新 `start = i - l/2`、`maxLen = l`。
   - `expand(s, i, i+1)`：偶数长度中心，若返回长度 `l > maxLen`，则更新 `start = i + 1 - l/2`、`maxLen = l`。
4. `expand(left, right)` 的实现：只要 `left >= 0`、`right < len(s)` 且 `s[left] == s[right]` 就继续向两侧走；循环结束时回文区间是 `(left, right)` 开区间，长度为 `right - left - 1`。
5. 返回 `s[start : start+maxLen]`。

### 复杂度分析

- **时间复杂度**: O(n²)，枚举 `2n - 1` 个中心，每个中心最多向两侧扩展 O(n) 步。
- **空间复杂度**: O(1)，只使用常数个变量（不计返回结果本身占用的空间）。

## 代码实现

```go
func LongestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }

    start, maxLen := 0, 1
    // 以 i（奇数长度回文）或 i、i+1 之间（偶数长度回文）为中心向两侧扩展
    for i := 0; i < len(s); i++ {
        if l := expand(s, i, i); l > maxLen {
            start, maxLen = i-l/2, l
        }
        if l := expand(s, i, i+1); l > maxLen {
            start, maxLen = i+1-l/2, l
        }
    }
    return s[start : start+maxLen]
}

func expand(s string, left, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right - left - 1
}
```

**执行过程示例**（`s = "babad"`）：

```
i=0: 中心 b   → "b"（长度1）
     中心 b|a → 无法扩展（长度0）
i=1: 中心 a   → 扩展到 "bab"（长度3），更新 start=0, maxLen=3
     中心 a|b → 无法扩展
i=2: 中心 b   → 扩展到 "aba"（长度3），不更新（未超过 maxLen）
     中心 b|a → 无法扩展
i=3: 中心 a   → "a"（长度1）
     中心 a|d → 无法扩展
返回 s[0:3] = "bab"
```
