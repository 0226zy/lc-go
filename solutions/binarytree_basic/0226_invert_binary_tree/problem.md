# 226. 翻转二叉树 (Invert Binary Tree)

## 题目描述

给你一棵二叉树的根节点 `root`，翻转这棵二叉树，并返回其根节点。

所谓翻转，就是把树中**每个节点的左右孩子指针交换**，整棵树做镜像变换。

### 示例 1

```
输入: root = [4,2,7,1,3,6,9]
输出: [4,7,2,9,6,3,1]
```

### 示例 2

```
输入: root = [2,1,3]
输出: [2,3,1]
```

### 示例 3

```
输入: root = []
输出: []
```

## 提示

- 树中节点数目在范围 `[0, 100]` 内
- `-100 <= Node.val <= 100`

## 题目解析

### 核心思路

二叉树递归的经典入门题。把"翻转整棵树"拆解为递归定义：

- **递归边界**：空树不需要翻转，直接返回 `nil`。
- **递归关系**：先把左子树翻转好、再把右子树翻转好，最后交换当前节点的左右孩子指针。

关键在于**指针交换必须发生在两次递归调用之后**（即后序位置），否则先交换再递归时 `root.Left` / `root.Right` 已经互换，虽然结果等价，但语义上容易混乱。本实现统一用"先递归、后交换"的写法：

```text
InvertTree(node) =
    若 node 为空 → 返回 nil
    node.Left  ← InvertTree(node.Left)
    node.Right ← InvertTree(node.Right)
    交换 node.Left 与 node.Right
    返回 node
```

也可以写成更常见的"先交换、再递归"（前序位置），两种写法都是 O(n)，只是操作时机不同。本目录采用后序写法，方便体会"自底向上构建答案"的递归思想。

### 算法步骤

1. 若 `root == nil`，返回 `nil`。
2. 递归翻转左子树：`root.Left = InvertTree(root.Left)`。
3. 递归翻转右子树：`root.Right = InvertTree(root.Right)`。
4. 交换 `root.Left` 与 `root.Right`，返回 `root`。

### 复杂度分析

- **时间复杂度**: O(n)，每个节点恰好访问一次。
- **空间复杂度**: O(h)，递归调用栈的深度等于树高 h；最坏情况（链式树）为 O(n)，平衡树为 O(log n)。

## 代码实现

```go
package invertbinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// InvertTree 翻转二叉树
// 给定二叉树根节点 root，交换每个节点的左右子树（镜像翻转），返回翻转后的根节点。
// 时间复杂度: O(n) 每个节点访问一次  空间复杂度: O(h) 递归栈深度，h 为树高
func InvertTree(root *datastructures.TreeNode) *datastructures.TreeNode {
	if root == nil {
		return nil
	}
	root.Left = InvertTree(root.Left)
	root.Right = InvertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
```

**执行过程示例**（`root = [4,2,7,1,3,6,9]`）：

```
原树:                  翻转后:
      4                      4
     / \                    / \
    2   7                  7   2
   / \ / \                / \ / \
  1  3 6  9              9  6 3  1

递归推演（自底向上）：
  1. 叶子 1、3、6、9 均无子节点，各自翻转后不变，逐层返回。
  2. 回到节点 2：左子已翻转为 1，右子已翻转为 3，交换 → 左 3、右 1。
  3. 回到节点 7：左子已翻转为 6，右子已翻转为 9，交换 → 左 9、右 6。
  4. 回到根节点 4：左子树根为 2、右子树根为 7，交换 → 左 7、右 2。
最终层序输出: [4,7,2,9,6,3,1]
```
