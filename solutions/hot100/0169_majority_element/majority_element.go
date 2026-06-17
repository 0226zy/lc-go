package majorityelement

// MajorityElement 多数元素
// 使用摩尔投票算法（Boyer-Moore Voting），在 O(n) 时间和 O(1) 空间内找到多数元素。
// 时间复杂度: O(n)  空间复杂度: O(1)
func MajorityElement(nums []int) int {
	ret, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			ret = num
		}
		if ret == num {
			count += 1
		} else {
			count -= 1
		}
	}
	return ret
}

// MajorityElementHashMap 哈希表解法
// 使用 map 统计每个数字出现次数，找到出现次数超过 n/2 的元素。
// 时间复杂度: O(n)  空间复杂度: O(n)
func MajorityElementHashMap(nums []int) int {
	// TODO: 实现哈希表解法
	panic("not implemented")
}
