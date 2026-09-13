# 104. 二叉树的最大深度 (Maximum Depth of Binary Tree)

## 题目描述

给定一个二叉树 `root`，返回其最大深度。

二叉树的 **最大深度** 是指从根节点到最远叶子节点的最长路径上的节点数。

**说明：** 叶子节点是指没有子节点的节点。

### 示例 1

```
输入: root = [3,9,20,null,null,15,7]
输出: 3
```

对应的二叉树：

```
    3
   / \
  9  20
    /  \
   15   7
```

### 示例 2

```
输入: root = [1,null,2]
输出: 2
```

## 提示

- 树中节点的数量在 `[0, 10^4]` 的范围内。
- `-100 <= Node.val <= 100`

## 题目解析

### 核心思路

这是二叉树 **后序遍历（DFS）** 的经典入门题，采用「分治 + 递归」的自底向上思路：

- **递归定义**：一棵树的最大深度 = `max(左子树最大深度, 右子树最大深度) + 1`。其中 `+1` 表示根节点自身占用一层。
- **基准情形**：空树（`root == nil`）的深度定义为 `0`，递归到此终止。
- **遍历顺序**：先分别递归求出左右子树的深度，再在父节点处合并结果，正是“左 → 右 → 根”的后序访问次序；信息由子节点向根节点逐层回传。
- 无需剪枝，每个节点恰好访问一次。等价地也可以用 BFS 层序遍历统计层数，但递归写法更简洁。

### 算法步骤

1. 若 `root == nil`，返回 `0`（空树深度为 0）。
2. 递归计算 `left = MaxDepth(root.Left)`。
3. 递归计算 `right = MaxDepth(root.Right)`。
4. 返回 `max(left, right) + 1`。

### 复杂度分析

- **时间复杂度**: O(n)，n 为节点总数，每个节点恰好访问一次。
- **空间复杂度**: O(h)，h 为树高，即递归调用栈的最大深度；平衡树为 O(log n)，最坏情况（链式树）退化为 O(n)。

## 代码实现

```go
package maximumdepthofbinarytree

import "github.com/0226zy/lc-go/pkg/datastructures"

// MaxDepth 二叉树的最大深度
// 给定一个二叉树的根节点 root，返回它的最大深度：即从根节点到最远叶子节点的最长路径上的节点个数。
// 时间复杂度: O(n)  每个节点恰好访问一次
// 空间复杂度: O(h)  递归调用栈的深度，h 为树高；最坏情况（链式树）退化为 O(n)
func MaxDepth(root *datastructures.TreeNode) int {
	if root == nil {
		return 0 // 空子树深度为 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	return max(left, right) + 1 // 当前节点深度 = 左右子树较大深度 + 1
}
```

**执行过程示例**（`root = [3,9,20,null,null,15,7]`）：

```
MaxDepth(3)
├─ MaxDepth(9)
│  ├─ MaxDepth(nil) = 0
│  └─ MaxDepth(nil) = 0
│  返回 max(0, 0) + 1 = 1
├─ MaxDepth(20)
│  ├─ MaxDepth(15)
│  │  ├─ MaxDepth(nil) = 0
│  │  └─ MaxDepth(nil) = 0
│  │  返回 max(0, 0) + 1 = 1
│  └─ MaxDepth(7)
│     ├─ MaxDepth(nil) = 0
│     └─ MaxDepth(nil) = 0
│     返回 max(0, 0) + 1 = 1
│  返回 max(1, 1) + 1 = 2
返回 max(1, 2) + 1 = 3
```

递归先一路下探到空节点返回 0，再自底向上逐层回传：叶子节点 9、15、7 的深度都是 1，节点 20 取左右较大值再加 1 得 2，最终根节点 3 得到 3。
