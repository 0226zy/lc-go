package binarytreerightsideview

import "github.com/0226zy/lc-go/pkg/datastructures"

// RightSideView 二叉树的右视图
// 返回站在二叉树右侧，从上到下能看到的节点值。
// 时间复杂度: O(n) 每个节点恰好入队出队一次  空间复杂度: O(n) 队列最多同时存放一层的节点
func RightSideView(root *datastructures.TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	queue := []*datastructures.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue) // 当前层的节点个数
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == size-1 { // 本层最后一个出队的节点，即最右侧节点
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
