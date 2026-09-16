package maximumnumberofones

import "sort"

// MaximumNumberOfOnes 矩阵中 1 的最大数量
// 有一个 width × height 的矩阵，每个单元格为 0 或 1，
// 且任意 sideLength × sideLength 的正方形子矩阵中 1 的数量不超过 maxOnes。
// 返回矩阵中最多可以有多少个 1。
// 思路：按下标 (i%sideLength, j%sideLength) 分同余类，每个子矩阵恰好包含每个
// 同余类中的一个格子，贪心地把格子数最多的 maxOnes 个同余类全部填 1。
// 时间复杂度: O(sideLength²·log sideLength)  空间复杂度: O(sideLength²)
func MaximumNumberOfOnes(width, height, sideLength, maxOnes int) int {
	// cnt 记录每个同余类 (r, c) 中的格子数
	cnt := make([]int, 0, sideLength*sideLength)
	for r := 0; r < sideLength; r++ {
		// 行下标 i ∈ [0, height) 且 i % sideLength == r 的行数
		rowCnt := 0
		if r < height {
			rowCnt = (height-1-r)/sideLength + 1
		}
		for c := 0; c < sideLength; c++ {
			// 列下标 j ∈ [0, width) 且 j % sideLength == c 的列数
			colCnt := 0
			if c < width {
				colCnt = (width-1-c)/sideLength + 1
			}
			cnt = append(cnt, rowCnt*colCnt)
		}
	}
	// 从大到小排序，取前 maxOnes 个同余类全部填 1
	sort.Sort(sort.Reverse(sort.IntSlice(cnt)))
	ans := 0
	for i := 0; i < maxOnes; i++ {
		ans += cnt[i]
	}
	return ans
}
