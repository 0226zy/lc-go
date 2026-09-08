package removeelement

// RemoveElement 移除元素
// 给定一个数组 nums 和一个值 val，移除数组中所有数值等于 val 的元素，
// 返回移除后数组的新长度（原数组前部需为保留下来的元素，顺序可保持）。
// 时间复杂度: O(n) 双指针单次遍历  空间复杂度: O(1) 原地操作
func RemoveElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
