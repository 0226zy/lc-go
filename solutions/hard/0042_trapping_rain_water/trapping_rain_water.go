package trappingrainwater

// Trap 接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算下雨之后能接多少雨水。
// 双指针解法
// 时间复杂度: O(n)  空间复杂度: O(1)
func Trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	left, right := 0, n-1
	leftMax, rightMax := height[left], height[right]
	ans := 0

	for left < right {
		if leftMax < rightMax {
			// 左侧最大值更小，说明左侧柱子上方的水量由 leftMax 决定
			left++
			if height[left] < leftMax {
				ans += leftMax - height[left]
			} else {
				leftMax = height[left]
			}
		} else {
			// 右侧最大值更小，说明右侧柱子上方的水量由 rightMax 决定
			right--
			if height[right] < rightMax {
				ans += rightMax - height[right]
			} else {
				rightMax = height[right]
			}
		}
	}
	return ans
}

// TrapDP 接雨水（动态规划解法）
// 预先计算每个位置左侧和右侧的最大高度，逐位置累加水量。
// 时间复杂度: O(n)  空间复杂度: O(n)
func TrapDP(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = maxInt(leftMax[i-1], height[i])
	}
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = maxInt(rightMax[i+1], height[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += minInt(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

// TrapStack 接雨水（单调栈解法）
// 维护一个下标单调递减栈，遇到更高的柱子时结算凹槽内的水量。
// 时间复杂度: O(n)  空间复杂度: O(n)
func TrapStack(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	stack := make([]int, 0, n)
	ans := 0
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			width := i - left - 1
			h := minInt(height[left], height[i]) - height[top]
			ans += width * h
		}
		stack = append(stack, i)
	}
	return ans
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
