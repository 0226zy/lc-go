package movezeroes

// MoveZeroes 移动零
// 给定一个数组 nums，将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 必须在不复制数组的情况下原地对数组进行操作。
// 时间复杂度: O(n) 单次遍历  空间复杂度: O(1) 常数空间
func MoveZeroes(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
