package maximumdistanceinarrays

// MaxDistance 数组列表中的最大距离
// 给定 m 个升序排序的数组，从两个不同的数组中各挑一个整数，返回能得到的最大差值 |a - b|。
// 时间复杂度: O(m)  空间复杂度: O(1)
func MaxDistance(arrays [][]int) int {
	// 用第 0 个数组的最小值和最大值初始化
	minVal := arrays[0][0]
	maxVal := arrays[0][len(arrays[0])-1]
	ans := 0

	// 遍历其余数组，用当前数组与「之前所有数组」的最小/最大值配对，
	// 保证配对的两个数一定来自不同的数组
	for i := 1; i < len(arrays); i++ {
		lo := arrays[i][0]                // 当前数组最小值（升序，首元素）
		hi := arrays[i][len(arrays[i])-1] // 当前数组最大值（升序，末元素）

		if d := abs(hi - minVal); d > ans {
			ans = d
		}
		if d := abs(maxVal - lo); d > ans {
			ans = d
		}

		if lo < minVal {
			minVal = lo
		}
		if hi > maxVal {
			maxVal = hi
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
