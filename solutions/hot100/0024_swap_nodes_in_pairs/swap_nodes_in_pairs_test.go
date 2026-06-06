package swapnodesinpairs

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

// TestSwapPairsIterative 测试迭代法
func TestSwapPairsIterative(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"示例1_四个节点", []int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{"示例2_空链表", []int{}, []int{}},
		{"示例3_单个节点", []int{1}, []int{1}},
		{"两个节点", []int{1, 2}, []int{2, 1}},
		{"三个节点_奇数长度", []int{1, 2, 3}, []int{2, 1, 3}},
		{"五个节点_奇数长度", []int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 5}},
		{"六个节点_偶数长度", []int{1, 2, 3, 4, 5, 6}, []int{2, 1, 4, 3, 6, 5}},
		{"所有节点值相同", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *datastructures.ListNode
			if len(tt.input) > 0 {
				head = datastructures.NewLinkedList(tt.input)
			}

			result := SwapPairsIterative(head)

			// 转换为切片进行比较
			var resultSlice []int
			if result != nil {
				resultSlice = result.ToSlice()
			}

			if !utils.EqualIntSlice(resultSlice, tt.expect) {
				t.Errorf("SwapPairsIterative(%v) = %v, want %v", tt.input, resultSlice, tt.expect)
			}
		})
	}
}

// TestSwapPairsRecursive 测试递归法
func TestSwapPairsRecursive(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		{"示例1_四个节点", []int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{"示例2_空链表", []int{}, []int{}},
		{"示例3_单个节点", []int{1}, []int{1}},
		{"两个节点", []int{1, 2}, []int{2, 1}},
		{"三个节点_奇数长度", []int{1, 2, 3}, []int{2, 1, 3}},
		{"五个节点_奇数长度", []int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 5}},
		{"六个节点_偶数长度", []int{1, 2, 3, 4, 5, 6}, []int{2, 1, 4, 3, 6, 5}},
		{"所有节点值相同", []int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *datastructures.ListNode
			if len(tt.input) > 0 {
				head = datastructures.NewLinkedList(tt.input)
			}

			result := SwapPairsRecursive(head)

			// 转换为切片进行比较
			var resultSlice []int
			if result != nil {
				resultSlice = result.ToSlice()
			}

			if !utils.EqualIntSlice(resultSlice, tt.expect) {
				t.Errorf("SwapPairsRecursive(%v) = %v, want %v", tt.input, resultSlice, tt.expect)
			}
		})
	}
}

// BenchmarkSwapPairsIterative 迭代法性能测试
func BenchmarkSwapPairsIterative(b *testing.B) {
	// 准备测试数据：100 个节点
	input := make([]int, 100)
	for i := range input {
		input[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 每次需要重新创建链表，因为原链表会被修改
		head := datastructures.NewLinkedList(input)
		SwapPairsIterative(head)
	}
}

// BenchmarkSwapPairsRecursive 递归法性能测试
func BenchmarkSwapPairsRecursive(b *testing.B) {
	// 准备测试数据：100 个节点
	input := make([]int, 100)
	for i := range input {
		input[i] = i + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 每次需要重新创建链表，因为原链表会被修改
		head := datastructures.NewLinkedList(input)
		SwapPairsRecursive(head)
	}
}
