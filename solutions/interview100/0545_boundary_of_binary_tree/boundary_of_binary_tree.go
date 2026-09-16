package boundaryofbinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// BoundaryOfBinaryTree 二叉树的边界
// 返回二叉树边界的节点值列表，边界由左边界（不含叶子）、所有叶子节点（从左到右）、
// 右边界（自底向上，不含叶子）三部分逆时针拼接而成。
// 时间复杂度: O(n) 每个节点最多被访问常数次  空间复杂度: O(n) 递归栈与结果存储
func BoundaryOfBinaryTree(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := []int{root.Val}
	// 根节点本身是叶子时，边界只有根
	if root.Left == nil && root.Right == nil {
		return ans
	}

	// 1. 左边界：从 root.Left 出发，优先走左孩子，不含叶子
	for cur := root.Left; cur != nil && (cur.Left != nil || cur.Right != nil); {
		ans = append(ans, cur.Val)
		if cur.Left != nil {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	// 2. 叶子边界：DFS 先左后右，按从左到右顺序收集所有叶子
	var collectLeaves func(node *datastructures.TreeNode)
	collectLeaves = func(node *datastructures.TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			ans = append(ans, node.Val)
			return
		}
		collectLeaves(node.Left)
		collectLeaves(node.Right)
	}
	collectLeaves(root)

	// 3. 右边界：从 root.Right 出发，优先走右孩子，不含叶子；收集后逆序输出
	var right []int
	for cur := root.Right; cur != nil && (cur.Left != nil || cur.Right != nil); {
		right = append(right, cur.Val)
		if cur.Right != nil {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
	}
	for i := len(right) - 1; i >= 0; i-- {
		ans = append(ans, right[i])
	}
	return ans
}
