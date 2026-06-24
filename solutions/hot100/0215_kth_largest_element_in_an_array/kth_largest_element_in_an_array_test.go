package kthlargestelementinanarray

import "testing"

func TestFindKthLargest(t *testing.T) {
t.Skip("未实现")
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "示例1：基本用例",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "示例2：有重复元素",
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		{
			name: "边界：k=1（最大值）",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    1,
			want: 6,
		},
		{
			name: "边界：k=n（最小值）",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    6,
			want: 1,
		},
		{
			name: "单元素",
			nums: []int{1},
			k:    1,
			want: 1,
		},
		{
			name: "包含负数",
			nums: []int{-1, 2, 0},
			k:    2,
			want: 0,
		},
		{
			name: "全部相同",
			nums: []int{5, 5, 5, 5},
			k:    2,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制切片避免原地修改
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			got := FindKthLargest(nums, tt.k)
			if got != tt.want {
				t.Errorf("FindKthLargest(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkFindKthLargest(b *testing.B) {
	nums := []int{3, 2, 1, 5, 6, 4}
	for i := 0; i < b.N; i++ {
		FindKthLargest(nums, 2)
	}
}
