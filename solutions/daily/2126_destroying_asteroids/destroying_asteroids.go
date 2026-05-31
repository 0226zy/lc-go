package destroyingasteroids

import "sort"

// AsteroidsDestroyed 摧毁小行星
// 排序 + 贪心，从小到大依次尝试摧毁
// 时间复杂度: O(n log n) 排序占主导  空间复杂度: O(1) 原地排序
func AsteroidsDestroyed(mass int, asteroids []int) bool {
	// 1. 排序（从小到大）
	sort.Ints(asteroids)

	// 2. 贪心碰撞
	curMass := mass
	for _, ast := range asteroids {
		if curMass >= ast {
			curMass += ast
		} else {
			return false
		}
	}
	return true
}
