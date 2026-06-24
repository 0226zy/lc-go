package jumpgame

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "示例1：可以到达",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "示例2：无法到达",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "边界：单元素",
			nums: []int{0},
			want: true,
		},
		{
			name: "边界：第一个元素为0",
			nums: []int{0, 2, 3},
			want: false,
		},
		{
			name: "一次跳到终点",
			nums: []int{5, 1, 1, 1, 1, 1},
			want: true,
		},
		{
			name: "刚好到达",
			nums: []int{1, 1, 1, 0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CanJump(tt.nums)
			if got != tt.want {
				t.Errorf("CanJump(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkCanJump(b *testing.B) {
	nums := []int{2, 3, 1, 1, 4}
	for i := 0; i < b.N; i++ {
		CanJump(nums)
	}
}
