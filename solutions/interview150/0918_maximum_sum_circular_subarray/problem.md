# 918. 环形子数组的最大和 (Maximum Sum Circular Subarray)

## 题目描述

给定一个长度为 `n` 的**环形整数数组** `nums` ，返回 `nums` 的非空 **子数组** 的最大可能和。

**环形数组** 意味着数组的末端将会与开头相连呈环状。形式上，`nums[i]` 的下一个元素是 `nums[(i + 1) % n]` ， `nums[i]` 的前一个元素是 `nums[(i - 1 + n) % n]` 。

**子数组** 最多只能包含固定缓冲区 `nums` 中的每个元素一次（形式上，对于子数组 `nums[i], nums[i + 1], ..., nums[j]` ，不存在 `i <= k1, k2 <= j` 其中 `k1 % n == k2 % n` ）。

### 示例 1

```
输入: nums = [1,-2,3,-2]
输出: 3
解释: 跨越环形的子数组最优为 [3, -2, 1]（索引 2 -> 3 -> 0），和为 2；
     不跨环的子数组最优为 [3]，和为 3。
     最大子数组和为 3。
```

### 示例 2

```
输入: nums = [5,-3,5]
输出: 10
解释: 从索引 1 -> 2 -> 0 的子数组为 [5, 5, -3]，和为 7。
     从索引 2 -> 0 -> 1 的子数组为 [5, -3, 5]，和为 7。
     两者和都为 7，最大子数组和为 10。
```

### 示例 3

```
输入: nums = [-3,-2,-3]
输出: -2
解释: 子数组 [-2] 的和为 -2，这是最大和的非空子数组。
```

## 提示

- `n == nums.length`
- `1 <= n <= 3 * 10^4`
- `-3 * 10^4 <= nums[i] <= 3 * 10^4`

## 题目解析

### 核心思路

环形数组的最大子数组和只有两种形态：

1. **和普通数组一样**，最优子数组不跨环尾/环首（中间某一段），答案就是普通 Kadane 的结果 `maxKadane`。
2. **跨越环尾和环首**，形如 `[nums[k], ..., nums[n-1], nums[0], ..., nums[j]]`。此时子数组在数组上是「掐头去尾」的两段。因为子数组不能为空且每个元素最多用一次，它等价于：**数组总和 - 中间连续一段（即“最小子数组和”）**。

所以答案是：

```
max( 普通最大子数组和, 数组总和 - 最小子数组和 )
```

**全负数的边界坑**：当所有元素都是负数时，`总和 - 最小子数组和 = 0`（最小子数组是整个数组），而题目要求子数组非空，答案必须是某个负数而不是 0。此时普通 Kadane 的结果就是正确答案，因此需要特判：**当 `maxKadane < 0` 时直接返回它**。

用一次遍历同时维护“以当前元素结尾的最大和”与“以当前元素结尾的最小和”，即可同时得到 `maxKadane` 和 `minKadane`。

### 算法步骤

1. 遍历数组，同时累加 `total`，并分别用 Kadane 维护：
   - `maxCur` / `maxSum`：最大子数组和（结尾必须含当前元素）
   - `minCur` / `minSum`：最小子数组和（结尾必须含当前元素）
2. 若 `maxSum < 0`，说明全是负数，返回 `maxSum`
3. 否则返回 `max(maxSum, total - minSum)`

### 复杂度分析

- **时间复杂度**: O(n)，一次遍历
- **空间复杂度**: O(1)，常数变量

## 代码实现

```go
func MaxSubarraySumCircular(nums []int) int {
    total := 0
    maxCur, maxSum := nums[0], nums[0] // 最大子数组和
    minCur, minSum := nums[0], nums[0] // 最小子数组和
    for _, x := range nums {
        total += x
        if maxCur+x > x {
            maxCur += x
        } else {
            maxCur = x
        }
        if maxCur > maxSum {
            maxSum = maxCur
        }
        if minCur+x < x {
            minCur += x
        } else {
            minCur = x
        }
        if minCur < minSum {
            minSum = minCur
        }
    }
    // 全是负数时，子数组非空，不能取“总和 - 最小和 = 0”
    if maxSum < 0 {
        return maxSum
    }
    // 跨环情形 = 总和 - 中间最小子数组和
    if rest := total - minSum; rest > maxSum {
        maxSum = rest
    }
    return maxSum
}
```

**执行过程示例**（`nums = [5,-3,5]`）：

```
遍历: total=7
  maxSum 的 Kadane: 5 -> max(2,5)=5 -> max(10,5)=10, 得 maxSum=10
  minSum 的 Kadane: 5 -> min(2,-3)=-3 -> min(2,5)=2,   得 minSum=-3
maxSum=10 >= 0, 跨环候选: total - minSum = 7-(-3)=10
答案: max(10, 10) = 10
```
