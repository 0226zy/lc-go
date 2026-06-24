package firstmissingpositive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
t.Skip("未实现")
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "示例1：基本用例",
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			name: "示例2：含负数",
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			name: "示例3：全大于n",
			nums: []int{7, 8, 9, 11, 12},
			want: 1,
		},
		{
			name: "边界：单个元素1",
			nums: []int{1},
			want: 2,
		},
		{
			name: "边界：单个元素非1",
			nums: []int{2},
			want: 1,
		},
		{
			name: "连续正整数",
			nums: []int{1, 2, 3, 4, 5},
			want: 6,
		},
		{
			name: "全为0",
			nums: []int{0, 0, 0},
			want: 1,
		},
		{
			name: "全为负数",
			nums: []int{-1, -2, -3},
			want: 1,
		},
		{
			name: "打乱且有缺失",
			nums: []int{2, 3, 5, 1, 6, 4, 8, 9},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制切片避免原地修改影响其他测试
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			got := FirstMissingPositive(nums)
			if got != tt.want {
				t.Errorf("FirstMissingPositive(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkFirstMissingPositive(b *testing.B) {
	nums := []int{3, 4, -1, 1}
	for i := 0; i < b.N; i++ {
		FirstMissingPositive(nums)
	}
}
