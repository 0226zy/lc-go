package singlenumberii

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [2,2,3,2]结果为3", []int{2, 2, 3, 2}, 3},
		{"示例2: [0,1,0,1,0,1,99]结果为99", []int{0, 1, 0, 1, 0, 1, 99}, 99},

		// 边界：单元素数组
		{"单元素数组[0]", []int{0}, 0},
		{"单元素数组[-1]", []int{-1}, -1},
		{"单元素数组[int32最大值]", []int{2147483647}, 2147483647},
		{"单元素数组[int32最小值]", []int{-2147483648}, -2147483648},

		// 边界：答案为 0，其余元素出现三次
		{"答案为0混入三次出现的7", []int{7, 7, 7, 0}, 0},

		// 负数场景：答案为负数
		{"答案为负数-3", []int{-3, 5, 5, 5}, -3},
		{"负数为出现三次的元素", []int{-2, -2, -2, 1}, 1},
		{"答案为int32最小值", []int{9, 9, 9, -2147483648}, -2147483648},

		// 混合正负数
		{"正负混合答案为-10", []int{4, -10, 4, 4, -7, -7, -7}, -10},

		// 较大规模场景
		{"三个大数混合", []int{30000, 500, 2500, 30000, 500, 2500, 30000, 500, 2500, 701}, 701},
		{"乱序排列", []int{43, 16, 45, 43, 89, 45, 16, 43, 45, 16}, 89},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleNumber(tt.nums); got != tt.want {
				t.Errorf("SingleNumber(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	// 构造 3*10^4 规模的数组：9999 个数各出现三次，外加一个只出现一次的数
	bigNums := make([]int, 0, 3*10000)
	for i := 0; i < 9999; i++ {
		bigNums = append(bigNums, i, i, i)
	}
	bigNums = append(bigNums, 123456789)

	benchmarks := []struct {
		name string
		nums []int
	}{
		{"小规模7个元素", []int{0, 1, 0, 1, 0, 1, 99}},
		{"含负数小规模", []int{-2, -2, -2, 1}},
		{"3万元素压力场景", bigNums},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SingleNumber(bm.nums)
			}
		})
	}
}
