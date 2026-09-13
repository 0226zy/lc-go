# 494. 目标和 (Target Sum)

## 题目描述

给你一个非负整数数组 `nums` 和一个整数 `target`。

向数组中的每个整数前添加 `'+'` 或 `'-'`，然后串联起所有整数，可以构造一个 **表达式**：

- 例如，`nums = [2, 1]`，可以在 `2` 之前添加 `'+'`，在 `1` 之前添加 `'-'`，然后串联起来得到表达式 `"+2-1"`。

返回可以通过上述方法构造的、运算结果等于 `target` 的不同 **表达式** 的数目。

### 示例 1

```
输入: nums = [1,1,1,1,1], target = 3
输出: 5
解释: 一共有 5 种方法让最终目标和为 3。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
```

### 示例 2

```
输入: nums = [1], target = 1
输出: 1
```

## 提示

- `1 <= nums.length <= 20`
- `0 <= nums[i] <= 1000`
- `0 <= sum(nums[i]) <= 1000`
- `-1000 <= target <= 1000`

## 题目解析

### 核心思路

本题要求给每个数分配 `+` 或 `-` 两种符号之一，统计最终和为 `target` 的方案数。每个数有 2 种选择，本质是枚举所有 `2^n` 种符号组合，是典型的 **回溯（穷举决策树）** 模型。

回溯三要素：

- **路径**：已处理的前 `i` 个数确定的符号，以及对应的累计和 `sum`。
- **选择列表**：对当前数 `nums[i]` 只有两种选择——取 `+nums[i]` 或取 `-nums[i]`。
- **结束条件**：`i == len(nums)`（所有数的符号都已确定），此时若 `sum == target` 则计 1 种方案，否则计 0。

纯回溯是 O(2^n) 的指数枚举。但本题 `n <= 20` 且所有数之和 `<= 1000`，搜索过程中会出现大量重复的 **(下标 i, 当前和 sum)** 状态——从同一个状态出发，后续能得到的方案数是完全相同的。因此叠加 **记忆化搜索**：用哈希表 `memo[(i, sum)]` 缓存每个状态的方案数，命中时直接返回，把指数级搜索压到多项式级别。

注意这里 `sum` 通过函数参数传递，天然实现了“做选择 / 撤销选择”：递归返回后调用方的 `sum` 不变，不需要显式还原现场。

### 算法步骤

1. 建立记忆化表 `memo`，键为 `(i, sum)`，值为该状态下的方案数。
2. 从 `backtrack(0, 0)` 开始搜索：
   - 若 `i == len(nums)`：`sum == target` 返回 1，否则返回 0。
   - 若 `(i, sum)` 命中缓存，直接返回缓存值。
   - 分别递归 `backtrack(i+1, sum+nums[i])`（取加号）和 `backtrack(i+1, sum-nums[i])`（取减号），方案数为两者之和。
   - 将结果写入 `memo` 后返回。
3. 返回 `backtrack(0, 0)` 的结果。

### 复杂度分析

- **时间复杂度**: O(n·S)，其中 n 为数组长度，S 为搜索过程中可达 `(i, sum)` 状态的不同取值个数。由于 `|sum| <= sum(nums) <= 1000`，每个下标 i 最多对应 O(2000) 种和，总状态数约 O(n·2000)，每个状态只计算一次。
- **空间复杂度**: O(n·S)，记忆化表开销；递归栈深度为 O(n)。

## 代码实现

```go
package targetsum

// FindTargetSumWays 目标和
// 给整数数组 nums 中的每个数添加 '+' 或 '-' 符号，统计有多少种符号组合
// 能使表达式结果等于 target。
// 时间复杂度: O(n*S)  n 为数组长度，S 为搜索过程中可达和的不同取值个数（最坏约 2*1000*n 级别）
// 空间复杂度: O(n*S)  记忆化表开销 + O(n) 递归栈深度
func FindTargetSumWays(nums []int, target int) int {
	// memo[(i, sum)] 记录“从下标 i 开始处理、当前累计和为 sum”时的方案数
	type state struct{ i, sum int }
	memo := make(map[state]int)

	var backtrack func(i, sum int) int
	backtrack = func(i, sum int) int {
		// 结束条件：所有数都已确定符号，累计和等于 target 则计 1 种方案
		if i == len(nums) {
			if sum == target {
				return 1
			}
			return 0
		}
		key := state{i, sum}
		if v, ok := memo[key]; ok {
			return v
		}
		// 选择列表：对 nums[i] 分别尝试加号与减号，无需显式还原（sum 通过参数传递）
		ways := backtrack(i+1, sum+nums[i]) + backtrack(i+1, sum-nums[i])
		memo[key] = ways
		return ways
	}
	return backtrack(0, 0)
}
```

**执行过程示例**（`nums = [1, 1], target = 0`）：

```
backtrack(0, 0)
├─ 取 +1 → backtrack(1, 1)
│   ├─ 取 +1 → backtrack(2, 2)：i 越界，sum=2 ≠ 0 → 0
│   └─ 取 -1 → backtrack(2, 0)：i 越界，sum=0 = 0 → 1
│   状态(1,1) 方案数 = 0 + 1 = 1（写入 memo）
└─ 取 -1 → backtrack(1, -1)
    ├─ 取 +1 → backtrack(2, 0)：命中 memo 之外，sum=0 = 0 → 1
    └─ 取 -1 → backtrack(2, -2)：sum=-2 ≠ 0 → 0
    状态(1,-1) 方案数 = 1 + 0 = 1（写入 memo）
总方案数 = 1 + 1 = 2，即 +1-1 和 -1+1 两种
```
