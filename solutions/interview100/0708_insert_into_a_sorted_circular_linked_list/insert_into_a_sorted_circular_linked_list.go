package insertintoasortedcircularlinkedlist

import "github.com/0226zy/lc-go/pkg/datastructures"

// Insert 循环有序列表的插入
// 给定升序循环链表中的任意节点 head 和待插入值 insertVal，
// 插入新值使链表仍保持循环有序，返回链表中的任意节点。
// 时间复杂度: O(n)  空间复杂度: O(1)
func Insert(head *datastructures.ListNode, insertVal int) *datastructures.ListNode {
	node := &datastructures.ListNode{Val: insertVal}
	// 边界：空链表，新节点自成一环
	if head == nil {
		node.Next = node
		return node
	}

	cur := head
	for {
		next := cur.Next
		// 情况一：常规位置 cur.Val <= insertVal <= next.Val
		if cur.Val <= insertVal && insertVal <= next.Val {
			break
		}
		// 情况二：cur 是最大值、next 是最小值的断点，
		// 新值成为新的最大值或最小值时插在断点处
		if cur.Val > next.Val && (insertVal >= cur.Val || insertVal <= next.Val) {
			break
		}
		cur = next
		// 转满一圈仍未找到（所有节点值相等），任意位置插入均可
		if cur == head {
			break
		}
	}
	node.Next = cur.Next
	cur.Next = node
	return head
}
