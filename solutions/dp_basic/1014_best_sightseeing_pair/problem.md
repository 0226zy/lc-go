# 1014. 最佳观光组合 (Best Sightseeing Pair)

## 题目描述

给你一个正整数数组 `values`，其中 `values[i]` 表示第 `i` 个观光景点的评分，并且两个景点 `i` 和 `j` 之间的 **距离** 为 `j - i`。

一对景点（`i < j`）组成的观光组合的得分为 `values[i] + values[j] + i - j`，也就是景点的评分之和 **减去** 它们两者之间的距离。

返回一对观光景点能取得的最高分。

### 示例

```
输入: values = [8,1,5,2,6]
输出: 11
解释: 选择 i = 0, j = 2，得分为 values[0] + values[2] + 0 - 2 = 8 + 5 - 2 = 11
```

```
输入: values = [1,2]
输出: 2
```

### 提示

- `2 <= values.length <= 5 * 10^4`
- `1 <= values[i] <= 1000`

## 四步拆解

**关键变形**：把得分公式按下标拆开——

```
values[i] + values[j] + i - j = (values[i] + i) + (values[j] - j)
```

对于固定的右端点 `j`，`(values[j] - j)` 是定值，想让总分最大，只要让 `j` 左边所有 `i` 中 `(values[i] + i)` 最大即可。这就把「枚举一对」变成了「维护前缀最优值」的 DP。

### 1. 状态定义

- `dp[j]` 表示**下标不超过 `j` 的所有位置中，`values[i] + i` 的最大值**，即「以 `j+1` 或更后的位置作为右端点时，左端点能贡献的最大值」。
- 最终答案：`max(dp[j-1] + values[j] - j)`（`j` 从 1 到 n-1），即枚举每个景点作为右端点，取所有配对得分的最大值。

### 2. 边界条件 + 遍历顺序

- base case：`dp[0] = values[0] + 0`（只有下标 0 可选）。
- 遍历顺序：**从 1 到 n-1 正序遍历**。`dp[j]` 依赖 `dp[j-1]`，而下标 `j` 作为右端点配对时需要的正是 `dp[j-1]`（左端点必须严格在 `j` 左边），正序遍历保证算到 `j` 时 `dp[j-1]` 已经求好。

### 3. 状态转移方程

```
dp[j] = max(dp[j-1], values[j] + j)
得分(j) = dp[j-1] + values[j] - j
答案 = max(得分(1), 得分(2), ..., 得分(n-1))
```

推导：考虑位置 `j` 是否加入「候选左端点」集合——要么不加入（沿用 `dp[j-1]`），要么加入（取 `values[j] + j`），两者取大。而当 `j` 充当右端点时，左端点只能来自 `0..j-1`，所以用 `dp[j-1]` 而不是 `dp[j]` 来计算得分。边界特殊情况：数组至少 2 个元素，`j` 从 1 开始枚举，天然满足 `i < j`。

### 4. 样例验证 + 代码实现

手动推导 `values = [8,1,5,2,6]`：

| j | values[j] | values[j]+j | dp[j] | 得分(j) = dp[j-1]+values[j]-j |
| --- | --- | --- | --- | --- |
| 0 | 8 | 8 | 8 | — |
| 1 | 1 | 2 | 8 | 8+1-1 = 8 |
| 2 | 5 | 7 | 8 | 8+5-2 = **11** |
| 3 | 2 | 5 | 8 | 8+2-3 = 7 |
| 4 | 6 | 10 | 10 | 8+6-4 = 10 |

最大得分为 11 ✅（与示例一致，对应 `i=0, j=2`）。

代码与上面的推导一一对应：

```go
// MaxScoreSightseeingPair 最佳观光组合（标准 DP 数组版）
func MaxScoreSightseeingPair(values []int) int {
	n := len(values)
	dp := make([]int, n)
	dp[0] = values[0] // base case：候选左端点只有下标 0
	ans := 0
	for j := 1; j < n; j++ {
		// j 作为右端点：左端点取 0..j-1 中 values[i]+i 的最大值
		if score := dp[j-1] + values[j] - j; score > ans {
			ans = score
		}
		// 更新前缀最优值，供后续右端点使用
		if v := values[j] + j; v > dp[j-1] {
			dp[j] = v
		} else {
			dp[j] = dp[j-1]
		}
	}
	return ans
}
```

## 空间优化版（不使用 DP 数组）

观察转移方程：`dp[j]` 只依赖 `dp[j-1]`，而且计算得分也只需要 `dp[j-1]`。因此完全不需要数组，用一个滚动变量 `best`（代表 `dp[j-1]`，即「当前位置左边 `values[i]+i` 的最大值」）即可，空间从 O(n) 降到 O(1)：

```go
// MaxScoreSightseeingPairOptimized 最佳观光组合（滚动变量空间优化版）
func MaxScoreSightseeingPairOptimized(values []int) int {
	ans := 0
	best := values[0] // 当前 j 左边 values[i]+i 的最大值，即 dp[j-1]
	for j := 1; j < len(values); j++ {
		if score := best + values[j] - j; score > ans {
			ans = score
		}
		if v := values[j] + j; v > best {
			best = v
		}
	}
	return ans
}
```

## 复杂度分析

| 版本 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- |
| DP 数组版 | O(n) | O(n) |
| 空间优化版 | O(n) | O(1) |

其中 `n` 为数组 `values` 的长度。两版都只遍历一次数组；由于配对得分天然是「前缀最优 + 当前值」的结构，滚动变量版是本题的最优解。
