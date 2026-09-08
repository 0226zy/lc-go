# 74. 搜索二维矩阵 (Search a 2D Matrix)

## 题目描述

给你一个满足下述两条属性的 `m x n` 整数矩阵：

- 每行中的整数从左到右按**非递减顺序**排列。
- 每行的第一个整数大于前一行的最后一个整数。

给你一个整数 `target` ，如果 `target` 在矩阵中，返回 `true` ；否则，返回 `false` 。要求时间复杂度为 O(log(m·n))。

### 示例 1

```
输入: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
输出: true
```

### 示例 2

```
输入: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
输出: false
```

## 提示

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 100`
- `-10^4 <= matrix[i][j], target <= 10^4`

## 题目解析

### 核心思路

两条性质合起来意味着：**把矩阵按行首尾相接，就是一个全局有序的数组**。例如示例 1 等价于数组 `[1,3,5,7,10,11,16,20,23,30,34,60]`。

因此直接套用 **lower_bound 二分模板**：把一维下标 `mid` 映射回二维坐标 `row = mid / n, col = mid % n`，照常二分即可，时间复杂度 O(log(m·n))，完全满足要求。

**扩展思考**（对第 240 题有用）：如果矩阵只满足“每行、每列各自有序”（不满足跨行有序），就不能拍平二分了，要从**左下角（或右上角）**出发走“楼梯”：当前数比 target 大就往左，比 target 小就往下，每次排除一行或一列，O(m+n)。本题的两种解法都给出，二分坐标映射为正解。

### 算法步骤（二分坐标映射）

1. `m, n` 为矩阵行列数，在虚拟一维数组 `[0, m*n)` 上做二分：`left, right = 0, m*n-1`
2. 当 `left <= right`：
   - `mid = (left + right) / 2`，取 `x = matrix[mid/n][mid%n]`
   - `x == target`，返回 `true`
   - `x < target`，`left = mid + 1`
   - `x > target`，`right = mid - 1`
3. 没找到，返回 `false`

### 复杂度分析

- **时间复杂度**: O(log(m·n))，对 m·n 个元素做二分
- **空间复杂度**: O(1)，常数变量

## 代码实现

```go
func SearchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    left, right := 0, m*n-1
    for left <= right {
        mid := (left + right) / 2
        x := matrix[mid/n][mid%n] // 一维下标映射回二维坐标
        if x == target {
            return true
        } else if x < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}
```

**执行过程示例**（`matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3`）：

```
虚拟数组: [1,3,5,7,10,11,16,20,23,30,34,60], m*n=12
left=0, right=11
mid=5:  matrix[1][1]=11 > 3,  right=4
mid=2:  matrix[0][2]=5  > 3,  right=1
mid=0:  matrix[0][0]=1  < 3,  left=1
mid=1:  matrix[0][1]=3  == 3, 返回 true
```
