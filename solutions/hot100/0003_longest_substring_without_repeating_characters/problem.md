# 3. 无重复字符的最长子串 (Longest Substring Without Repeating Characters)

## 题目描述

给定一个字符串 `s`，请你找出其中不含有重复字符的 **最长子串** 的长度。

### 示例 1

```
输入：s = "abcabcbb"
输出：3
解释：因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

### 示例 2

```
输入：s = "bbbbb"
输出：1
解释：因为无重复字符的最长子串是 "b"，所以其长度为 1。
```

### 示例 3

```
输入：s = "pwwkew"
输出：3
解释：因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是子串的长度，"pwke" 是一个子序列，不是子串。
```

## 提示

- `0 <= s.length <= 5 * 10^4`
- `s` 由英文字母、数字、符号和空格组成

## 题目解析

### 核心思路

这道题是 **滑动窗口** 的经典应用：

1. **滑动窗口**：维护一个窗口 `[left, right)`，保证窗口内的字符不重复。右指针不断右移扩大窗口，当遇到重复字符时，左指针右移收缩窗口，直到窗口内无重复字符。

2. **哈希表/数组记录**：用哈希表（或长度 128 的数组）记录每个字符最近出现的位置，当发现重复时可以直接跳转左指针，避免逐步移动。

3. **关键点**：左指针只能往右移，不能回退，这保证了整体 O(n) 的时间复杂度。

### 算法步骤

1. 初始化左指针 `left = 0`，最大长度 `maxLen = 0`
2. 用哈希表 `charIndex` 记录字符最近出现的索引
3. 右指针 `right` 从 0 遍历到末尾：
   - 若 `s[right]` 已在哈希表中且位置 >= `left`，则更新 `left = charIndex[s[right]] + 1`
   - 更新 `charIndex[s[right]] = right`
   - 更新 `maxLen = max(maxLen, right - left + 1)`
4. 返回 `maxLen`

### 复杂度分析

- **时间复杂度**: O(n)，其中 n 为字符串长度，每个字符最多被访问两次
- **空间复杂度**: O(min(m, n))，其中 m 为字符集大小（ASCII 为 128）

## 代码实现

### 最优解：滑动窗口 + 固定数组

用 `[128]int` 数组替代哈希表，数组分配在栈上，零堆分配。`pos[c]` 存储字符 `c` 最近出现的位置 +1（利用零值天然表示"未出现"，省去额外判断）。

```go
func lengthOfLongestSubstring(s string) int {
    var pos [128]int // pos[c] = 字符 c 最近出现位置 +1，0 表示未出现
    maxLen := 0
    left := 0
    for right := 0; right < len(s); right++ {
        if p := pos[s[right]]; p > left {
            left = p // 直接跳到重复字符的下一个位置
        }
        pos[s[right]] = right + 1
        if l := right - left + 1; l > maxLen {
            maxLen = l
        }
    }
    return maxLen
}
```

**执行过程示例**（`s = "abcabcbb"`）：

```
right=0 'a': pos['a']=0, left=0, pos['a']=1, len=1, maxLen=1
right=1 'b': pos['b']=0, left=0, pos['b']=2, len=2, maxLen=2
right=2 'c': pos['c']=0, left=0, pos['c']=3, len=3, maxLen=3
right=3 'a': pos['a']=1>0, left=1, pos['a']=4, len=3, maxLen=3
right=4 'b': pos['b']=2>1, left=2, pos['b']=5, len=3, maxLen=3
right=5 'c': pos['c']=3>2, left=3, pos['c']=6, len=3, maxLen=3
right=6 'b': pos['b']=5>3, left=5, pos['b']=7, len=2, maxLen=3
right=7 'b': pos['b']=7>5, left=7, pos['b']=8, len=1, maxLen=3
结果：3
```

**性能优势**：
- `0 B/op, 0 allocs`：固定数组在栈上分配，无堆内存分配
- 单次遍历，无函数调用开销
- 直接数组索引 O(1)，比 map 的哈希查找更快
