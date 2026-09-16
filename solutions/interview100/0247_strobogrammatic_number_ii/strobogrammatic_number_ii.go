package strobogrammaticnumberii

// FindStrobogrammatic 中心对称数 II
// 给定整数 n，返回所有长度为 n 的中心对称数（旋转 180 度后与原数相同的数字）。
// 返回顺序任意。
// 时间复杂度: O(5^(n/2) * n)  空间复杂度: O(n) 不计返回值
func FindStrobogrammatic(n int) []string {
	// 旋转 180 度后互相映射的数字对
	pairs := [][2]byte{{'0', '0'}, {'1', '1'}, {'8', '8'}, {'6', '9'}, {'9', '6'}}

	ans := make([]string, 0)
	buf := make([]byte, n)
	var dfs func(left, right int)
	dfs = func(left, right int) {
		// 所有位置填完，得到一个合法解
		if left > right {
			ans = append(ans, string(buf))
			return
		}
		for _, p := range pairs {
			// 中间位（奇数长度）必须放旋转后与自身相同的数字
			if left == right && p[0] != p[1] {
				continue
			}
			// 多位数不能有前导零
			if left == 0 && n > 1 && p[0] == '0' {
				continue
			}
			buf[left], buf[right] = p[0], p[1]
			dfs(left+1, right-1)
		}
	}
	dfs(0, n-1)
	return ans
}
