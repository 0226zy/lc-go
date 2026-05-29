package minelement

import "testing"

func TestMinElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "官方示例1",
			nums: []int{10, 12, 13, 14},
			want: 1,
		},
		{
			name: "官方示例2",
			nums: []int{1, 2, 3, 4},
			want: 1,
		},
		{
			name: "官方示例3",
			nums: []int{999, 19, 199},
			want: 10,
		},
		{
			name: "单元素",
			nums: []int{5},
			want: 5,
		},
		{
			name: "最大值边界",
			nums: []int{10000},
			want: 1,
		},
		{
			name: "全为个位数",
			nums: []int{9, 8, 7, 6},
			want: 6,
		},
		{
			name: "包含0的数位",
			nums: []int{101, 202, 303},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinElement(tt.nums)
			if got != tt.want {
				t.Errorf("MinElement(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkMinElement(b *testing.B) {
	nums := []int{10, 12, 13, 14}
	for i := 0; i < b.N; i++ {
		MinElement(nums)
	}
}
