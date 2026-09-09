# 0152. 乘积最大子数组 (Maximum Product Subarray)

## 题目描述

给你一个整数数组 `nums`，请你找出数组中乘积最大的**非空连续子数组**（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

测试用例的答案是一个 **32 位**整数。

### 示例

```
输入: nums = [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
```

```
输入: nums = [-2,0,-1]
输出: 0
解释: 结果不能为 2，因为 [-2,-1] 不是连续子数组。选择子数组 [0]，乘积为 0。
```

### 提示

- `1 <= nums.length <= 2 * 10^4`
- `-10 <= nums[i] <= 10`
- `nums` 的任何子数组的乘积都保证是一个 **32 位**整数

## 四步拆解

### 1. 状态定义

如果照抄 0053「最大子数组和」的定义「`dp[i]` = 以 `nums[i]` 结尾的最大乘积」，会发现转移不出来：`[-2, 3, -4]` 中，前一个最大乘积 `-2×3 = -6` 再乘 `-4` 反而变成最大的 `24`。**负数会把最大变最小、最小变最大**，所以必须同时记录两个状态：

- `dpMax[i]` 表示**以 `nums[i]` 结尾**的连续子数组的**最大**乘积；
- `dpMin[i]` 表示**以 `nums[i]` 结尾**的连续子数组的**最小**乘积。

最终答案：`max(dpMax[0], dpMax[1], ..., dpMax[n-1])`。注意答案不是 `dpMax[n-1]`——最大乘积的子数组可以在任何位置结尾，所以要全程记录最大值。

### 2. 边界条件 + 遍历顺序

- base case：`dpMax[0] = dpMin[0] = nums[0]`（只有一个元素时，最大、最小乘积都是它自己）。
- 遍历顺序：**从 1 到 n-1 正序遍历**。`dpMax[i]` 和 `dpMin[i]` 只依赖 `i-1` 处的两个状态，正序保证子问题先求解。

### 3. 状态转移方程

```
dpMax[i] = max(nums[i], dpMax[i-1] * nums[i], dpMin[i-1] * nums[i])
dpMin[i] = min(nums[i], dpMin[i-1] * nums[i], dpMax[i-1] * nums[i])
```

推导：考虑以 `nums[i]` 结尾的乘积最大子数组，它只有两类形态——

1. **单独成段**：就是 `nums[i]` 自己（前面的乘积不如不要，比如前面乘积为 0 或异号）；
2. **接在前一段后面**：值为「前一段乘积 × `nums[i]`」。`nums[i]` 为正时接 `dpMax[i-1]` 最大；`nums[i]` 为负时反而是接 `dpMin[i-1]`（最小负值翻身变最大）。

不用判断 `nums[i]` 的符号，直接把三个候选（自己、接最大、接最小）取 `max` 即可；`dpMin[i]` 同理取 `min`。边界特殊情况：`nums[i] = 0` 时两个状态都被重置为 0（任何段乘以 0 都是 0），公式天然覆盖，无需特判。

### 4. 样例验证 + 代码实现

手动推导 `nums = [2,3,-2,4]`：

| i | 0 | 1 | 2 | 3 |
| --- | --- | --- | --- | --- |
| nums[i] | 2 | 3 | -2 | 4 |
| dpMax[i] | 2 | 6 | -2 | 4 |
| dpMin[i] | 2 | 3 | -12 | -48 |

- `i=1`：`dpMax[1] = max(3, 2×3, 2×3) = 6`，`dpMin[1] = 3`。
- `i=2`：候选为 `-2`、`6×(-2) = -12`、`3×(-2) = -6` → `dpMax[2] = -2`（单独成段），`dpMin[2] = -12`。
- `i=3`：候选为 `4`、`(-2)×4 = -8`、`(-12)×4 = -48` → `dpMax[3] = 4`，`dpMin[3] = -48`。
- 全程 `dpMax` 的最大值为 `max(2, 6, -2, 4) = 6` ✅（与示例一致）。

代码与上面的推导一一对应：

```go
// MaxProduct 乘积最大子数组（标准 DP 数组版）
func MaxProduct(nums []int) int {
	n := len(nums)
	dpMax := make([]int, n)
	dpMin := make([]int, n)
	dpMax[0], dpMin[0] = nums[0], nums[0]
	ans := nums[0]
	for i := 1; i < n; i++ {
		dpMax[i] = max(nums[i], max(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]))
		dpMin[i] = min(nums[i], min(dpMin[i-1]*nums[i], dpMax[i-1]*nums[i]))
		ans = max(ans, dpMax[i])
	}
	return ans
}
```

## 空间优化版（不使用 DP 数组）

观察转移方程：`dpMax[i]`、`dpMin[i]` 只依赖 `i-1` 处的两个值，更早的状态不会再被用到。因此用两个滚动变量 `curMax`、`curMin` 代替整个数组即可，空间从 O(n) 降到 O(1)。注意更新时要先用临时变量保存旧的 `curMax`，否则算 `curMin` 时会读到被覆盖的新值：

```go
// MaxProductOptimized 乘积最大子数组（滚动变量空间优化版）
func MaxProductOptimized(nums []int) int {
	curMax, curMin, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		prevMax := curMax // 先保存旧值，再更新
		curMax = max(nums[i], max(curMax*nums[i], curMin*nums[i]))
		curMin = min(nums[i], min(curMin*nums[i], prevMax*nums[i]))
		ans = max(ans, curMax)
	}
	return ans
}
```

## 复杂度分析

| 版本 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- |
| DP 数组版 | O(n) | O(n) |
| 空间优化版 | O(n) | O(1) |
