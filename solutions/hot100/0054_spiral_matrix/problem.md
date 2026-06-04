# 54. 螺旋矩阵 (Spiral Matrix)

## 题目描述

给你一个 `m` 行 `n` 列的矩阵 `matrix` ，请按照 **顺时针螺旋顺序** ，返回矩阵中的所有元素。

### 示例 1

```
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
```

### 示例 2

```
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]
```

## 提示

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 10`
- `-100 <= matrix[i][j] <= 100`

## 题目解析

### 核心思路

这道题要求按照顺时针螺旋顺序返回矩阵中的所有元素。

**关键观察**：螺旋遍历可以分解为四个方向：右→下→左→上，每完成一圈就缩小边界。

**最优解法：边界收缩法**
1. 定义四个边界：top（上边界）、bottom（下边界）、left（左边界）、right（右边界）
2. 按照右→下→左→上的顺序遍历：
   - 从左到右遍历上边界，然后 top++
   - 从上到下遍历右边界，然后 right--
   - 如果 top <= bottom，从右到左遍历下边界，然后 bottom--
   - 如果 left <= right，从下到上遍历左边界，然后 left++
3. 重复步骤2，直到所有元素都被遍历

### 算法步骤（边界收缩法）

1. 初始化边界：top = 0, bottom = m-1, left = 0, right = n-1
2. 当 top <= bottom 且 left <= right 时：
   - 从左到右遍历上边界（第 top 行）
   - top++（上边界下移）
   - 从上到下遍历右边界（第 right 列）
   - right--（右边界左移）
   - 如果 top <= bottom，从右到左遍历下边界（第 bottom 行）
   - bottom--（下边界上移）
   - 如果 left <= right，从下到上遍历左边界（第 left 列）
   - left++（左边界右移）

### 复杂度分析

- **时间复杂度**: O(mn)，遍历矩阵中的所有元素
- **空间复杂度**: O(1)，除了输出数组外，只使用常数额外空间

## 代码实现

### 最优解：边界收缩法

```go
func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }
    
    m, n := len(matrix), len(matrix[0])
    result := make([]int, 0, m*n)
    
    top, bottom := 0, m-1
    left, right := 0, n-1
    
    for top <= bottom && left <= right {
        // 从左到右遍历上边界
        for j := left; j <= right; j++ {
            result = append(result, matrix[top][j])
        }
        top++
        
        // 从上到下遍历右边界
        for i := top; i <= bottom; i++ {
            result = append(result, matrix[i][right])
        }
        right--
        
        // 如果还有行，从右到左遍历下边界
        if top <= bottom {
            for j := right; j >= left; j-- {
                result = append(result, matrix[bottom][j])
            }
            bottom--
        }
        
        // 如果还有列，从下到上遍历左边界
        if left <= right {
            for i := bottom; i >= top; i-- {
                result = append(result, matrix[i][left])
            }
            left++
        }
    }
    
    return result
}
```

**执行过程示例**（`matrix = [[1,2,3],[4,5,6],[7,8,9]]`）：

```
初始: top=0, bottom=2, left=0, right=2, result=[]

第1圈:
从左到右遍历上边界: [1,2,3] → result=[1,2,3], top=1
从上到下遍历右边界: [6,9] → result=[1,2,3,6,9], right=1
从右到左遍历下边界: [8,7] → result=[1,2,3,6,9,8,7], bottom=1
从下到上遍历左边界: [4] → result=[1,2,3,6,9,8,7,4], left=1

第2圈:
从左到右遍历上边界: [5] → result=[1,2,3,6,9,8,7,4,5], top=2
此时 top=2 > bottom=1，循环结束

结果: [1,2,3,6,9,8,7,4,5]
```

**性能优势**：
- O(mn) 时间复杂度，一次遍历完成
- O(1) 空间复杂度（除输出数组外），原地遍历
- 代码逻辑清晰，易于理解和实现
