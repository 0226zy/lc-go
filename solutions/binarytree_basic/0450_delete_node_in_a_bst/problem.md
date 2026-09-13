# 450. 删除二叉搜索树中的节点 (Delete Node in a BST)

## 题目描述

给定一个二叉搜索树的根节点 `root` 和一个值 `key`，删除 BST 中值为 `key` 的节点，并保证删除后仍是合法的 BST。返回删除后树的根节点引用（可能与原根不同）。

一般来说，删除节点可分为两种情况：

- 节点是叶子节点，直接删除；
- 节点不是叶子节点且有一个或两个孩子，则用它的**子节点**或**中序后继 / 前驱**替换后再删除。

### 示例 1

```
输入: root = [5,3,6,2,4,null,7], key = 3
输出: [5,4,6,2,null,null,7]
解释: 值为 3 的节点有两个孩子，用中序后继 4 替换后删除原 4。
```

### 示例 2

```
输入: root = [5,3,6,2,4,null,7], key = 0
输出: [5,3,6,2,4,null,7]
解释: 树中没有值为 0 的节点，树保持不变。
```

### 示例 3

```
输入: root = [], key = 0
输出: []
```

## 提示

- 节点数范围 `[0, 10^4]`
- `-10^5 <= Node.val <= 10^5`
- 每个节点值唯一
- `root` 是合法的二叉搜索树
- `-10^5 <= key <= 10^5`

## 题目解析

### 核心思路

先按 BST 性质定位到目标节点（`key < root.Val` 往左，`key > root.Val` 往右），再按孩子数量分三种情况删除：

1. **无左孩子**（叶节点或只有右孩子）：用右孩子顶替当前节点，返回 `root.Right`。
2. **只有左孩子**：用左孩子顶替，返回 `root.Left`。
3. **左右都有**：取**右子树中的最小节点**（中序后继）的值覆盖当前节点，再在右子树中递归删除该后继（后继最多只有右孩子，会落入情况 1/2）。

把「删除」写成返回新子树根的递归函数，父节点用 `root.Left = DeleteNode(...)` / `root.Right = DeleteNode(...)` 接回，边界统一，代码最干净。

### 算法步骤

1. `root == nil` → 返回 `nil`。
2. `key < root.Val` → `root.Left = DeleteNode(root.Left, key)`。
3. `key > root.Val` → `root.Right = DeleteNode(root.Right, key)`。
4. `key == root.Val`：
   - 左空 → 返回右；右空 → 返回左；
   - 否则找右子树最左节点 `successor`，`root.Val = successor.Val`，再 `root.Right = DeleteNode(root.Right, successor.Val)`。
5. 返回 `root`。

### 复杂度分析

- **时间复杂度**: O(h)，h 为树高；平衡时 O(log n)，退化成链时 O(n)
- **空间复杂度**: O(h)，递归栈深度

## 代码实现

```go
func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = DeleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = DeleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		root.Val = successor.Val
		root.Right = DeleteNode(root.Right, successor.Val)
	}
	return root
}
```
