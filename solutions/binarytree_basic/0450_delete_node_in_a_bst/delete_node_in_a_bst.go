package deletenodeinabst

import "github.com/0226zy/lc-go/pkg/datastructures"

// DeleteNode 删除二叉搜索树中的节点
// 给定一棵二叉搜索树的根节点 root 和一个值 key，删除 BST 中值等于 key 的节点，
// 并返回删除后仍然合法的 BST 的根节点；若 key 不存在则原树不变。
// 时间复杂度: O(h) h 为树高，平衡时 O(log n)，退化为链时 O(n)  空间复杂度: O(h) 递归栈深度
func DeleteNode(root *datastructures.TreeNode, key int) *datastructures.TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = DeleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = DeleteNode(root.Right, key)
	} else {
		// 找到待删除节点，分三种情况处理
		if root.Left == nil {
			// 情况一/二：无左孩子（叶节点或只有右孩子），用右孩子顶替
			return root.Right
		}
		if root.Right == nil {
			// 情况二：只有左孩子，用左孩子顶替
			return root.Left
		}
		// 情况三：双子节点，取右子树的最小节点（中序后继）顶替
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		root.Val = successor.Val
		// 在右子树中递归删除后继节点
		root.Right = DeleteNode(root.Right, successor.Val)
	}
	return root
}
