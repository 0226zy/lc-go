# 125. 验证回文串 (Valid Palindrome)

## 题目描述

如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样，则可以认为该短语是一个**回文串**。

字母和数字都属于字母数字字符。

给你一个字符串 `s`，如果它是**回文串**，返回 `true`；否则，返回 `false`。

### 示例 1

```
输入: s = "A man, a plan, a canal: Panama"
输出: true
解释: "amanaplanacanalpanama" 是回文串。
```

### 示例 2

```
输入: s = "race a car"
输出: false
解释: "raceacar" 不是回文串。
```

### 示例 3

```
输入: s = " "
输出: true
解释: 在移除非字母数字字符之后，s 是一个空字符串 "" 。
由于空字符串正着反着读都一样，所以是回文串。
```

## 提示

- `1 <= s.length <= 2 * 10^5`
- `s` 仅由可打印的 ASCII 字符组成

## 题目解析

### 核心思路

本题的关键在于两点：

1. **只关心字母和数字字符**，空格、标点等其他字符一律跳过；
2. **字母比较时忽略大小写**，数字不需要处理。

最直接的解法是“先过滤、再判断”：遍历一遍字符串，把所有字母数字字符（统一转小写）收集到新字符串，再判断新字符串是否回文。这样思路简单，但需要 O(n) 的额外空间。

更优的做法是**双指针原地判断**：

- 左指针 `left` 从头开始，右指针 `right` 从尾开始；
- 两个指针各自跳过非字母数字字符，直到都指向有效字符；
- 比较两个有效字符（大写字母先转小写），不等则直接返回 `false`；
- 相等则两个指针同时向中间收缩，直到相遇。

由于 `s` 只含可打印 ASCII 字符，可以**按字节处理**，无需 `unicode` 包：判断范围 `'0'-'9'`、`'a'-'z'`、`'A'-'Z'` 即可，转小写只需对大写字母加上 `'a' - 'A'` 的偏移。

一个容易出错的细节：题目要求数字也算有效字符，所以像 `"0P"` 这样的输入，过滤后是 `"0p"`，不是回文，答案应为 `false`，不要把数字也一起跳过。

本题属于**“对撞双指针”**模板，核心是“跳过无效字符，两端向中间夹击比较”。

### 算法步骤

1. 初始化 `left = 0`，`right = len(s) - 1`。
2. 循环 while `left < right`：
   - 移动 `left`，跳过所有非字母数字字符（注意保持 `left < right`）。
   - 移动 `right`，跳过所有非字母数字字符（同样保持 `left < right`）。
   - 比较 `toLower(s[left])` 与 `toLower(s[right])`，不等则返回 `false`。
   - 相等则 `left++`、`right--`。
3. 循环结束说明所有有效字符两两匹配，返回 `true`。

### 复杂度分析

- **时间复杂度**: O(n)，每个字符最多被两个指针各访问一次，n 为字符串长度。
- **空间复杂度**: O(1)，只使用常数个指针变量，原地判断。

## 代码实现

```go
func IsPalindrome(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
        // 左指针跳过非字母数字字符
        for left < right && !isAlphanumeric(s[left]) {
            left++
        }
        // 右指针跳过非字母数字字符
        for left < right && !isAlphanumeric(s[right]) {
            right--
        }
        // 大小写不敏感比较两个有效字符
        if toLower(s[left]) != toLower(s[right]) {
            return false
        }
        left++
        right--
    }
    return true
}

// isAlphanumeric 判断字节是否为字母或数字
func isAlphanumeric(c byte) bool {
    return c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

// toLower 将大写字母转成小写，其余字符原样返回
func toLower(c byte) byte {
    if c >= 'A' && c <= 'Z' {
        return c + ('a' - 'A')
    }
    return c
}
```

**执行过程示例**（`s = "A man, a plan, a canal: Panama"`）：

```
left/right 各自跳过空格、逗号、冒号等无效字符后，依次比较：
a↔a, m↔m, a↔a, n↔n, a↔a, p↔p, l↔l, a↔a, n↔n, a↔a, c↔c 全部相等
两指针相遇，返回 true
```

**边界示例**（`s = " "`）：左指针右移后与右指针交错，循环直接结束，空有效串视为回文，返回 `true`。
