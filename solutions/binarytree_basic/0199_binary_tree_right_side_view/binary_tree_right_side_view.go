package rightsiderview

import "github.com/0226zy/lc-go/pkg/datastructures"

// RightSideView 二叉树的右视图（BFS 层序遍历）
// 站在二叉树右侧，从上到下能看到的节点值，即每一层最靠右的节点。
// 时间复杂度: O(n) 每个节点恰好入队、出队各一次  空间复杂度: O(n) 队列最多同时存放一整层的节点
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

// RightSideViewDFS 二叉树的右视图（DFS 先右后左）
// 深度优先遍历时优先递归右子树，每一层首次被访问到的节点就是该层最右侧节点。
// 时间复杂度: O(n) 每个节点恰好访问一次  空间复杂度: O(h) 递归栈深度，h 为树高，最坏 O(n)
func RightSideViewDFS(root *datastructures.TreeNode) []int {
	var result []int
	var dfs func(node *datastructures.TreeNode, depth int)
	dfs = func(node *datastructures.TreeNode, depth int) {
		if node == nil {
			return
		}
		// depth == len(result) 说明本层还没有记录过节点，
		// 而当前节点是本层最先被访问到的（先右后左），即最右侧节点
		if depth == len(result) {
			result = append(result, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return result
}
