# 701. 二叉搜索树中的插入操作 (Insert into a Binary Search Tree)

## 题目描述

给定二叉搜索树（BST）的根节点 `root` 和要插入树中的值 `val`，将值插入二叉搜索树。返回插入后二叉搜索树的根节点。输入数据 **保证**，新值和原始二叉搜索树中的任意节点值都不同。

**注意**，可能存在多种有效的插入方式，只要树在插入后仍保持为二叉搜索树即可。你可以返回 **任意有效的结果**。

### 示例 1

```
输入: root = [4,2,7,1,3], val = 5
输出: [4,2,7,1,3,5]
解释: 另一个满足题目要求、通过验收的树是 [5,2,7,1,3,null,null,4]。
```

### 示例 2

```
输入: root = [40,20,60,10,30,50,70], val = 25
输出: [40,20,60,10,30,50,70,null,null,25]
```

### 示例 3

```
输入: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
输出: [4,2,7,1,3,5]
```

## 提示

- 树中的节点数在 `[0, 10^4]` 范围内
- `-10^8 <= Node.val <= 10^8`
- 所有值 `Node.val` 是 **独一无二** 的
- `-10^8 <= val <= 10^8`
- 保证 `val` 在原始 BST 中不存在

## 题目解析

### 核心思路

本题是二叉搜索树的最基本操作之一：**沿查找路径一路下探，直到落在一个空位上，把新节点挂在那里**。

BST 的性质：左子树所有值 < 根值 < 右子树所有值。因此插入位置是唯一确定的"空叶子挂点"：

- **递归定义**：`InsertIntoBST(root, val)` 表示"在以 root 为根的 BST 中插入 val，返回插入后的子树根"。
- **递归方向**：`val < root.Val` 就递归进左子树，否则递归进右子树（题目保证 val 不会等于任何已有值，不会出现相等分支）。
- **终止条件**：下探到 `nil` 空位时，说明找到了唯一合适的挂点，直接 `new(TreeNode{Val: val})` 返回给上层，由父节点把它接到左/右孩子上。
- **回溯返回**：每一层递归把子树的返回值重新赋给 `root.Left` / `root.Right`，然后原样返回 `root`——这样整棵树的结构逐层拼回去，根节点不变（除非原树为空）。

因为 val 不存在于树中，所以不存在"查找到相等值"的中途返回；也不需要旋转等平衡操作，新节点一定挂在叶子位置。

### 算法步骤

1. 若 `root == nil`，说明找到了插入位置：返回值为 `val` 的新节点（空树时新节点就是整棵树的根）。
2. 若 `val < root.Val`，递归 `InsertIntoBST(root.Left, val)`，把结果赋给 `root.Left`。
3. 否则递归 `InsertIntoBST(root.Right, val)`，把结果赋给 `root.Right`。
4. 返回 `root`。

### 复杂度分析

- **时间复杂度**: O(h)，h 为树高。平衡树时 h = O(log n)；树退化为链表时 h = O(n)。
- **空间复杂度**: O(h)，递归调用栈的深度，最坏（链式树）为 O(n)。若改为迭代写法，额外空间可降为 O(1)。

## 代码实现

```go
package insertintobst

import "github.com/0226zy/lc-go/pkg/datastructures"

// InsertIntoBST 二叉搜索树中的插入操作
// 给定二叉搜索树的根节点 root 和一个值 val，将该值插入到二叉搜索树中，
// 保持二叉搜索树性质不变，返回插入后的根节点。保证新值不与原树中任何值重复。
// 时间复杂度: O(h)  h 为树高，平衡时为 O(log n)，链式退化为 O(n)  空间复杂度: O(h) 递归栈深度
func InsertIntoBST(root *datastructures.TreeNode, val int) *datastructures.TreeNode {
	if root == nil {
		return &datastructures.TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = InsertIntoBST(root.Left, val)
	} else {
		root.Right = InsertIntoBST(root.Right, val)
	}
	return root
}
```

**执行过程示例**（`root = [4,2,7,1,3], val = 5`）：

```
InsertIntoBST(4, 5): 5 > 4，递归进右子树
  InsertIntoBST(7, 5): 5 < 7，递归进左子树
    InsertIntoBST(nil, 5): 空位！创建节点 5 并返回
  节点 7 的 Left 接上新节点 5，返回 7
节点 4 的 Right 接回 7（指针不变），返回 4

结果树: [4,2,7,1,3,5]
      4
     / \
    2   7
   / \ /
  1  3 5
```
