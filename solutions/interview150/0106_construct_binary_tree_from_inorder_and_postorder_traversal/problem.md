# 106. 从中序与后序遍历序列构造二叉树 (Construct Binary Tree from Inorder and Postorder Traversal)

## 题目描述

给定两个整数数组 `inorder` 和 `postorder`，其中 `inorder` 是二叉树的**中序遍历**，`postorder` 是同一棵树的**后序遍历**，请构造二叉树并返回其根节点。

### 示例 1

```
输入: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出: [3,9,20,null,null,15,7]
```

### 示例 2

```
输入: inorder = [-1], postorder = [-1]
输出: [-1]
```

## 提示

- `1 <= inorder.length <= 3000`
- `postorder.length == inorder.length`
- `-3000 <= inorder[i], postorder[i] <= 3000`
- `inorder` 和 `postorder` 的值均**互不相同**
- 保证 `inorder` 和 `postorder` 是同一棵二叉树的中序遍历和后序遍历

## 题目解析

### 核心思路

本题与 [105. 从前序与中序遍历序列构造二叉树](./../0105_construct_binary_tree_from_preorder_and_inorder_traversal/problem.md) 是同一类模板，区别只在**根的位置**：

- **后序遍历**：左子树 → 右子树 → 根，所以 `postorder[最后一个元素]` 是（当前这棵树的）根节点
- **中序遍历**：左子树 → 根 → 右子树，根的左边是左子树、右边是右子树

递归还原的思路完全一致：每拿到一段后序序列，最后一个元素就是根；拿根的值去中序序列里定位，把中序切成两半；左半段的长度决定左子树大小，然后对两段子序列递归。

**区间如何划分**（设中序区间为 `[il, ir]`，后序区间为 `[pl, pr]`）：

1. 根节点值 = `postorder[pr]`，在中序中的位置为 `k`，左子树大小 `leftSize = k - il`
2. 左子树：中序 `[il, k-1]`，后序 `[pl, pl+leftSize-1]`
3. 右子树：中序 `[k+1, ir]`，后序 `[pl+leftSize, pr-1]`

**注意**：必须先递归构造右子树还是左子树？都行——因为这里用下标区间定位，不依赖"先序指针向后推进"的顺序，左右子树的区间都是从长度算出来的，互不干扰。这点和 105 题的写法略有不同，请对比体会。

同样用**哈希表**预处理"值 → 中序下标"，把每次找根从 O(n) 降到 O(1)。

### 算法步骤

1. 特判：序列为空返回 `nil`
2. 预处理：建立哈希表 `indexMap`，记录每个值在中序序列中的下标
3. 递归函数 `build(il, ir, pl, pr)`：
   - 若 `il > ir`，返回 `nil`（空区间）
   - 根值 = `postorder[pr]`，创建根节点
   - 用哈希表找到根在中序中的位置 `k`，`leftSize = k - il`
   - 递归构造左子树：`build(il, k-1, pl, pl+leftSize-1)`
   - 递归构造右子树：`build(k+1, ir, pl+leftSize, pr-1)`
   - 返回根节点
4. 返回 `build(0, n-1, 0, n-1)`

### 复杂度分析

- **时间复杂度**: O(n)，每个节点创建一次；哈希表定位 O(1)
- **空间复杂度**: O(n)，哈希表 O(n) + 递归栈 O(h)

## 代码实现

```go
// BuildTree 递归 + 哈希表实现
func BuildTree(inorder, postorder []int) *datastructures.TreeNode {
    n := len(inorder)
    indexMap := make(map[int]int, n)
    for i, v := range inorder {
        indexMap[v] = i
    }
    var build func(il, ir, pl, pr int) *datastructures.TreeNode
    build = func(il, ir, pl, pr int) *datastructures.TreeNode {
        if il > ir {
            return nil
        }
        root := &datastructures.TreeNode{Val: postorder[pr]}
        k := indexMap[postorder[pr]]
        leftSize := k - il
        root.Left = build(il, k-1, pl, pl+leftSize-1)
        root.Right = build(k+1, ir, pl+leftSize, pr-1)
        return root
    }
    return build(0, n-1, 0, n-1)
}
```

**执行过程示例**（`inorder = [9,3,15,20,7]`, `postorder = [9,15,7,20,3]`）：

```
中序:   [9, 3, 15, 20, 7]
后序:   [9, 15, 7, 20, 3]   根 = 3（后序末元素）
中序中 3 左边 [9] 是左子树（leftSize=1），右边 [15,20,7] 是右子树

          3
   后序[9]/   \后序[15,7,20]
中序[9]       中序[15,20,7]

左子树递归：后序末元素 = 9，中序中 9 左右皆空 → 9 是叶子
右子树递归：后序末元素 = 20，中序中 20 左边 [15] 右边 [7] → 15、7 是叶子

最终树: [3,9,20,null,null,15,7]
```
