package subarrayproductlessthank

import "testing"

func TestNumSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		// LeetCode 官方示例
		{
			"示例1: 乘积等于k需严格排除",
			[]int{10, 5, 2, 6},
			100,
			8,
		},
		{
			"示例2: k为0直接返回0",
			[]int{1, 2, 3},
			0,
			0,
		},

		// 边界：k 为 1，任何子数组乘积都 >= 1
		{
			"k为1直接返回0",
			[]int{1, 1, 1},
			1,
			0,
		},

		// 边界：空数组
		{
			"空数组",
			[]int{},
			100,
			0,
		},

		// 边界：单元素数组
		{
			"单元素乘积小于k",
			[]int{5},
			10,
			1,
		},
		{
			"单元素乘积等于k",
			[]int{10},
			10,
			0,
		},
		{
			"单元素乘积大于k",
			[]int{100},
			10,
			0,
		},

		// 全为 1：所有子数组乘积都是 1，k=2 时全部合法
		{
			"全为1且k为2",
			[]int{1, 1, 1, 1},
			2,
			10,
		},

		// 每个元素单独都 >= k，答案为 0
		{
			"所有元素都不小于k",
			[]int{10, 20, 30},
			5,
			0,
		},

		// 乘积恰好等于 k 的子数组不计入
		{
			"子数组乘积恰好等于k",
			[]int{2, 5},
			10,
			2,
		},

		// 一般情况：数组中存在多种合法组合
		{
			"混合情况",
			[]int{1, 2, 3, 4},
			10,
			7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumSubarrayProductLessThanK(tt.nums, tt.k); got != tt.want {
				t.Errorf("NumSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNumSubarrayProductLessThanK(b *testing.B) {
	// 构造 3*10^4 规模的全 1 数组（题目数据上限）
	nums := make([]int, 30000)
	for i := range nums {
		nums[i] = 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NumSubarrayProductLessThanK(nums, 1<<30)
	}
}
