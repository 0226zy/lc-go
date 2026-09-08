package longestconsecutivesequence

import (
	"math/rand"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// LeetCode 官方示例
		{"示例1: [100,4,200,1,3,2]最长序列[1,2,3,4]长度为4", []int{100, 4, 200, 1, 3, 2}, 4},
		{"示例2: [0,3,7,2,5,8,4,6,0,1]最长序列[0..8]长度为9", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},

		// 边界：空数组与单元素
		{"空数组返回0", []int{}, 0},
		{"单元素返回1", []int{42}, 1},

		// 边界：重复元素
		{"全部重复元素返回1", []int{5, 5, 5, 5}, 1},
		{"重复元素不影响连续长度", []int{1, 2, 2, 3, 3, 3, 4}, 4},

		// 边界：完全连续 / 完全不连续
		{"已排序完全连续", []int{1, 2, 3, 4, 5}, 5},
		{"逆序完全连续", []int{5, 4, 3, 2, 1}, 5},
		{"完全不连续各自孤立", []int{10, 30, 50, 70}, 1},

		// 边界：负数与正负混合
		{"全负数连续序列", []int{-5, -3, -4, -1, -2}, 5},
		{"跨正负零的连续序列", []int{-2, -1, 0, 1, 2}, 5},

		// 边界：极值
		{"含int32最小值", []int{-2147483648, -2147483647, -2147483646}, 3},
		{"含int32最大值", []int{2147483645, 2147483646, 2147483647}, 3},

		// 易错：多个连续段取最长
		{"两个连续段取较长者", []int{1, 2, 3, 100, 101}, 3},
		{"断续间隔的两个连续段", []int{0, 1, 2, 5, 6, 7, 8}, 4},
		{"起点之前存在空隙不合并", []int{1, 3, 4, 5, 6}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsecutive(tt.nums); got != tt.want {
				t.Errorf("LongestConsecutive(%v) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

// TestLongestConsecutiveLargeInput 压力测试：10^5 个随机打乱的连续整数
func TestLongestConsecutiveLargeInput(t *testing.T) {
	const n = 100000
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i - n/2 // 包含负数与正数的连续区间 [-50000, 49999]
	}
	rand.Shuffle(n, func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	if got := LongestConsecutive(nums); got != n {
		t.Errorf("LongestConsecutive(10^5随机打乱连续整数) = %v, want %v", got, n)
	}
}

func BenchmarkLongestConsecutive(b *testing.B) {
	// 基准数据：10^5 个随机打乱的连续整数，模拟最坏输入规模
	const n = 100000
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	rand.Shuffle(n, func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	benchmarks := []struct {
		name string
		nums []int
	}{
		{"10万个随机打乱的连续整数", nums},
		{"10万个相同元素", make([]int, n)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LongestConsecutive(bm.nums)
			}
		})
	}
}
