package sortcolors

// SortColors 颜色分类
// 使用荷兰国旗问题的三指针解法，将 0 移到前端、2 移到末端，1 留在中间。
// 时间复杂度: O(n)  空间复杂度: O(1)
func SortColors(nums []int) {
	// 三指针：p0 指向下一个 0 的位置，p2 指向下一个 2 的位置
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; {
		switch nums[i] {
		case 0:
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
			i++
		case 1:
			i++
		case 2:
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
			// 不递增 i，因为交换过来的元素还未检查
		}
	}
}

// SortColorsCounting 计数排序解法
// 先统计 0、1、2 的个数，再按顺序重写数组。
// 时间复杂度: O(n)  空间复杂度: O(1)
func SortColorsCounting(nums []int) {
	// 统计 0、1、2 的个数，再按顺序重写
	var cnt [3]int
	for _, v := range nums {
		cnt[v]++
	}
	idx := 0
	for v := 0; v < 3; v++ {
		for k := 0; k < cnt[v]; k++ {
			nums[idx] = v
			idx++
		}
	}
}
