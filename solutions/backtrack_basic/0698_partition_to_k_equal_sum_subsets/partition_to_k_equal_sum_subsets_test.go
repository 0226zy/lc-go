package partitiontokequalsumsubsets

import "testing"

func TestCanPartitionKSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: nums=[4,3,2,3,5,2,1],k=4",
			[]int{4, 3, 2, 3, 5, 2, 1}, 4, true},
		{"示例2: nums=[1,2,3,4],k=3",
			[]int{1, 2, 3, 4}, 3, false},

		// 边界：k 等于数组长度，每个数各自成一组，必须全部相等
		{"k等于数组长度且全部相等", []int{2, 2, 2, 2}, 4, true},
		{"k等于数组长度但元素不等", []int{1, 2, 3, 4}, 4, false},
		// 边界：k 为 1，恒可以（数组非空）
		{"k为1", []int{1, 2, 3}, 1, true},
		// 边界：总和不能整除 k
		{"总和不能整除k", []int{1, 2, 3, 5}, 3, false},
		// 边界：最大元素超过单桶容量
		{"最大元素超过target", []int{2, 2, 2, 2, 4}, 4, false},
		// 边界：单元素数组
		{"单元素k为1", []int{7}, 1, true},
		// 需要回溯才能找到解的场景：贪心放法会卡住
		{"需要回溯的场景", []int{4, 3, 2, 3, 5, 2, 1}, 4, true},
		{"可均分的重复元素", []int{3, 3, 3, 3}, 2, true},
		{"不可均分需整体失败", []int{2, 2, 2, 3}, 3, false},
		// 较大输入，依赖剪枝才能快速通过
		{"较大输入可划分", []int{10, 10, 10, 7, 7, 7, 7, 7, 7, 6, 6, 6}, 3, true},
		{"较大输入不可划分", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums) // 函数内部会排序，复制一份避免用例间相互影响
			if got := CanPartitionKSubsets(nums, tt.k); got != tt.want {
				t.Errorf("CanPartitionKSubsets(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

func BenchmarkCanPartitionKSubsets(b *testing.B) {
	benchmarks := []struct {
		name string
		nums []int
		k    int
	}{
		{"官方示例k=4", []int{4, 3, 2, 3, 5, 2, 1}, 4},
		{"不可划分k=3", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3},
		{"n=16最坏可划分", []int{4, 3, 2, 3, 5, 2, 1, 4, 3, 2, 3, 5, 2, 1, 4, 5}, 4},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				nums := make([]int, len(bm.nums))
				copy(nums, bm.nums)
				CanPartitionKSubsets(nums, bm.k)
			}
		})
	}
}
