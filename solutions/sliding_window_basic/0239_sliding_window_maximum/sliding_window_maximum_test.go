package slidingwindowmaximum

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "示例1：基本用例",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "示例2：k=1",
			nums: []int{1, -1},
			k:    1,
			want: []int{1, -1},
		},
		{
			name: "边界：k等于数组长度",
			nums: []int{1, 3, 2},
			k:    3,
			want: []int{3},
		},
		{
			name: "所有元素相同",
			nums: []int{7, 7, 7, 7},
			k:    2,
			want: []int{7, 7, 7},
		},
		{
			name: "递减序列",
			nums: []int{5, 4, 3, 2, 1},
			k:    3,
			want: []int{5, 4, 3},
		},
		{
			name: "递增序列",
			nums: []int{1, 2, 3, 4, 5},
			k:    3,
			want: []int{3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxSlidingWindow(tt.nums, tt.k)
			if got == nil {
				t.Skip("未实现")
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxSlidingWindow(b *testing.B) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	for i := 0; i < b.N; i++ {
		MaxSlidingWindow(nums, 3)
	}
}
