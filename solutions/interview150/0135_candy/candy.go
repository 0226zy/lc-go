package candy

// Candy 分发糖果
// n 个孩子站成一排，每人有一个评分 ratings[i]。
// 规则：每个孩子至少分到 1 个糖果；评分更高的孩子必须比相邻孩子分到更多糖果。
// 求满足条件所需的最少糖果总数。
// 时间复杂度: O(n) 两次遍历  空间复杂度: O(n) 辅助数组
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
	// 取 max(candies[i], candies[i+1]+1)，同时累加总数
	total := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] && candies[i+1]+1 > candies[i] {
			candies[i] = candies[i+1] + 1
		}
		total += candies[i]
	}
	return total
}
