package trappingrainwater

// Trap 接雨水（标准 DP 数组版）
// leftMax[i]、rightMax[i] 分别表示下标 i 左/右侧（含自身）的最高柱子，
// 位置 i 的接水量 = min(leftMax[i], rightMax[i]) - height[i]。
// 时间复杂度: O(n)  空间复杂度: O(n)
func Trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	// 正序填左侧最高值
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	// 倒序填右侧最高值
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	// 每个位置的接水量由左右最高值中较矮的一侧决定
	ans := 0
	for i := 0; i < n; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

// TrapAlternative 接雨水（双指针版）
// 接水量取决于左右最高值中较小的一侧，用双指针从两端向中间收缩，
// 哪侧短板就先结算哪侧，省去两个 DP 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func TrapAlternative(height []int) int {
	left, right := 0, len(height)-1
	lMax, rMax := 0, 0
	ans := 0
	for left < right {
		lMax = max(lMax, height[left])
		rMax = max(rMax, height[right])
		if lMax < rMax { // 左侧是短板，结算 left 位置
			ans += lMax - height[left]
			left++
		} else { // 右侧是短板，结算 right 位置
			ans += rMax - height[right]
			right--
		}
	}
	return ans
}
