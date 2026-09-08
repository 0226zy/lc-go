package addbinary

// AddBinary 二进制求和
// 给定两个二进制字符串 a 和 b，返回它们的和（以二进制字符串形式输出）。
// 模拟竖式加法：从低位向高位逐位相加，处理进位。
// 时间复杂度: O(max(m, n))  空间复杂度: O(max(m, n))
func AddBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	res := make([]byte, 0, max(len(a), len(b))+1)

	// 从两个字符串的末位（二进制最低位）开始逐位相加
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		// sum 只可能是 0/1/2/3，除以 2 的商是进位，余数是当前位
		res = append(res, byte('0'+sum%2))
		carry = sum / 2
	}

	// res 中是低位在前，需要原地反转成高位在前
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}
