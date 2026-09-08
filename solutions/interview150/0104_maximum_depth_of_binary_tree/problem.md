# 104. 二叉树的最大深度 (Maximum Depth of Binary Tree)

## 题目描述

给定一个二叉树 `root`，返回其最大深度。

二叉树的**最大深度**是指从根节点到最远叶子节点的最长路径上的节点总数。

### 示例 1

```
输入: root = [3,9,20,null,null,15,7]
输出: 3
```

### 示例 2

```
输入: root = [1,null,2]
输出: 2
```

## 提示

- 树中节点的数量在 `[0, 10^4]` 范围内
- `-100 <= Node.val <= 100`

## 题目解析

### 核心思路

这是一道经典的**二叉树递归遍历**入门题。

思考方式很自然：整棵树的最大深度是多少？先看根节点自己算一层，再看左右两棵子树——左边最深的深度和右边最深的深度，取较大者，再加 1（根这一层），就是整棵树的深度。

这个思路本身就是递归的定义：

```
maxDepth(树) = max(maxDepth(左子树), maxDepth(右子树)) + 1
```

递归的边界：空树的深度为 0。这是递归能停下来的关键。

### 算法步骤

1. 如果 `root` 为 `nil`，返回 0（空树深度为 0）
2. 递归计算左子树深度 `leftDepth`
3. 递归计算右子树深度 `rightDepth`
4. 返回 `max(leftDepth, rightDepth) + 1`

### 复杂度分析

- **时间复杂度**: O(n)，每个节点恰好被访问一次（n 为节点数）
- **空间复杂度**: O(h)，递归栈深度为树的高度 h；最坏情况（退化成链）为 O(n)，平衡树为 O(log n)

## 代码实现

```go
// MaxDepth 递归实现
func MaxDepth(root *datastructures.TreeNode) int {
    if root == nil {
        return 0
    }
    leftDepth := MaxDepth(root.Left)
    rightDepth := MaxDepth(root.Right)
    return max(leftDepth, rightDepth) + 1
}
```

**执行过程示例**（`root = [3,9,20,null,null,15,7]`）：

```
          3
         / \
        9   20
           /  \
          15   7

MaxDepth(3) = max(MaxDepth(9), MaxDepth(20)) + 1
MaxDepth(9)  = max(MaxDepth(nil), MaxDepth(nil)) + 1 = 1
MaxDepth(20) = max(MaxDepth(15), MaxDepth(7)) + 1
MaxDepth(15) = 1, MaxDepth(7) = 1  => MaxDepth(20) = 2
MaxDepth(3)  = max(1, 2) + 1 = 3
```
