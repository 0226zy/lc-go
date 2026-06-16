package searchinsertposition

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "示例1-目标存在",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			name:   "示例2-目标插入中间",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			name:   "示例3-目标大于所有元素",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			name:   "目标小于所有元素",
			nums:   []int{2, 3, 5, 6},
			target: 0,
			want:   0,
		},
		{
			name:   "目标等于第一个元素",
			nums:   []int{1, 3, 5, 6},
			target: 1,
			want:   0,
		},
		{
			name:   "目标等于最后一个元素",
			nums:   []int{1, 3, 5, 6},
			target: 6,
			want:   3,
		},
		{
			name:   "单元素-找到",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "单元素-未找到且更小",
			nums:   []int{5},
			target: 3,
			want:   0,
		},
		{
			name:   "单元素-未找到且更大",
			nums:   []int{5},
			target: 7,
			want:   1,
		},
		{
			name:   "负数数组",
			nums:   []int{-5, -3, -1, 0, 2},
			target: -2,
			want:   2,
		},
		{
			name:   "目标在数组末尾前插入",
			nums:   []int{1, 2, 4, 6, 7},
			target: 5,
			want:   3,
		},
		{
			name:   "空数组",
			nums:   []int{},
			target: 5,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchInsert(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("SearchInsert() = %d, want %d", got, tt.want)
			}
		})
	}
}

func BenchmarkSearchInsert(b *testing.B) {
	nums := []int{1, 3, 5, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30}
	target := 15
	for i := 0; i < b.N; i++ {
		_ = SearchInsert(nums, target)
	}
}
