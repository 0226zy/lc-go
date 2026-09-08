package twosum

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
	}{
		// LeetCode 官方示例
		{"示例1: nums=[2,7,11,15], target=9", []int{2, 7, 11, 15}, 9},
		{"示例2: nums=[3,2,4], target=6", []int{3, 2, 4}, 6},
		{"示例3: nums=[3,3], target=6", []int{3, 3}, 6},

		// 边界：数组长度最小值 2
		{"两元素数组直接命中", []int{1, 2}, 3},

		// 边界：含负数与 0
		{"负数与正数配对", []int{-3, 4, 3, 90}, 0},
		{"两个负数配对", []int{-1, -2, -3, -4, -5}, -8},
		{"目标值为0且答案在末尾", []int{5, 75, -5}, 0},

		// 边界：两个相同的值配对（不能误用同一元素两次，需真实出现两次）
		{"两个相同值组成答案", []int{6, 3, 6}, 12},

		// 边界：答案位于数组首尾两端
		{"答案分布在首尾", []int{0, 4, 3, 0}, 0},

		// 边界：极大/极小数值（约束范围内 ±10^9）
		{"极大值与极小值配对", []int{-1000000000, 500, 1000000000}, 0},
		{"极大target命中", []int{1, 999999999, 999999998}, 1999999997},

		// 易错：先存后查导致用到自己，本实现边查边存可正确跳过
		{"补数等于自身但只有一份时不命中自身", []int{5, 3, 7}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoSum(tt.nums, tt.target)
			// 下标顺序任意：校验和正确、两个下标不同且在合法范围内
			if got == nil {
				t.Fatalf("TwoSum(%v, %d) = nil, want 有效下标对", tt.nums, tt.target)
			}
			if len(got) != 2 || got[0] == got[1] {
				t.Fatalf("TwoSum(%v, %d) = %v, 应为两个不同下标", tt.nums, tt.target, got)
			}
			if got[0] < 0 || got[0] >= len(tt.nums) || got[1] < 0 || got[1] >= len(tt.nums) {
				t.Fatalf("TwoSum(%v, %d) = %v, 下标越界", tt.nums, tt.target, got)
			}
			if tt.nums[got[0]]+tt.nums[got[1]] != tt.target {
				t.Errorf("TwoSum(%v, %d) = %v, nums[got] 之和 = %d, want %d",
					tt.nums, tt.target, got, tt.nums[got[0]]+tt.nums[got[1]], tt.target)
			}
		})
	}
}

// TestTwoSumExactAnswer 对本实现返回的确定下标做精确断言（题目保证答案唯一）
func TestTwoSumExactAnswer(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"示例1精确答案", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"示例2精确答案", []int{3, 2, 4}, 6, []int{1, 2}},
		{"示例3精确答案", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.nums, tt.target); !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	// 压力场景：构造长度为 10^4 的数组，让哈希表尽量多积累元素后再命中
	const n = 10000
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i*2 + 1 // 奇数序列 1,3,5,...
	}
	target := nums[0] + nums[n-1] // 命中位置靠后，接近扫描半个数组

	b.Run("万级数组命中靠后", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TwoSum(nums, target)
		}
	})

	b.Run("小数组答案在开头", func(b *testing.B) {
		small := []int{3, 3, 1, 2, 4}
		for i := 0; i < b.N; i++ {
			TwoSum(small, 6)
		}
	})
}
