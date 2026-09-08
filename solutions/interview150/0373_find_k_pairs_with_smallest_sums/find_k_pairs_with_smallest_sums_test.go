package findkpairswithsmallestsums

import (
	"reflect"
	"testing"
)

func TestKSmallestPairs(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		k     int
		want  [][]int
	}{
		// LeetCode 官方示例
		{"示例1", []int{1, 7, 11}, []int{2, 4, 6}, 3, [][]int{{1, 2}, {1, 4}, {1, 6}}},
		{"示例2: 含重复元素", []int{1, 1, 2}, []int{1, 2, 3}, 2, [][]int{{1, 1}, {1, 1}}},

		// 边界：k=1
		{"k=1", []int{1, 2}, []int{3, 4}, 1, [][]int{{1, 3}}},
		// 边界：k 等于数对总数
		{"k等于总数", []int{1, 2}, []int{3, 4}, 4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}}},
		// 边界：单元素数组
		{"nums1单元素", []int{5}, []int{1, 2, 3}, 2, [][]int{{5, 1}, {5, 2}}},
		{"nums2单元素", []int{1, 2, 3}, []int{0}, 3, [][]int{{1, 0}, {2, 0}, {3, 0}}},
		// 边界：空数组
		{"nums1为空", []int{}, []int{1, 2}, 2, nil},
		{"nums2为空", []int{1, 2}, []int{}, 2, nil},
		// 边界：k 超过数对总数
		{"k超过总数", []int{1}, []int{2}, 5, [][]int{{1, 2}}},
		// 边界：负数
		{"负数数组", []int{-10, -5}, []int{-5, 0}, 2, [][]int{{-10, -5}, {-10, 0}}},
		// 边界：nums1 很长但 k 很小
		{"nums1很长k很小", []int{1, 2, 3, 4, 5, 6, 7, 8}, []int{1}, 3, [][]int{{1, 1}, {2, 1}, {3, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KSmallestPairs(tt.nums1, tt.nums2, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KSmallestPairs(%v, %v, %d) = %v, want %v",
					tt.nums1, tt.nums2, tt.k, got, tt.want)
			}
		})
	}
}

// TestKSmallestPairsSorted 验证输出按和升序（题意要求返回和最小的 k 对）
func TestKSmallestPairsSorted(t *testing.T) {
	nums1 := []int{1, 3, 5, 7, 9}
	nums2 := []int{2, 4, 6, 8, 10}
	k := 10
	got := KSmallestPairs(nums1, nums2, k)
	if len(got) != k {
		t.Fatalf("长度 = %d, want %d", len(got), k)
	}
	for i := 1; i < len(got); i++ {
		if got[i-1][0]+got[i-1][1] > got[i][0]+got[i][1] {
			t.Errorf("第 %d 对与第 %d 对和逆序: %v, %v", i-1, i, got[i-1], got[i])
		}
	}
}

func BenchmarkKSmallestPairs(b *testing.B) {
	// 构造两个 1e5 规模的升序数组，取 k=1e4
	n := 100000
	nums1 := make([]int, n)
	nums2 := make([]int, n)
	for i := 0; i < n; i++ {
		nums1[i] = i
		nums2[i] = 2 * i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KSmallestPairs(nums1, nums2, 10000)
	}
}
