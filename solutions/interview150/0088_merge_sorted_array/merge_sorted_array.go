package mergesortedarray

// Merge 合并两个有序数组
// 给定两个按非递减顺序排列的整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，
// 使 nums1 成为按非递减顺序排列的数组。nums1 的尾部有足够的空间容纳 nums2。
// 时间复杂度: O(m + n) 双指针单次反向遍历  空间复杂度: O(1) 原地合并
func Merge(nums1 []int, m int, nums2 []int, n int) {
	// i、j 分别指向 nums1 和 nums2 的最后一个有效元素
	i, j := m-1, n-1
	// write 指向 nums1 尾部，是下一个元素应该写入的位置
	write := m + n - 1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[write] = nums1[i]
			i--
		} else {
			nums1[write] = nums2[j]
			j--
		}
		write--
	}
}
