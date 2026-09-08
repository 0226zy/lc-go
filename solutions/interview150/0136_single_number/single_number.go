package singlenumber

// SingleNumber 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次外，其余每个元素均出现两次，找出那个只出现一次的元素。
// 利用异或性质：a^a=0、a^0=a、异或满足交换律，所有成对元素相互抵消，剩下的就是答案。
// 时间复杂度: O(n)  空间复杂度: O(1)
func SingleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}
