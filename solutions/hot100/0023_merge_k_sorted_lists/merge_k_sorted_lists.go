package mergeksortedlists

import "github.com/0226zy/lc-go/pkg/datastructures"

// MinHeap 手写最小堆，参考 container/heap 的接口设计
type MinHeap struct {
	data []*datastructures.ListNode
}

// NewMinHeap 创建并初始化最小堆
func NewMinHeap() *MinHeap {
	return &MinHeap{data: make([]*datastructures.ListNode, 0)}
}

// Len 返回堆的大小
func (h *MinHeap) Len() int {
	return len(h.data)
}

// Less 比较两个节点的值
func (h *MinHeap) Less(i, j int) bool {
	return h.data[i].Val < h.data[j].Val
}

// Swap 交换两个节点的位置
func (h *MinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// Push 向堆中插入一个节点，O(log k)
func (h *MinHeap) Push(node *datastructures.ListNode) {
	h.data = append(h.data, node)
	h.heapifyUp(h.Len() - 1)
}

// Pop 从堆中弹出最小值，O(log k)
func (h *MinHeap) Pop() *datastructures.ListNode {
	if h.Len() == 0 {
		return nil
	}
	// 将堆顶与最后一个元素交换
	minNode := h.data[0]
	lastIdx := h.Len() - 1
	h.data[0] = h.data[lastIdx]
	h.data = h.data[:lastIdx]
	// 下沉调整
	h.heapifyDown(0)
	return minNode
}

// heapifyUp 自下而上调整堆
func (h *MinHeap) heapifyUp(idx int) {
	for idx > 0 {
		parent := (idx - 1) / 2
		if !h.Less(idx, parent) {
			break
		}
		h.Swap(idx, parent)
		idx = parent
	}
}

// heapifyDown 自上而下调整堆
func (h *MinHeap) heapifyDown(idx int) {
	n := h.Len()
	for {
		left := 2*idx + 1
		right := 2*idx + 2
		smallest := idx

		if left < n && h.Less(left, smallest) {
			smallest = left
		}
		if right < n && h.Less(right, smallest) {
			smallest = right
		}
		if smallest == idx {
			break
		}
		h.Swap(idx, smallest)
		idx = smallest
	}
}

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
	// 创建最小堆
	h := NewMinHeap()

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
