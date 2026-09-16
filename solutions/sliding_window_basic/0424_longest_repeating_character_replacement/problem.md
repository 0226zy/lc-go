# 424. 替换后的最长重复字符 (Longest Repeating Character Replacement)

## 题目描述

给你一个字符串 `s` 和一个整数 `k`。你可以选择字符串中的任一字符，并将其改变为任何其他的大写英文字母。该操作最多可执行 `k` 次。

在执行上述操作后，返回 **包含相同字母的最长子串的长度**。

### 示例 1

```
输入: s = "ABAB", k = 2
输出: 4
解释: 用两个 'A' 替换两个 'B'，反之亦然，得到的串全部相同，长度为 4。
```

### 示例 2

```
输入: s = "AABABBA", k = 1
输出: 4
解释: 将中间的一个 'A' 替换为 'B'，字符串变为 "AABBBBA"。
子串 "BBBB" 有最长的重复字母，答案为 4。
```

## 提示

- `1 <= s.length <= 10^5`
- `s` 仅由大写英文字母组成
- `0 <= k <= s.length`

## 题目解析

### 核心思路

本题是 **滑动窗口（Sliding Window）** 的经典应用。

关键观察：对于一个窗口 `[left, right]`，要把它变成全部相同字符的子串，需要替换的次数是：

```
窗口长度 - 窗口内出现次数最多的字符的个数
```

即：**留下出现最多的字符不动，把其余字符全部改成它**。替换次数越少越好，所以最优策略就是朝着窗口内的众数字符替换。

因此维护一个滑动窗口：

- 用计数数组 `count[26]` 记录窗口内每个大写字符的出现次数；
- 用 `maxCount` 记录窗口内出现次数最多的字符个数（即历史上遇到过的最大众数频次）；
- 当 `窗口长度 - maxCount > k` 时，说明即使把其余字符全部替换也超出预算，窗口不合法，收缩左边界；
- 每次右边界扩张后，用当前窗口长度更新答案。

**为什么 `maxCount` 收缩时不需要回退？** 当左边界右移导致众数字符离开窗口时，`maxCount` 可能短暂“虚高”，但这只会让窗口暂时不被额外收缩。窗口长度一旦超过历史最优，就必须伴随着 `maxCount` 的真实增长（否则 `right - left + 1 - maxCount > k` 会立即收缩），因此不会影响答案的正确性，只是减少了不必要的收缩，是一种常见的“懒惰收缩”技巧。

### 算法步骤

1. 初始化 `count [26]int` 计数数组、`left = 0`、`maxCount = 0`、`ans = 0`。
2. 右指针 `right` 从 0 遍历到 `len(s)-1`：
   - 将 `s[right]` 加入窗口：`count[s[right]-'A']++`，并更新 `maxCount`。
   - 若 `right - left + 1 - maxCount > k`：窗口不合法，将 `s[left]` 移出窗口（计数减一），`left++`。
   - 更新答案：`ans = max(ans, right - left + 1)`。
3. 返回 `ans`。

### 复杂度分析

- **时间复杂度**: O(n)。`right` 和 `left` 各自单调右移，每个字符最多进出窗口一次；计数更新和最大值维护都是 O(1)（字符集大小固定为 26）。
- **空间复杂度**: O(|Σ|) = O(1)。`count` 数组大小固定为 26，与输入规模无关。

## 代码实现

```go
func LongestRepeatingCharacterReplacement(s string, k int) int {
    var count [26]int
    left, maxCount, ans := 0, 0, 0
    for right := 0; right < len(s); right++ {
        // 右端字符加入窗口，更新窗口内最大众数频次
        count[s[right]-'A']++
        if count[s[right]-'A'] > maxCount {
            maxCount = count[s[right]-'A']
        }
        // 需要替换的次数超过 k，收缩左边界
        if right-left+1-maxCount > k {
            count[s[left]-'A']--
            left++
        }
        // 窗口一定合法，更新答案
        if right-left+1 > ans {
            ans = right - left + 1
        }
    }
    return ans
}
```

**执行过程示例**（示例 2：`s = "AABABBA", k = 1`）：

```
right=0 'A': count[A]=1, maxCount=1, 窗口"A"，窗口长 1, ans=1
right=1 'A': count[A]=2, maxCount=2, 窗口"AA"，窗口长 2, ans=2
right=2 'B': count[B]=1, maxCount=2, 窗口"AAB"，2-...=3-2=1 <= k, ans=3
right=3 'A': count[A]=3, maxCount=3, 窗口"AABA"，4-3=1 <= k, ans=4
right=4 'B': count[B]=2, maxCount=3, 窗口"AABAB"，5-3=2 > 1 → 收缩 left
             移出 'A'，窗口"ABAB"，4-3=1 <= k, ans=4
right=5 'B': count[B]=3, maxCount=3, 窗口"ABABB"，5-3=2 > 1 → 收缩 left
             移出 'B'，窗口"BABB"，4-3=1 <= k, ans=4
right=6 'A': count[A]=3, maxCount=3, 窗口"BABBA"，5-3=2 > 1 → 收缩 left
             移出 'B'，窗口"ABBA"，4-3=1 <= k, ans=4
结果: 4（对应子串 "AABA" 或 "ABBA" 替换一个字符后全部相同）
```
