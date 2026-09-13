# 101. 对称二叉树 (Symmetric Tree)

## 题目描述

给你一个二叉树的根节点 `root`，检查它是否轴对称。

换句话说：把根节点的左子树和右子树拿出来，这两棵子树必须**互为镜像**——不仅对应位置的节点值要相等，而且结构上要左右翻转对应（一棵的左子树对应另一棵的右子树）。

### 示例 1

```
输入: root = [1,2,2,3,4,4,3]
输出: true
```

### 示例 2

```
输入: root = [1,2,2,null,3,null,3]
输出: false
解释: 左右子树的值分布虽然相同，但 3 都出现在内侧（都靠近根），结构上不对称。
```

## 提示

- 树中节点数目在范围 `[1, 1000]` 内
- `-100 <= Node.val <= 100`

## 题目解析

### 核心思路

"对称"可以拆成"左右两棵子树互为镜像"。判断镜像本身又是一个递归问题——这与判断"两棵树是否相同"几乎同构，唯一区别在于子树的**对应关系是交叉的**：

- 判断"两树相同"：`左对左、右对右` 递归；
- 判断"两树镜像"：**`左对右、右对左`** 递归。

于是定义递归函数 `isMirror(a, b)`，判断以 `a`、`b` 为根的两棵子树是否互为镜像，终止与递归条件为：

1. `a`、`b** 同时为空 → 两棵空树，镜像成立，返回 `true`；
2. 恰好一个为空 → 结构不同，返回 `false`（剪枝：不必再往下比较）；
3. `a.Val != b.Val` → 值不同，返回 `false`（剪枝）；
4. 否则继续递归比较 `isMirror(a.Left, b.Right)` 与 `isMirror(a.Right, b.Left)`，两者都成立才算镜像。

递归定义汇总为：

```
isMirror(a, b) = (a、b 同为 nil)
              或 (a、b 均非 nil 且 a.Val == b.Val
                  且 isMirror(a.Left, b.Right)
                  且 isMirror(a.Right, b.Left))
```

入口处理：空树（`root == nil`）视为对称，直接返回 `true`；否则把根节点的左右孩子作为一对镜像比较的起始点，调用 `isMirror(root.Left, root.Right)`。

### 算法步骤

1. 若 `root == nil`，返回 `true`。
2. 调用递归函数 `isMirror(a, b)`（初始传入 `root.Left`、`root.Right`）：
   - `a`、`b` 都为空 → 返回 `true`；
   - 恰有一个为空 → 返回 `false`；
   - `a.Val != b.Val` → 返回 `false`；
   - 返回 `isMirror(a.Left, b.Right) && isMirror(a.Right, b.Left)`。
3. 将递归结果作为整棵树是否对称的答案。

### 复杂度分析

- **时间复杂度**: O(n)。每个节点最多参与一次值比较；一旦发现不匹配就提前返回，实际比较次数可能更少。
- **空间复杂度**: O(h)。开销来自递归调用栈，h 为树高；最坏情况（退化为链式树）h = n，即 O(n)。

## 代码实现

```go
package symmetrictree

import "github.com/0226zy/lc-go/pkg/datastructures"

// IsSymmetric 对称二叉树
// 给定一棵二叉树的根节点 root，判断它是否关于根节点左右对称（左子树与右子树互为镜像）。
// 时间复杂度: O(n) 每个节点最多比较一次  空间复杂度: O(h) 递归栈深度为树高 h
func IsSymmetric(root *datastructures.TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// isMirror 递归判断以 a、b 为根的两棵子树是否互为镜像：
// 要求 a、b 同时为空或同时非空且值相等，并且 a 的左子树对应 b 的右子树、a 的右子树对应 b 的左子树。
func isMirror(a, b *datastructures.TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return isMirror(a.Left, b.Right) && isMirror(a.Right, b.Left)
}
```

**执行过程示例**（`root = [1,2,2,3,4,4,3]`）：

```
        1
       / \
      2   2
     / \ / \
    3  4 4  3

IsSymmetric(root): root 非空，调用 isMirror(左2, 右2)
isMirror(2, 2): 值相等，继续比较
  ├─ isMirror(2的左=3, 2'的右=3): 值相等
  │    ├─ isMirror(nil, nil) = true        （外侧：3的左 对 3的右）
  │    └─ isMirror(nil, nil) = true        （内侧：3的右 对 3的左）
  │    → true
  └─ isMirror(2的右=4, 2'的左=4): 值相等
       ├─ isMirror(nil, nil) = true
       └─ isMirror(nil, nil) = true
       → true
整体结果: true
```

再看一个反例（`root = [1,2,2,null,3,null,3]`）：

```
isMirror(2, 2): 值相等
  ├─ isMirror(nil, 3): 一空一非空 → false   （左2 没有左孩子，右2 却有右孩子）
  └─ 短路，不再比较
整体结果: false
```
