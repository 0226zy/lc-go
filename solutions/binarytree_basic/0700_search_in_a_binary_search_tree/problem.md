# 700. 二叉搜索树中的搜索 (Search in a Binary Search Tree)

## 题目描述

给定二叉搜索树（BST）的根节点 `root` 和一个整数值 `val`。

你需要在 BST 中找到节点值等于 `val` 的节点。返回以该节点为根的子树。如果节点不存在，则返回 `nil`。

### 示例 1

```
输入: root = [4,2,7,1,3], val = 2
输出: [2,1,3]
```

### 示例 2

```
输入: root = [4,2,7,1,3], val = 5
输出: []
```

## 提示

- 树中节点数在 `[1, 5000]` 范围内
- `1 <= Node.val <= 10^7`
- `root` 是二叉搜索树
- `1 <= val <= 10^7`

## 题目解析

### 核心思路

本题利用 **二叉搜索树的性质进行定向查找**，是 BST 最基础的应用：

> 对任意节点 `node`，其左子树中所有节点的值都 **小于** `node.val`，右子树中所有节点的值都 **大于** `node.val`。

因此每一步都能做出确定的方向决策，而不是像普通二叉树那样两边都要搜：

- 若 `val == node.val`，当前节点就是目标，直接返回；
- 若 `val < node.val`，目标只可能在 **左子树** 中，丢弃右子树；
- 若 `val > node.val`，目标只可能在 **右子树** 中，丢弃左子树。

递归定义：`searchBST(node)` 表示「在以 `node` 为根的 BST 中查找值为 `val` 的节点」，返回找到的节点或 `nil`。每递归一层就排除掉一半（理想平衡时）的节点，所以也叫「二叉搜索树上的二分查找」。

由于每一步只有一条路径继续向下，完全不需要回溯；用迭代写也一样简洁，递归版更清晰、贴合树形问题的表达方式。

### 算法步骤

1. 若 `root == nil` 或 `root.Val == val`，返回 `root`（两种情况合一：`nil` 表示未找到，命中则返回该节点）。
2. 若 `val < root.Val`，递归进入左子树：`return SearchBST(root.Left, val)`。
3. 否则递归进入右子树：`return SearchBST(root.Right, val)`。

### 复杂度分析

- **时间复杂度**: O(h)，h 为树高。平衡 BST 时为 O(log n)，退化为链式树时最坏 O(n)。
- **空间复杂度**: O(h)，递归调用栈深度，最坏 O(n)。

## 代码实现

```go
package searchbst

import (
	"github.com/0226zy/lc-go/pkg/datastructures"
)

// SearchBST 二叉搜索树中的搜索
// 给定二叉搜索树的根节点 root 和一个值 val，
// 返回树中值等于 val 的节点（以其为根的子树），不存在则返回 nil。
// 时间复杂度: O(h) h 为树高，平衡时 O(log n)  空间复杂度: O(h) 递归栈深度
func SearchBST(root *datastructures.TreeNode, val int) *datastructures.TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return SearchBST(root.Left, val)
	}
	return SearchBST(root.Right, val)
}
```

**执行过程示例**（`root = [4,2,7,1,3], val = 2`）：

```
树结构:
      4
     / \
    2   7
   / \
  1   3

SearchBST(4, 2): 2 < 4 → 进入左子树
  SearchBST(2, 2): 2 == 2 → 命中，返回节点 2
结果: 以节点 2 为根的子树 [2,1,3]
```

**执行过程示例**（`root = [4,2,7,1,3], val = 5`）：

```
SearchBST(4, 5): 5 > 4 → 进入右子树
  SearchBST(7, 5): 5 < 7 → 进入左子树
    SearchBST(nil, 5): root == nil → 返回 nil
结果: nil（未找到）
```
