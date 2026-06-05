package linkedlistcycleii

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// createCycleLinkedList 创建带环链表，返回头节点和环的入口节点
// pos 为环入口索引，-1 表示无环
func createCycleLinkedList(vals []int, pos int) (head *datastructures.ListNode, cycleEntry *datastructures.ListNode) {
	head = datastructures.NewLinkedList(vals)
	if pos < 0 || head == nil {
		return head, nil
	}

	// 找到环的入口节点
	curr := head
	for i := 0; curr != nil; i++ {
		if i == pos {
			cycleEntry = curr
		}
		if curr.Next == nil {
			// 将尾节点连接到环入口
			if cycleEntry != nil {
				curr.Next = cycleEntry
			}
			break
		}
		curr = curr.Next
	}

	return head, cycleEntry
}

func TestDetectCycleHash(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		pos          int
		expectCycle  bool // 是否有环
	}{
		// LeetCode 官方示例
		{"示例1_有环_pos=1", []int{3, 2, 0, -4}, 1, true},
		{"示例2_有环_pos=0", []int{1, 2}, 0, true},
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
			head, cycleEntry := createCycleLinkedList(tt.input, tt.pos)
			result := DetectCycleHash(head)

			if !tt.expectCycle {
				// 无环，期望返回 nil
				if result != nil {
					t.Errorf("DetectCycleHash() = %v, want nil", result.Val)
				}
			} else {
				// 有环，期望返回环的入口节点（指针比较）
				if result != cycleEntry {
					if result == nil {
						t.Errorf("DetectCycleHash() = nil, want cycle entry (pos=%d)", tt.pos)
					} else {
						t.Errorf("DetectCycleHash() = node with val %d, want cycle entry with val %d (pointer mismatch)", result.Val, cycleEntry.Val)
					}
				}
			}
		})
	}
}

func TestDetectCycleTwoPointers(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		pos          int
		expectCycle  bool
	}{
		// LeetCode 官方示例
		{"示例1_有环_pos=1", []int{3, 2, 0, -4}, 1, true},
		{"示例2_有环_pos=0", []int{1, 2}, 0, true},
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
			head, cycleEntry := createCycleLinkedList(tt.input, tt.pos)
			result := DetectCycleTwoPointers(head)

			if !tt.expectCycle {
				// 无环，期望返回 nil
				if result != nil {
					t.Errorf("DetectCycleTwoPointers() = %v, want nil", result.Val)
				}
			} else {
				// 有环，期望返回环的入口节点（指针比较）
				if result != cycleEntry {
					if result == nil {
						t.Errorf("DetectCycleTwoPointers() = nil, want cycle entry (pos=%d)", tt.pos)
					} else {
						t.Errorf("DetectCycleTwoPointers() = node with val %d, want cycle entry with val %d (pointer mismatch)", result.Val, cycleEntry.Val)
					}
				}
			}
		})
	}
}

func TestDetectCycleMark(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		pos          int
		expectCycle  bool
	}{
		// LeetCode 官方示例
		{"示例1_有环_pos=1", []int{3, 2, 0, -4}, 1, true},
		{"示例2_有环_pos=0", []int{1, 2}, 0, true},
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
			head, cycleEntry := createCycleLinkedList(tt.input, tt.pos)
			result := DetectCycleMark(head)

			if !tt.expectCycle {
				// 无环，期望返回 nil
				if result != nil {
					t.Errorf("DetectCycleMark() = %v, want nil", result.Val)
				}
			} else {
				// 有环，期望返回环的入口节点（指针比较）
				if result != cycleEntry {
					if result == nil {
						t.Errorf("DetectCycleMark() = nil, want cycle entry (pos=%d)", tt.pos)
					} else {
						t.Errorf("DetectCycleMark() = node with val %d, want cycle entry with val %d (pointer mismatch)", result.Val, cycleEntry.Val)
					}
				}
			}
		})
	}
}

func BenchmarkDetectCycleHash(b *testing.B) {
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
				// 每次重新创建链表
				head, _ := createCycleLinkedList(tc.input, tc.pos)
				DetectCycleHash(head)
			}
		})
	}
}

func BenchmarkDetectCycleTwoPointers(b *testing.B) {
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
				// 每次重新创建链表
				head, _ := createCycleLinkedList(tc.input, tc.pos)
				DetectCycleTwoPointers(head)
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
