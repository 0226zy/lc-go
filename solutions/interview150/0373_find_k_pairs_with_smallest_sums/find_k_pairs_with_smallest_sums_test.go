package findkpairswithsmallestsums

import (
	"reflect"
	"sort"
	"testing"
)

// pairSum 返回数对和；和相同则按字典序比较，便于断言多答案等价
func pairSum(p []int) int { return p[0] + p[1] }

// equalKPairs 判断 got 是否是和最小的 k 个数对（和相同的对之间顺序不限）
func equalKPairs(got, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}
	if got == nil && want == nil {
		return true
	}
	g := append([][]int(nil), got...)
	w := append([][]int(nil), want...)
	less := func(a [][]int) func(i, j int) bool {
		return func(i, j int) bool {
			si, sj := a[i][0]+a[i][1], a[j][0]+a[j][1]
			if si != sj {
				return si < sj
			}
			if a[i][0] != a[j][0] {
				return a[i][0] < a[j][0]
			}
			return a[i][1] < a[j][1]
		}
	}
	sort.Slice(g, less(g))
	sort.Slice(w, less(w))
	return reflect.DeepEqual(g, w)
}

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
		// 边界：k 等于数对总数（和相同的对顺序不限）
		{"k等于总数", []int{1, 2}, []int{3, 4}, 4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}}},
		// 边界：单元素数组
		{"nums1单元素", []int{5}, []int{1, 2, 3}, 2, [][]int{{5, 1}, {5, 2}}},
		{"nums2单元素", []int{1, 2, 3}, []int{0}, 3, [][]int{{1, 0}, {2, 0}, {3, 0}}},
		// 边界：空数组
		{"nums1为空", []int{}, []int{1, 2}, 2, nil},
		{"nums2为空", []int{1, 2}, []int{}, 2, nil},
		// 边界：k 超过数对总数
		{"k超过总数", []int{1}, []int{2}, 5, [][]int{{1, 2}}},
		// 边界：负数（和同为 -10 的两对均可）
		{"负数数组", []int{-10, -5}, []int{-5, 0}, 2, [][]int{{-10, -5}, {-10, 0}}},
		// 边界：nums1 很长但 k 很小
		{"nums1很长k很小", []int{1, 2, 3, 4, 5, 6, 7, 8}, []int{1}, 3, [][]int{{1, 1}, {2, 1}, {3, 1}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KSmallestPairs(tt.nums1, tt.nums2, tt.k)
			// 空结果：实现返回 nil，期望也是 nil
			if tt.want == nil {
				if got != nil && len(got) != 0 {
					t.Errorf("KSmallestPairs(%v, %v, %d) = %v, want nil",
						tt.nums1, tt.nums2, tt.k, got)
				}
				return
			}
			if !equalKPairs(got, tt.want) {
				// 再校验：got 的 k 个和应都不超过 want 中最大和（前 k 小）
				if len(got) != len(tt.want) {
					t.Errorf("KSmallestPairs(%v, %v, %d) = %v, want %v",
						tt.nums1, tt.nums2, tt.k, got, tt.want)
					return
				}
				gSums := make([]int, len(got))
				wSums := make([]int, len(tt.want))
				for i := range got {
					gSums[i] = pairSum(got[i])
					wSums[i] = pairSum(tt.want[i])
				}
				sort.Ints(gSums)
				sort.Ints(wSums)
				if !reflect.DeepEqual(gSums, wSums) {
					t.Errorf("KSmallestPairs(%v, %v, %d) sums=%v, want sums=%v (pairs got=%v want=%v)",
						tt.nums1, tt.nums2, tt.k, gSums, wSums, got, tt.want)
				}
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
