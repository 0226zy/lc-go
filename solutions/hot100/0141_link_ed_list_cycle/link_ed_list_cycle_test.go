package linkedlistcycle

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestHasCycleHash(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		pos    int // 环入口索引，-1 表示无环
		expect bool
	}{
		// LeetCode 官方示例
		{"示例1_有环", []int{3, 2, 0, -4}, 1, true},
		{"示例2_有环", []int{1, 2}, 0, true},
		{"示例3_无环", []int{1}, -1, false},

		// 边界情况
		{"空链表", []int{}, -1, false},
		{"单个节点_无环", []int{1}, -1, false},
		{"单个节点_自环", []int{1}, 0, true},

		// 特殊值
		{"包含负数_有环", []int{-1, -2, -3}, 1, true},
		{"包含零_有环", []int{0, 1, 0}, 0, true},
		{"全部相同值_有环", []int{1, 1, 1}, 0, true},

		// 长链表
		{"长链表_有环", makeRange(1, 100), 50, true},
		{"长链表_无环", makeRange(1, 100), -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewCycleLinkedList(tt.input, tt.pos)
			result := HasCycleHash(head)

			if result != tt.expect {
				t.Errorf("HasCycleHash() = %v, want %v", result, tt.expect)
			}
		})
	}
}

func TestHasCycleTwoPointers(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		pos    int
		expect bool
	}{
		// LeetCode 官方示例
		{"示例1_有环", []int{3, 2, 0, -4}, 1, true},
		{"示例2_有环", []int{1, 2}, 0, true},
		{"示例3_无环", []int{1}, -1, false},

		// 边界情况
		{"空链表", []int{}, -1, false},
		{"单个节点_无环", []int{1}, -1, false},
		{"单个节点_自环", []int{1}, 0, true},

		// 特殊值
		{"包含负数_有环", []int{-1, -2, -3}, 1, true},
		{"包含零_有环", []int{0, 1, 0}, 0, true},
		{"全部相同值_有环", []int{1, 1, 1}, 0, true},

		// 长链表
		{"长链表_有环", makeRange(1, 100), 50, true},
		{"长链表_无环", makeRange(1, 100), -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewCycleLinkedList(tt.input, tt.pos)
			result := HasCycleTwoPointers(head)

			if result != tt.expect {
				t.Errorf("HasCycleTwoPointers() = %v, want %v", result, tt.expect)
			}
		})
	}
}

func TestHasCycleMark(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		pos    int
		expect bool
	}{
		// LeetCode 官方示例
		{"示例1_有环", []int{3, 2, 0, -4}, 1, true},
		{"示例2_有环", []int{1, 2}, 0, true},
		{"示例3_无环", []int{1}, -1, false},

		// 边界情况
		{"空链表", []int{}, -1, false},
		{"单个节点_无环", []int{1}, -1, false},
		{"单个节点_自环", []int{1}, 0, true},

		// 特殊值
		{"包含负数_有环", []int{-1, -2, -3}, 1, true},
		{"包含零_有环", []int{0, 1, 0}, 0, true},
		{"全部相同值_有环", []int{1, 1, 1}, 0, true},

		// 长链表
		{"长链表_有环", makeRange(1, 100), 50, true},
		{"长链表_无环", makeRange(1, 100), -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 注意：标记法会修改链表值，所以每个测试用例需要创建新的链表
			head := datastructures.NewCycleLinkedList(tt.input, tt.pos)
			result := HasCycleMark(head)

			if result != tt.expect {
				t.Errorf("HasCycleMark() = %v, want %v", result, tt.expect)
			}
		})
	}
}

func BenchmarkHasCycleHash(b *testing.B) {
	// 准备测试数据
	testCases := []struct {
		name   string
		input  []int
		pos    int
	}{
		{"短链表_有环", []int{1, 2, 3, 4, 5}, 2},
		{"中等链表_有环", makeRange(1, 100), 50},
		{"长链表_有环", makeRange(1, 1000), 500},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// 每次重新创建链表，因为 Hash 法不会修改链表
				head := datastructures.NewCycleLinkedList(tc.input, tc.pos)
				HasCycleHash(head)
			}
		})
	}
}

func BenchmarkHasCycleTwoPointers(b *testing.B) {
	// 准备测试数据
	testCases := []struct {
		name   string
		input  []int
		pos    int
	}{
		{"短链表_有环", []int{1, 2, 3, 4, 5}, 2},
		{"中等链表_有环", makeRange(1, 100), 50},
		{"长链表_有环", makeRange(1, 1000), 500},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// 每次重新创建链表，因为快慢指针法不会修改链表
				head := datastructures.NewCycleLinkedList(tc.input, tc.pos)
				HasCycleTwoPointers(head)
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
