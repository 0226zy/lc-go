package trappingrainwater

// Trap 接雨水
// 给定 n 个非负整数表示柱状图各柱子的高度，下雨之后能接多少雨水。
// 时间复杂度: O(n) 单次遍历  空间复杂度: O(1) 常数空间
func Trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	total := 0
	for left < right {
		// 矮的一侧决定水位；移动矮侧指针，若当前柱比该侧最高矮则能接水
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				total += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				total += rightMax - height[right]
			}
			right--
		}
	}
	return total
}

// TrapDP 接雨水的动态规划解法（对照）
// 对每个位置，雨水量 = min(左侧最大高度, 右侧最大高度) - height[i]（若为正）。
// 时间复杂度: O(n) 三次遍历  空间复杂度: O(n) 两个辅助数组
func TrapDP(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	total := 0
	for i := 1; i < n-1; i++ {
		total += min(leftMax[i], rightMax[i]) - height[i]
	}
	return total
}

// TrapStack 接雨水的单调栈解法（对照）
// 栈中保存高度递增的下标，遇到更高的柱子时，以栈顶为谷底弹出并计算积水。
// 时间复杂度: O(n)  空间复杂度: O(n) 栈
func TrapStack(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	stack := make([]int, 0, n)
	total := 0
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			bottom := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			leftIdx := stack[len(stack)-1]
			width := i - leftIdx - 1
			boundedHeight := min(height[leftIdx], height[i]) - bottom
			total += width * boundedHeight
		}
		stack = append(stack, i)
	}
	return total
}
