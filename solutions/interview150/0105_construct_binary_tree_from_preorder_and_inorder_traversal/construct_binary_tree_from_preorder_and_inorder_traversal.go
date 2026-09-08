package constructbinarytreefrompreorderandinordertraversal

import "github.com/0226zy/lc-go/pkg/datastructures"

// BuildTree 从前序与中序遍历序列构造二叉树
// 给定二叉树的先序遍历 preorder 和中序遍历 inorder（值互不相同），还原二叉树并返回根节点。
// 时间复杂度: O(n) 每个节点创建一次  空间复杂度: O(n) 哈希表 + 递归栈 O(h)
func BuildTree(preorder, inorder []int) *datastructures.TreeNode {
	n := len(preorder)
	indexMap := make(map[int]int, n)
	for i, v := range inorder {
		indexMap[v] = i
	}

	// build 用先序区间 [pl, pr] 与中序区间 [il, ir] 还原子树
	var build func(pl, pr, il, ir int) *datastructures.TreeNode
	build = func(pl, pr, il, ir int) *datastructures.TreeNode {
		if pl > pr {
			return nil
		}
		root := &datastructures.TreeNode{Val: preorder[pl]}
		k := indexMap[preorder[pl]] // 根在中序中的位置
		leftSize := k - il          // 左子树节点数
		root.Left = build(pl+1, pl+leftSize, il, k-1)
		root.Right = build(pl+leftSize+1, pr, k+1, ir)
		return root
	}
	return build(0, n-1, 0, n-1)
}
