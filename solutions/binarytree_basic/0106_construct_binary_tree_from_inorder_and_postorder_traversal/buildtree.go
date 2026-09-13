package buildtree

import "github.com/0226zy/lc-go/pkg/datastructures"

// BuildTree 从中序与后序遍历序列构造二叉树
// 给定两个整数数组 inorder 和 postorder（值互不相同），其中 inorder 是二叉树的中序遍历，
// postorder 是同一棵树的后序遍历，还原这棵二叉树并返回其根节点。
// 时间复杂度: O(n) 每个节点恰好创建一次  空间复杂度: O(n) 哈希表开销 + O(h) 递归栈
func BuildTree(inorder []int, postorder []int) *datastructures.TreeNode {
	// 哈希表记录每个值在中序遍历中的位置，避免递归时线性查找根
	pos := make(map[int]int, len(inorder))
	for i, v := range inorder {
		pos[v] = i
	}

	idx := len(postorder) - 1 // 后序末元素是根，从末尾往前依次取子树根

	// build 用中序区间 [lo, hi] 还原一棵子树，全局指针 idx 同步消耗后序序列
	var build func(lo, hi int) *datastructures.TreeNode
	build = func(lo, hi int) *datastructures.TreeNode {
		if lo > hi {
			return nil
		}
		root := &datastructures.TreeNode{Val: postorder[idx]}
		idx--
		k := pos[root.Val]
		// 后序倒序时右子树的根先出现，因此必须先建右子树再建左子树
		root.Right = build(k+1, hi)
		root.Left = build(lo, k-1)
		return root
	}
	return build(0, len(inorder)-1)
}
