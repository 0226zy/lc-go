package containerwithmostwater

// MaxArea 盛最多水的容器
// 给定长度为 n 的整数数组 height，数组元素代表 n 条竖线的高度，
// 找出其中两条线，使它们与 x 轴构成的容器可以容纳最多的水，返回最大储水量。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MaxArea(height []int) int {
	left, right := 0, len(height)-1
	ans := 0

	for left < right {
		// 容器面积 = 底宽 * 较矮一侧的高度（短板效应）
		width := right - left
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		if area := width * h; area > ans {
			ans = area
		}

		// 双指针向中间收缩：移动较矮一侧的指针才有可能找到更高的短板，
		// 移动较高一侧只会让宽度变小而高度受限，面积不可能变大
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return ans
}
