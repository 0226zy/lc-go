package removeduplicatesfromsortedarray

// RemoveDuplicates 删除有序数组中的重复项
// 给定一个升序排列的数组 nums，删除重复出现的元素，使每个元素只出现一次，
// 返回移除后数组的新长度（原数组前部需为去重后的元素，顺序保持不变）。
// 时间复杂度: O(n) 双指针单次遍历  空间复杂度: O(1) 原地操作
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1 // slow 指向下一个不重复元素应该写入的位置
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
