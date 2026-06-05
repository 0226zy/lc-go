package intersectionoftwolinkedlists

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		name   string
		listA  []int
		listB  []int
		skipA  int
		skipB  int
		want   int // 相交节点的值，0 表示不相交
	}{
		// LeetCode 官方示例
		{"示例1: 相交在8", []int{4, 1, 8, 4, 5}, []int{5, 6, 1, 8, 4, 5}, 2, 3, 8},
		{"示例2: 相交在2", []int{1, 9, 1, 2, 4}, []int{3, 2, 4}, 3, 1, 2},
		{"示例3: 不相交", []int{2, 6, 4}, []int{1, 5}, 3, 2, 0},

		// 边界：空链表
		{"空链表A", nil, []int{1, 2, 3}, 0, 0, 0},
		{"空链表B", []int{1, 2, 3}, nil, 0, 0, 0},
		{"两个空链表", nil, nil, 0, 0, 0},

		// 边界：单节点相交
		{"单节点相交", []int{1}, []int{1}, 0, 0, 1},
		{"单节点不相交", []int{1}, []int{2}, 1, 1, 0},

		// 边界：首节点相交
		{"首节点相交", []int{1, 2, 3}, []int{1, 4, 5}, 0, 0, 1},

		// 边界：尾节点相交
		{"尾节点相交", []int{1, 2, 3}, []int{4, 5, 3}, 2, 2, 3},

		// 较大链表
		{"长链表相交", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{99, 98, 7, 8, 9, 10}, 6, 2, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var headA, headB *datastructures.ListNode
			var intersectNode *datastructures.ListNode

			if tt.want == 0 {
				// 不相交
				headA = datastructures.NewLinkedList(tt.listA)
				headB = datastructures.NewLinkedList(tt.listB)
			} else {
				// 创建相交链表
				headA, headB, intersectNode = createIntersectingLists(tt.listA, tt.listB, tt.skipA, tt.skipB)
			}

			got := GetIntersectionNode(headA, headB)

			if tt.want == 0 {
				// 不相交
				if got != nil {
					t.Errorf("GetIntersectionNode() = %v, want nil", got.Val)
				}
			} else {
				// 相交
				if got == nil || got.Val != tt.want {
					if got == nil {
						t.Errorf("GetIntersectionNode() = nil, want %d", tt.want)
					} else {
						t.Errorf("GetIntersectionNode() = %d, want %d", got.Val, tt.want)
					}
				}
				// 验证返回的是同一个节点（地址相同）
				if got != intersectNode {
					t.Errorf("GetIntersectionNode() 返回的节点地址不正确")
				}
			}
		})
	}
}

// createIntersectingLists 创建两个相交的链表
// listA: 链表A的完整值（包含相交部分）
// listB: 链表B的完整值（包含相交部分）
// skipA: 链表A中从头部到相交节点的节点数
// skipB: 链表B中从头部到相交节点的节点数
// 返回: headA, headB, intersectNode（相交节点）
func createIntersectingLists(listA, listB []int, skipA, skipB int) (*datastructures.ListNode, *datastructures.ListNode, *datastructures.ListNode) {
	if len(listA) == 0 || len(listB) == 0 {
		return nil, nil, nil
	}

	// 创建链表A的非相交部分
	var headA, tailA *datastructures.ListNode
	if skipA > 0 {
		headA = &datastructures.ListNode{Val: listA[0]}
		tailA = headA
		for i := 1; i < skipA; i++ {
			tailA.Next = &datastructures.ListNode{Val: listA[i]}
			tailA = tailA.Next
		}
	}

	// 创建相交部分（公共部分）
	intersectNode := &datastructures.ListNode{Val: listA[skipA]}
	tail := intersectNode
	for i := skipA + 1; i < len(listA); i++ {
		tail.Next = &datastructures.ListNode{Val: listA[i]}
		tail = tail.Next
	}

	// 连接链表A
	if tailA != nil {
		tailA.Next = intersectNode
	} else {
		headA = intersectNode
	}

	// 创建链表B
	headB := &datastructures.ListNode{Val: listB[0]}
	tailB := headB
	for i := 1; i < skipB; i++ {
		tailB.Next = &datastructures.ListNode{Val: listB[i]}
		tailB = tailB.Next
	}
	// 连接相交部分
	tailB.Next = intersectNode

	return headA, headB, intersectNode
}

func BenchmarkGetIntersectionNode(b *testing.B) {
	benchmarks := []struct {
		name  string
		listA []int
		listB []int
		skipA int
		skipB int
	}{
		{"相交在中间", []int{1, 2, 3, 4, 5}, []int{6, 7, 3, 4, 5}, 2, 2},
		{"不相交", []int{1, 2, 3}, []int{4, 5, 6}, 3, 3},
		{"长链表", generateList(100), generateList(100), 50, 50},
		{"超长链表", generateList(1000), generateList(1000), 500, 500},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				headA, headB, _ := createIntersectingLists(bm.listA, bm.listB, bm.skipA, bm.skipB)
				GetIntersectionNode(headA, headB)
			}
		})
	}
}

// generateList 生成长度为 n 的链表 [1,2,...,n]
func generateList(n int) []int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i + 1
	}
	return list
}
