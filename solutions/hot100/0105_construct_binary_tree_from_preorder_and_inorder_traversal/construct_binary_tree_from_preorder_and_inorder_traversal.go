package buildtreeprein

import "github.com/0226zy/lc-go/pkg/datastructures"

type TreeNode = datastructures.TreeNode

// BuildTree 从前序与中序遍历序列构造二叉树
// 给定两个整数数组 preorder 和 inorder，其中 preorder 是二叉树的先序遍历，
// inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。
// 时间复杂度: O(?)  空间复杂度: O(?)
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	n := len(inorder)
	rootIdx := 0
	for rootIdx < n && inorder[rootIdx] != root.Val {
		rootIdx++
	}
	leftNum:=rootIdx
	root.Left = BuildTree(preorder[1:leftNum+1], inorder[0:rootIdx])
	root.Right = BuildTree(preorder[leftNum+1:], inorder[rootIdx+1:])
	return root
}
