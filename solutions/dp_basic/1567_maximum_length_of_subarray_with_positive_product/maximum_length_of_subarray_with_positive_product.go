package maximumlengthsubarraypositiveproduct

// GetMaxLen 乘积为正数的最长子数组长度（标准 DP 数组版）
// pos[i] 表示以 nums[i] 结尾、乘积为正的最长子数组长度，
// neg[i] 表示以 nums[i] 结尾、乘积为负的最长子数组长度。
// 乘正数符号不变，乘负数符号翻转（正接负变负、负接负变正），遇 0 全部归零。
// 时间复杂度: O(n)  空间复杂度: O(n)
func GetMaxLen(nums []int) int {
	n := len(nums)
	pos := make([]int, n) // pos[i]：以 i 结尾乘积为正的最长长度
	neg := make([]int, n) // neg[i]：以 i 结尾乘积为负的最长长度
	// base case
	if nums[0] > 0 {
		pos[0] = 1
	} else if nums[0] < 0 {
		neg[0] = 1
	}
	ans := pos[0]
	for i := 1; i < n; i++ {
		switch {
		case nums[i] > 0:
			pos[i] = pos[i-1] + 1
			if neg[i-1] > 0 { // 负积存在，接上正数仍为负
				neg[i] = neg[i-1] + 1
			}
		case nums[i] < 0:
			if neg[i-1] > 0 { // 负负得正
				pos[i] = neg[i-1] + 1
			}
			neg[i] = pos[i-1] + 1
		default:
			// nums[i] == 0：乘积为 0，两个状态都是 0
		}
		ans = max(ans, pos[i])
	}
	return ans
}

// GetMaxLenOptimized 乘积为正数的最长子数组长度（滚动变量空间优化版）
// pos[i]、neg[i] 只依赖前一个位置的状态，用两个滚动变量代替 dp 数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func GetMaxLenOptimized(nums []int) int {
	pos, neg := 0, 0 // 滚动变量，分别代表 pos[i-1]、neg[i-1]
	ans := 0
	for _, x := range nums {
		switch {
		case x > 0:
			// 正数不改变符号：正积变长；负积存在时也变长
			pos++
			if neg > 0 {
				neg++
			}
		case x < 0:
			// 负数翻转符号：旧负积变正积，旧正积变负积
			newPos, newNeg := 0, pos+1
			if neg > 0 {
				newPos = neg + 1
			}
			pos, neg = newPos, newNeg
		default:
			// 乘积为 0，一切归零重新开始
			pos, neg = 0, 0
		}
		ans = max(ans, pos)
	}
	return ans
}
