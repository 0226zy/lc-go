package wigglesort

// WiggleSort 摆动排序
// 将整数数组 nums 原地重排为 nums[0] <= nums[1] >= nums[2] <= nums[3]... 的摆动序列。
// 思路：一次遍历，偶数下标保证 nums[i] <= nums[i+1]，奇数下标保证 nums[i] >= nums[i+1]，违反就交换。
// 时间复杂度: O(n)  空间复杂度: O(1)
func WiggleSort(nums []int) {
	for i := 0; i+1 < len(nums); i++ {
		// 偶数下标要求「小」：nums[i] <= nums[i+1]
		// 奇数下标要求「大」：nums[i] >= nums[i+1]
		if (i%2 == 0 && nums[i] > nums[i+1]) || (i%2 == 1 && nums[i] < nums[i+1]) {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
}
