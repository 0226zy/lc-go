package earliestfinish

import "math"

// EarliestFinishTime 陆地和水上项目最早完成时间 I
// 从每类项目中恰好体验一个，顺序可任意，返回最早完成时间。
// 时间复杂度: O(n + m)  空间复杂度: O(1)
func EarliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	planA := calc(landStartTime, landDuration, waterStartTime, waterDuration)
	planB := calc(waterStartTime, waterDuration, landStartTime, landDuration)
	if planA < planB {
		return planA
	}
	return planB
}

// calc 计算先玩第一类、再玩第二类的最早完成时间
func calc(s1 []int, d1 []int, s2 []int, d2 []int) int {
	minEnd := math.MaxInt64
	for i := 0; i < len(s1); i++ {
		end := s1[i] + d1[i]
		if end < minEnd {
			minEnd = end
		}
	}

	minFinish := math.MaxInt64
	for j := 0; j < len(s2); j++ {
		start := s2[j]
		if start < minEnd {
			start = minEnd
		}
		finish := start + d2[j]
		if finish < minFinish {
			minFinish = finish
		}
	}
	return minFinish
}
