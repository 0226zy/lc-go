# 121. 买卖股票的最佳时机 (Best Time to Buy and Sell Stock)

## 题目描述

给定一个数组 `prices`，它的第 `i` 个元素 `prices[i]` 表示一支股票第 `i` 天的价格。

你只能选择**某一天**买入这只股票，并选择在**未来的某一天**卖出（先买后卖，不能当天买当天卖）。你最多只能完成**一笔**交易。返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 `0`。

### 示例 1

```
输入: prices = [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（价格 = 1）买入，在第 5 天（价格 = 6）卖出，利润 = 6 - 1 = 5。
     注意利润不能是 7 - 1 = 6，因为卖出价格必须大于买入价格。
```

### 示例 2

```
输入: prices = [7,6,4,3,1]
输出: 0
解释: 股价一路下跌，不做任何交易，最大利润为 0。
```

## 提示

- `1 <= prices.length <= 10^5`
- `0 <= prices[i] <= 10^4`

## 题目解析

### 核心思路

这道题的本质是：在数组中找到一对 `(i, j)`，满足 `i < j` 且 `prices[j] - prices[i]` 最大。暴力做法是枚举所有 `(i, j)` 组合，O(n²)，10⁵ 的规模会超时。

换个角度想：**如果我们在第 `j` 天卖出，最优的买入时机一定是第 `0` 到 `j-1` 天中价格最低的那一天**。也就是说，在遍历到第 `j` 天时，只要知道"历史最低价"，就能立刻算出"今天卖出的最大利润"。

于是只需要一遍遍历：

- 用一个变量 `minPrice` 维护**从左到右看到的最低价格**（最佳买入点）；
- 对每一天的价格 `price`，计算"以历史最低价买入、今天卖出"的利润 `price - minPrice`，更新最大利润。

为什么这样一定是对的？因为任何一笔合法交易都是"某个低点买入，之后的某个高点卖出"。遍历到卖出日 `j` 时，`minPrice` 恰好是 `j` 之前所有日子里的最低价，所以 `price[j] - minPrice` 就是所有以 `j` 为卖出日的交易中的最大利润。遍历完所有 `j`，取最大值，必然覆盖全局最优解。

本题属于 **"遍历过程中维护前缀最小值/最大值"** 的经典模型。

### 算法步骤

1. 初始化 `minPrice = prices[0]`，`maxProfit = 0`；
2. 从左到右遍历每一天的价格 `price`：
   - 如果 `price < minPrice`，说明出现了更优的买入时机，更新 `minPrice = price`；
   - 否则计算 `profit = price - minPrice`，如果 `profit > maxProfit` 则更新 `maxProfit`；
3. 返回 `maxProfit`。

### 复杂度分析

- **时间复杂度**: O(n)，只遍历一次数组
- **空间复杂度**: O(1)，只用两个变量

## 代码实现

```go
func MaxProfit(prices []int) int {
	minPrice := prices[0] // 历史最低价
	maxProfit := 0        // 最大利润
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}
```

**执行过程示例**（`prices = [7,1,5,3,6,4]`）：

```
初始: minPrice=7, maxProfit=0
price=7: 不小于 minPrice，利润 0，不更新
price=1: 小于 minPrice，更新 minPrice=1
price=5: 利润 5-1=4 > 0，更新 maxProfit=4
price=3: 利润 3-1=2 < 4，不更新
price=6: 利润 6-1=5 > 4，更新 maxProfit=5
price=4: 利润 4-1=3 < 5，不更新
结果: 5
```
