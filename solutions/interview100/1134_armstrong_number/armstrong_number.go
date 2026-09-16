package armstrongnumber

// IsArmstrong 阿姆斯特朗数
// 判断 n 是否为阿姆斯特朗数：k 位数每位数字的 k 次幂之和等于它本身。
// 时间复杂度: O(k^2)，k 为位数（最多 8）  空间复杂度: O(1)
func IsArmstrong(n int) bool {
	// 第一步：数出位数 k
	k := 0
	for tmp := n; tmp > 0; tmp /= 10 {
		k++
	}

	// 第二步：累加每位数字的 k 次幂
	sum := 0
	for tmp := n; tmp > 0; tmp /= 10 {
		d := tmp % 10
		p := 1
		for i := 0; i < k; i++ {
			p *= d
		}
		sum += p
		// 提前剪枝：和已经超过 n，不可能是阿姆斯特朗数
		if sum > n {
			return false
		}
	}
	return sum == n
}
