package removeduplicatesfromsortedarrayii

// RemoveDuplicates 删除有序数组中的重复项 II
// 给定一个升序排列的数组 nums，使每个元素最多出现两次，删除多余的重复项，
// 返回移除后数组的新长度（原数组前部需为处理后的元素，顺序保持不变）。
// 时间复杂度: O(n) 双指针单次遍历  空间复杂度: O(1) 原地操作
func RemoveDuplicates(nums []int) int {
	// slow 指向下一个元素应该写入的位置，同时保证 nums[slow-2] 之前的元素
	// 与新写入的元素不相等，即每个元素最多保留两份
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if slow < 2 || nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
