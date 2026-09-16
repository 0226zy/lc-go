package missingelementinsortedarray

import "testing"

func TestMissingElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: k=1", []int{4, 7, 9, 10}, 1, 5},
		{"示例2: k=3", []int{4, 7, 9, 10}, 3, 8},
		{"示例3: 缺失数超出数组末尾", []int{1, 2, 4}, 3, 6},

		// 边界：单元素数组，缺失数全部在数组之外
		{"单元素数组", []int{3}, 1, 4},
		{"单元素数组大k", []int{3}, 10, 13},

		// 边界：无缺失（连续数组），缺失数都在末尾之后
		{"连续数组k=1", []int{1, 2, 3, 4}, 1, 5},
		{"连续数组k=3", []int{1, 2, 3, 4}, 3, 7},

		// 边界：第 k 个缺失数恰好等于最后一个数组元素之前的位置
		{"缺失数紧邻末元素前", []int{1, 5}, 3, 4},
		{"缺失数跳过末元素", []int{1, 5}, 4, 6},

		// 边界：k 恰好等于某段缺失数的总数
		{"k恰好等于段内缺失总数", []int{2, 5, 6}, 2, 4},

		// 边界：极端值
		{"大数值", []int{9999999, 10000000}, 1, 10000001},

		// 典型场景
		{"缺失数在中间某段", []int{1, 2, 4, 5, 10}, 4, 8},
		{"首段就有足够缺失", []int{1, 10, 11}, 5, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MissingElement(tt.nums, tt.k); got != tt.want {
				t.Errorf("MissingElement(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkMissingElement(b *testing.B) {
	// 构造 50000 个元素、间隔为 3 的升序数组
	nums := make([]int, 50000)
	for i := range nums {
		nums[i] = 1 + i*3
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MissingElement(nums, 100000)
	}
}
