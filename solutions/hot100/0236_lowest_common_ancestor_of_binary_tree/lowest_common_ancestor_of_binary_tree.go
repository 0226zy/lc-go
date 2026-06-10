package lowestcommonancestorofbinarytree

import (
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// LowestCommonAncestor 二叉树的最近公共祖先（递归 DFS）
// 在左右子树中分别搜索 p 和 q，如果两边都找到则当前节点为 LCA。
// 时间复杂度: O(n)  空间复杂度: O(h)
func LowestCommonAncestor(root, p, q *datastructures.TreeNode) *datastructures.TreeNode {
	if root == nil {
		return nil
	}
	// 如果当前节点是 p 或 q，直接返回
	if root == p || root == q {
		return root
	}
	// 在左右子树中搜索
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	// 如果 p 和 q 分别在左右子树中，当前节点就是 LCA
	if left != nil && right != nil {
		return root
	}
	// 如果只有一边找到，返回那一边的结果
	if left != nil {
		return left
	}
	return right
}

// LowestCommonAncestorIterative 二叉树的最近公共祖先（存储父节点 + 集合）
// 先 DFS 记录每个节点的父节点，再从 p 向上遍历记录路径，最后从 q 向上找第一个交集。
// 时间复杂度: O(n)  空间复杂度: O(n)
func LowestCommonAncestorIterative(root, p, q *datastructures.TreeNode) *datastructures.TreeNode {
	// 存储每个节点的父节点
	parent := make(map[*datastructures.TreeNode]*datastructures.TreeNode)
	stack := []*datastructures.TreeNode{root}
	parent[root] = nil

	// DFS 遍历，记录父节点（使用显式栈）
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			parent[node.Left] = node
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			stack = append(stack, node.Right)
		}
	}

	// 从 p 向上遍历，记录路径
	ancestors := make(map[*datastructures.TreeNode]bool)
	for curr := p; curr != nil; curr = parent[curr] {
		ancestors[curr] = true
	}

	// 从 q 向上遍历，找到第一个在路径中的节点
	for curr := q; curr != nil; curr = parent[curr] {
		if ancestors[curr] {
			return curr
		}
	}
	return nil
}
