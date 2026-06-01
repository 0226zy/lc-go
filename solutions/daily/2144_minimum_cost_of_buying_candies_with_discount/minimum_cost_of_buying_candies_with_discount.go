package mincostcandies

import "sort"

// MinimumCost 打折购买糖果的最小开销
// 每购买两颗糖果，第三颗免费（免费糖果价格不超过两颗中较便宜的那颗）。
// 返回购买所有糖果的最小开销。
// 时间复杂度: O(n log n)  空间复杂度: O(1)
func MinimumCost(cost []int) int {

	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})

	ret := 0
	for i, c := range cost {
		if (i+1)%3 != 0 {
			ret += c
		}
	}
	return ret
}
