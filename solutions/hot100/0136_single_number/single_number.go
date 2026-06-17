package singlenumber

// SingleNumber 只出现一次的数字
// 利用异或运算的性质：a ^ a = 0，a ^ 0 = a，所有数字异或后即为只出现一次的数字。
// 时间复杂度: O(n)  空间复杂度: O(1)
func SingleNumber(nums []int) int {
	ret := nums[0]
	for _, num := range nums[1:] {
		ret = ret ^ num
	}
	return ret
}

// SingleNumberHashMap 哈希表解法
// 使用 map 统计每个数字出现次数，找到出现次数为 1 的数字。
// 时间复杂度: O(n)  空间复杂度: O(n)
func SingleNumberHashMap(nums []int) int {
	dict := map[int]int{}
	for _, num := range nums {
		dict[num]++
	}
	for num, count := range dict {
		if count == 1 {
			return num
		}
	}
	return -1
}
