package singlenumberii

// SingleNumber 只出现一次的数字 II
// 给定一个整数数组 nums，除某个元素仅出现一次外，其余每个元素都恰出现三次，找出并返回那个只出现一次的元素。
// 时间复杂度: O(n)  空间复杂度: O(1)
func SingleNumber(nums []int) int {
	var ans int32
	// 按位统计：第 i 位上 1 出现的总次数对 3 取模，即为答案在第 i 位的取值
	for i := 0; i < 32; i++ {
		var cnt int32
		for _, num := range nums {
			// 转成 int32 按 32 位补码处理负数，右移后取最低位累加
			cnt += (int32(num) >> i) & 1
		}
		if cnt%3 != 0 {
			ans |= 1 << i
		}
	}
	// ans 以 int32 存储，符号位已正确还原，直接转回 int
	return int(ans)
}
