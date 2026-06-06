package sortlist

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

func TestSortListTopDown(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		// 示例 1
		{"示例1", []int{4, 2, 1, 3}, []int{1, 2, 3, 4}},

		// 示例 2
		{"示例2", []int{-1, 5, 3, 4, 0}, []int{-1, 0, 3, 4, 5}},

		// 示例 3：空链表
		{"空链表", nil, []int{}},

		// 边界：单节点
		{"单节点", []int{1}, []int{1}},

		// 边界：两个节点
		{"两个节点_逆序", []int{2, 1}, []int{1, 2}},
		{"两个节点_已排序", []int{1, 2}, []int{1, 2}},

		// 正常：已排序
		{"已排序", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},

		// 正常：逆序
		{"逆序", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},

		// 正常：有重复元素
		{"有重复元素", []int{3, 1, 2, 1, 3}, []int{1, 1, 2, 3, 3}},

		// 边界：全为负数的逆序
		{"全负数_逆序", []int{-1, -3, -2, -5, -4}, []int{-5, -4, -3, -2, -1}},

		// 边界：混合正负数
		{"混合正负数", []int{0, -1, 2, -2, 1}, []int{-2, -1, 0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *datastructures.ListNode
			if tt.nums != nil {
				head = datastructures.NewLinkedList(tt.nums)
			}

			result := SortListTopDown(head)
			resultSlice := []int{}
			if result != nil {
				resultSlice = result.ToSlice()
			}

			if !utils.EqualIntSlice(resultSlice, tt.expected) {
				t.Errorf("SortListTopDown(%v) = %v, expected %v",
					tt.nums, resultSlice, tt.expected)
			}
		})
	}
}

func TestSortListBottomUp(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		// 示例 1
		{"示例1", []int{4, 2, 1, 3}, []int{1, 2, 3, 4}},

		// 示例 2
		{"示例2", []int{-1, 5, 3, 4, 0}, []int{-1, 0, 3, 4, 5}},

		// 示例 3：空链表
		{"空链表", nil, []int{}},

		// 边界：单节点
		{"单节点", []int{1}, []int{1}},

		// 边界：两个节点
		{"两个节点_逆序", []int{2, 1}, []int{1, 2}},
		{"两个节点_已排序", []int{1, 2}, []int{1, 2}},

		// 正常：已排序
		{"已排序", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},

		// 正常：逆序
		{"逆序", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},

		// 正常：有重复元素
		{"有重复元素", []int{3, 1, 2, 1, 3}, []int{1, 1, 2, 3, 3}},

		// 边界：全为负数的逆序
		{"全负数_逆序", []int{-1, -3, -2, -5, -4}, []int{-5, -4, -3, -2, -1}},

		// 边界：混合正负数
		{"混合正负数", []int{0, -1, 2, -2, 1}, []int{-2, -1, 0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *datastructures.ListNode
			if tt.nums != nil {
				head = datastructures.NewLinkedList(tt.nums)
			}

			result := SortListBottomUp(head)
			resultSlice := []int{}
			if result != nil {
				resultSlice = result.ToSlice()
			}

			if !utils.EqualIntSlice(resultSlice, tt.expected) {
				t.Errorf("SortListBottomUp(%v) = %v, expected %v",
					tt.nums, resultSlice, tt.expected)
			}
		})
	}
}

// BenchmarkSortListTopDown 自顶向下归并排序性能测试
func BenchmarkSortListTopDown(b *testing.B) {
	// 准备测试数据：1000 个随机节点
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = 1000 - i // 逆序，最坏情况
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		head := datastructures.NewLinkedList(nums)
		result := SortListTopDown(head)
		_ = result // 避免未使用变量警告
	}
}

// BenchmarkSortListBottomUp 自底向上归并排序性能测试
func BenchmarkSortListBottomUp(b *testing.B) {
	// 准备测试数据：1000 个随机节点
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = 1000 - i // 逆序，最坏情况
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		head := datastructures.NewLinkedList(nums)
		result := SortListBottomUp(head)
		_ = result // 避免未使用变量警告
	}
}
