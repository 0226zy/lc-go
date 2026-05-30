package threesum

import (
	"reflect"
	"sort"
	"testing"
)

// equal2DIntSlice 比较两个二维整数切片是否相等（忽略内部切片顺序）
func equal2DIntSlice(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	// 对每个内部切片排序后再比较
	sortInner := func(s [][]int) [][]int {
		cp := make([][]int, len(s))
		for i, v := range s {
			inner := make([]int, len(v))
			copy(inner, v)
			sort.Ints(inner)
			cp[i] = inner
		}
		// 对外层按字典序排序
		sort.Slice(cp, func(i, j int) bool {
			for k := 0; k < len(cp[i]) && k < len(cp[j]); k++ {
				if cp[i][k] != cp[j][k] {
					return cp[i][k] < cp[j][k]
				}
			}
			return len(cp[i]) < len(cp[j])
		})
		return cp
	}
	aSorted := sortInner(a)
	bSorted := sortInner(b)
	return reflect.DeepEqual(aSorted, bSorted)
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "官方示例1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "官方示例2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "官方示例3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "空数组",
			nums: []int{},
			want: [][]int{},
		},
		{
			name: "不足三个元素",
			nums: []int{1, 2},
			want: [][]int{},
		},
		{
			name: "全正数",
			nums: []int{1, 2, 3},
			want: [][]int{},
		},
		{
			name: "全负数",
			nums: []int{-1, -2, -3},
			want: [][]int{},
		},
		{
			name: "多个重复元素",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "两正一负",
			nums: []int{-2, 1, 1},
			want: [][]int{{-2, 1, 1}},
		},
		{
			name: "大量重复零",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "复杂重复",
			nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			want: [][]int{{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThreeSum(tt.nums)
			if !equal2DIntSlice(got, tt.want) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkThreeSum(b *testing.B) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	for i := 0; i < b.N; i++ {
		ThreeSum(nums)
	}
}
