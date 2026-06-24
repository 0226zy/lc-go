package jumpgameii

import "testing"

func TestJump(t *testing.T) {
t.Skip("未实现")
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "示例1：基本用例",
			nums: []int{2, 3, 1, 1, 4},
			want: 2,
		},
		{
			name: "示例2：另一种路径",
			nums: []int{2, 3, 0, 1, 4},
			want: 2,
		},
		{
			name: "边界：单元素",
			nums: []int{0},
			want: 0,
		},
		{
			name: "边界：已到达",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "每步最大跳跃",
			nums: []int{1, 1, 1, 1},
			want: 3,
		},
		{
			name: "一次大跳到达",
			nums: []int{5, 1, 1, 1, 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Jump(tt.nums)
			if got != tt.want {
				t.Errorf("Jump(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkJump(b *testing.B) {
	nums := []int{2, 3, 1, 1, 4}
	for i := 0; i < b.N; i++ {
		Jump(nums)
	}
}
