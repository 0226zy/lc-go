# 131. 分割回文串 (Palindrome Partitioning)

## 题目描述

给你一个字符串 `s`，请你将 `s` 分割成一些子串，使每个子串都是 **回文串** 。返回 `s` 所有可能的分割方案。

**回文串** 是正着读和反着读都一样的字符串。

### 示例 1

```
输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
```

### 示例 2

```
输入：s = "a"
输出：[["a"]]
```

## 提示

- `1 <= s.length <= 16`
- `s` 仅由小写英文字母组成

## 题目解析

### 核心思路

这是 **回溯“分割”模板**：把字符串的每个字符间隙看作一个“可切/不可切”的决策点，用递归枚举所有切法，只保留每段都是回文串的方案。

回溯三要素：

- **路径**：已经切出的回文子串列表 `path`，以及当前切割起点 `start`。
- **选择列表**：下一个子串的结束位置 `end`（`start <= end < n`），即所有候选子串 `s[start:end+1]`。
- **结束条件**：`start == len(s)`，说明整串被切完且每段都是回文，把 `path` 的拷贝收入结果。

**剪枝**：枚举 `end` 时先判断 `s[start:end+1]` 是否为回文串，不是就直接跳过——非回文的段不可能出现在合法方案里，也就不必向下递归。判断回文用双指针，单次 O(n)。

关键点：递归返回后要执行 `path = path[:len(path)-1]` 撤销选择，且收集答案时必须 **拷贝** `path`，否则后续修改会污染已收集的结果。

### 算法步骤

1. 初始化空结果集 `result` 与空路径 `path`。
2. 从 `start = 0` 开始回溯：
   - 若 `start == len(s)`，把 `path` 的拷贝加入 `result`，返回。
   - 枚举 `end` 从 `start` 到 `len(s)-1`：
     - 若 `s[start:end+1]` 不是回文串，`continue` 跳过；
     - 把 `s[start:end+1]` 加入 `path`；
     - 递归 `backtrack(end + 1)`；
     - 回溯：弹出末尾子串。
3. 返回 `result`。

### 复杂度分析

- **时间复杂度**: O(n × 2^n)。n 个字符共有 n-1 个间隙，每个间隙可切/不切，最坏有 2^(n-1) 种切法；每种切法判断回文与拷贝结果的代价为 O(n)。
- **空间复杂度**: O(n)，递归栈最深 n 层，`path` 最长存 n 个单字符子串（不计输出本身）。

## 代码实现

```go
package palindromepartitioning

// Partition 分割回文串
// 给定字符串 s，将 s 分割成若干子串，使每个子串都是回文串，返回所有可能的分割方案。
// 时间复杂度: O(n × 2^n)  n 为字符串长度，每个位置都可切/不切  空间复杂度: O(n) 递归栈深度（不计输出）
func Partition(s string) [][]string {
	var result [][]string
	path := []string{}

	// backtrack 表示当前准备从 s[start:] 中切出下一个回文子串
	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			// 已经切到字符串末尾，当前 path 是一种完整分割方案，拷贝后收集
			partition := make([]string, len(path))
			copy(partition, path)
			result = append(result, partition)
			return
		}
		// 枚举下一个子串的结束位置 end，子串为 s[start:end+1]
		for end := start; end < len(s); end++ {
			if !isPalindrome(s, start, end) {
				continue // 不是回文串不能作为一段，直接跳过
			}
			path = append(path, s[start:end+1])
			backtrack(end + 1)
			path = path[:len(path)-1] // 回溯：撤销本次切割
		}
	}
	backtrack(0)
	return result
}

// isPalindrome 双指针判断 s[left:right+1] 是否为回文串
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
```

**执行过程示例**（`s = "aab"`）：

```
backtrack(start=0) path=[]
  end=0: "a" 是回文 → path=["a"]
    backtrack(start=1)
      end=1: "a" 是回文 → path=["a","a"]
        backtrack(start=2)
          end=2: "b" 是回文 → path=["a","a","b"]
            backtrack(start=3) → start==len(s)，收集 ["a","a","b"]
          回溯 → path=["a","a"]
      回溯 → path=["a"]
      end=2: "ab" 不是回文 → 跳过
  回溯 → path=[]
  end=1: "aa" 是回文 → path=["aa"]
    backtrack(start=2)
      end=2: "b" 是回文 → path=["aa","b"]
        backtrack(start=3) → 收集 ["aa","b"]
      回溯 → path=["aa"]
  回溯 → path=[]
  end=2: "aab" 不是回文 → 跳过
结果: [["a","a","b"],["aa","b"]]
```
