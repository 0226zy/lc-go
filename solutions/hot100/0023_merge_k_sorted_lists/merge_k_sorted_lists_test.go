package mergeksortedlists

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

// TestMergeKListsDivideConquer 测试分治合并法
func TestMergeKListsDivideConquer(t *testing.T) {
	tests := []struct {
		name   string
		lists  [][]int
		expect []int
	}{
		{
			name:   "示例1: 三个链表",
			lists:  [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			expect: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:   "空输入",
			lists:  [][]int{},
			expect: []int{},
		},
		{
			name:   "单个链表",
			lists:  [][]int{{1, 2, 3}},
			expect: []int{1, 2, 3},
		},
		{
			name:   "两个链表",
			lists:  [][]int{{1, 3, 5}, {2, 4, 6}},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "包含空链表",
			lists:  [][]int{{1, 2, 3}, {}, {4, 5, 6}},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "所有链表都为空",
			lists:  [][]int{{}, {}, {}},
			expect: []int{},
		},
		{
			name:   "链表数组中有 nil",
			lists:  [][]int{{1, 2, 3}, nil, {4, 5, 6}},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "单个空链表",
			lists:  [][]int{{}},
			expect: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建链表数组
			lists := make([]*datastructures.ListNode, len(tt.lists))
			for i, vals := range tt.lists {
				if vals == nil {
					lists[i] = nil
				} else if len(vals) == 0 {
					lists[i] = nil
				} else {
					lists[i] = datastructures.NewLinkedList(vals)
				}
			}

			// 执行合并
			result := MergeKListsDivideConquer(lists)

			// 验证结果
			resultSlice := listToSlice(result)
			if !utils.EqualIntSlice(resultSlice, tt.expect) {
				t.Errorf("MergeKListsDivideConquer() = %v, expect %v", resultSlice, tt.expect)
			}
		})
	}
}

// TestMergeKListsSequential 测试顺序合并法
func TestMergeKListsSequential(t *testing.T) {
	tests := []struct {
		name   string
		lists  [][]int
		expect []int
	}{
		{
			name:   "示例1: 三个链表",
			lists:  [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			expect: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:   "空输入",
			lists:  [][]int{},
			expect: []int{},
		},
		{
			name:   "单个链表",
			lists:  [][]int{{1, 2, 3}},
			expect: []int{1, 2, 3},
		},
		{
			name:   "两个链表",
			lists:  [][]int{{1, 3, 5}, {2, 4, 6}},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "包含空链表",
			lists:  [][]int{{1, 2, 3}, {}, {4, 5, 6}},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建链表数组
			lists := make([]*datastructures.ListNode, len(tt.lists))
			for i, vals := range tt.lists {
				if vals == nil {
					lists[i] = nil
				} else if len(vals) == 0 {
					lists[i] = nil
				} else {
					lists[i] = datastructures.NewLinkedList(vals)
				}
			}

			// 执行合并
			result := MergeKListsSequential(lists)

			// 验证结果
			resultSlice := listToSlice(result)
			if !utils.EqualIntSlice(resultSlice, tt.expect) {
				t.Errorf("MergeKListsSequential() = %v, expect %v", resultSlice, tt.expect)
			}
		})
	}
}

// TestMergeKListsMinHeap 测试最小堆法
func TestMergeKListsMinHeap(t *testing.T) {
	tests := []struct {
		name   string
		lists  [][]int
		expect []int
	}{
		{
			name:   "示例1: 三个链表",
			lists:  [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			expect: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:   "空输入",
			lists:  [][]int{},
			expect: []int{},
		},
		{
			name:   "单个链表",
			lists:  [][]int{{1, 2, 3}},
			expect: []int{1, 2, 3},
		},
		{
			name:   "两个链表",
			lists:  [][]int{{1, 3, 5}, {2, 4, 6}},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:   "包含空链表",
			lists:  [][]int{{1, 2, 3}, {}, {4, 5, 6}},
			expect: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建链表数组
			lists := make([]*datastructures.ListNode, len(tt.lists))
			for i, vals := range tt.lists {
				if vals == nil {
					lists[i] = nil
				} else if len(vals) == 0 {
					lists[i] = nil
				} else {
					lists[i] = datastructures.NewLinkedList(vals)
				}
			}

			// 执行合并
			result := MergeKListsMinHeap(lists)

			// 验证结果
			resultSlice := listToSlice(result)
			if !utils.EqualIntSlice(resultSlice, tt.expect) {
				t.Errorf("MergeKListsMinHeap() = %v, expect %v", resultSlice, tt.expect)
			}
		})
	}
}

// BenchmarkMergeKListsDivideConquer 基准测试 - 分治合并法
func BenchmarkMergeKListsDivideConquer(b *testing.B) {
	// 创建 100 个链表，每个链表 100 个节点
	lists := make([]*datastructures.ListNode, 100)
	for i := 0; i < 100; i++ {
		vals := make([]int, 100)
		for j := 0; j < 100; j++ {
			vals[j] = i*100 + j
		}
		lists[i] = datastructures.NewLinkedList(vals)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeKListsDivideConquer(lists)
	}
}

// BenchmarkMergeKListsSequential 基准测试 - 顺序合并法
func BenchmarkMergeKListsSequential(b *testing.B) {
	// 创建 100 个链表，每个链表 100 个节点
	lists := make([]*datastructures.ListNode, 100)
	for i := 0; i < 100; i++ {
		vals := make([]int, 100)
		for j := 0; j < 100; j++ {
			vals[j] = i*100 + j
		}
		lists[i] = datastructures.NewLinkedList(vals)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeKListsSequential(lists)
	}
}

// BenchmarkMergeKListsMinHeap 基准测试 - 最小堆法
func BenchmarkMergeKListsMinHeap(b *testing.B) {
	// 创建 100 个链表，每个链表 100 个节点
	lists := make([]*datastructures.ListNode, 100)
	for i := 0; i < 100; i++ {
		vals := make([]int, 100)
		for j := 0; j < 100; j++ {
			vals[j] = i*100 + j
		}
		lists[i] = datastructures.NewLinkedList(vals)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeKListsMinHeap(lists)
	}
}

// listToSlice 将链表转换为切片（辅助函数）
func listToSlice(head *datastructures.ListNode) []int {
	var result []int
	for curr := head; curr != nil; curr = curr.Next {
		result = append(result, curr.Val)
	}
	return result
}
