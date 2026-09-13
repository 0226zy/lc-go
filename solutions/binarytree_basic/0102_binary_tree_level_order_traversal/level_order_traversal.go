package levelorder

import "github.com/0226zy/lc-go/pkg/datastructures"

// LevelOrder 二叉树的层序遍历
// 给定二叉树的根节点 root，从上到下、从左到右逐层返回每层节点的值。
// 时间复杂度: O(n) 每个节点恰好入队、出队各一次  空间复杂度: O(n) 队列最多同时存放一层的节点
func LevelOrder(root *datastructures.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue) // 当前层节点个数：层与层边界的唯一依据
		level := make([]int, 0, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}
