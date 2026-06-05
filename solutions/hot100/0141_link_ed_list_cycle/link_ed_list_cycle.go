package linkedlistcycle

import "github.com/0226zy/lc-go/pkg/datastructures"

// HasCycleHash 方法一：哈希表法
// 使用哈希表记录访问过的节点，如果再次遇到则说明有环
// 时间复杂度: O(n)  空间复杂度: O(n)
func HasCycleHash(head *datastructures.ListNode) bool {
	visited := make(map[*datastructures.ListNode]bool)
	for curr := head; curr != nil; curr = curr.Next {
		if visited[curr] {
			return true // 再次访问，说明有环
		}
		visited[curr] = true
	}
	return false // 遍历到 nil，无环
}

// HasCycleTwoPointers 方法二：快慢指针法（Floyd 判圈算法）
// 快指针每次走两步，慢指针每次走一步，如果有环则两者一定会相遇
// 时间复杂度: O(n)  空间复杂度: O(1)
func HasCycleTwoPointers(head *datastructures.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next          // 慢指针走一步
		fast = fast.Next.Next    // 快指针走两步
		if slow == fast {        // 相遇，说明有环
			return true
		}
	}
	return false // fast 到达尾部，无环
}

// HasCycleMark 方法三：标记法（修改链表节点值）
// 将访问过的节点值改为特殊标记值，如果再次遇到标记值则说明有环
// 注意：此方法会修改链表原始数据，不推荐在实际项目中使用
// 时间复杂度: O(n)  空间复杂度: O(1)
func HasCycleMark(head *datastructures.ListNode) bool {
	const marker = -100001 // 题目给定节点值范围是 [-10^5, 10^5]，用范围外的值作为标记
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val == marker {
			return true // 再次遇到标记值，说明有环
		}
		curr.Val = marker // 标记为已访问
	}
	return false // 遍历到 nil，无环
}
