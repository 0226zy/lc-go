package binarytreeverticalordertraversal

import "github.com/0226zy/lc-go/pkg/datastructures"

// VerticalOrder 二叉树的垂直遍历
// 给定二叉树根节点，按列（从左到右）分组返回节点值；同一列内按从上往下的顺序排列。
// 根节点列号为 0，左孩子列号 -1，右孩子列号 +1。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(n)
func VerticalOrder(root *datastructures.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	// cols 记录每一列的节点值（按 BFS 访问顺序，即从上往下）
	cols := make(map[int][]int)
	minCol, maxCol := 0, 0

	// 队列元素同时携带节点与其列号
	type item struct {
		node *datastructures.TreeNode
		col  int
	}
	queue := []item{{node: root, col: 0}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		cols[cur.col] = append(cols[cur.col], cur.node.Val)
		if cur.col < minCol {
			minCol = cur.col
		}
		if cur.col > maxCol {
			maxCol = cur.col
		}

		if cur.node.Left != nil {
			queue = append(queue, item{node: cur.node.Left, col: cur.col - 1})
		}
		if cur.node.Right != nil {
			queue = append(queue, item{node: cur.node.Right, col: cur.col + 1})
		}
	}

	// 按列号从小到大拼接结果
	result := make([][]int, 0, maxCol-minCol+1)
	for c := minCol; c <= maxCol; c++ {
		result = append(result, cols[c])
	}
	return result
}
