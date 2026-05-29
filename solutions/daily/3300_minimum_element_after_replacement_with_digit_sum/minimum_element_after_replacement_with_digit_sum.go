package minelement

// MinElement 替换后最小元素
// 将数组中每个元素替换为其数位和，返回替换后的最小元素。
// 时间复杂度:   空间复杂度:
func MinElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := int(^uint64(0) >> 1)
	for _, num := range nums {
		sum := 0
		for ; num > 0; num = num / 10 {
			sum += num % 10
		}
		if sum < res {
			res = sum
		}
	}
	return res

}
