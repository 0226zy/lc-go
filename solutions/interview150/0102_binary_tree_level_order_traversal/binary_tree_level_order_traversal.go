package binarytreelevelordertraversal

import "github.com/0226zy/lc-go/pkg/datastructures"

// LevelOrder 二叉树的层序遍历
// 返回二叉树按层序遍历的结果，每一层的节点值放在一个切片中。
// 时间复杂度: O(n) 每个节点恰好入队出队一次  空间复杂度: O(n) 队列最多同时存放一层的节点
func LevelOrder(root *datastructures.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue) // 当前层的节点个数
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
