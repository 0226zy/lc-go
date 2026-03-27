package easy

// TwoSum 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，
// 找出和为目标值的两个整数的下标。
// 时间复杂度: O(n)  空间复杂度: O(n)
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
