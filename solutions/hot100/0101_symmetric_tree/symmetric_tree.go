package symmetrictree

import (
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// IsSymmetric 对称二叉树（递归 DFS）
// 给你一个二叉树的根节点 root，检查它是否轴对称。
// 时间复杂度: O(n)  空间复杂度: O(h)
func IsSymmetric(root *datastructures.TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// isMirror 判断两棵树是否互为镜像
func isMirror(p, q *datastructures.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}

// IsSymmetricIterative 对称二叉树（迭代，队列）
// 使用队列每次出队两个节点进行比较。
// 时间复杂度: O(n)  空间复杂度: O(n)
func IsSymmetricIterative(root *datastructures.TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*datastructures.TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		p := queue[0]
		q := queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Left, q.Right, p.Right, q.Left)
	}
	return true
}
