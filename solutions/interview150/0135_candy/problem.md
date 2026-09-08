# 135. 分发糖果 (Candy)

## 题目描述

`n` 个孩子站成一排，每个孩子有一个**评分** `ratings[i]`。

按照以下要求，给每个孩子分发糖果：

- 每个孩子至少分配到 `1` 个糖果。
- 相邻两个孩子评分更高的那一个必须分到更多的糖果。

求最少需要准备多少糖果。

### 示例 1

```
输入: ratings = [1,0,2]
输出: 5
解释: 可以分别给这三个孩子分发 2、1、2 颗糖果。
```

### 示例 2

```
输入: ratings = [1,2,2]
输出: 4
解释: 可以分别给这三个孩子分发 1、2、1 颗糖果。
      第三个孩子只得到 1 颗糖果，满足题目要求。
```

## 提示

- `n == ratings.length`
- `1 <= n <= 2 * 10^4`
- `0 <= ratings[i] <= 2 * 10^4`

## 题目解析

### 核心思路

两个约束方向：孩子既要和**左邻**比，又要和**右邻**比。一个位置同时受两边影响，很难一次处理。

经典的拆解办法：**把「同时满足两边」拆成「分别满足两边」**。

- 只考虑左邻：从左往右扫一遍，若 `ratings[i] > ratings[i-1]`，就令 `candies[i] = candies[i-1] + 1`，否则给 1 颗。这样所有「比左边高」的孩子都满足要求。
- 只考虑右邻：从右往左扫一遍，若 `ratings[i] > ratings[i+1]`，就令 `candies[i] = max(candies[i], candies[i+1] + 1)`。这样所有「比右边高」的孩子也满足要求。

每个孩子的糖果数取的是两种要求中的**较大值**，所以两个约束同时成立；而在每个方向上我们都只在「必须加」时才加 1，因此总数最少。

这属于**贪心 + 两次遍历**模型，与「135 类左右约束」问题（如加油站、除自身乘积）同宗：一次遍历只负责一个方向，互不干扰。

### 算法步骤

1. 创建数组 `candies`，所有元素初始化为 1。
2. **从左往右**遍历：若 `ratings[i] > ratings[i-1]`，令 `candies[i] = candies[i-1] + 1`。
3. **从右往左**遍历：若 `ratings[i] > ratings[i+1]`，令 `candies[i] = max(candies[i], candies[i+1] + 1)`。
4. 累加 `candies` 求和返回。

### 复杂度分析

- **时间复杂度**: O(n)，两次遍历
- **空间复杂度**: O(n)，辅助数组 `candies`

## 代码实现

```go
func Candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	// 初始化每人 1 个糖果
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	// 第一遍（从左往右）：保证评分高于左邻的孩子糖果更多
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// 第二遍（从右往左）：保证评分高于右邻的孩子糖果更多
	total := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] && candies[i+1]+1 > candies[i] {
			candies[i] = candies[i+1] + 1
		}
		total += candies[i]
	}
	return total
}
```

**执行过程示例**（`ratings = [1,0,2]`）：

```
初始: candies = [1, 1, 1]

第一遍（从左往右）:
  i=1: ratings[1]=0 < ratings[0]=1，不变
  i=2: ratings[2]=2 > ratings[1]=0，candies[2] = 1 + 1 = 2
  得到 [1, 1, 2]

第二遍（从右往左）:
  i=1: ratings[1]=0 < ratings[2]=2，不变
  i=0: ratings[0]=1 > ratings[1]=0，candies[0] = max(1, 1+1) = 2
  得到 [2, 1, 2]

总数: 2 + 1 + 2 = 5
```
