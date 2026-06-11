package maxtotalsubarray

import "testing"

func TestMaxTotalSubarrayValue(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int64
	}{
		{
			name: "示例1",
			nums: []int{1, 3, 2},
			k:    2,
			want: 4,
		},
		{
			name: "示例2",
			nums: []int{4, 2, 5, 1},
			k:    3,
			want: 12,
		},
		{
			name: "单元素数组",
			nums: []int{5},
			k:    1,
			want: 0,
		},
		{
			name: "所有元素相同",
			nums: []int{3, 3, 3, 3},
			k:    5,
			want: 0,
		},
		{
			name: "两个元素选最大差",
			nums: []int{5, 1},
			k:    1,
			want: 4,
		},
		{
			name: "重复选择同一子数组",
			nums: []int{1, 5},
			k:    100,
			want: 400,
		},
		{
			name: "递增数组",
			nums: []int{1, 2, 3, 4, 5},
			k:    1,
			want: 4,
		},
		{
			name: "递减数组",
			nums: []int{5, 4, 3, 2, 1},
			k:    2,
			want: 8,
		},
		{
			name: "大数值",
			nums: []int{0, 1000000000},
			k:    3,
			want: 3000000000,
		},
		{
			name: "k为1选全局最大差",
			nums: []int{7, 2, 9, 4, 5},
			k:    1,
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxTotalSubarrayValue(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("MaxTotalSubarrayValue() = %d, want %d", got, tt.want)
			}
		})
	}
}

func BenchmarkMaxTotalSubarrayValue_Small(b *testing.B) {
	nums := []int{4, 2, 5, 1}
	k := 3
	for i := 0; i < b.N; i++ {
		_ = MaxTotalSubarrayValue(nums, k)
	}
}

func BenchmarkMaxTotalSubarrayValue_LargeK(b *testing.B) {
	nums := []int{1, 100}
	k := 100000
	for i := 0; i < b.N; i++ {
		_ = MaxTotalSubarrayValue(nums, k)
	}
}
