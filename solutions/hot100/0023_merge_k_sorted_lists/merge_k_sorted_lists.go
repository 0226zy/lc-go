package mergeksortedlists

import (
	"github.com/0226zy/lc-go/pkg/algorithms/sorts"
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// MergeKListsDivideConquer 合并 K 个升序链表 - 分治合并法（最优解）
// 使用分治思想，将 k 个链表两两配对合并
// 时间复杂度: O(n log k)  空间复杂度: O(log k)
func MergeKListsDivideConquer(lists []*datastructures.ListNode) *datastructures.ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	// 分治：将问题分解成两个子问题
	mid := len(lists) / 2
	left := MergeKListsDivideConquer(lists[:mid])
	right := MergeKListsDivideConquer(lists[mid:])

	// 合并两个有序链表
	return mergeTwoLists(left, right)
}

// mergeTwoLists 合并两个有序链表（辅助函数）
// 时间复杂度: O(m+n)  空间复杂度: O(1)
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

// MergeKListsSequential 合并 K 个升序链表 - 顺序合并法
// 逐个合并链表：先合并前两个，再将结果与第三个合并
// 时间复杂度: O(k² * n)  空间复杂度: O(1)
func MergeKListsSequential(lists []*datastructures.ListNode) *datastructures.ListNode {
	if len(lists) == 0 {
		return nil
	}

	result := lists[0]
	for i := 1; i < len(lists); i++ {
		result = mergeTwoLists(result, lists[i])
	}

	return result
}

// MergeKListsMinHeap 合并 K 个升序链表 - 最小堆法（手写堆）
// 使用最小堆维护所有链表的头节点，每次取出最小值
// 时间复杂度: O(n log k)  空间复杂度: O(k)
func MergeKListsMinHeap(lists []*datastructures.ListNode) *datastructures.ListNode {
	// 创建最小堆，使用 pkg/algorithms/sorts 包的通用最小堆
	h := sorts.NewMinHeap(func(a, b *datastructures.ListNode) bool {
		return a.Val < b.Val
	})

	// 将每个非空链表的头节点加入堆
	for _, list := range lists {
		if list != nil {
			h.Push(list)
		}
	}

	// 创建哑节点
	dummy := &datastructures.ListNode{}
	curr := dummy

	// 循环直到堆为空
	for h.Len() > 0 {
		// 从堆中取出最小节点
		minNode := h.Pop()

		// 将最小节点加入结果链表
		curr.Next = minNode
		curr = curr.Next

		// 如果最小节点有下一个节点，将其加入堆
		if minNode.Next != nil {
			h.Push(minNode.Next)
		}
	}

	return dummy.Next
}
