package majorityelement

// MajorityElement 多数元素
// 给定一个大小为 n 的数组 nums，找出其中的多数元素。多数元素是出现次数
// 严格大于 ⌊n / 2⌋ 的元素。题目保证多数元素一定存在。
// 时间复杂度: O(n) 一次遍历  空间复杂度: O(1) 常数空间
func MajorityElement(nums []int) int {
	candidate := 0
	count := 0 // count 表示当前候选人与已遍历元素之间的“净票数”
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// MajorityElementHash 多数元素（哈希计数对照解法）
// 用哈希表统计每个元素出现的次数，一旦超过 n/2 立即返回。
// 时间复杂度: O(n)  空间复杂度: O(n)
func MajorityElementHash(nums []int) int {
	half := len(nums) / 2
	counts := make(map[int]int)
	for _, v := range nums {
		counts[v]++
		if counts[v] > half {
			return v
		}
	}
	return 0 // 题目保证存在多数元素，不会走到这里
}
