package containerwithmostwater

// MaxArea 盛最多水的容器
// 给定一个长度为 n 的整数数组 height，找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		width := right - left
		var h int
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		area := width * h
		if area > res {
			res = area
		}
	}
	return res
}
