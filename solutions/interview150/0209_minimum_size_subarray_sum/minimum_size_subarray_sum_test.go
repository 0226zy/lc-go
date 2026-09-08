package minimumsizesubarraysum

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		// LeetCode 官方示例
		{"示例1: target=7, nums=[2,3,1,2,4,3]", 7, []int{2, 3, 1, 2, 4, 3}, 2},
		{"示例2: target=4, nums=[1,4,4]", 4, []int{1, 4, 4}, 1},
		{"示例3: target=11, nums 全为1", 11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},

		// 边界：单元素满足
		{"单元素恰好等于target", 5, []int{5}, 1},
		// 边界：单元素不满足
		{"单元素小于target", 5, []int{3}, 0},
		// 边界：整个数组的和刚好等于 target
		{"整个数组刚好达标", 10, []int{2, 3, 5}, 3},
		// 边界：所有元素之和都小于 target
		{"总和不足target", 100, []int{1, 2, 3, 4}, 0},
		// 边界：大元素在末尾
		{"末尾大元素", 15, []int{1, 1, 1, 1, 15}, 1},
		// 边界：大元素在开头
		{"开头大元素", 15, []int{15, 1, 1, 1, 1}, 1},
		// 用例：需要跳过中间的较小窗口
		{"需要收缩到更短窗口", 7, []int{2, 3, 1, 2, 4, 3, 1}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinSubArrayLen(tt.target, tt.nums); got != tt.want {
				t.Errorf("MinSubArrayLen(%d, %v) = %d, want %d", tt.target, tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkMinSubArrayLen(b *testing.B) {
	benchmarks := []struct {
		name   string
		target int
		nums   []int
	}{
		{"len=10", 15, []int{2, 3, 1, 2, 4, 3, 1, 5, 6, 7}},
		{"len=100", 50, generateNums(100)},
		{"len=1000", 500, generateNums(1000)},
		{"len=10000", 5000, generateNums(10000)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinSubArrayLen(bm.target, bm.nums)
			}
		})
	}
}

// generateNums 生成长度为 n 的数组，元素在 1~10 之间循环
func generateNums(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i%10 + 1
	}
	return nums
}
