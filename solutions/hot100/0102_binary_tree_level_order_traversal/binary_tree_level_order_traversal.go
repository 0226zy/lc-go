package binarytreelevelordertraversal

import "github.com/0226zy/lc-go/pkg/datastructures"

type TreeNode = datastructures.TreeNode

// LevelOrder 二叉树的层序遍历
// 给你一个二叉树，请你返回其按层序遍历得到的节点值。（即逐层地，从左到右访问所有节点）。
// 时间复杂度: O(n)  空间复杂度: O(n)
func LevelOrder(root *TreeNode) [][]int {
	if root==nil{
		return nil
	}
	queue := []*TreeNode{root}
	ret := [][]int{}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}
		for i := 0; i < levelSize; i++ {
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
		ret = append(ret, level)
	}
	return ret
}
