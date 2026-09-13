package twosumii

// TwoSum 两数之和 II - 输入有序数组
// 给定已按非递减顺序排列的整数数组 numbers，从中找出两个数使它们的和等于目标数 target。
// 返回这两个数的下标（下标从 1 开始），且只能使用每个下标一次。
// 对撞双指针：左指针从小端、右指针从大端向中间逼近。
// 时间复杂度: O(n)  空间复杂度: O(1)
func TwoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		switch {
		case sum == target:
			// 题目要求返回从 1 开始的下标
			return []int{left + 1, right + 1}
		case sum < target:
			left++ // 和偏小，左指针右移增大和
		default:
			right-- // 和偏大，右指针左移减小和
		}
	}
	// 题目保证恰好存在一组解，此处兜底
	return []int{-1, -1}
}
