# 14. 最长公共前缀 (Longest Common Prefix)

## 题目描述

给定一个字符串数组 `strs`，返回这些字符串的最长公共前缀。如果不存在公共前缀，返回空字符串 `""`。

### 示例 1

```
输入: strs = ["flower","flow","flight"]
输出: "fl"
```

### 示例 2

```
输入: strs = ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
```

### 提示

- `1 <= strs.length <= 200`
- `0 <= strs[i].length <= 200`
- `strs[i]` 仅由小写英文字母组成

## 题目解析

### 核心思路

这道题属于**纵向扫描**类题目。最直接的想法：以第一个字符串为“基准答案”，然后逐个与后面的字符串比较，不断把前缀“剪短”，直到它能同时是所有字符串的前缀为止。

关键洞察：
- 最长公共前缀的长度，**不可能超过数组中最短的字符串**。
- 只要发现某个位置上字符不一致，该位置及其后面的部分就全部不属于公共前缀。

另一种思路是“纵向扫描”：固定列 `j`，逐列比较所有字符串的同一列字符，遇到不同就返回前 `j` 个字符。两种解法复杂度相同，这里采用更直观的“逐个缩短前缀”法。

### 算法步骤

1. 如果数组为空，直接返回 `""`。
2. 取 `strs[0]` 作为初始 `prefix`。
3. 从第二个字符串开始，逐个比较：
   - 只要 `prefix` 不是当前字符串的前缀，就不断删掉 `prefix` 的最后一个字符（缩短）。
   - 如果缩短到空，直接返回 `""`。
4. 遍历完所有字符串后，剩下的 `prefix` 就是答案。

### 复杂度分析

- **时间复杂度**: O(n × m)，n 为字符串个数，m 为最短字符串长度。最坏情况下每个字符串的每个字符都要比较一次
- **空间复杂度**: O(1)，没有额外分配空间（`prefix` 是原字符串的切片视图）

## 代码实现

```go
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !hasPrefix(strs[i], prefix) {
			if prefix == "" {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func hasPrefix(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	return s[:len(prefix)] == prefix
}
```

**执行过程示例**（`strs = ["flower","flow","flight"]`）：

```
初始 prefix = "flower"
与 "flow" 比较:
  "flower" 不是 "flow" 的前缀 → 缩短为 "flowe"
  "flowe"  不是 "flow" 的前缀 → 缩短为 "flow"
  "flow"   是   "flow" 的前缀 → 停止，prefix = "flow"
与 "flight" 比较:
  "flow"   不是 "flight" 的前缀 → 缩短为 "flo"
  "flo"    不是 "flight" 的前缀 → 缩短为 "fl"
  "fl"     是   "flight" 的前缀 → 停止，prefix = "fl"

结果: "fl"
```
