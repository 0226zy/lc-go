package houserobberii

import "testing"

var robCases = []struct {
	name string
	nums []int
	want int
}{
	{name: "示例1：首尾冲突", nums: []int{2, 3, 2}, want: 3},
	{name: "示例2：偷1和3", nums: []int{1, 2, 3, 1}, want: 4},
	{name: "示例3：三间取最大", nums: []int{1, 2, 3}, want: 3},
	{name: "只有一间房", nums: []int{7}, want: 7},
	{name: "两间房取大者", nums: []int{1, 2}, want: 2},
	{name: "全部为空房", nums: []int{0, 0, 0, 0}, want: 0},
	{name: "首尾都大只能取一", nums: []int{10, 1, 1, 10}, want: 11},
	{name: "放弃首间更优", nums: []int{1, 3, 1, 3, 100}, want: 103},
	{name: "长数组", nums: []int{2, 1, 4, 8, 9, 6, 3, 5}, want: 20},
}

func TestRob(t *testing.T) {
	for _, tt := range robCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rob(tt.nums); got != tt.want {
				t.Errorf("Rob(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestRobOptimized(t *testing.T) {
	for _, tt := range robCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := RobOptimized(tt.nums); got != tt.want {
				t.Errorf("RobOptimized(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkRob(b *testing.B) {
	nums := []int{2, 7, 9, 3, 1, 5, 8, 4, 6, 2}
	for i := 0; i < b.N; i++ {
		Rob(nums)
	}
}

func BenchmarkRobOptimized(b *testing.B) {
	nums := []int{2, 7, 9, 3, 1, 5, 8, 4, 6, 2}
	for i := 0; i < b.N; i++ {
		RobOptimized(nums)
	}
}
