# 0304. 二维区域和检索 - 矩阵不可变 (Range Sum Query 2D - Immutable)

## 题目描述

给定一个二维矩阵 `matrix`，处理以下类型的多个查询：

计算其子矩形范围内元素的**总和**，该子矩阵的**左上角**为 `(row1, col1)`，**右下角**为 `(row2, col2)`。

实现 `NumMatrix` 类：

- `NumMatrix(int[][] matrix)` 用整数矩阵 `matrix` 初始化对象。
- `int sumRegion(int row1, int col1, int row2, int col2)` 返回左上角 `(row1, col1)`、右下角 `(row2, col2)` 所描述的子矩阵的元素总和。

### 示例

```
输入:
["NumMatrix", "sumRegion", "sumRegion", "sumRegion"]
[[[[3, 0, 1, 4, 2],
   [5, 6, 3, 2, 1],
   [1, 2, 0, 1, 5],
   [4, 1, 0, 1, 7],
   [1, 0, 3, 0, 5]]], [2, 1, 4, 3], [1, 1, 2, 2], [1, 2, 2, 4]]
输出: [null, 8, 11, 12]
解释:
NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]);
numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素和)
numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素和)
numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素和)
```

### 提示

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 200`
- `-10^4 <= matrix[i][j] <= 10^4`
- `0 <= row1 <= row2 < m`
- `0 <= col1 <= col2 < n`
- 最多调用 `10^4` 次 `sumRegion`

## 四步拆解

### 1. 状态定义

这是一道典型的**二维前缀和**问题，前缀和本质上就是 DP：

- `pre[i][j]` 表示**以 `(0,0)` 为左上角、`(i-1, j-1)` 为右下角**的矩形区域元素和（下标比矩阵多 1，留出哨兵行列）。
- 最终答案（每次查询）：用容斥原理由 4 个 `pre` 值算出。

### 2. 边界条件 + 遍历顺序

- base case：`pre[0][j] = pre[i][0] = 0`（第 0 行、第 0 列为哨兵，空区域和为 0），这样矩阵第一行/第一列也能套用统一公式。
- 遍历顺序：**按行从上到下、每行从左到右**。`pre[i][j]` 依赖正上方、正左方、左上方三个状态，正序遍历时三者都已求好。

### 3. 状态转移方程

预处理（构造前缀和）：

```
pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i][j-1][j-1] + matrix[i-1][j-1]
```

推导：大矩形 = 上方矩形 + 左方矩形 − 左上重叠部分（被加了两次，减回一次）+ 当前格。这是二维的容斥。

查询时同理容斥：

```
sumRegion(r1, c1, r2, c2)
  = pre[r2+1][c2+1] - pre[r1][c2+1] - pre[r2+1][c1] + pre[r1][c1]
```

推导：目标矩形 = 大矩形 − 上方多余部分 − 左方多余部分 + 左上角被减了两次的部分（加回一次）。边界特殊情况（`row1=0` 或 `col1=0`）由哨兵行列天然覆盖，无需特判。

### 4. 样例验证 + 代码实现

用小矩阵手动推导：

```
matrix = 1 2
         3 4
```

pre 表（含哨兵行/列）：

| pre | j=0 | j=1 | j=2 |
| --- | --- | --- | --- |
| i=0 | 0 | 0 | 0 |
| i=1 | 0 | 1 | 3 |
| i=2 | 0 | 4 | 10 |

- `pre[1][1] = 0 + 0 - 0 + 1 = 1`
- `pre[1][2] = pre[0][2] + pre[1][1] - pre[0][1] + 2 = 3`
- `pre[2][2] = pre[1][2] + pre[2][1] - pre[1][1] + 4 = 3 + 4 - 1 + 4 = 10`（= 1+2+3+4 ✅）

查询 `sumRegion(0, 1, 1, 1)`（即第二列的 2 + 4 = 6）：

```
pre[2][2] - pre[0][2] - pre[2][1] + pre[0][1] = 10 - 0 - 4 + 0 = 6 ✅
```

```go
// NumMatrix 二维前缀和
type NumMatrix struct {
	pre [][]int
}

// Constructor 预处理 O(m*n)
func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	pre := make([][]int, m+1)
	for i := range pre {
		pre[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{pre: pre}
}

// SumRegion 容斥查询 O(1)
func (nm *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	return nm.pre[row2+1][col2+1] - nm.pre[row1][col2+1] - nm.pre[row2+1][col1] + nm.pre[row1][col1]
}
```

## 为什么必须用数组（无法 O(1) 空间）

本题是"多次查询、矩阵不变"的场景，核心思路就是**空间换时间**：预处理时把整个前缀和表存下来，之后每次查询只做 4 次数组访问。如果每次查询都现算区域和（O(m·n) 一次），10⁴ 次查询 × 200×200 的矩阵会到 4×10⁸ 次运算，远超合理范围。前缀和表是算法本体的一部分，不存在"滚动变量优化"的版本——这正是前缀和类题目与普通一维 DP 的区别。

退一步说，如果非要省空间，可以只缓存**行方向的一维前缀和**（每行一个数组），查询时按行累加，空间不变（仍是 O(m·n) 总量）但查询降到 O(m)；这在二维前缀和的基础上没有实际收益，标准做法就是上面的二维前缀和。

## 复杂度分析

| 阶段 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- |
| 预处理（Constructor） | O(m·n) | O(m·n) |
| 单次查询（SumRegion） | O(1) | O(1) |
| 暴力对照：单次查询 | O(m·n) | O(1) |
