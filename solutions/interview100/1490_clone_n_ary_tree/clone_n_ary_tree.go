package clonenarytree

// Node N 叉树节点定义
type Node struct {
	Val      int
	Children []*Node
}

// CloneTree 克隆 N 叉树
// 给定一棵 N 叉树的根节点 root，返回它的深拷贝：
// 新树与原树结构和值完全相同，但不共享任何节点指针。
// 时间复杂度: O(n) n 为节点数  空间复杂度: O(h) h 为树高（递归调用栈）
func CloneTree(root *Node) *Node {
	if root == nil {
		return nil
	}
	clone := &Node{Val: root.Val}
	if len(root.Children) > 0 {
		clone.Children = make([]*Node, len(root.Children))
		for i, child := range root.Children {
			clone.Children[i] = CloneTree(child)
		}
	}
	return clone
}
