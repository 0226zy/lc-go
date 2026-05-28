package longestincreasingsubsequence

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "官方示例1",
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			name: "官方示例2",
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			name: "官方示例3",
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},
		{
			name: "空输入",
			nums: []int{},
			want: 0,
		},
		{
			name: "单元素",
			nums: []int{1},
			want: 1,
		},
		{
			name: "严格递增",
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "严格递减",
			nums: []int{5, 4, 3, 2, 1},
			want: 1,
		},
		{
			name: "乱序",
			nums: []int{1, 2, 4, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLIS(tt.nums)
			if got != tt.want {
				t.Errorf("LengthOfLIS(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkLengthOfLIS(b *testing.B) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	for i := 0; i < b.N; i++ {
		LengthOfLIS(nums)
	}
}
