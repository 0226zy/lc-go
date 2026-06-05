package linkedlistcycleii

import "github.com/0226zy/lc-go/pkg/datastructures"

// DetectCycleHash 方法一：哈希表法
// 使用哈希表记录访问过的节点，第一个重复的节点就是环的入口
// 时间复杂度: O(n)  空间复杂度: O(n)
func DetectCycleHash(head *datastructures.ListNode) *datastructures.ListNode {
	visited := make(map[*datastructures.ListNode]bool)
	for curr := head; curr != nil; curr = curr.Next {
		if visited[curr] {
			return curr // 再次访问，说明是环的入口
		}
		visited[curr] = true
	}
	return nil // 遍历到 nil，无环
}

// DetectCycleTwoPointers 方法二：快慢指针法（Floyd 算法）
// 分两个阶段：1. 找到相遇点；2. 找到环的入口
// 时间复杂度: O(n)  空间复杂度: O(1)
func DetectCycleTwoPointers(head *datastructures.ListNode) *datastructures.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 阶段一：找到相遇点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break // 找到相遇点
		}
	}

	// 如果 fast 或 fast.Next 为 nil，说明无环
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 阶段二：找到环的入口
	slow = head // 将 slow 移回头部
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow // 相遇点就是环的入口
}

// DetectCycleMark 方法三：标记法（修改链表节点值）
// 将访问过的节点值改为特殊标记值，第一个再次遇到标记值的节点就是环的入口
// 注意：此方法会修改链表原始数据，不推荐在实际项目中使用
// 时间复杂度: O(n)  空间复杂度: O(1)
func DetectCycleMark(head *datastructures.ListNode) *datastructures.ListNode {
	const marker = -100001 // 题目给定节点值范围是 [-10^5, 10^5]，用范围外的值作为标记

	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val == marker {
			return curr // 再次遇到标记值，说明是环的入口
		}
		curr.Val = marker // 标记为已访问
	}

	return nil // 遍历到 nil，无环
}
