package maximumdepthofbinarytree

import (
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// MaxDepth 二叉树的最大深度（递归 DFS）
// 给定一个二叉树 root，返回其最大深度。
// 时间复杂度: O(n)  空间复杂度: O(h)
func MaxDepth(root *datastructures.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// MaxDepthBFS 二叉树的最大深度（BFS 层序遍历）
// 使用队列逐层遍历，每遍历完一层深度加 1。
// 时间复杂度: O(n)  空间复杂度: O(w)
func MaxDepthBFS(root *datastructures.TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*datastructures.TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}
