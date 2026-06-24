package nextpermutation

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"官方示例1", []int{1, 2, 3}, []int{1, 3, 2}},
		{"官方示例2 最大排列", []int{3, 2, 1}, []int{1, 2, 3}},
		{"官方示例3", []int{1, 1, 5}, []int{1, 5, 1}},
		{"单元素", []int{1}, []int{1}},
		{"两元素升序", []int{1, 2}, []int{2, 1}},
		{"两元素降序", []int{2, 1}, []int{1, 2}},
		{"重复元素", []int{2, 2, 1}, []int{1, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)
			NextPermutation(nums)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Errorf("NextPermutation(%v) = %v, want %v", tt.nums, nums, tt.want)
			}
		})
	}
}

func TestNextPermutation_Sequence(t *testing.T) {
	// 连续调用应产生字典序的所有排列
	nums := []int{1, 2, 3}
	seen := map[[3]int]bool{{1, 2, 3}: true}
	for i := 0; i < 5; i++ {
		NextPermutation(nums)
		var key [3]int
		copy(key[:], nums)
		if seen[key] {
			t.Errorf("产生重复排列: %v", nums)
		}
		seen[key] = true
	}
	if len(seen) != 6 {
		t.Errorf("应产生 6 个排列，实际 %d", len(seen))
	}
}

func TestReverse(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	reverse(nums, 1, 4)
	want := []int{1, 5, 4, 3, 2}
	if !reflect.DeepEqual(nums, want) {
		t.Errorf("reverse(1,4) = %v, want %v", nums, want)
	}
}

func BenchmarkNextPermutation(b *testing.B) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i
	}
	for i := 0; i < b.N; i++ {
		NextPermutation(nums)
	}
}
