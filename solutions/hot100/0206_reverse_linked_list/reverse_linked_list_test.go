package reverselinkedlist

import (
	"fmt"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		expect []int
	}{
		// LeetCode 官方示例
		{"示例1_五个节点", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"示例2_两个节点", []int{1, 2}, []int{2, 1}},
		{"示例3_空链表", []int{}, []int{}},

		// 边界情况
		{"单个节点", []int{1}, []int{1}},
		{"三个节点", []int{1, 2, 3}, []int{3, 2, 1}},

		// 特殊值
		{"包含负数", []int{-1, -2, -3}, []int{-3, -2, -1}},
		{"包含零", []int{0, 1, 0}, []int{0, 1, 0}},
		{"相同值", []int{1, 1, 1}, []int{1, 1, 1}},

		// 长链表
		{"长链表_10个节点", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var head *datastructures.ListNode
			if len(tt.input) > 0 {
				head = datastructures.NewLinkedList(tt.input)
			}

			result := ReverseList(head)

			// 转换为切片进行比较
			var resultSlice []int
			if result != nil {
				resultSlice = result.ToSlice()
			}

			if !utils.EqualIntSlice(resultSlice, tt.expect) {
				t.Errorf("ReverseList(%v) = %v, want %v", tt.input, resultSlice, tt.expect)
			}
		})
	}
}

func BenchmarkReverseList(b *testing.B) {
	// 准备测试数据
	testCases := [][]int{
		makeRange(1, 10),    // 10 个节点
		makeRange(1, 100),   // 100 个节点
		makeRange(1, 1000),  // 1000 个节点
		makeRange(1, 10000), // 10000 个节点
	}

	for _, tc := range testCases {
		b.Run(ttName(len(tc)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				head := datastructures.NewLinkedList(tc)
				ReverseList(head)
			}
		})
	}
}

// makeRange 生成 [start, end] 的整数切片
func makeRange(start, end int) []int {
	result := make([]int, end-start+1)
	for i := range result {
		result[i] = start + i
	}
	return result
}

// ttName 根据节点数生成测试名称
func ttName(n int) string {
	return fmt.Sprintf("链表长度_%d", n)
}
