# 1183. 矩阵中 1 的最大数量

> 难度：困难 ｜ 分类：数学 ｜ 尊享面试 100 题 · 第 100 题
> 链接：https://leetcode.cn/problems/maximum-number-of-ones/

## 题目描述

有一个大小为 `width × height` 的矩阵 `M`，矩阵中的每个单元格只能是 `0` 或 `1`。

矩阵 `M` 中每个大小为 `sideLength × sideLength` 的正方形子矩阵中，`1` 的数量不得超过 `maxOnes`。

请你设计一个算法，计算矩阵中最多可以有多少个 `1`。

### 示例 1

```
输入: width = 3, height = 3, sideLength = 2, maxOnes = 1
输出: 4
解释: 在 3×3 的矩阵中，任意 2×2 子矩阵都不能包含超过 1 个 1。
     一种最优填法（1 可以填在下标 (0,0)、(0,2)、(2,0)、(2,2) 这 4 个位置）：
     1 0 1
     0 0 0
     1 0 1
     可以验证任意 2×2 子矩阵都恰好只有 1 个 1。
```

### 示例 2

```
输入: width = 3, height = 3, sideLength = 2, maxOnes = 2
输出: 6
解释: 一种最优填法：
     1 0 1
     1 0 1
     1 0 1
     任意 2×2 子矩阵都恰好包含 2 个 1，矩阵共 6 个 1。
```

### 提示

- `1 <= width, height <= 100`
- `1 <= sideLength <= min(width, height)`
- `1 <= maxOnes <= sideLength * sideLength`

## 思路解析

### 核心思路

关键观察：把矩阵中的每个格子按下标 `(i % sideLength, j % sideLength)` 分成 `sideLength × sideLength` 个**同余类**。任意一个 `sideLength × sideLength` 的正方形子矩阵，行下标恰好覆盖连续的 `sideLength` 个整数（模 `sideLength` 的余数恰好各出现一次），列下标同理，因此**每个子矩阵恰好包含每个同余类中的一个格子**。

于是题目转化为：每个同余类对应子矩阵中的一个位置，「子矩阵中 1 的数量 ≤ maxOnes」等价于「被填 1 的同余类至多有 maxOnes 个」（把某个同余类的所有格子全部填 1，不会违反约束，因为任何子矩阵只取该类的一个格子）。要最大化 1 的总数，只需统计每个同余类包含的格子数，贪心地把格子数最多的 `maxOnes` 个同余类全部填 1。

### 算法步骤

1. 对每个余数对 `(r, c)`（`0 <= r, c < sideLength`），统计该同余类的格子数：
   `cnt[r][c] = (行下标 i 满足 i % sideLength == r 的行数) × (列下标 j 满足 j % sideLength == c 的列数)`，
   其中行数为 `(height - 1 - r) / sideLength + 1`，列数同理（用宽度计算）。
2. 把所有 `cnt` 值放入切片，从大到小排序。
3. 取前 `maxOnes` 个求和，即为答案。

### 复杂度分析

- **时间复杂度**: O(sideLength² · log sideLength)，只需统计 `sideLength²` 个同余类并排序，与 `width × height` 无关
- **空间复杂度**: O(sideLength²)，存放各同余类的格子数

## 代码实现

```go
func MaximumNumberOfOnes(width, height, sideLength, maxOnes int) int {
	// cnt[i] 记录每个同余类 (r, c) 中的格子数
	cnt := make([]int, 0, sideLength*sideLength)
	for r := 0; r < sideLength; r++ {
		// 行下标 i ∈ [0, height) 且 i % sideLength == r 的行数
		rowCnt := 0
		if r < height {
			rowCnt = (height-1-r)/sideLength + 1
		}
		for c := 0; c < sideLength; c++ {
			colCnt := 0
			if c < width {
				colCnt = (width-1-c)/sideLength + 1
			}
			cnt = append(cnt, rowCnt*colCnt)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cnt)))
	ans := 0
	for i := 0; i < maxOnes; i++ {
		ans += cnt[i]
	}
	return ans
}
```
