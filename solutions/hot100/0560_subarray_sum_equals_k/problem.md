# 560. 和为 K 的子数组 (Subarray Sum Equals K)

## 题目描述

给你一个整数数组 `nums` 和一个整数 `k`，请你统计并返回 **该数组中和为 `k` 的子数组的个数**。

子数组是数组中元素的连续非空序列。

### 示例 1

```
输入：nums = [1,1,1], k = 2
输出：2
解释：前缀和 0→1→2→3，preSum=0 出现 1 次，preSum=1 出现 2 次，preSum=2 出现 3 次。
当 preSum[j] - preSum[i] = 2 时，即 preSum[i] = preSum[j] - 2。
答案为 2，对应子数组 [1,1]（从索引 0 开始）和 [1,1]（从索引 1 开始）。
```

### 示例 2

```
输入：nums = [1,2,3], k = 3
输出：2
解释：前缀和 0→1→3→6，答案为 2，对应子数组 [3] 和 [1,2]。
```

## 提示

- `1 <= nums.length <= 2 * 10^4`
- `-1000 <= nums[i] <= 1000`
- `-10^7 <= k <= 10^7`

## 题目解析

### 核心思路

这道题使用 **前缀和 + 哈希表** 的经典解法：

1. **前缀和概念**：定义 `preSum[i]` 为 `nums[0..i-1]` 的和（前缀和数组长度为 n+1）。
   - 子数组 `nums[i..j]` 的和 = `preSum[j+1] - preSum[i]`

2. **问题转化**：寻找满足 `preSum[j] - preSum[i] = k` 的 `(i, j)` 对数，即 `preSum[i] = preSum[j] - k`。

3. **哈希表优化**：用哈希表记录每个前缀和出现的次数，遍历时直接查询 `preSum - k` 的出现次数，累加答案。

### 算法步骤

1. 初始化哈希表 `count`，`count[0] = 1`（空前缀的和为 0）
2. 初始化 `preSum = 0`，`ans = 0`
3. 遍历数组：
   - `preSum += nums[i]`
   - 若 `preSum - k` 在哈希表中，则 `ans += count[preSum - k]`
   - `count[preSum]++`
4. 返回 `ans`

### 复杂度分析

- **时间复杂度**: O(n)，单次遍历
- **空间复杂度**: O(n)，哈希表存储前缀和

## 代码实现

### 最优解：前缀和 + 哈希表

```go
func subarraySum(nums []int, k int) int {
    count := make(map[int]int)
    count[0] = 1 // 空前缀
    ans := 0
    preSum := 0
    for _, num := range nums {
        preSum += num
        if c, ok := count[preSum-k]; ok {
            ans += c
        }
        count[preSum]++
    }
    return ans
}
```

**执行过程示例**（`nums = [1,1,1], k = 2`）：

```
i=0: num=1, preSum=1, preSum-k=-1 不在 map 中, count={0:1, 1:1}, ans=0
i=1: num=1, preSum=2, preSum-k=0 在 map 中(c=1), count={0:1, 1:1, 2:1}, ans=1
i=2: num=1, preSum=3, preSum-k=1 在 map 中(c=1), count={0:1, 1:1, 2:1, 3:1}, ans=2
结果：2
```

**性能优势**：
- O(n) 时间复杂度，比暴力解法的 O(n²) 快得多
- 只需一次遍历，空间换时间
- 哈希表查询 O(1)，整体高效
