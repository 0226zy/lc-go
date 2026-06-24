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
		{
			name: "含负数",
			nums: []int{-2, -1, 0, 1, 2},
			want: 5,
		},
		{
			name: "负数乱序",
			nums: []int{3, -2, 0, -1, 2},
			want: 3,
		},
		{
			name: "上界规模",
			nums: func() []int {
				nums := make([]int, 2500)
				for i := range nums {
					nums[i] = i
				}
				return nums
			}(),
			want: 2500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLIS(tt.nums); got != tt.want {
				t.Errorf("LengthOfLIS(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func TestLengthOfLISBinary(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"官方示例1", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"官方示例2", []int{0, 1, 0, 3, 2, 3}, 4},
		{"官方示例3", []int{7, 7, 7, 7, 7, 7, 7}, 1},
		{"空输入", []int{}, 0},
		{"单元素", []int{1}, 1},
		{"严格递增", []int{1, 2, 3, 4, 5}, 5},
		{"严格递减", []int{5, 4, 3, 2, 1}, 1},
		{"乱序", []int{1, 2, 4, 3}, 3},
		{"含负数", []int{-2, -1, 0, 1, 2}, 5},
		{"负数乱序", []int{3, -2, 0, -1, 2}, 3},
		{"全相同", []int{5, 5, 5, 5}, 1},
		{"交替升降", []int{1, 3, 2, 4, 3, 5, 4, 6}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLISBinary(tt.nums); got != tt.want {
				t.Errorf("LengthOfLISBinary(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

// TestConsistency 验证两种实现结果一致
func TestConsistency(t *testing.T) {
	cases := [][]int{
		{10, 9, 2, 5, 3, 7, 101, 18},
		{0, 1, 0, 3, 2, 3},
		{7, 7, 7, 7, 7, 7, 7},
		{1},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{-2, -1, 0, 1, 2},
		{3, -2, 0, -1, 2},
		{1, 3, 2, 4, 3, 5, 4, 6},
	}
	for _, nums := range cases {
		a := LengthOfLIS(nums)
		b := LengthOfLISBinary(nums)
		if a != b {
			t.Errorf("两种实现结果不一致: nums=%v, DP=%d, Binary=%d", nums, a, b)
		}
	}
}

func BenchmarkLengthOfLIS(b *testing.B) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	for i := 0; i < b.N; i++ {
		LengthOfLIS(nums)
	}
}

func BenchmarkLengthOfLISBinary(b *testing.B) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	for i := 0; i < b.N; i++ {
		LengthOfLISBinary(nums)
	}
}

func BenchmarkLengthOfLISLarge(b *testing.B) {
	nums := make([]int, 2500)
	for i := range nums {
		nums[i] = (i*7 + 3) % 10000 // 伪随机
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LengthOfLIS(nums)
	}
}

func BenchmarkLengthOfLISBinaryLarge(b *testing.B) {
	nums := make([]int, 2500)
	for i := range nums {
		nums[i] = (i*7 + 3) % 10000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LengthOfLISBinary(nums)
	}
}
