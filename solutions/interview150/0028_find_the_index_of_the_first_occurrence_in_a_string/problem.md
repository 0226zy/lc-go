# 28. 找出字符串中第一个匹配项的下标 (Find the Index of the First Occurrence in a String)

## 题目描述

给你两个字符串 `haystack` 和 `needle` ，请你在 `haystack` 字符串中找出 `needle` 字符串的第一个匹配项的下标（下标从 0 开始）。如果 `needle` 不是 `haystack` 的一部分，则返回 `-1`。

### 示例 1

```
输入: haystack = "sadbutsad", needle = "sad"
输出: 0
解释: "sad" 在下标 0 和 6 处匹配，第一个匹配项的下标是 0。
```

### 示例 2

```
输入: haystack = "leetcode", needle = "leeto"
输出: -1
解释: "leeto" 没有在 "leetcode" 中出现。
```

### 提示

- `1 <= haystack.length, needle.length <= 10^4`
- `haystack` 和 `needle` 仅由小写英文字符组成

## 题目解析

### 核心思路

这道题就是经典的**字符串单模式匹配**问题：在一个文本串 `haystack` 中查找模式串 `needle` 第一次出现的位置。

最朴素的解法是暴力双重循环：外层枚举起始下标 `i`，内层逐字符比较 `haystack[i+j]` 和 `needle[j]`，一旦失配就整体后移一位重新比较。最坏情况下（比如 `haystack = "aaaa...a"`，`needle = "aaa...b"`）时间复杂度是 O(mn)，每个位置都要重新比较一长串。

**KMP 算法**的核心观察是：暴力算法之所以慢，是因为它每次失配后把模式串只后移一位，大量信息被浪费了。当在位置 `j` 失配时，模式串前 `j` 个字符已经和文本串匹配过了，这些匹配信息本身告诉我们「模式串的前缀中，最长的一段能同时作为后缀拼在当前位置」的长度——也就是可以直接把模式串向后滑动到合适的位置，而不是退回到只移一位。

为了利用这个信息，先对模式串做一次预处理，求出 `next` 数组（也称作 `lps`，longest prefix-suffix）：

- `next[j]` 表示：`needle[0..j]` 这个子串中，**最长的、既是前缀又是后缀**的真子串长度。例如 `"abab"` 的最长相等前后缀是 `"ab"`，长度为 2。
- 匹配时在 `haystack[i] != needle[j]` 失配，就让 `j = next[j-1]`，相当于把模式串向右滑动，跳过不可能匹配的位置，而 `i` 不回退。

KMP 把最坏时间复杂度从 O(mn) 降到 O(m+n)，是本题的标准最优解。

### 算法步骤

1. **预处理 next 数组**：
   - `next[0] = 0`（单个字符没有真前后缀）
   - 用双指针 `i`（遍历后缀）和 `j`（遍历前缀）扫描模式串：
     - 若 `needle[i] == needle[j]`，说明前后缀可各延长一位，`next[i] = j + 1`，两个指针都前进；
     - 若不等，说明当前前缀长度接不上，利用已求出的 `next[j-1]` 回退 `j`，继续尝试更短的前缀；直到 `j` 回退到 0 还不匹配，则 `next[i] = 0`。
2. **主匹配**：
   - 双指针 `i` 遍历 `haystack`，`j` 遍历 `needle`；
   - 字符相等则两个指针都前进；
   - 失配时 `i` 不动，`j = next[j-1]` 回退；
   - 若 `j` 到达 `len(needle)`，说明匹配成功，返回 `i - len(needle) + 1`。
3. 遍历结束仍未匹配成功，返回 `-1`。

### 复杂度分析

- **时间复杂度**: O(m + n)，m 为 `haystack` 长度，n 为 `needle` 长度。预处理 next 数组 O(n)，主匹配 O(m)，且主匹配中 `i` 不回退。
- **空间复杂度**: O(n)，存储 next 数组。

## 代码实现

### 主解：KMP

```go
// StrStr KMP 匹配：返回 needle 在 haystack 中第一次出现的下标，不存在返回 -1
func StrStr(haystack, needle string) int {
    n := len(needle)
    if n == 0 {
        return 0
    }
    // 1. 预处理 next 数组：next[i] 表示 needle[0..i] 的最长相等真前后缀长度
    next := make([]int, n)
    for i, j := 1, 0; i < n; i++ {
        for j > 0 && needle[i] != needle[j] {
            j = next[j-1] // 回退到更短的前缀继续尝试
        }
        if needle[i] == needle[j] {
            j++
        }
        next[i] = j
    }
    // 2. 主匹配：i 不回退，失配时 j 按 next 回退
    for i, j := 0, 0; i < len(haystack); i++ {
        for j > 0 && haystack[i] != needle[j] {
            j = next[j-1]
        }
        if haystack[i] == needle[j] {
            j++
        }
        if j == n {
            return i - n + 1
        }
    }
    return -1
}
```

### 附：对照实现

```go
// StrStrBuiltin 基于 strings.Index 的实现（内部使用 Rabin-Karp 等优化）
func StrStrBuiltin(haystack, needle string) int {
    return strings.Index(haystack, needle)
}
```

**执行过程示例**（`haystack = "sadbutsad"`, `needle = "sad"`）：

```
needle = "sad"，next 数组：
  子串 "s": next[0] = 0
  子串 "sa": 无相等前后缀, next[1] = 0
  子串 "sad": 无相等前后缀, next[2] = 0

主匹配:
i=0 's' == 's' → j=1
i=1 'a' == 'a' → j=2
i=2 'd' == 'd' → j=3 == len(needle)，匹配成功，返回 i-n+1 = 2-3+1 = 0
```

**性能优势**：KMP 保证最坏 O(m+n)，比暴力匹配的 O(mn) 在最坏输入（大量重复字符）下快得多。
