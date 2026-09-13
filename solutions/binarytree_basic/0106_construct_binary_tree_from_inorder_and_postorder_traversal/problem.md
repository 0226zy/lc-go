# 106. 从中序与后序遍历序列构造二叉树 (Construct Binary Tree from Inorder and Postorder Traversal)

## 题目描述

给定两个整数数组 `inorder` 和 `postorder`，其中 `inorder` 是二叉树的中序遍历，`postorder` 是同一棵树的后序遍历，请你构造并返回这棵 **二叉树**。

### 示例 1

```
输入: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
输出: [3,9,20,null,null,15,7]
```

```
        3
       / \
      9   20
         /  \
        15   7
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
- `inorder` 和 `postorder` 都由 **不同** 的值组成
- `postorder` 中每一个值都在 `inorder` 中
- `inorder` **保证** 为树的中序遍历
- `postorder` **保证** 为树的后序遍历

## 题目解析

### 核心思路

两种遍历序列的分工非常明确：

- **后序遍历**（左 → 右 → 根）：序列的**最后一个元素一定是当前子树的根**。
- **中序遍历**（左 → 根 → 右）：找到根的位置 `k` 后，`k` **左边是左子树、右边是右子树**，同时还能得到左子树的节点个数。

于是还原过程就是一个标准的「分治递归」：每确定一个根，就把问题拆成「还原左子树」和「还原右子树」两个更小的同类问题。

实现上有两个关键点：

1. **哈希表加速定位根**：每次递归都要在中序序列里找根的位置，预处理一个 `值 -> 下标` 的哈希表，把每次 O(n) 的查找降到 O(1)。
2. **从后序末尾倒序取根，必须先建右子树**：后序倒序（根 → 右 → 左）时，紧跟着根出现的是**右子树**的根。用一个全局指针 `idx` 从 `postorder` 末尾往前消耗元素，取到根后**先递归建右子树、再递归建左子树**，指针才能与两种序列的划分严格对齐。

「状态携带方式」：递归函数只需要携带中序区间的左右端点 `[lo, hi]`，后序序列的消耗交给全局指针 `idx`，无需同时维护两套区间。

### 算法步骤

1. 预处理：遍历 `inorder`，建立 `值 -> 中序下标` 的哈希表 `pos`。
2. 令 `idx = len(postorder) - 1`，从后序末尾开始取根。
3. 定义递归函数 `build(lo, hi)`，还原中序区间 `[lo, hi]` 对应的子树：
   - 若 `lo > hi`，区间为空，返回 `nil`；
   - 取 `postorder[idx]` 作为根节点的值，创建根节点，`idx--`；
   - 查哈希表得到根在中序中的位置 `k = pos[root.Val]`；
   - **先**递归 `root.Right = build(k+1, hi)`（右子树的根在后序倒序中先出现）；
   - **再**递归 `root.Left = build(lo, k-1)`；
   - 返回根节点。
4. 返回 `build(0, len(inorder)-1)`。

### 复杂度分析

- **时间复杂度**: O(n)。每个节点恰好创建一次，每次递归中哈希表查找为 O(1)，总共 n 次。
- **空间复杂度**: O(n)。哈希表占用 O(n)；递归栈深度为树高 O(h)，最坏（链式树）为 O(n)。

## 代码实现

```go
package buildtree

import "github.com/0226zy/lc-go/pkg/datastructures"

// BuildTree 从中序与后序遍历序列构造二叉树
// 给定两个整数数组 inorder 和 postorder（值互不相同），其中 inorder 是二叉树的中序遍历，
// postorder 是同一棵树的后序遍历，还原这棵二叉树并返回其根节点。
// 时间复杂度: O(n) 每个节点恰好创建一次  空间复杂度: O(n) 哈希表开销 + O(h) 递归栈
func BuildTree(inorder []int, postorder []int) *datastructures.TreeNode {
	// 哈希表记录每个值在中序遍历中的位置，避免递归时线性查找根
	pos := make(map[int]int, len(inorder))
	for i, v := range inorder {
		pos[v] = i
	}

	idx := len(postorder) - 1 // 后序末元素是根，从末尾往前依次取子树根

	// build 用中序区间 [lo, hi] 还原一棵子树，全局指针 idx 同步消耗后序序列
	var build func(lo, hi int) *datastructures.TreeNode
	build = func(lo, hi int) *datastructures.TreeNode {
		if lo > hi {
			return nil
		}
		root := &datastructures.TreeNode{Val: postorder[idx]}
		idx--
		k := pos[root.Val]
		// 后序倒序时右子树的根先出现，因此必须先建右子树再建左子树
		root.Right = build(k+1, hi)
		root.Left = build(lo, k-1)
		return root
	}
	return build(0, len(inorder)-1)
}
```

**执行过程示例**（`inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]`）：

```
build(0, 4):
  idx=4 取 postorder[4]=3 为根，idx=3；中序 k=1
  先建右子树 build(2, 4):
    idx=3 取 postorder[3]=20 为根，idx=2；中序 k=3
    先建右子树 build(4, 4):
      idx=2 取 postorder[2]=7 为根，idx=1；中序 k=4
      右子树 build(5, 4): 空，返回 nil
      左子树 build(4, 3): 空，返回 nil
      → 返回节点 7
    左子树 build(2, 2):
      idx=1 取 postorder[1]=15 为根，idx=0；中序 k=2
      右子树 build(3, 2): 空，返回 nil
      左子树 build(2, 1): 空，返回 nil
      → 返回节点 15
    → 返回节点 20（左 15、右 7）
  再建左子树 build(0, 0):
    idx=0 取 postorder[0]=9 为根，idx=-1；中序 k=0
    右子树 build(1, 0): 空，返回 nil
    左子树 build(0, -1): 空，返回 nil
    → 返回节点 9
  → 返回根节点 3（左 9、右 20）

最终树:
        3
       / \
      9   20
         /  \
        15   7
```
