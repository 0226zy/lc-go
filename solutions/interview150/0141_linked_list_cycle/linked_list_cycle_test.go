package linkedlistcycle

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pos  int // 环入口索引，-1 表示无环
		want bool
	}{
		// LeetCode 官方示例
		{"示例1: [3,2,0,-4] 环在索引1", []int{3, 2, 0, -4}, 1, true},
		{"示例2: [1,2] 环在索引0", []int{1, 2}, 0, true},
		{"示例3: [1] 无环", []int{1}, -1, false},

		// 边界：空链表
		{"空链表", []int{}, -1, false},

		// 边界：单节点无环
		{"单节点无环", []int{7}, -1, false},

		// 边界：单节点自环
		{"单节点自环", []int{7}, 0, true},

		// 边界：两个节点无环
		{"两节点无环", []int{1, 2}, -1, false},

		// 边界：尾节点指向尾节点自身
		{"尾节点自环", []int{1, 2, 3, 4}, 3, true},

		// 边界：环入口在头部
		{"环入口在头部", []int{5, 6, 7, 8}, 0, true},

		// 边界：较长链表无环
		{"长链表无环", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1, false},

		// 边界：较长链表环在中间
		{"长链表环在中间", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, true},

		// 边界：负值节点有环
		{"负值节点有环", []int{-10, -5, 0, -4}, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewCycleLinkedList(tt.vals, tt.pos)
			if got := HasCycle(head); got != tt.want {
				t.Errorf("HasCycle(%v, pos=%d) = %v, want %v", tt.vals, tt.pos, got, tt.want)
			}
		})
	}
}

func BenchmarkHasCycle(b *testing.B) {
	// 无环长链表，快指针需走到末尾，是最长路径
	head := datastructures.NewCycleLinkedList(makeRange(10000), -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasCycle(head)
	}
}

func BenchmarkHasCycleWithCycle(b *testing.B) {
	// 整链成环（pos=0），慢指针需走完环外+一圈才相遇
	head := datastructures.NewCycleLinkedList(makeRange(10000), 0)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasCycle(head)
	}
}

// makeRange 生成 [1, n] 的连续值切片
func makeRange(n int) []int {
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		vals[i] = i + 1
	}
	return vals
}
