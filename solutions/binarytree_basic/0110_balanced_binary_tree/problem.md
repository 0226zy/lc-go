# 110. 平衡二叉树 (Balanced Binary Tree)

## 题目描述

给定一个二叉树，判断它是否是 **高度平衡** 的二叉树。

本题中，一棵高度平衡二叉树定义为：一个二叉树的 **每个节点** 的左右两个子树的高度差的绝对值不超过 1。

### 示例 1

```
输入: root = [3,9,20,null,null,15,7]
输出: true
```

```
      3
     / \
    9  20
       / \
      15  7
```

### 示例 2

```
输入: root = [1,2,2,3,3,null,null,4,4]
输出: false
```

```
          1
         / \
        2   2
       / \
      3   3
     / \
    4   4
```

### 示例 3

```
输入: root = []
输出: true
```

## 提示

- 树中的节点数在范围 `[0, 5000]` 内
- `-10^4 <= Node.val <= 10^4`

## 题目解析

### 核心思路

判断一棵树是否平衡，需要知道每个节点 **左右子树的高度**。关键在于选择遍历顺序：

- **自顶向下**（先算当前节点高度差，再递归子树）：每个节点的高度会被父节点重复计算，时间复杂度退化为 O(n²)（链式树时）。
- **自底向上**（后序遍历：左 → 右 → 根）：递归返回子树高度的同时判断平衡性，每个节点只访问一次，时间 O(n)。

本题采用 **自底向上携带高度** 的写法，并用哨兵值 `-1` 表示"该子树已不平衡"，实现 **提前剪枝**：

- 递归函数 `height(node)` 返回以 `node` 为根的子树高度；若该子树不平衡，返回 `-1`。
- 一旦左子树返回 `-1`，右子树就不用再算了，直接把 `-1` 逐层向上传播到根。
- 空树高度定义为 `0`，空树是平衡树。

### 算法步骤

1. 定义递归函数 `height(node) int`：
   - 若 `node == nil`，返回 `0`；
   - 递归求左子树高度 `left`，若 `left == -1` 直接返回 `-1`（剪枝）；
   - 递归求右子树高度 `right`，若 `right == -1` 直接返回 `-1`（剪枝）；
   - 若 `|left - right| > 1`，返回 `-1`；
   - 否则返回 `max(left, right) + 1`。
2. 主函数 `IsBalanced` 返回 `height(root) != -1`。

### 复杂度分析

- **时间复杂度**: O(n)，n 为节点数，每个节点只被访问一次；一旦发现不平衡立即提前返回。
- **空间复杂度**: O(h)，递归调用栈的深度，h 为树高；最坏情况（链式树）为 O(n)，平衡树为 O(log n)。

## 代码实现

```go
package isbalanced

import "github.com/0226zy/lc-go/pkg/datastructures"

// IsBalanced 平衡二叉树
// 给定一棵二叉树，判断它是否是高度平衡的二叉树：
// 即每个节点的左右两个子树的高度差的绝对值不超过 1。
// 时间复杂度: O(n) 每个节点只访问一次  空间复杂度: O(h) 递归栈深度，h 为树高
func IsBalanced(root *datastructures.TreeNode) bool {
	return height(root) != -1
}

// height 自底向上计算以 node 为根的子树高度；
// 一旦发现某个子树不平衡，直接返回 -1 并逐层向上传播，实现提前剪枝。
func height(node *datastructures.TreeNode) int {
	if node == nil {
		return 0
	}
	left := height(node.Left)
	if left == -1 {
		return -1 // 左子树不平衡，整棵树不必再算
	}
	right := height(node.Right)
	if right == -1 {
		return -1 // 右子树不平衡，整棵树不必再算
	}
	if abs(left-right) > 1 {
		return -1 // 当前节点左右子树高度差超过 1
	}
	return max(left, right) + 1
}
```

**执行过程示例**（`root = [1,2,2,3,3,null,null,4,4]`，即示例 2）：

```
height(1):
  height(2左):
    height(3左):
      height(4左): 左右均为空 → 返回 1
      height(4右): 左右均为空 → 返回 1
      |1-1| ≤ 1 → 3左 高度 = max(1,1)+1 = 2
    height(3右): 左右均为空 → 返回 1
    |2-1| = 1 ≤ 1 → 2左 高度 = max(2,1)+1 = 3
  height(2右): 左右均为空 → 返回 1
  |3-1| = 2 > 1 → 根节点不平衡，返回 -1
IsBalanced = (-1 != -1) = false
```

再看示例 1 `[3,9,20,null,null,15,7]`：

```
height(9)  = 1（叶子）
height(15) = 1，height(7) = 1 → height(20) = max(1,1)+1 = 2
height(3)  = max(1,2)+1 = 3，各节点高度差均 ≤ 1
IsBalanced = (3 != -1) = true
```
