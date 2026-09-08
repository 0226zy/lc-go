package medianoftwosortedarrays

import (
	"math"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	const eps = 1e-5
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		// LeetCode 官方示例
		{"示例1: 奇数长度", []int{1, 3}, []int{2}, 2.0},
		{"示例2: 偶数长度", []int{1, 2}, []int{3, 4}, 2.5},

		// 边界：一个数组为空
		{"nums1为空", []int{}, []int{1}, 1.0},
		{"nums2为空", []int{2}, []int{}, 2.0},
		{"空数组与长数组奇数", []int{}, []int{1, 2, 3}, 2.0},
		{"空数组与长数组偶数", []int{}, []int{1, 2, 3, 4}, 2.5},

		// 边界：单元素 + 单元素（奇偶交叉验证交换逻辑）
		{"两个单元素", []int{1}, []int{2}, 1.5},
		{"两个相同单元素", []int{2}, []int{2}, 2.0},

		// 边界：完全不相交的两段
		{"nums1全小于nums2", []int{1, 2}, []int{3, 4, 5}, 3.0},
		{"nums1全大于nums2", []int{4, 5, 6}, []int{1, 2, 3}, 3.5},

		// 边界：交错数组
		{"交错奇数", []int{1, 3}, []int{2, 4, 5}, 3.0},
		{"交错偶数", []int{1, 3, 5}, []int{2, 4, 6}, 3.5},

		// 边界：含重复元素
		{"含重复元素", []int{1, 2, 2}, []int{2, 2, 2}, 2.0},
		{"全相同元素", []int{3, 3}, []int{3, 3, 3}, 3.0},

		// 边界：负数和极值
		{"含负数", []int{-5, -3}, []int{-2, 0}, -2.5},
		{"跨零", []int{-1000000}, []int{1000000}, 0.0},
		{"最小约束值", []int{-1000000, -1000000}, []int{-1000000}, -1000000.0},

		// 边界：nums1 恰好为较长数组（触发交换）
		{"nums1较长且为奇数", []int{1, 2, 3, 4, 5}, []int{6, 7}, 4.0},
		{"nums1较长且为偶数", []int{1, 2, 3, 4}, []int{5, 6}, 3.5},

		// 中位数完全落在其中一个数组内
		{"中位数在nums2", []int{1, 2}, []int{3, 4, 5, 6, 7, 8}, 4.5},
		{"中位数在nums1", []int{3, 4, 5, 6, 7, 8}, []int{1, 2}, 4.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMedianSortedArrays(tt.nums1, tt.nums2)
			if math.Abs(got-tt.want) > eps {
				t.Errorf("FindMedianSortedArrays(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}

// TestFindMedianSortedArraysBruteForce 用朴素合并法对拍验证大随机用例
func TestFindMedianSortedArraysBruteForce(t *testing.T) {
	const eps = 1e-5
	type args struct{ nums1, nums2 []int }
	cases := []args{
		{[]int{2, 2, 4, 4}, []int{2, 2, 4, 4}},
		{[]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10, 12, 14}},
		{[]int{10, 20, 30, 40, 50, 60}, []int{5, 15, 25}},
		{[]int{-10, -8, -6, -4}, []int{-9, -7, -5, -3, -2}},
		{[]int{1, 1, 1, 1, 1}, []int{2, 2, 2}},
	}
	for i, c := range cases {
		got := FindMedianSortedArrays(c.nums1, c.nums2)
		want := bruteForceMedian(c.nums1, c.nums2)
		if math.Abs(got-want) > eps {
			t.Errorf("对拍用例%d: FindMedianSortedArrays(%v, %v) = %v, 朴素法 = %v", i, c.nums1, c.nums2, got, want)
		}
	}
}

// bruteForceMedian 朴素解法：合并排序后取中位数，作为对拍基准
func bruteForceMedian(nums1, nums2 []int) float64 {
	merged := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}
	merged = append(merged, nums1[i:]...)
	merged = append(merged, nums2[j:]...)
	n := len(merged)
	if n%2 == 1 {
		return float64(merged[n/2])
	}
	return float64(merged[n/2-1]+merged[n/2]) / 2.0
}

func BenchmarkFindMedianSortedArrays(b *testing.B) {
	benchmarks := []struct {
		name  string
		nums1 []int
		nums2 []int
	}{
		{"m=10_n=10", makeRange(0, 10), makeRange(50, 10)},
		{"m=500_n=500", makeRange(0, 500), makeRange(1000, 500)},
		{"m=1000_n=1000", makeRange(0, 1000), makeRange(2000, 1000)},
		{"m=1_n=1000", makeRange(0, 1), makeRange(2000, 1000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindMedianSortedArrays(bm.nums1, bm.nums2)
			}
		})
	}
}

// makeRange 生成从 start 开始、长度为 n 的递增数组
func makeRange(start, n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = start + i
	}
	return nums
}
