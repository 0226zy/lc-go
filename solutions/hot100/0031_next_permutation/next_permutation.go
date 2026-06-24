package nextpermutation

// NextPermutation 下一个排列
// 原地将 nums 修改为字典序下一个更大的排列；若已是最大排列则重排为升序。
// 时间复杂度: O(n)  空间复杂度: O(1)
func NextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// 1. 从后往前找第一个下降点 i，满足 nums[i] < nums[i+1]
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		// 2. 从后往前找第一个 j，满足 nums[j] > nums[i]
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 3. 反转 i+1 到末尾，使其升序
	reverse(nums, i+1, n-1)
}

// reverse 反转 nums[i:j]
func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
