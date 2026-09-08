package twosum

// TwoSum 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值 target 的两个整数，并返回它们的下标（下标顺序任意）。
// 每种输入只对应一个答案，数组中同一个元素不能使用两次。
// 时间复杂度: O(n)  空间复杂度: O(n)
func TwoSum(nums []int, target int) []int {
	// 哈希表记录「数值 -> 下标」，遍历时边查边存
	indexByValue := make(map[int]int, len(nums))
	for i, num := range nums {
		// 若补数 target-num 之前已经出现过，直接组成答案返回
		if j, ok := indexByValue[target-num]; ok {
			return []int{j, i}
		}
		indexByValue[num] = i
	}
	// 题目保证答案存在，走不到这里；返回 nil 保持健壮性
	return nil
}
