package partitionequalsubsetsum

// CanPartition 分割等和子集（0-1 背包求可行性）
// 判断数组能否分成两个和相等的子集，转化为能否凑成 sum/2。
// 时间复杂度: O(n*target)  空间复杂度: O(target)
func CanPartition(nums []int) bool {
	sum := 0
	maxNum := 0
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}

	// sum 为奇数，无法平分
	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	// 最大元素超过 target，无法凑成（所有数为正）
	if maxNum > target {
		return false
	}

	// 最大元素恰好等于 target，直接成功
	if maxNum == target {
		return true
	}

	// 0-1 背包求可行性：dp[j] = 能否凑成和 j
	dp := make([]bool, target+1)
	dp[0] = true // 和为 0 总能凑成（不选任何元素）

	for _, num := range nums {
		// 倒序遍历，保证每个元素只选一次
		for j := target; j >= num; j-- {
			if dp[j-num] {
				dp[j] = true
			}
		}
		// 提前剪枝：已凑成 target，无需继续
		if dp[target] {
			return true
		}
	}
	return dp[target]
}

// CanPartition2D 分割等和子集（二维 DP 版本，便于理解）
// dp[i][j] = 前 i 个元素能否凑成和 j。
// 时间复杂度: O(n*target)  空间复杂度: O(n*target)
func CanPartition2D(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	n := len(nums)
	// dp[i][j] = 前 i 个元素能否凑成和 j
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true // 和为 0 总能凑成
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			// 不选 nums[i-1]
			dp[i][j] = dp[i-1][j]
			// 选 nums[i-1]（若容量足够）
			if j >= nums[i-1] && dp[i-1][j-nums[i-1]] {
				dp[i][j] = true
			}
		}
	}
	return dp[n][target]
}

// CanPartitionBitset 分割等和子集（位运算优化）
// 用整数位图代替 bool 数组，利用位运算并行加速。
// 时间复杂度: O(n*target/64)  空间复杂度: O(target/64)
func CanPartitionBitset(nums []int) bool {
	sum := 0
	maxNum := 0
	for _, num := range nums {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if maxNum > target {
		return false
	}
	if maxNum == target {
		return true
	}

	// bitset: 第 j 位为 1 表示能凑成和 j
	// 用 []uint64 模拟位图
	words := (target + 64) / 64
	bits := make([]uint64, words)
	bits[0] = 1 // bit 0 = 1（和为 0 可凑）

	for _, num := range nums {
		// 左移 num 位：dp = dp | (dp << num)
		// 分段处理 64 位整数数组
		newBits := make([]uint64, words)
		for i := 0; i < words; i++ {
			newBits[i] = bits[i] // 不选
		}
		// 选：左移 num 位
		wordShift := num / 64
		bitShift := num % 64
		for i := 0; i < words; i++ {
			src := i - wordShift
			if src < 0 {
				continue
			}
			var val uint64
			val = bits[src] << bitShift
			if bitShift > 0 && src-1 >= 0 {
				val |= bits[src-1] >> (64 - bitShift)
			}
			newBits[i] |= val
		}
		bits = newBits

		// 提前剪枝
		if (bits[target/64]>>(target%64))&1 == 1 {
			return true
		}
	}
	return (bits[target/64]>>(target%64))&1 == 1
}
