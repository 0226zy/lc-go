package threesum

import "sort"

// ThreeSum 三数之和
// 给定整数数组 nums，返回所有满足 nums[i]+nums[j]+nums[k]==0（i、j、k 互不相同）且不重复的三元组，三元组内部按升序输出。
// 时间复杂度: O(n^2)  空间复杂度: O(log n)（排序栈空间，不计输出）
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	result := make([][]int, 0)

	for i := 0; i < n-2; i++ {
		// 剪枝：排序后 nums[i] > 0，其后所有数非负，三数之和不可能为 0
		if nums[i] > 0 {
			break
		}
		// i 去重：与上一个首元素相同，以它为起点的三元组上一轮已找过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 对撞双指针：在 nums[i+1..n-1] 中找两数之和等于 -nums[i]
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			switch {
			case sum < 0:
				left++ // 和太小，需要更大的数
			case sum > 0:
				right-- // 和太大，需要更小的数
			default:
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 左右指针去重：跳过所有与当前值相同的元素
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}
