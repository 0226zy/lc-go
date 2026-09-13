package insertintobst

import "github.com/0226zy/lc-go/pkg/datastructures"

// InsertIntoBST 二叉搜索树中的插入操作
// 给定二叉搜索树的根节点 root 和一个值 val，将该值插入到二叉搜索树中，
// 保持二叉搜索树性质不变，返回插入后的根节点。保证新值不与原树中任何值重复。
// 时间复杂度: O(h)  h 为树高，平衡时为 O(log n)，链式退化为 O(n)  空间复杂度: O(h) 递归栈深度
func InsertIntoBST(root *datastructures.TreeNode, val int) *datastructures.TreeNode {
	if root == nil {
		return &datastructures.TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = InsertIntoBST(root.Left, val)
	} else {
		root.Right = InsertIntoBST(root.Right, val)
	}
	return root
}
