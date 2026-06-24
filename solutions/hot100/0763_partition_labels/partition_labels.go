package partitionlabels

// PartitionLabels 划分字母区间
// 将字符串划分为尽可能多的片段，同一字母最多出现在一个片段中，返回每个片段的长度。
// 时间复杂度: O(n)  空间复杂度: O(1)
func PartitionLabels(s string) []int {
	start, end := 0, 0
	lastPos := make(map[byte]int, len(s))
	for i, c := range s {
		lastPos[byte(c)] = i
	}
	ret := []int{}
	for i, c := range s {
		end = max(end, lastPos[byte(c)])
		if i == end {
			ret = append(ret, end-start+1)
			start = end + 1
		}
	}

	return ret
}
