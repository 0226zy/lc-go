package binarytreerightsideview

import (
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// RightSideView 二叉树的右视图（BFS 层序遍历）
// 给定一个二叉树根节点，返回从右侧所能看到的节点值（每层最右节点）。
// 时间复杂度: O(n)  空间复杂度: O(w)
func RightSideView(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == size-1 {
				result = append(result, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}

// RightSideViewDFS 二叉树的右视图（DFS 递归）
// 按照「根 → 右 → 左」顺序遍历，每层第一个被访问的节点即为右视图节点。
// 时间复杂度: O(n)  空间复杂度: O(h)
func RightSideViewDFS(root *datastructures.TreeNode) []int {
	var result []int
	dfs(root, 0, &result)
	return result
}

func dfs(node *datastructures.TreeNode, depth int, result *[]int) {
	if node == nil {
		return
	}
	if depth == len(*result) {
		*result = append(*result, node.Val)
	}
	dfs(node.Right, depth+1, result)
	dfs(node.Left, depth+1, result)
}
