package longestincreasingsubsequence

// LengthOfLIS 最长递增子序列（动态规划）
// 给定一个整数数组 nums，找到其中最长严格递增子序列的长度。
// dp[i] 表示以 nums[i] 结尾的 LIS 长度，dp[i] = max(dp[j]+1) for j<i 且 nums[j]<nums[i]
// 时间复杂度: O(n^2)  空间复杂度: O(n)
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// LengthOfLISBinary 最长递增子序列（耐心排序 + 二分查找）
// 维护一个 tails 数组，tails[k] 表示长度为 k+1 的递增子序列的最小末尾元素。
// 遍历 nums，对每个元素在 tails 中二分查找第一个 >= 它的位置并替换；
// 若比所有 tails 元素都大，则追加（扩展最长长度）。
// 时间复杂度: O(n log n)  空间复杂度: O(n)
func LengthOfLISBinary(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// tails[k] = 长度为 k+1 的递增子序列的最小末尾元素
	tails := make([]int, 0, len(nums))
	tails = append(tails, nums[0])

	for i := 1; i < len(nums); i++ {
		// 当前元素比 tails 末尾大，可直接扩展最长子序列
		if nums[i] > tails[len(tails)-1] {
			tails = append(tails, nums[i])
			continue
		}
		// 二分查找第一个 >= nums[i] 的位置，替换之（保证同长度末尾最小）
		idx := lowerBound(tails, nums[i])
		tails[idx] = nums[i]
	}
	return len(tails)
}

// lowerBound 在有序数组中找第一个 >= target 的位置
func lowerBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
