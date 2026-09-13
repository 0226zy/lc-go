package partitiontokequalsumsubsets

import "sort"

// CanPartitionKSubsets 划分为k个相等的子集
// 给定一个整数数组 nums 和一个正整数 k，判断能否把数组划分成 k 个非空子集，使每个子集的元素和都相等。
// 时间复杂度: O(k^n) 最坏为指数级，n 为数组长度；排序与桶剪枝可显著削减搜索量
// 空间复杂度: O(n) 递归栈深度不超过 n 层（不计 k 个桶的 O(k) 辅助空间）
func CanPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// 总和不能整除 k，直接不可能均分
	if k <= 0 || sum%k != 0 {
		return false
	}
	target := sum / k

	// 降序排列：大数先放，能更早触发"桶放不下"的剪枝
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	// 最大的数超过单个桶容量，必然无法均分
	if nums[0] > target {
		return false
	}

	buckets := make([]int, k) // buckets[i] 表示第 i 个桶当前已装入的元素和

	// backtrack 尝试把 nums[index] 放进某个桶，成功放完所有数则返回 true
	var backtrack func(index int) bool
	backtrack = func(index int) bool {
		// 所有数都放完，且每个桶都恰好为 target，划分成功
		if index == len(nums) {
			return true
		}
		for i := 0; i < k; i++ {
			// 剪枝1：当前桶放不下 nums[index]
			if buckets[i]+nums[index] > target {
				continue
			}
			// 剪枝2：与上一个等量的桶等价，重复尝试没有意义
			if i > 0 && buckets[i] == buckets[i-1] {
				continue
			}
			buckets[i] += nums[index]
			if backtrack(index + 1) {
				return true
			}
			buckets[i] -= nums[index]
		}
		return false
	}
	return backtrack(0)
}
