package totalwaviness

// TotalWaviness 范围内总波动值 II
// 返回区间 [num1, num2] 内所有数字的波动值之和。
// 时间复杂度: TODO  空间复杂度: TODO
func TotalWaviness(num1 int, num2 int) int64 {

	ret := int64(0)
	for n := num1; n <= num2; n++ {
		nums := []int{}
		for tmp := n; tmp > 1; tmp = tmp / 10 {
			nums = append(nums, tmp%10)
		}

		left, right := 0, len(nums)-1
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
		res := waviness(nums)
		res += ret
	}
	return ret
}

func waviness(nums []int) int64 {
	if len(nums) < 3 {
		return 0
	}
	n := len(nums)
	ret := int64(0)
	for i := 1; i < n-1; i++ {
		if nums[i] > nums[i+1] && nums[i] > nums[i-1] {
			ret++
		}
		if nums[i] < nums[i+1] && nums[i] < nums[i-1] {
			ret++
		}
	}
	return ret
}
