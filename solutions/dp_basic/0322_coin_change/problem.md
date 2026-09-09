# 0322. 零钱兑换 (Coin Change)

## 题目描述

给你一个整数数组 `coins`，表示不同面额的硬币；以及一个整数 `amount`，表示总金额。

计算并返回可以凑成总金额所需的**最少的硬币个数**。如果没有任何一种硬币组合能组成总金额，返回 `-1`。

你可以认为每种硬币的数量是无限的。

### 示例

```
输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
```

```
输入: coins = [2], amount = 3
输出: -1
```

```
输入: coins = [1], amount = 0
输出: 0
```

### 提示

- `1 <= coins.length <= 12`
- `1 <= coins[i] <= 2^31 - 1`
- `0 <= amount <= 10^4`

## 四步拆解

### 1. 状态定义

- `dp[i]` 表示**凑成金额 `i` 所需的最少硬币个数**。
- 最终答案：`dp[amount]`。如果 `dp[amount]` 仍是初始的「不可能」标记，说明凑不出，返回 `-1`。

为什么这样定义？凑金额 `i` 的最后一步一定是「放了某枚面额为 `c` 的硬币」，那么在此之前需要凑出金额 `i - c`。也就是说大问题可以拆成「金额更小的同类问题」，一维数组足够描述全部状态。这其实是**完全背包**求最小物品数的经典模型（每种硬币可无限使用）。

### 2. 边界条件 + 遍历顺序

- base case：
  - `dp[0] = 0`（金额 0 不需要任何硬币）；
  - 其余 `dp[i]` 初始化为一个「不可能」的标记。这里用 `amount + 1`：因为最多用 `amount` 枚面额 1 的硬币就能凑出 `amount`，任何合法答案都不会超过 `amount`，所以 `amount + 1` 一定比所有合法答案大，可以安全地当作无穷大参与取 min。
- 遍历顺序：**外层正序遍历金额 `i` 从 1 到 `amount`，内层遍历每种硬币**。`dp[i]` 依赖的 `dp[i-c]` 下标都比 `i` 小，正序保证它们已经求好。本题求的是「最少个数」而非「方案数」，取 min 运算与硬币遍历顺序无关，所以内外层可以互换，这里统一用外层金额、内层硬币的写法。

### 3. 状态转移方程

```
dp[i] = min(dp[i], dp[i-c] + 1)   对所有满足 c <= i 的硬币 c
```

推导：考虑凑金额 `i` 的最后一枚硬币，它必然是某种面额 `c`（且 `c <= i`）。放了这枚硬币之前，需要凑出金额 `i - c`，最少要 `dp[i-c]` 枚，加上这最后一枚就是 `dp[i-c] + 1`。枚举所有可能的 `c`，取最小值即可。

边界特殊情况：`dp[i-c]` 等于 `amount + 1`（表示 `i-c` 凑不出）时，`dp[i-c] + 1` 为 `amount + 2`，在取 min 时天然被淘汰，不会污染结果，无需特判。所有硬币都试完 `dp[i]` 仍是 `amount + 1`，说明金额 `i` 凑不出。

### 4. 样例验证 + 代码实现

手动推导 `coins = [1,2,5], amount = 11` 的部分过程：

| i | 0 | 1 | 2 | 3 | 4 | 5 | ... | 10 | 11 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| dp[i] | 0 | 1 | 1 | 2 | 2 | 1 | ... | 2 | 3 |

- `dp[1] = dp[0]+1 = 1`（只能放面额 1）
- `dp[2] = min(dp[1]+1, dp[0]+1) = 1`（放一枚面额 2）
- `dp[3] = min(dp[2]+1, dp[1]+1) = 2`（2+1）
- `dp[5] = min(dp[4]+1, dp[3]+1, dp[0]+1) = 1`（放一枚面额 5）
- `dp[10] = dp[5]+1 = 2`（5+5）
- `dp[11] = min(dp[10]+1, dp[9]+1, dp[6]+1) = 3`（5+5+1）✅（与示例一致）

代码与上面的推导一一对应：

```go
// CoinChange 零钱兑换（标准 DP 数组版）
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // amount+1 大于任何合法答案，当作无穷大
	}
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if c <= i {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1 // 凑不出
	}
	return dp[amount]
}
```

## 备选版：广度优先搜索（BFS）

把问题看成一张图：每个金额是一个节点，金额 `x` 可以用一枚硬币 `c` 走到 `x - c`。那么「凑出 `amount` 的最少硬币数」就是从 `amount` 到 `0` 的**最短路径长度**，这正是 BFS 的拿手好戏——从 `amount` 出发逐层扩展，第一次到达 0 时的层数就是答案。BFS 同样需要一个 `visited` 数组防止重复入队，所以空间并没有省，只是换了一种视角，适合作为对照理解「DP 本质上是 DAG 上的最短/最长路径」。

```go
// CoinChangeAlternative 零钱兑换（BFS 版）
func CoinChangeAlternative(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	visited := make([]bool, amount+1)
	queue := []int{amount}
	visited[amount] = true
	steps := 0
	for len(queue) > 0 {
		steps++ // 每扩展一层，相当于多用一枚硬币
		next := []int{}
		for _, cur := range queue {
			for _, c := range coins {
				rem := cur - c
				if rem == 0 {
					return steps // 第一次到达 0，层数即最少硬币数
				}
				if rem > 0 && !visited[rem] {
					visited[rem] = true
					next = append(next, rem)
				}
			}
		}
		queue = next
	}
	return -1
}
```

## 复杂度分析

| 版本 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- |
| DP 数组版 | O(amount × len(coins)) | O(amount) |
| BFS 版 | O(amount × len(coins)) | O(amount) |
