# 54. 螺旋矩阵 (Spiral Matrix)

## 题目描述

给你一个 `m x n` 的矩阵 `matrix`，请按照**顺时针螺旋顺序**，返回矩阵中的所有元素。

### 示例 1

```
输入: matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出: [1,2,3,6,9,8,7,4,5]
```

### 示例 2

```
输入: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
```

### 提示

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 10`
- `-100 <= matrix[i][j] <= 100`

## 题目解析

### 核心思路

螺旋遍历就是**一圈一圈地剥洋葱**：从最外圈开始，沿「上 → 右 → 下 → 左」四个方向依次走，每走完一条边，就把对应的边界往内收缩一格，直到边界交叉、所有元素取完。

这类题属于「**边界收缩模拟**」模板。难点不在思路，而在两个容易踩的坑：

1. **边界收缩的时机**：走完上边一行后 `top++`，走完右边一列后 `right--`——收缩必须紧跟在对应边的遍历之后，否则会用旧边界重复取值。
2. **奇数行/列的收尾**：当矩阵收缩到只剩一行或一列时，「向左」「向上」两步会遍历到已经访问过的元素。判断方法很简单：走完右边一列后，如果 `top > bottom` 说明没有下一行了，跳过向左；同理 `left > right` 时跳过向上。

以示例 1 为例走一遍：

```
第 1 圈: 右 1→2→3，下 6→9，左 8→7，上 4
第 2 圈: 收缩后只剩中心 5，右走完 top>bottom，下走完 right<left，结束
```

### 算法步骤

1. 初始化四条边界：`top=0`、`bottom=m-1`、`left=0`、`right=n-1`，结果切片 `result`。
2. 当 `top <= bottom && left <= right` 时循环：
   - **向右**：从 `left` 到 `right` 取第 `top` 行，`top++`。
   - **向下**：从 `top` 到 `bottom` 取第 `right` 列，`right--`。
   - **向左**（若 `top <= bottom`）：从 `right` 到 `left` 取第 `bottom` 行，`bottom--`。
   - **向上**（若 `left <= right`）：从 `bottom` 到 `top` 取第 `left` 列，`left++`。
3. 循环结束，返回 `result`。

### 复杂度分析

- **时间复杂度**: O(m × n)——每个元素恰好被访问一次。
- **空间复杂度**: O(1)——不计返回结果切片，只使用四个边界变量。

## 代码实现

```go
func SpiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return nil
    }
    m, n := len(matrix), len(matrix[0])
    result := make([]int, 0, m*n)
    top, bottom := 0, m-1
    left, right := 0, n-1
    for top <= bottom && left <= right {
        for j := left; j <= right; j++ { // 向右
            result = append(result, matrix[top][j])
        }
        top++
        for i := top; i <= bottom; i++ { // 向下
            result = append(result, matrix[i][right])
        }
        right--
        if top <= bottom { // 向左（防单行重复）
            for j := right; j >= left; j-- {
                result = append(result, matrix[bottom][j])
            }
            bottom--
        }
        if left <= right { // 向上（防单列重复）
            for i := bottom; i >= top; i-- {
                result = append(result, matrix[i][left])
            }
            left++
        }
    }
    return result
}
```

**执行过程示例**（示例 2，`matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]`）：

```
初始: top=0 bottom=2 left=0 right=3
向右: 取 1 2 3 4, top=1
向下: 取 8 12, right=2
向左: 取 11 10 9, bottom=1
向上: 取 5, left=1
第 2 圈: top=1 bottom=1 left=1 right=2
向右: 取 6 7, top=2 > bottom=1，循环结束
结果: [1,2,3,4,8,12,11,10,9,5,6,7] ✓
```
