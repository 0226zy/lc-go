package medianoftwosortedarrays

// FindMedianSortedArrays 寻找两个正序数组的中位数
// 给定两个正序数组 nums1 和 nums2，返回它们合并后的中位数，要求 O(log(m+n))。
// 核心思想：在两个数组中各找一个切割位置 i、j（i+j = (m+n+1)/2），使得
// nums1[i-1] <= nums2[j] 且 nums2[j-1] <= nums1[i]，即左半最大值 <= 右半最小值；
// 对较短的数组二分切割位置 i 即可。
// 时间复杂度: O(log(min(m,n)))  空间复杂度: O(1)
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 保证 nums1 是较短的数组，使二分区间 [0, m] 有意义
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	half := (m + n + 1) / 2 // 左半部分应包含的元素个数（奇数时左半多一个）

	left, right := 0, m
	for left <= right {
		i := left + (right-left)/2 // 防溢出写法
		j := half - i              // j 由 i 唯一确定，且因 m <= n 必有 0 <= j <= n
		if i < m && nums2[j-1] > nums1[i] {
			left = i + 1 // i 太小，左半的最大值越界了
		} else if i > 0 && nums1[i-1] > nums2[j] {
			right = i - 1 // i 太大
		} else {
			// 切割正确：求左半最大值 maxLeft（i==0 或 j==0 表示某侧为空）
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else if nums1[i-1] > nums2[j-1] {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = nums2[j-1]
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft) // 奇数：左半多一个元素，maxLeft 就是中位数
			}
			// 偶数：再求右半最小值 minRight（i==m 或 j==n 表示某侧为空）
			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else if nums1[i] < nums2[j] {
				minRight = nums1[i]
			} else {
				minRight = nums2[j]
			}
			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0 // 题目保证 m+n >= 1，实际不会走到这里
}
