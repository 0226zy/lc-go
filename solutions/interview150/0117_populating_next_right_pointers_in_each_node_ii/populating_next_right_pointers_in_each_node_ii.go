package populatingnextright

// Node 带 Next 指针的二叉树节点
// pkg/datastructures 中没有该类型，按题目要求在本包内定义
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// Connect 填充每个节点的下一个右侧节点指针 II
// 给定一棵二叉树的根节点 root，将每个节点的 next 指针指向其同一层右侧的下一个节点；
// 如果该节点已经是本层最右侧节点，则 next 置为 nil。
// 要求使用 O(1) 额外空间（递归栈不计）。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(1) 只使用常数额外指针
func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	// curr 指向当前层的最左节点
	curr := root
	for curr != nil {
		// dummy 是下一层的虚拟头节点，tail 是下一层已连接链表的尾节点
		dummy := &Node{}
		tail := dummy
		// 沿着当前层的 next 链逐个访问，把下一层的节点按从左到右顺序串起来
		for node := curr; node != nil; node = node.Next {
			if node.Left != nil {
				tail.Next = node.Left
				tail = tail.Next
			}
			if node.Right != nil {
				tail.Next = node.Right
				tail = tail.Next
			}
		}
		// 进入下一层
		curr = dummy.Next
	}
	return root
}
