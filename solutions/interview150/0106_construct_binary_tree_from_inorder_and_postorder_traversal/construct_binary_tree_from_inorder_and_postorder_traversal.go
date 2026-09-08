package constructbinarytreefrominorderandpostordertraversal

import "github.com/0226zy/lc-go/pkg/datastructures"

// BuildTree 从中序与后序遍历序列构造二叉树
// 给定二叉树的中序遍历 inorder 和后序遍历 postorder（值互不相同），还原二叉树并返回根节点。
// 时间复杂度: O(n) 每个节点创建一次  空间复杂度: O(n) 哈希表 + 递归栈 O(h)
func BuildTree(inorder, postorder []int) *datastructures.TreeNode {
	n := len(inorder)
	indexMap := make(map[int]int, n)
	for i, v := range inorder {
		indexMap[v] = i
	}

	// build 用中序区间 [il, ir] 与后序区间 [pl, pr] 还原子树
	var build func(il, ir, pl, pr int) *datastructures.TreeNode
	build = func(il, ir, pl, pr int) *datastructures.TreeNode {
		if il > ir {
			return nil
		}
		root := &datastructures.TreeNode{Val: postorder[pr]} // 后序末元素为根
		k := indexMap[postorder[pr]]                         // 根在中序中的位置
		leftSize := k - il                                   // 左子树节点数
		root.Left = build(il, k-1, pl, pl+leftSize-1)
		root.Right = build(k+1, ir, pl+leftSize, pr-1)
		return root
	}
	return build(0, n-1, 0, n-1)
}
