# 6. Z 字形变换 (Zigzag Conversion)

## 题目描述

将一个给定字符串 `s` 按照给定的行数 `numRows` ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 `"PAYPALISHIRING"`，行数为 3 时，排列如下：

```
P   A   H   N
A P L S I I G
Y   I   R
```

之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如 `"PAHNAPLSIIGYIR"`。

### 示例 1

```
输入: s = "PAYPALISHIRING", numRows = 3
输出: "PAHNAPLSIIGYIR"
```

### 示例 2

```
输入: s = "PAYPALISHIRING", numRows = 4
输出: "PINALSIGYAHRPI"
解释:
P     I    N
A   L S  I G
Y A   H R
P     I
```

### 示例 3

```
输入: s = "A", numRows = 1
输出: "A"
```

### 提示

- `1 <= s.length <= 1000`
- `s` 由英文字母（小写和大写）、`','` 和 `'.'` 组成
- `1 <= numRows <= 1000`

## 题目解析

### 核心思路

这道题属于**找规律模拟**类题目。表面上是“摆 Z 字形”，但真正考察的是：**不真的去摆这个二维数组，能不能直接算出每一行包含原字符串的哪些下标？**

观察 Z 字形排列会发现，字符下标的变化是有固定周期的：

- 从第 0 行垂直向下走到第 `numRows-1` 行，再斜向上走回第 0 行，这一个来回称为一个**周期**，周期长度 = `2 * numRows - 2`。
- 在每个周期里：
  - **首行（0）和末行（numRows-1）**各只出现 1 次：下标分别为 `i` 和 `i + numRows - 1`（i 为周期起点）。
  - **中间行 `r`** 会出现 2 次：一次是垂直向下经过（下标 `i + r`），一次是斜向上经过（下标 `i + cycle - r`）。

所以只要按行遍历，按上述公式直接取出字符拼接即可，无需模拟二维矩阵。

### 算法步骤

1. 特判：`numRows == 1` 或 `numRows >= len(s)` 时，直接返回原字符串。
2. 计算周期长度 `cycle = 2 * numRows - 2`。
3. 逐行 `row` 从 0 到 `numRows-1` 遍历：
   - 周期内向下字符下标为 `i = row, row + cycle, row + 2*cycle, ...`，依次追加 `s[i]`。
   - 如果 `row` 是中间行，还要追加斜向上的字符 `s[i + cycle - 2*row]`（下标越界则跳过）。
4. 拼接结果就是答案。

### 复杂度分析

- **时间复杂度**: O(n)。每个字符恰好被访问一次
- **空间复杂度**: O(n)。输出字符串本身需要 O(n) 空间

## 代码实现

```go
func Convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	var sb strings.Builder
	sb.Grow(len(s))
	cycle := 2*numRows - 2

	for row := 0; row < numRows; row++ {
		for i := row; i < len(s); i += cycle {
			sb.WriteByte(s[i]) // 向下经过的字符
			// 中间行还有一个向上斜线经过的字符
			if row > 0 && row < numRows-1 {
				j := i + cycle - 2*row
				if j < len(s) {
					sb.WriteByte(s[j])
				}
			}
		}
	}
	return sb.String()
}
```

**执行过程示例**（`s = "PAYPALISHIRING"`, `numRows = 3`，`cycle = 4`）：

```
下标: 0 1 2 3 4 5 6 7 8 9 10 11 12 13
字符: P A Y P A L I S H I R  I  N  G

row=0: 下标 0,4,8,12        → P, A, H, N
row=1: 下标 1,5,9,13 + 斜线 3,7,11 → A, L, I, G, P, S, I
row=2: 下标 2,6,10          → Y, I, R

结果: "PAHNAPLSIIGYIR"
```

**补充：模拟法**

另一种做法是创建 `numRows` 个 `strings.Builder`，遍历字符串时用一个变量记录当前行和方向（到底/到顶时转向），把每个字符追加到对应行的 Builder 里，最后按行拼接。代码更直观但常数更大，面试时两种讲法都可以。
