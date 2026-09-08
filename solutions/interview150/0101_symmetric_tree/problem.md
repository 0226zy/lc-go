# 101. 对称二叉树 (Symmetric Tree)

## 题目描述

给你一个二叉树的根节点 `root`，检查它是否轴对称。

即：以根节点的左右孩子为根的两棵子树，是否互为镜像（一棵的左子树对应另一棵的右子树，且值相等）。

### 示例 1

```
输入: root = [1,2,2,3,4,4,3]
输出: true
```

### 示例 2

```
输入: root = [1,2,2,null,3,null,3]
输出: false
```

## 提示

- 树中节点数目在 `[1, 1000]` 范围内
- `-100 <= Node.val <= 100`

## 题目解析

### 核心思路

本题不是比较"两棵相同的树"，而是比较"两棵互为镜像的树"。这是 [100. 相同的树](./../0100_same_tree/problem.md) 的变体。

判断两棵子树 `left` 和 `right` 互为镜像，要满足：

1. 两个根要么都为空（都是空树，镜像成立），要么都非空（一空一非空 → 不是镜像）
2. 两个根的值相等
3. **`left` 的左子树与 `right` 的右子树互为镜像**（注意交叉对应！）
4. **`left` 的右子树与 `right` 的左子树互为镜像**

对比第 100 题，唯一的区别就是第 3、4 步把"左对左、右对右"换成了"左对右、右对左"。所以镜像判断也是一个递归：

```
isMirror(a, b) = (a、b 同为空) 或 (a、b 都非空 且 a.Val == b.Val
                 且 isMirror(a.Left, b.Right) 且 isMirror(a.Right, b.Left))
```

入口：空树算对称；否则判断 `isMirror(root.Left, root.Right)`。

### 算法步骤

1. 空树直接返回 `true`
2. 递归函数 `check(a, b)`：
   - `a`、`b` 都为空 → `true`
   - 恰好一个为空 → `false`
   - 值不相等 → `false`
   - 返回 `check(a.Left, b.Right) && check(a.Right, b.Left)`
3. 返回 `check(root.Left, root.Right)`

### 复杂度分析

- **时间复杂度**: O(n)，每个节点最多访问一次
- **空间复杂度**: O(h)，递归栈深度为树高 h

## 代码实现

```go
// IsSymmetric 递归实现
func IsSymmetric(root *datastructures.TreeNode) bool {
    if root == nil {
        return true
    }
    return check(root.Left, root.Right)
}

func check(a, b *datastructures.TreeNode) bool {
    if a == nil && b == nil {
        return true
    }
    if a == nil || b == nil {
        return false
    }
    if a.Val != b.Val {
        return false
    }
    return check(a.Left, b.Right) && check(a.Right, b.Left)
}
```

**执行过程示例**（`root = [1,2,2,3,4,4,3]`）：

```
        1
       / \
      2   2
     / \ / \
    3  4 4  3

check(2, 2): 值相等
  ├─ check(2的左=3, 2的右=3): 值相等
  │    ├─ check(nil, nil) = true
  │    └─ check(nil, nil) = true   → true
  └─ check(2的右=4, 2的左=4): 同理 → true
结果: true
```
