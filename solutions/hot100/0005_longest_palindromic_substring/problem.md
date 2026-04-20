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

---

### 解法二：动态规划

#### 思路

定义 `dp[i][j]` 表示 `s[i..j]` 是否为回文。状态转移：

```
dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
```

**初始条件**：
- 长度 1：`dp[i][i] = true`（单字符必回文）
- 长度 2：`dp[i][i+1] = (s[i] == s[i+1])`

**填表顺序**：按子串长度从 3 递增到 n，这样计算 `dp[i][j]` 时 `dp[i+1][j-1]` 已就绪。

在填表过程中记录最长回文的起始位置和长度。

#### 复杂度

- **时间**：O(n²)，双重循环
- **空间**：O(n²)，`dp` 二维数组

与中心扩展相比，DP 多了 O(n²) 的空间开销，但思路更通用，适合作为 DP 练习。

#### 代码

```go
func longestPalindromeDP(s string) string {
    n := len(s)
    if n <= 1 {
        return s
    }

    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
        dp[i][i] = true
    }

    start, maxLen := 0, 1

    // 长度 2
    for i := 0; i < n-1; i++ {
        if s[i] == s[i+1] {
            dp[i][i+1] = true
            start, maxLen = i, 2
        }
    }

    // 长度 3 ~ n
    for length := 3; length <= n; length++ {
        for i := 0; i <= n-length; i++ {
            j := i + length - 1
            if s[i] == s[j] && dp[i+1][j-1] {
                dp[i][j] = true
                if length > maxLen {
                    start, maxLen = i, length
                }
            }
        }
    }

    return s[start : start+maxLen]
}
```

#### DP 填表过程示例（`s = "babad"`）

```
      b  a  b  a  d
  b [ T  .  T  .  . ]
  a [ .  T  .  T  . ]
  b [ .  .  T  .  . ]
  a [ .  .  .  T  . ]
  d [ .  .  .  .  T ]

长度2: 无相邻相同字符
长度3: dp[0][2]=T (s[0]='b'==s[2]='b', dp[1][1]=T) → "bab"
       dp[1][3]=T (s[1]='a'==s[3]='a', dp[2][2]=T) → "aba"
长度4: dp[0][3]=F (s[0]='b'!=s[3]='a')
       dp[1][4]=F (s[1]='a'!=s[4]='d')
长度5: dp[0][4]=F (s[0]='b'!=s[4]='d')
结果: maxLen=3, start=0 → "bab"
```
