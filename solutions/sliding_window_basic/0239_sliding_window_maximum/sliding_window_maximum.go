package slidingwindowmaximum

// MaxSlidingWindow 滑动窗口最大值
// 返回每个滑动窗口的最大值。
// 思路：单调递减队列，队首始终是当前窗口最大值的下标。
// 时间复杂度: O(n)  空间复杂度: O(k)
func MaxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return nil
	}

	// deque 存下标，对应 nums 值保持单调递减
	deque := []int{}
	result := make([]int, 0, n-k+1)

	for i := 0; i < n; i++ {
		// 1. 移除队首超出窗口的下标
		for len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}

		// 2. 从队尾移除所有比当前元素小的，保持单调递减
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// 3. 当前下标入队
		deque = append(deque, i)

		// 4. 窗口形成后，队首即为最大值
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
