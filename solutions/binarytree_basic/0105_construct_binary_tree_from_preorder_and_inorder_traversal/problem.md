# 105. 从前序与中序遍历序列构造二叉树 (Construct Binary Tree from Preorder and Inorder Traversal)

## 题目描述

给定两个整数数组 `preorder` 和 `inorder`，其中 `preorder` 是二叉树的**前序遍历**结果，`inorder` 是同一棵树的**中序遍历**结果，请根据这两个遍历序列还原这棵二叉树，并返回其根节点。

题目保证两个序列中的节点值**互不相同**，且一定对应同一棵合法的二叉树。

### 示例 1

```
输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]
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
输入: preorder = [-1], inorder = [-1]
输出: [-1]
```

## 提示

- `1 <= preorder.length <= 3000`
- `inorder.length == preorder.length`
- `-3000 <= preorder[i], inorder[i] <= 3000`
- `preorder` 和 `inorder` 均由**互不相同**的值组成
- 保证 `preorder` 和 `inorder` 是同一棵二叉树的前序遍历和中序遍历

## 题目解析

### 核心思路

要还原二叉树，先想清楚两种遍历各自"泄露"了什么信息：

- **前序遍历**的顺序是「根 → 左子树 → 右子树」，因此当前（子）树的根永远是前序序列中**最先出现的那个元素**。
- **中序遍历**的顺序是「左子树 → 根 → 右子树」，因此一旦知道根在中序序列中的位置 `k`，`k` 左边整段就是左子树的中序，`k` 右边整段就是右子树的中序。

把两条信息拼起来就得到递归定义：

1. 用前序序列的当前元素确定根节点；
2. 在中序序列里定位根，把"左子树包含哪些节点、右子树包含哪些节点"划分开；
3. 左子树的节点数 `leftSize` 也决定了前序序列中接下来的 `leftSize` 个元素属于左子树，其余的属于右子树；
4. 对左、右子树分别递归执行同样的还原过程。

**状态携带方式**：递归只携带中序区间的左右边界 `[il, ir]`，再用一个**全局前序下标指针 `preIdx`** 顺序消费前序序列。因为前序是"先根再左后右"，所以每递归一层，`preorder[preIdx]` 恰好就是当前子树的根——不需要再传前序区间边界。

**性能要点（为什么用哈希表）**：如果每次都在中序区间里线性扫描找根，单次查找是 O(n)，整棵树退化为 O(n²)。题目保证节点值互不相同，因此可以预处理一张「值 → 中序下标」的哈希表，把定位根的代价降到 O(1)，整体复杂度降为 O(n)。

**终止与剪枝**：中序区间为空（`il > ir`）时直接返回 `nil`，这就是唯一的终止条件；空区间对应空子树，不需要额外剪枝。

### 算法步骤

1. 特判：`preorder` 为空时返回 `nil`。
2. 预处理：遍历 `inorder`，建立哈希表 `pos`，`pos[v]` 表示值 `v` 在中序序列中的下标。
3. 定义递归函数 `build(il, ir int) *TreeNode`，表示还原中序区间为 `[il, ir]` 的子树：
   - 若 `il > ir`，返回 `nil`（空子树）；
   - 取 `rootVal = preorder[preIdx]`，令 `preIdx` 前进一格（消费掉当前根）；
   - 创建根节点，查表得到根在中序中的位置 `k = pos[rootVal]`；
   - **先**递归左子树 `build(il, k-1)`，**再**递归右子树 `build(k+1, ir)`——顺序不能颠倒，因为前序序列中左子树的节点排在右子树前面；
   - 返回根节点。
4. 返回 `build(0, n-1)`。

### 复杂度分析

- **时间复杂度**: O(n)。每个节点恰好被创建一次，定位根节点通过哈希表 O(1) 完成。
- **空间复杂度**: O(n)。哈希表占用 O(n)；递归调用栈深度等于树高 h，最坏（链式树）为 O(n)，最好（平衡树）为 O(log n)。

## 代码实现

```go
package buildtree

import "github.com/0226zy/lc-go/pkg/datastructures"

// BuildTree 从前序与中序遍历序列构造二叉树
// 给定二叉树的前序遍历 preorder 和中序遍历 inorder（节点值互不相同），还原二叉树并返回根节点。
// 时间复杂度: O(n)  空间复杂度: O(n)
func BuildTree(preorder []int, inorder []int) *datastructures.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// pos[v] 记录值 v 在中序遍历中的下标，O(1) 定位根
	pos := make(map[int]int, len(inorder))
	for i, v := range inorder {
		pos[v] = i
	}

	preIdx := 0 // 前序序列的全局消费指针，始终指向当前子树的根
	var build func(il, ir int) *datastructures.TreeNode
	build = func(il, ir int) *datastructures.TreeNode {
		if il > ir { // 中序区间为空，说明是空子树
			return nil
		}
		rootVal := preorder[preIdx]
		preIdx++
		root := &datastructures.TreeNode{Val: rootVal}
		k := pos[rootVal] // 根把中序区间切成左右两半
		// 必须先左后右：前序序列中左子树的节点排在右子树之前
		root.Left = build(il, k-1)
		root.Right = build(k+1, ir)
		return root
	}
	return build(0, len(inorder)-1)
}
```

**执行过程示例**（`preorder = [3,9,20,15,7]`, `inorder = [9,3,15,20,7]`）：

```
预处理哈希表 pos: {9:0, 3:1, 15:2, 20:3, 7:4}

build(0, 4):  取 preorder[0]=3，preIdx→1，k=pos[3]=1
              左区间 [0,0]，右区间 [2,4]
  ├─ build(0, 0): 取 preorder[1]=9，preIdx→2，k=0
  │    ├─ build(0, -1): il>ir，返回 nil   （9 的左孩子为空）
  │    └─ build(1, 0):  il>ir，返回 nil   （9 的右孩子为空）
  │    返回节点 9（叶子）
  └─ build(2, 4): 取 preorder[2]=20，preIdx→3，k=pos[20]=3
       ├─ build(2, 2): 取 preorder[3]=15，preIdx→4，k=2
       │    左右区间均为空 → 返回节点 15（叶子）
       └─ build(4, 4): 取 preorder[4]=7，preIdx→5，k=4
            左右区间均为空 → 返回节点 7（叶子）
       返回节点 20（左 15，右 7）

返回根节点 3，最终树为:

    3
   / \
  9  20
    /  \
   15   7

即层序表示 [3,9,20,null,null,15,7]，与示例 1 一致。
```
