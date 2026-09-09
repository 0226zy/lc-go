package deleteandearn

import "testing"

var deleteAndEarnCases = []struct {
	name string
	nums []int
	want int
}{
	{name: "示例1", nums: []int{3, 4, 2}, want: 6},
	{name: "示例2", nums: []int{2, 2, 3, 3, 3, 4}, want: 9},
	{name: "单元素", nums: []int{5}, want: 5},
	{name: "全相同全拿", nums: []int{3, 3, 3, 3}, want: 12},
	{name: "相邻二选一选大的", nums: []int{1, 2}, want: 2},
	{name: "隔一个全拿", nums: []int{1, 1, 3, 3}, want: 8},
	{name: "间隔数值互不冲突", nums: []int{1, 3, 5}, want: 9},
	{name: "大数值与2不相邻可同拿", nums: []int{1, 2, 100}, want: 102},
	{name: "分散数值", nums: []int{8, 10, 4, 9, 1, 3, 5, 9, 4, 10}, want: 37},
}

func TestDeleteAndEarn(t *testing.T) {
	for _, tt := range deleteAndEarnCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteAndEarn(tt.nums); got != tt.want {
				t.Errorf("DeleteAndEarn(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestDeleteAndEarnOptimized(t *testing.T) {
	for _, tt := range deleteAndEarnCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeleteAndEarnOptimized(tt.nums); got != tt.want {
				t.Errorf("DeleteAndEarnOptimized(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkDeleteAndEarn(b *testing.B) {
	nums := []int{8, 10, 4, 9, 1, 3, 5, 9, 4, 10}
	for i := 0; i < b.N; i++ {
		DeleteAndEarn(nums)
	}
}

func BenchmarkDeleteAndEarnOptimized(b *testing.B) {
	nums := []int{8, 10, 4, 9, 1, 3, 5, 9, 4, 10}
	for i := 0; i < b.N; i++ {
		DeleteAndEarnOptimized(nums)
	}
}
