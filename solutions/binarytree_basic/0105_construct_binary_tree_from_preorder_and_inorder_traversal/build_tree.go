package buildtree

import "github.com/0226zy/lc-go/pkg/datastructures"

// BuildTree 从前序与中序遍历序列构造二叉树
// 给定二叉树的前序遍历 preorder 和中序遍历 inorder（节点值互不相同），还原二叉树并返回根节点。
// 时间复杂度: O(n)  空间复杂度: O(n)
func BuildTree(preorder []int, inorder []int) *datastructures.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// pos[v] 记录值 v 在中序遍历中的下标，O(1) 定位根
	pos := make(map[int]int, len(inorder))
	for i, v := range inorder {
		pos[v] = i
	}

	preIdx := 0 // 前序序列的全局消费指针，始终指向当前子树的根
	var build func(il, ir int) *datastructures.TreeNode
	build = func(il, ir int) *datastructures.TreeNode {
		if il > ir { // 中序区间为空，说明是空子树
			return nil
		}
		rootVal := preorder[preIdx]
		preIdx++
		root := &datastructures.TreeNode{Val: rootVal}
		k := pos[rootVal] // 根把中序区间切成左右两半
		// 必须先左后右：前序序列中左子树的节点排在右子树之前
		root.Left = build(il, k-1)
		root.Right = build(k+1, ir)
		return root
	}
	return build(0, len(inorder)-1)
}
