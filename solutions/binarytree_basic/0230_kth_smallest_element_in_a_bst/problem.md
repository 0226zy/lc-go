# 230. 二叉搜索树中第 K 小的元素 (Kth Smallest Element in a BST)

## 题目描述

给定一个二叉搜索树的根节点 `root`，和一个整数 `k`，请你设计一个算法查找其中第 `k` 小的元素（从 1 开始计数）。

### 示例 1

```
输入: root = [3,1,4,null,2], k = 1
输出: 1
```

对应的二叉搜索树：

```
    3
   / \
  1   4
   \
    2
```

### 示例 2

```
输入: root = [5,3,6,2,4,null,null,1], k = 3
输出: 3
```

对应的二叉搜索树：

```
      5
     / \
    3   6
   / \
  2   4
 /
1
```

## 提示

- 树中的节点数为 `n`
- `1 <= k <= n <= 10^4`
- `0 <= Node.val <= 10^4`

**进阶**：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？

## 题目解析

### 核心思路

二叉搜索树（BST）最重要的性质是：**中序遍历（左 → 根 → 右）得到的序列恰好是递增有序的**。因此「第 k 小的元素」等价于「中序遍历过程中访问到的第 k 个节点」。

基于这一点，我们不需要先把整棵树遍历完存成数组再取下标，而是在中序遍历的同时维护一个计数器：

- **遍历顺序**：严格按照 左子树 → 当前节点 → 右子树 的顺序递归。
- **状态携带方式**：用闭包外的变量 `count` 记录已访问的节点个数，`ans` 记录答案。
- **剪枝要点**：一旦 `count == k`，说明当前节点就是第 k 小，记录答案后立即沿递归链一路返回（递归函数返回 `true` 表示「已找到」），后续节点全部不再访问，做到**提前终止**。

这样平均情况下只访问 `h + k` 个节点（h 为树高），而不是全部 n 个节点。

### 算法步骤

1. 初始化计数器 `count = 0` 和答案变量 `ans`。
2. 定义递归中序遍历函数 `inorder`，返回 `bool` 表示是否已经找到答案：
   - 节点为空，返回 `false`；
   - 先递归左子树，若左子树已找到答案，直接返回 `true`（剪枝）；
   - 访问当前节点：`count++`；若 `count == k`，令 `ans = node.Val`，返回 `true`；
   - 否则递归右子树，返回其结果。
3. 从根节点开始遍历，结束后返回 `ans`。

### 复杂度分析

- **时间复杂度**: O(h + k)。从根走到中序起点需要 O(h)，之后最多再访问 k 个节点；最坏情况（k = n 且树退化为链）为 O(n)。
- **空间复杂度**: O(h)。递归调用栈的深度等于树高 h；平衡树为 O(log n)，退化为链时最坏 O(n)。

## 代码实现

```go
package kthsmallest

import "github.com/0226zy/lc-go/pkg/datastructures"

// KthSmallest 二叉搜索树中第 K 小的元素
// 给定二叉搜索树的根节点 root 和整数 k，返回其中第 k 小的元素（从 1 开始计数）。
// 利用 BST 中序遍历结果递增的性质，遍历时计数，数到第 k 个节点立即提前终止。
// 时间复杂度: O(h + k)  空间复杂度: O(h)
func KthSmallest(root *datastructures.TreeNode, k int) int {
	count := 0 // 已访问的节点个数
	ans := 0   // 第 k 小的节点值

	// inorder 中序遍历，返回值表示是否已经找到第 k 小（找到则向上层层返回，实现剪枝）
	var inorder func(node *datastructures.TreeNode) bool
	inorder = func(node *datastructures.TreeNode) bool {
		if node == nil {
			return false
		}
		// 左子树中已找到则直接返回，不再访问当前节点和右子树
		if inorder(node.Left) {
			return true
		}
		count++
		if count == k { // 中序第 k 个被访问的节点就是第 k 小
			ans = node.Val
			return true
		}
		return inorder(node.Right)
	}

	inorder(root)
	return ans
}
```

**执行过程示例**（`root = [5,3,6,2,4,null,null,1], k = 3`）：

```
中序遍历从根 5 开始，一路向左下钻：
inorder(5) -> inorder(3) -> inorder(2) -> inorder(1) -> inorder(nil) 返回 false
访问 1: count=1，不等于 k=3 -> inorder(nil)（1 的右子树）返回 false
回到 2: 访问 2: count=2，不等于 k -> inorder(nil)（2 的右子树）返回 false
回到 3: 访问 3: count=3 == k=3 -> ans=3，返回 true
inorder(5) 收到左子树返回 true，直接返回 true
（节点 4、5、6 及其子树全部不再访问，提前剪枝）
结果: 3
```
