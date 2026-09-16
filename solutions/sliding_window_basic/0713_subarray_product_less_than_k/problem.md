# 713. 乘积小于 K 的子数组 (Subarray Product Less Than K)

## 题目描述

给你一个整数数组 `nums` 和一个整数 `k`，请你返回乘积 **严格小于** `k` 的 **连续子数组** 的个数。

### 示例 1

```
输入: nums = [10,5,2,6], k = 100
输出: 8
解释: 8 个乘积小于 100 的连续子数组为:
[10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]
注意 [10,5,2] 的乘积为 100，严格小于 k 才合法，因此不计入答案。
```

### 示例 2

```
输入: nums = [1,2,3], k = 0
输出: 0
解释: nums 中所有元素都是正整数，任何子数组的乘积都 >= 1，不可能严格小于 0。
```

## 提示

- `1 <= nums.length <= 3 * 10^4`
- `1 <= nums[i] <= 1000`
- `0 <= k <= 10^6`

## 题目解析

### 核心思路

本题是 **滑动窗口** 的经典应用。与「和为 K 的子数组」不同，本题所有元素都是 **正整数**，这意味着窗口的乘积随窗口扩大单调不减、随窗口收缩单调不增，因此可以用双指针维护一个乘积始终小于 `k` 的窗口 `[left, right]`。

关键观察：

- 用 `right` 逐个向右扫描，把 `nums[right]` 乘入窗口乘积 `product`。
- 当 `product >= k` 时，不断移出 `nums[left]`（即 `product /= nums[left]`）并右移 `left`，直到窗口乘积重新小于 `k`。
- 收缩完成后，窗口 `[left, right]` 满足乘积 < k，此时**以 right 结尾的合法子数组恰好有 `right - left + 1` 个**：`[right, right]`、`[right-1, right]`、…、`[left, right]`。把它们累加到答案中。

为什么 `right - left + 1` 不重不漏？每次只统计「以当前 right 结尾」的合法子数组，不同的 right 对应的子数组互不相同；而以 right 结尾的子数组只要起点落在 `[left, right]` 内，乘积一定小于 `k`（因为正数越乘越大，窗口本身已满足条件，其子窗口只会更小）。

**特判**：当 `k <= 1` 时，由于所有元素都是正整数，任何子数组乘积都 >= 1，不可能严格小于 `k`，直接返回 0（同时也避免后面除法收缩逻辑的边界问题）。

### 算法步骤

1. 若 `k <= 1`，直接返回 0。
2. 初始化 `left = 0`、`product = 1`、`count = 0`。
3. `right` 从 0 遍历到 `n-1`：
   - 把 `nums[right]` 乘入 `product`。
   - 当 `product >= k` 时，循环执行 `product /= nums[left]; left++`，直到 `product < k`。
   - 累加 `right - left + 1` 到 `count`。
4. 返回 `count`。

### 复杂度分析

- **时间复杂度**: O(n)。`right` 和 `left` 各自最多向右移动 n 步，整体为线性扫描。
- **空间复杂度**: O(1)。只使用了常数个变量。

## 代码实现

```go
func NumSubarrayProductLessThanK(nums []int, k int) int {
    // 所有元素为正整数，k <= 1 时不可能存在乘积严格小于 k 的子数组
    if k <= 1 {
        return 0
    }
    count, left, product := 0, 0, 1
    for right := 0; right < len(nums); right++ {
        // 右指针扩张窗口，乘入新元素
        product *= nums[right]
        // 乘积不小于 k，收缩左边界直到恢复合法
        for product >= k {
            product /= nums[left]
            left++
        }
        // 以 right 结尾且乘积小于 k 的子数组共有 right-left+1 个
        count += right - left + 1
    }
    return count
}
```

**执行过程示例**（示例 1：`nums = [10,5,2,6], k = 100`）：

```
right=0: product=10  <100, left=0, count += 0-0+1=1  → count=1   ([10])
right=1: product=50  <100, left=0, count += 1-0+1=2  → count=3   ([5], [10,5])
right=2: product=100 >=100 → 除以 nums[0]=10，left=1, product=10
         product=10 <100, left=1, count += 2-1+1=2  → count=5   ([2], [5,2])
right=3: product=60  <100, left=1, count += 3-1+1=3  → count=8   ([6], [2,6], [5,2,6])
最终结果: 8
```

> 观察：`right=2` 时乘积恰好等于 100，由于题目要求**严格小于**，需要收缩左边界；收缩到 `left=1` 后 `[5,2]` 乘积为 10，恢复合法。
