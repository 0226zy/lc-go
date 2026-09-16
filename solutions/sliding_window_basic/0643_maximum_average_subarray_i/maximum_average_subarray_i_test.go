package maximumaveragesubarrayi

import (
	"math"
	"testing"
)

// eps 浮点数比较误差上限
const eps = 1e-5

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		// LeetCode 官方示例
		{
			"示例1: 最大平均数 51/4",
			[]int{1, 12, -5, -6, 50, 3}, 4,
			12.75,
		},
		{
			"示例2: 单元素数组",
			[]int{5}, 1,
			5.0,
		},

		// 边界：空输入
		{
			"空数组",
			[]int{}, 1,
			0.0,
		},
		{
			"空数组且k为0",
			[]int{}, 0,
			0.0,
		},

		// 边界：k 等于数组长度，唯一窗口
		{
			"k等于数组长度",
			[]int{1, 2, 3, 4}, 4,
			2.5,
		},

		// 边界：k 为 1，最大平均数即数组最大值
		{
			"k为1取最大元素",
			[]int{3, 7, 1, 9, 4}, 1,
			9.0,
		},

		// 边界：全部元素相同
		{
			"全部元素相同",
			[]int{6, 6, 6, 6, 6}, 3,
			6.0,
		},

		// 边界：全部为负数，最大窗口和为负
		{
			"全部为负数",
			[]int{-1, -2, -3, -4}, 2,
			-1.5,
		},

		// 边界：极端值（题目约束内的最小/最大值）
		{
			"极端最小值-10000",
			[]int{-10000, -10000, -10000}, 3,
			-10000.0,
		},
		{
			"极端最大值10000",
			[]int{10000, 10000, 10000}, 2,
			10000.0,
		},
		{
			"极端值混合",
			[]int{10000, -10000, 10000, -10000, 10000}, 3,
			10000.0 / 3.0,
		},

		// 一般情况：最大窗口在数组开头
		{
			"最大窗口在开头",
			[]int{10, 10, 1, 1, 1}, 2,
			10.0,
		},
		// 一般情况：最大窗口在数组结尾
		{
			"最大窗口在结尾",
			[]int{1, 1, 1, 10, 10}, 2,
			10.0,
		},
		// 一般情况：结果为非整数
		{
			"平均数为非整数",
			[]int{0, 1, 1, 3, 3}, 4,
			2.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMaxAverage(tt.nums, tt.k)
			if math.Abs(got-tt.want) > eps {
				t.Errorf("FindMaxAverage(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

// makeNums 生成长度为 n、全部填充 fill 的数组
func makeNums(n, fill int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = fill
	}
	return nums
}

func BenchmarkFindMaxAverage(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
		k    int
	}{
		{"长度1000窗口10", 1000, 10},
		{"长度10000窗口100", 10000, 100},
		{"长度100000窗口1000", 100000, 1000},
	}

	for _, bm := range benchmarks {
		nums := makeNums(bm.n, 1)
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FindMaxAverage(nums, bm.k)
			}
		})
	}
}
