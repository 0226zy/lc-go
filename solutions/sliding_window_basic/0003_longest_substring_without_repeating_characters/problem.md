# 3. 无重复字符的最长子串 (Longest Substring Without Repeating Characters)

## 题目描述

给定一个字符串 `s`，请你找出其中**不含有重复字符的** **最长子串**的长度。

### 示例 1

```
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

### 示例 2

```
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
```

### 示例 3

```
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是子串的长度，"pwke" 是一个子序列，不是子串。
```

### 提示

- `0 <= s.length <= 5 * 10^4`
- `s` 由英文字母、数字、符号和空格组成

## 题目解析

### 核心思路

子串 = 连续的一段，这决定了可以用**滑动窗口**：维护一个「窗口内没有重复字符」的区间 `[left, right]`，用 `right` 逐个扫字符；一旦新字符和窗口里的某个字符重复，就把 `left` 跳过去。

关键是 `left` 该跳到哪。朴素的想法是 `left` 一步一步右移，直到重复消失——但那样可能退化成 O(n²)。注意到：如果窗口里字符 `c` 上一次出现在位置 `idx`，那么任何以 `idx` 或更早位置开头的子串都必然包含两个 `c`，可以直接跳过。于是用哈希表 `last` 记录每个字符**最近一次出现的位置**，遇到重复时直接把 `left` 一次性跳到 `idx + 1`，两个指针都只向右走，整体 O(n)。

这就是「**可变长度滑动窗口 + 哈希表记录位置**」的模板，和 209 题的区别在于：209 收缩的条件是「和超标」，本题收缩的条件是「出现重复字符」。

### 算法步骤

1. 初始化空哈希表 `last`（字符 → 最近一次出现的下标），`left = 0`，`ans = 0`。
2. `right` 从左到右遍历字符串：
   - 取当前字符 `c = s[right]`；
   - 若 `c` 之前出现过，且它的上次位置 `idx` 还在当前窗口内（`idx >= left`），则把 `left` 跳到 `idx + 1`；
   - 更新 `last[c] = right`；
   - 用窗口长度 `right - left + 1` 更新 `ans`。
3. 返回 `ans`。

### 复杂度分析

- **时间复杂度**: O(n)，每个字符最多进窗口一次、`left` 最多跳 n 次
- **空间复杂度**: O(|Σ|)，|Σ| 为字符集大小，哈希表最多存字符集那么多键

## 代码实现

```go
func LengthOfLongestSubstring(s string) int {
	// last 记录每个字符最近一次出现的位置（下标）
	last := make(map[byte]int)
	left, ans := 0, 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		// 若字符 c 已出现在当前窗口 [left, right) 内，则把左边界直接跳到它后面
		if idx, ok := last[c]; ok && idx >= left {
			left = idx + 1
		}
		last[c] = right
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
```

**执行过程示例**（`s = "pwwkew"`）：

```
初始: left=0, ans=0
right=0 'p': 未见过, last[p]=0, 窗口长1, ans=1
right=1 'w': 未见过, last[w]=1, 窗口长2, ans=2
right=2 'w': last[w]=1 在窗口内 → left=2, last[w]=2, 窗口长1, ans=2
right=3 'k': 未见过, last[k]=3, 窗口长2 ("wk"), ans=2
right=4 'e': 未见过, last[e]=4, 窗口长3 ("wke"), ans=3
right=5 'w': last[w]=2 已不在窗口内(2 < left=3), 不跳; last[w]=5, 窗口长3 ("kew"), ans=3
结果: 3
```
