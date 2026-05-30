# 53. 最大子数组和 (Maximum Subarray)

## 题目描述

给你一个整数数组 `nums`，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

**子数组** 是数组中的一个连续部分。

### 示例 1

```
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6。
```

### 示例 2

```
输入：nums = [1]
输出：1
```

### 示例 3

```
输入：nums = [5,4,-1,7,8]
输出：23
```

## 提示

- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`

## 题目解析

### 核心思路

这道题是 **动态规划** 的经典入门题，有两种主流解法：

1. **动态规划（Kadane 算法）**：`dp[i]` 表示以 `nums[i]` 结尾的最大子数组和。状态转移：`dp[i] = max(nums[i], dp[i-1] + nums[i])`。空间优化后只需 O(1)。

2. **分治法**：将数组分成左右两半，最大子数组和要么在左半、要么在右半、要么跨越中点。时间复杂度 O(n log n)，空间复杂度 O(log n)（递归栈）。

本题最优解是 **Kadane 算法**，O(n) 时间、O(1) 空间。

### 算法步骤（Kadane 算法）

1. 初始化 `maxSum = nums[0]`（答案），`curSum = 0`（当前子数组和）
2. 遍历数组：
   - `curSum = max(nums[i], curSum + nums[i])`（要么从头开始，要么继续累加）
   - `maxSum = max(maxSum, curSum)`（更新答案）
3. 返回 `maxSum`

### 复杂度分析

- **时间复杂度**: O(n)，单次遍历
- **空间复杂度**: O(1)，只使用常数额外空间

## 代码实现

### 最优解：Kadane 算法（动态规划空间优化）

```go
func maxSubArray(nums []int) int {
    maxSum := nums[0]
    curSum := 0
    for _, num := range nums {
        if curSum > 0 {
            curSum += num
        } else {
            curSum = num
        }
        if curSum > maxSum {
            maxSum = curSum
        }
    }
    return maxSum
}
```

**执行过程示例**（`nums = [-2,1,-3,4,-1,2,1,-5,4]`）：

```
i=0: num=-2, curSum=-2, maxSum=-2
i=1: num=1,  curSum=1,  maxSum=1
i=2: num=-3, curSum=-2, maxSum=1
i=3: num=4,  curSum=4,  maxSum=4
i=4: num=-1, curSum=3,  maxSum=4
i=5: num=2,  curSum=5,  maxSum=5
i=6: num=1,  curSum=6,  maxSum=6
i=7: num=-5, curSum=1,  maxSum=6
i=8: num=4,  curSum=5,  maxSum=6
结果：6
```

**性能优势**：
- O(n) 时间复杂度，比暴力解法 O(n²) 快得多
- O(1) 空间复杂度，原地计算，零额外分配
- 单次遍历，缓存友好
