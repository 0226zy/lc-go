package combinationsumiv

import "testing"

var combinationSum4Cases = []struct {
	name   string
	nums   []int
	target int
	want   int
}{
	{name: "官方示例1", nums: []int{1, 2, 3}, target: 4, want: 7},
	{name: "官方示例2凑不出", nums: []int{9}, target: 3, want: 0},
	{name: "单元素恰好整除", nums: []int{1}, target: 5, want: 1},
	{name: "单元素目标为1", nums: []int{1}, target: 1, want: 1},
	{name: "目标为0", nums: []int{1, 2}, target: 0, want: 1},
	{name: "顺序敏感排列计数", nums: []int{2, 1, 3}, target: 5, want: 13},
	{name: "空数组", nums: []int{}, target: 3, want: 0},
}

func TestCombinationSum4(t *testing.T) {
	for _, tt := range combinationSum4Cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinationSum4(tt.nums, tt.target); got != tt.want {
				t.Errorf("CombinationSum4(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func BenchmarkCombinationSum4(b *testing.B) {
	nums := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		CombinationSum4(nums, 32)
	}
}
