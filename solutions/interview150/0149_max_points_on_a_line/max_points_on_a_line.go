package maxpointsontaline

// MaxPoints 直线上最多的点数
// 给定平面上的点集 points，返回落在同一条直线上的最大点数。
// 以每个点为基准，统计其余点相对该点的斜率；斜率用 gcd 约分后的 (dx, dy) 表示并哈希。
// 注意：垂直线（dx=0）、水平线（dy=0）、重合点需要单独处理；符号统一使 (dx,dy) 唯一。
// 时间复杂度: O(n^2)  空间复杂度: O(n)
func MaxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	ans := 1
	for i := 0; i < n; i++ {
		// 只统计 i 之后的点，避免重复枚举；斜率 -> 共线点数
		slopes := make(map[[2]int]int)
		duplicates := 0 // 与 points[i] 完全重合的点数
		localMax := 0   // 过 points[i] 的某条直线上（不含 i 与重合点）的最多其它点数

		for j := i + 1; j < n; j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]
			if dx == 0 && dy == 0 {
				duplicates++
				continue
			}
			g := gcd(abs(dx), abs(dy))
			dx /= g
			dy /= g
			// 统一符号：保证 dx > 0；若垂直（dx==0）则保证 dy > 0
			if dx < 0 || (dx == 0 && dy < 0) {
				dx = -dx
				dy = -dy
			}
			key := [2]int{dx, dy}
			slopes[key]++
			if slopes[key] > localMax {
				localMax = slopes[key]
			}
		}

		// +1 计入基准点 i，再加上重合点
		if cur := localMax + duplicates + 1; cur > ans {
			ans = cur
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
