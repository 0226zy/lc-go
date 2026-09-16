package wigglesort

import (
	"sort"
	"testing"
)

// isWiggle 校验数组是否满足 nums[0] <= nums[1] >= nums[2] <= nums[3]... 的摆动条件
func isWiggle(nums []int) bool {
	for i := 0; i+1 < len(nums); i++ {
		if i%2 == 0 && nums[i] > nums[i+1] {
			return false
		}
		if i%2 == 1 && nums[i] < nums[i+1] {
			return false
		}
	}
	return true
}

// isPermutation 校验重排后数组与原数组元素完全一致（多重集合意义下）
func isPermutation(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := make([]int, len(a))
	bCopy := make([]int, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	sort.Ints(aCopy)
	sort.Ints(bCopy)
	for i := range aCopy {
		if aCopy[i] != bCopy[i] {
			return false
		}
	}
	return true
}

func TestWiggleSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		// LeetCode 官方示例
		{"示例1: [3,5,2,1,6,4]", []int{3, 5, 2, 1, 6, 4}},
		{"示例2: [6,6,5,6,3,8] 本身已满足", []int{6, 6, 5, 6, 3, 8}},

		// 边界：空输入与单元素
		{"空数组", []int{}},
		{"单元素", []int{1}},
		{"两个元素升序", []int{1, 2}},
		{"两个元素降序", []int{2, 1}},

		// 边界：大量重复元素
		{"全部相同", []int{7, 7, 7, 7, 7}},
		{"含重复需交换", []int{5, 5, 3, 3, 1, 1}},

		// 典型场景
		{"严格升序", []int{1, 2, 3, 4, 5}},
		{"严格降序", []int{5, 4, 3, 2, 1}},
		{"含零与极值", []int{0, 10000, 0, 10000, 0}},
		{"乱序大数组", []int{9, 3, 7, 1, 8, 2, 6, 4, 5, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origin := make([]int, len(tt.nums))
			copy(origin, tt.nums)

			WiggleSort(tt.nums)

			if !isWiggle(tt.nums) {
				t.Errorf("WiggleSort(%v) 结果 %v 不满足摆动条件", origin, tt.nums)
			}
			if !isPermutation(tt.nums, origin) {
				t.Errorf("WiggleSort(%v) 结果 %v 不是原数组的重排", origin, tt.nums)
			}
		})
	}
}

func BenchmarkWiggleSort(b *testing.B) {
	// 构造 5 万元素的数组
	nums := make([]int, 50000)
	for i := range nums {
		nums[i] = (i * 7919) % 10001
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WiggleSort(nums)
	}
}
