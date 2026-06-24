package houserobber

import "testing"

func TestRob(t *testing.T) {
t.Skip("未实现")
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "示例1：4间房",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "示例2：5间房",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			name: "边界：单间房",
			nums: []int{5},
			want: 5,
		},
		{
			name: "边界：两间房",
			nums: []int{2, 3},
			want: 3,
		},
		{
			name: "全为0",
			nums: []int{0, 0, 0},
			want: 0,
		},
		{
			name: "相等金额",
			nums: []int{2, 2, 2, 2},
			want: 4,
		},
		{
			name: "经典间隔偷",
			nums: []int{2, 1, 1, 2},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Rob(tt.nums)
			if got != tt.want {
				t.Errorf("Rob(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkRob(b *testing.B) {
	nums := []int{2, 7, 9, 3, 1}
	for i := 0; i < b.N; i++ {
		Rob(nums)
	}
}
