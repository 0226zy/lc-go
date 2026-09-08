package findpeakelement

// FindPeakElement 寻找峰值
// 峰值是严格大于左右邻居的元素（边界外视为 -∞，相邻元素互不相等）。
// 二分爬坡：只要处于上坡（nums[mid] < nums[mid+1]）就往右走，否则峰值在左侧（含 mid），
// 最终收敛处必为峰值。
// 时间复杂度: O(log n) 二分  空间复杂度: O(1) 常数变量
func FindPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[mid+1] {
			left = mid + 1 // 处于上坡，峰值在右侧
		} else {
			right = mid // mid 本身是下坡起点，峰值在左侧（含 mid）
		}
	}
	return left
}
