package mergeksortedlists

import (
	"container/heap"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// MergeKLists 合并 K 个升序链表
// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
// 核心思路：分治——把 k 个链表两两配对合并，一轮之后剩 k/2 个，
// 重复直到只剩 1 个。每次都调用「合并两个有序链表」。
// 时间复杂度: O(N log k)，N 为所有节点总数，共 log k 轮合并
// 空间复杂度: O(log k) 递归栈深度
func MergeKLists(lists []*datastructures.ListNode) *datastructures.ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeRange(lists, 0, len(lists)-1)
}

// mergeRange 分治合并 lists[low..high]
func mergeRange(lists []*datastructures.ListNode, low, high int) *datastructures.ListNode {
	if low == high {
		return lists[low]
	}
	mid := low + (high-low)/2
	left := mergeRange(lists, low, mid)
	right := mergeRange(lists, mid+1, high)
	return mergeTwoLists(left, right)
}

// mergeTwoLists 合并两个升序链表
func mergeTwoLists(l1, l2 *datastructures.ListNode) *datastructures.ListNode {
	dummy := &datastructures.ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}
	return dummy.Next
}

// MergeKListsWithHeap 合并 K 个升序链表（最小堆解法，对照）
// 把 k 个链表的头节点放进最小堆，每次弹出最小节点接到结果链上，
// 并把该节点的下一个节点推入堆。堆中始终最多 k 个元素。
// 时间复杂度: O(N log k)  空间复杂度: O(k) 堆的容量
func MergeKListsWithHeap(lists []*datastructures.ListNode) *datastructures.ListNode {
	h := &nodeHeap{}
	heap.Init(h)
	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}
	dummy := &datastructures.ListNode{}
	curr := dummy
	for h.Len() > 0 {
		node := heap.Pop(h).(*datastructures.ListNode)
		curr.Next = node
		curr = curr.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return dummy.Next
}

// nodeHeap 以节点值排序的最小堆
type nodeHeap []*datastructures.ListNode

func (h nodeHeap) Len() int            { return len(h) }
func (h nodeHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x interface{}) { *h = append(*h, x.(*datastructures.ListNode)) }
func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
