# 105. 从前序与中序遍历序列构造二叉树 (Construct Binary Tree from Preorder and Inorder Traversal)

## 题目描述

给定两个整数数组 `preorder` 和 `inorder`，其中 `preorder` 是二叉树的**先序遍历**，`inorder` 是同一棵树的**中序遍历**，请构造二叉树并返回其根节点。

### 示例 1

```
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]
```

### 示例 2

```
输入: preorder = [-1], inorder = [-1]
输出: [-1]
```

## 提示

- `1 <= preorder.length <= 3000`
- `inorder.length == preorder.length`
- `-3000 <= preorder[i], inorder[i] <= 3000`
- `preorder` 和 `inorder` 的值均**互不相同**
- 保证 `preorder` 和 `inorder` 是同一棵二叉树的前序遍历和中序遍历

## 题目解析

### 核心思路

这是"遍历序列还原二叉树"的经典模板题，关键在于吃透两种遍历的性质：

- **先序遍历**：根 → 左子树 → 右子树，所以 `preorder[0]` 永远是（当前这棵树的）根节点
- **中序遍历**：左子树 → 根 → 右子树，所以**根的左边是左子树的所有节点，根的右边是右子树的所有节点**

于是可以递归地还原：每拿到一段先序序列，第一个元素就是根；拿这个根的值去中序序列里找位置，这个位置把中序序列切成两半，左半段的长度 `leftSize` 决定了左子树有多少个节点；相应地，先序序列去掉根之后，前 `leftSize` 个元素是左子树、剩下的是右子树。对两段子序列递归同样的操作即可。

**区间如何划分**（设先序区间为 `[pl, pr]`，中序区间为 `[il, ir]`）：

1. 根节点值 = `preorder[pl]`，在中序中的位置为 `k`，左子树大小 `leftSize = k - il`
2. 左子树：先序 `[pl+1, pl+leftSize]`，中序 `[il, k-1]`
3. 右子树：先序 `[pl+leftSize+1, pr]`，中序 `[k+1, ir]`

**优化的关键点**：每次在中序序列里线性找根是 O(n)，总复杂度退化为 O(n²)。因为题目保证值互不相同，可以预先用一个**哈希表**记录"值 → 中序下标"，把找根变成 O(1)，整体降到 O(n)。

### 算法步骤

1. 特判：序列为空返回 `nil`
2. 预处理：建立哈希表 `indexMap`，记录每个值在中序序列中的下标
3. 递归函数 `build(pl, pr, il, ir)`：
   - 若 `pl > pr`，返回 `nil`（空区间）
   - 根值 = `preorder[pl]`，创建根节点
   - 用哈希表找到根在中序中的位置 `k`，`leftSize = k - il`
   - 递归构造左子树：`build(pl+1, pl+leftSize, il, k-1)`
   - 递归构造右子树：`build(pl+leftSize+1, pr, k+1, ir)`
   - 返回根节点
4. 返回 `build(0, n-1, 0, n-1)`

### 复杂度分析

- **时间复杂度**: O(n)，每个节点创建一次；哈希表定位 O(1)
- **空间复杂度**: O(n)，哈希表 O(n) + 递归栈 O(h)

## 代码实现

```go
// BuildTree 递归 + 哈希表实现
func BuildTree(preorder, inorder []int) *datastructures.TreeNode {
    n := len(preorder)
    indexMap := make(map[int]int, n)
    for i, v := range inorder {
        indexMap[v] = i
    }
    var build func(pl, pr, il, ir int) *datastructures.TreeNode
    build = func(pl, pr, il, ir int) *datastructures.TreeNode {
        if pl > pr {
            return nil
        }
        root := &datastructures.TreeNode{Val: preorder[pl]}
        k := indexMap[preorder[pl]]
        leftSize := k - il
        root.Left = build(pl+1, pl+leftSize, il, k-1)
        root.Right = build(pl+leftSize+1, pr, k+1, ir)
        return root
    }
    return build(0, n-1, 0, n-1)
}
```

**执行过程示例**（`preorder = [3,9,20,15,7]`, `inorder = [9,3,15,20,7]`）：

```
先序: [3, 9, 20, 15, 7]   根 = 3（先序首元素）
中序: [9, 3, 15, 20, 7]   3 左边 [9] 是左子树，右边 [15,20,7] 是右子树
                            leftSize = 1

        3
       / \
先序[9]/   \先序[20,15,7]
中序[9]     中序[15,20,7]

左子树递归：根=9，中序中 9 左右皆空 → 9 是叶子
右子树递归：根=20（先序首），中序中 20 左边 [15] 右边 [7] → 15、7 是叶子

最终树: [3,9,20,null,null,15,7]
```
