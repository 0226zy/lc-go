# 100. 相同的树 (Same Tree)

## 题目描述

给你两棵二叉树的根节点 `p` 和 `q`，编写一个函数来检验这两棵树是否相同。

如果两棵树在**结构上相同**，并且**节点上的值也相同**，则认为这两棵树相同。

### 示例 1

```
输入: p = [1,2,3], q = [1,2,3]
输出: true
```

### 示例 2

```
输入: p = [1,2], q = [1,null,2]
输出: false
```

### 示例 3

```
输入: p = [1,2,1], q = [1,1,2]
输出: false
```

## 提示

- 两棵树上的节点数目都在范围 `[0, 100]` 内
- `-10^4 <= Node.val <= 10^4`

## 题目解析

### 核心思路

这又是一道**二叉树递归遍历**模板题，只是递归函数这次要同时接收两棵树。

判断两棵树是否相同，可以拆成三个条件：

1. 两棵树的根节点要么都是空，要么都不是空（一空一非空 → 不同）
2. 两个根节点的值相等
3. 左子树相同 **且** 右子树相同（递归处理子问题）

注意条件 1 还暗藏一个细节：如果两个根节点**都是空**，返回 `true`（两棵空树当然相同）；一空一非空返回 `false`。然后比较根值，再递归比较左右子树。只要把这三个条件按顺序写清楚，代码就是递归定义本身。

### 算法步骤

1. 如果 `p` 和 `q` 都为 `nil`，返回 `true`（都是空树，相同）
2. 如果恰好一个为 `nil`，返回 `false`（结构不同）
3. 如果 `p.Val != q.Val`，返回 `false`（值不同）
4. 递归判断 `IsSameTree(p.Left, q.Left)` 和 `IsSameTree(p.Right, q.Right)`，两者都为 `true` 才返回 `true`

### 复杂度分析

- **时间复杂度**: O(min(n, m))，最坏情况下需要比较所有节点（n、m 为两棵树的节点数，某处不同时会提前返回）
- **空间复杂度**: O(min(h1, h2))，递归栈深度为较矮那棵树的高度

## 代码实现

```go
// IsSameTree 递归实现
func IsSameTree(p, q *datastructures.TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    if p.Val != q.Val {
        return false
    }
    return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}
```

**执行过程示例**（`p = [1,2]`, `q = [1,null,2]`）：

```
p:   1      q:   1
    /           \
   2             2

IsSameTree(1, 1): 值相等
  ├─ IsSameTree(2, nil): p 非空 q 空 → false
  └─ 短路返回 false（右子树无需再比较）
结果: false
```
