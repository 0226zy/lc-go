package medianoftwosortedarrays

import "math"

// FindMedianSortedArrays 寻找两个正序数组的中位数
// 给定两个大小分别为 m 和 n 的正序数组，找出并返回这两个正序数组的中位数。
// 时间复杂度: O(log(min(m,n)))  空间复杂度: O(1)
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	left, right := 0, m
	leftTotal := (m + n + 1) / 2
	for left <= right {
		i := (left + right) / 2
		j := leftTotal - i

		var nums1LeftMax, nums1RightMin float64
		var nums2LeftMax, nums2RightMin float64

		if i == 0 {
			nums1LeftMax = math.Inf(-1)
		} else {
			nums1LeftMax = float64(nums1[i-1])
		}

		if i == m {
			nums1RightMin = math.Inf(1)
		} else {
			nums1RightMin = float64(nums1[i])
		}

		if j == 0 {
			nums2LeftMax = math.Inf(-1)
		} else {
			nums2LeftMax = float64(nums2[j-1])
		}
		if j == n {
			nums2RightMin = math.Inf(1)
		} else {
			nums2RightMin = float64(nums2[j])
		}

		// find
		if nums1LeftMax <= nums2RightMin &&
			nums2LeftMax <= nums1RightMin {

			// 数组长度奇数
			if (m+n)%2 == 1 {
				return math.Max(nums1LeftMax, nums2LeftMax)
			}
			// 偶数
			leftMax := math.Max(nums1LeftMax, nums2LeftMax)
			rightMin := math.Min(nums1RightMin, nums2RightMin)
			return (leftMax + rightMin) / 2.0
		}

		// 二分
		if nums1LeftMax > nums2RightMin {
			// 太右，往左
			right = i - 1
		} else {
			// 太左，往右
			left = i + 1
		}
	}
	return 0.0
}
