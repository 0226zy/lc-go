# 0918. 环形子数组的最大和 (Maximum Sum Circular Subarray)

## 题目描述

给定一个长度为 `n` 的**环形**整数数组 `nums`（`nums[n-1]` 的下一个元素是 `nums[0]`），返回 `nums` 的非空**子数组**的最大可能和。

环形数组意味着子数组的端点可以相连，例如子数组 `nums[i], nums[i+1], ..., nums[n-1], nums[0], ..., nums[j]` 是合法的，但每个元素**最多只能使用一次**。

### 示例

```
输入: nums = [1,-2,3,-2]
输出: 3
解释: 子数组 [3] 有最大和 3。
```

```
输入: nums = [5,-3,5]
输出: 10
解释: 子数组 [5,5]（首尾相接）有最大和 5 + 5 = 10。
```

```
输入: nums = [-3,-2,-3]
输出: -2
解释: 子数组 [-2] 有最大和 -2（子数组不能为空，只能选最大的单个元素）。
```

### 提示

- `n == nums.length`
- `1 <= n <= 3 * 10^4`
- `-3 * 10^4 <= nums[i] <= 3 * 10^4`

## 四步拆解

**关键观察**：环形子数组只有两种形态：

1. **不跨首尾**：就是普通子数组，答案是 0053 最大子数组和（Kadane 算法）的结果 `maxSum`；
2. **跨越首尾**：子数组 = 整段数组 − 中间挖掉的一段。要让"剩下的和"最大，就要让"挖掉的中间一段"和最小，即 `total - minSum`（`total` 为数组总和，`minSum` 为最小子数组和）。

所以答案 = `max(maxSum, total - minSum)`。唯一例外：如果数组**全为负数**，`minSum == total` 会让 `total - minSum = 0`，对应"挖掉全部"的空子数组，不合法；此时答案就是不跨界的 `maxSum`（最大的那个负数）。

### 1. 状态定义

借用 Kadane 的"以 i 结尾"定义，同时维护最大、最小两套状态：

- `maxDP[i]` 表示**以 `nums[i]` 结尾**的（非空）子数组的**最大和**；
- `minDP[i]` 表示**以 `nums[i]` 结尾**的（非空）子数组的**最小和**。

最终答案：令 `maxSum = max(maxDP)`、`minSum = min(minDP)`、`total = sum(nums)`，则
- 若 `maxSum < 0`（全负数），答案为 `maxSum`；
- 否则答案为 `max(maxSum, total - minSum)`。

### 2. 边界条件 + 遍历顺序

- base case：`maxDP[0] = minDP[0] = nums[0]`（以第一个元素结尾的子数组只有它自己）。
- 遍历顺序：**从 1 到 n-1 正序遍历**。`maxDP[i]` 只依赖 `maxDP[i-1]`，`minDP[i]` 只依赖 `minDP[i-1]`，正序保证前一个位置先求好。

### 3. 状态转移方程

```
maxDP[i] = max(nums[i], maxDP[i-1] + nums[i])   // 要么另起一段，要么接在前一段后面
minDP[i] = min(nums[i], minDP[i-1] + nums[i])   // 同理求最小
```

推导：考虑以 `nums[i]` 结尾的子数组，只有两种选择——**从 `nums[i]` 重新开始一段**（和为 `nums[i]`），或**接在以 `i-1` 结尾的最优子数组后面**（和为 `maxDP[i-1] + nums[i]`）。取较大者即最大和；同理取较小者即最小和。

边界特殊情况：全负数组时 `total - minSum = 0` 对应空子数组，必须排除，直接返回 `maxSum`（因为此时所有元素为负，`maxSum` 就是最大单元素）。

### 4. 样例验证 + 代码实现

手动推导 `nums = [5, -3, 5]`，`total = 7`：

| i | nums[i] | maxDP[i] | minDP[i] |
| --- | --- | --- | --- |
| 0 | 5 | 5 | 5 |
| 1 | -3 | max(-3, 5-3) = 2 | min(-3, 5-3) = -3 |
| 2 | 5 | max(5, 2+5) = 7 | min(5, -3+5) = 2 |

- `maxSum = 7`，`minSum = -3`；
- `total - minSum = 7 - (-3) = 10`（对应跨界子数组 [5,5]，挖掉中间的 -3）；
- 答案 `max(7, 10) = 10`，与示例输出一致 ✅

代码与上面的推导一一对应：

```go
// MaxSubarraySumCircular 环形子数组的最大和（标准 DP 数组版）
func MaxSubarraySumCircular(nums []int) int {
	n := len(nums)
	maxDP := make([]int, n) // 以 i 结尾的子数组最大和
	minDP := make([]int, n) // 以 i 结尾的子数组最小和
	maxDP[0], minDP[0] = nums[0], nums[0]
	total, maxSum, minSum := nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		maxDP[i] = max(nums[i], maxDP[i-1]+nums[i])
		minDP[i] = min(nums[i], minDP[i-1]+nums[i])
		total += nums[i]
		maxSum = max(maxSum, maxDP[i])
		minSum = min(minSum, minDP[i])
	}
	if maxSum < 0 { // 全负数：total-minSum 会得到空子数组，不合法
		return maxSum
	}
	return max(maxSum, total-minSum)
}
```

## 空间优化版（不使用 DP 数组）

观察转移方程：`maxDP[i]`、`minDP[i]` 都只依赖 `i-1` 处的状态，且我们只需要它们的整体最大值/最小值。因此用几个滚动变量代替数组即可：用 `curMax`、`curMin` 表示"以当前元素结尾"的最大/最小和，用 `maxSum`、`minSum` 记录历史最优，空间从 O(n) 降到 O(1)：

```go
// MaxSubarraySumCircularOptimized 环形子数组的最大和（滚动变量空间优化版）
func MaxSubarraySumCircularOptimized(nums []int) int {
	total, curMax, curMin := nums[0], nums[0], nums[0]
	maxSum, minSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curMax = max(nums[i], curMax+nums[i])
		curMin = min(nums[i], curMin+nums[i])
		maxSum = max(maxSum, curMax)
		minSum = min(minSum, curMin)
		total += nums[i]
	}
	if maxSum < 0 {
		return maxSum
	}
	return max(maxSum, total-minSum)
}
```

## 复杂度分析

| 版本 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- |
| DP 数组版 | O(n) | O(n) |
| 空间优化版 | O(n) | O(1) |

**进阶**：跨界形态还有另一种经典解法——「前缀和 + 单调队列」维护滑动窗口最小前缀和，同样可以 O(n) 求解，是 0053 系列向环形/定长子数组推广的重要技巧。
