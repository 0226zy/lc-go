package longestconsecutivesequence

// LongestConsecutive 最长连续序列
// 给定一个未排序的整数数组 nums，返回其中最长的连续元素序列的长度，要求时间复杂度为 O(n)。
// 序列中的元素在原数组中可以不连续出现，但数值必须逐次相差 1（如 [1,2,3,4] 长度为 4）。
// 时间复杂度: O(n)  空间复杂度: O(n)
func LongestConsecutive(nums []int) int {
	// 用哈希集合去重，并提供 O(1) 的存在性判断
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}

	longest := 0
	for num := range set {
		// 只有 num-1 不存在时，num 才是某个连续序列的起点
		// 非起点的数会被它所属序列的起点向上扩展时覆盖到，直接跳过
		if _, ok := set[num-1]; ok {
			continue
		}

		// 从起点 num 开始向上扩展，统计当前连续序列长度
		currentLen := 1
		for {
			if _, ok := set[num+1]; !ok {
				break
			}
			num++
			currentLen++
		}
		if currentLen > longest {
			longest = currentLen
		}
	}
	return longest
}
