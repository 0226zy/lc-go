package encodenarytreetobinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// Node N 叉树节点定义
type Node struct {
	Val      int
	Children []*Node
}

// Codec N 叉树与二叉树互转的编解码器
// 编码采用「左孩子右兄弟」表示法：
// N 叉树节点的第一个孩子成为二叉节点的左孩子，其余兄弟沿左孩子的右链依次排列；
// 解码为逆过程，沿左孩子的右链逐一收回所有孩子。
type Codec struct {
}

// Constructor 创建编解码器实例
func Constructor() Codec {
	return Codec{}
}

// encode 将 N 叉树编码为二叉树
// 时间复杂度: O(n) 每个节点转换一次  空间复杂度: O(h) 递归栈深度
func (c *Codec) encode(root *Node) *datastructures.TreeNode {
	if root == nil {
		return nil
	}
	head := &datastructures.TreeNode{Val: root.Val}
	// 用虚拟节点把孩子们串成一条右链，链头挂到 Left 上
	dummy := &datastructures.TreeNode{}
	cur := dummy
	for _, child := range root.Children {
		cur.Right = c.encode(child)
		cur = cur.Right
	}
	head.Left = dummy.Right
	return head
}

// decode 将二叉树解码回 N 叉树
// 时间复杂度: O(n) 每个节点转换一次  空间复杂度: O(h) 递归栈深度
func (c *Codec) decode(root *datastructures.TreeNode) *Node {
	if root == nil {
		return nil
	}
	head := &Node{Val: root.Val}
	// 左孩子的右链即原 N 叉树节点的所有孩子
	for cur := root.Left; cur != nil; cur = cur.Right {
		head.Children = append(head.Children, c.decode(cur))
	}
	return head
}
