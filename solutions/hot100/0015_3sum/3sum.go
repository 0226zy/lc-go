package threesum

import "sort"

// ThreeSum 三数之和
// 给定一个整数数组 nums，找出所有和为 0 且不重复的三元组。
//
// 核心思路：先对数组排序，然后固定第一个数，使用双指针在剩余区间中寻找另外两个数。
// 通过排序后的连续性，在遍历和指针移动时跳过重复元素，实现自然去重。
//
// 时间复杂度: O(n²)  空间复杂度: O(log n)（排序的栈空间，不计结果存储）
func ThreeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < n-2; i++ {
		// 剪枝1：当前最小三数之和大于0，后续不可能找到
		if nums[i]+nums[i+1]+nums[i+2] > 0 {
			break
		}
		// 剪枝2：当前数与最大两数之和小于0，当前数太小，继续往后找更大的数
		if nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		// 去重：跳过重复的第一个数
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 去重：跳过重复的左指针元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 去重：跳过重复的右指针元素
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
