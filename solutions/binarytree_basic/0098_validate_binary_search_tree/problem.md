# 98. 验证二叉搜索树 (Validate Binary Search Tree)

## 题目描述

给你一个二叉树的根节点 `root`，判断其是否是一个有效的二叉搜索树。

有效二叉搜索树定义如下：

- 节点的左子树只包含 **严格小于** 当前节点的数。
- 节点的右子树只包含 **严格大于** 当前节点的数。
- 所有左子树和右子树自身必须也是二叉搜索树。

### 示例 1

```
输入: root = [2,1,3]
输出: true
```

```
    2
   / \
  1   3
```

### 示例 2

```
输入: root = [5,1,4,null,null,3,6]
输出: false
解释: 根节点的值是 5，但是右子节点的值是 4。
```

```
    5
   / \
  1   4
     / \
    3   6
```

## 提示

- 树中节点数目范围在 `[1, 10^4]` 内
- `-2^31 <= Node.val <= 2^31 - 1`

## 题目解析

### 核心思路

本题的陷阱在于：**BST 约束不是「父子」之间的局部关系，而是「祖先链」上的全局区间约束**。

例如 `[10,5,15,null,null,6,20]` 中，节点 `6` 是 `15` 的左孩子，满足 `6 < 15`，但它在根 `10` 的右子树里，必须大于 `10`——只比较父子关系会误判为合法。

**解法：递归携带合法上下界（开区间）**。定义递归函数 `checkBounds(node, lower, upper)`，含义是：

> `node` 的值必须严格落在 `(lower, upper)` 开区间内，且其左右子树也分别满足各自收紧后的区间。

递归时把当前节点值作为新的边界传递给子树：

- 左孩子继承区间 `(lower, val)`：右边界被当前值收紧；
- 右孩子继承区间 `(val, upper)`：左边界被当前值收紧。

由于题目值域可达 `±2^31`，用 `int64` 类型的 `math.MinInt64 / math.MaxInt64` 作为初始哨兵边界，避免与节点极值冲突。

**等价解法：中序遍历递增性**。BST 的中序遍历结果必然严格递增，因此也可以中序遍历并检查相邻元素大小。本题采用上下界递归，中序思路在下方「算法步骤」后补充说明。

### 算法步骤

1. 调用 `checkBounds(root, MinInt64, MaxInt64)`。
2. 递归过程：
   - 节点为空 → 返回 `true`（空树视为合法 BST）；
   - 若 `val <= lower` 或 `val >= upper` → 返回 `false`；
   - 否则递归检查 `checkBounds(node.Left, lower, val)` 与 `checkBounds(node.Right, val, upper)`，两者都合法才合法。
3. 短路特性：`&&` 一旦遇到左子树非法，右子树不再检查，提前剪枝。

**中序遍历等价思路**：中序遍历（左→根→右）BST 得到严格递增序列。实现时维护一个「前一个节点值」的变量，遍历时若当前值不大于前一个值即判定非法；无需存整个数组，可用递归闭包或迭代栈实现。

### 复杂度分析

- **时间复杂度**: O(n)，n 为节点数，每个节点恰好访问一次；遇到违规时 `&&` 短路可提前返回。
- **空间复杂度**: O(h)，递归栈深度，h 为树高；平衡树 O(log n)，退化为链式树时 O(n)。

## 代码实现

```go
package isvalidbst

import (
	"math"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// IsValidBST 验证二叉搜索树
// 判断给定二叉树是否为合法的二叉搜索树：左子树所有节点值严格小于根，右子树所有节点值严格大于根，且左右子树也必须分别是 BST。
// 采用携带合法上下界的递归：每个节点必须落在 (lower, upper) 开区间内，递归时将当前节点值收紧为子树的新边界。
// 时间复杂度: O(n) 每个节点恰好访问一次  空间复杂度: O(h) 递归栈深度，h 为树高
func IsValidBST(root *datastructures.TreeNode) bool {
	// 使用 int64 边界，避免节点值取 int32 极值时与哨兵冲突
	return checkBounds(root, math.MinInt64, math.MaxInt64)
}

// checkBounds 递归检查 node 是否落在 (lower, upper) 开区间内，
// 并向下递归：左子树继承 (lower, 当前值)，右子树继承 (当前值, upper)
func checkBounds(node *datastructures.TreeNode, lower, upper int64) bool {
	if node == nil {
		return true
	}
	val := int64(node.Val)
	if val <= lower || val >= upper { // 超出祖先限定的区间，直接判定非法
		return false
	}
	return checkBounds(node.Left, lower, val) && checkBounds(node.Right, val, upper)
}
```

**执行过程示例**（`root = [10,5,15,null,null,6,20]`，期望输出 `false`）：

```
树结构:
        10
       /  \
      5    15
          /  \
         6    20

递归推演（区间均为开区间）：
checkBounds(10, (-∞, +∞))        → 10 在区间内，继续
├─ checkBounds(5,  (-∞, 10))     → 5 在 (-∞,10) 内，左右为空，返回 true
└─ checkBounds(15, (10, +∞))     → 15 在 (10,+∞) 内，继续
   ├─ checkBounds(6,  (10, 15))  → 6 <= 下界 10，越界！返回 false
   └─ (&& 短路，右子树 20 不再检查)
结果: false
```

对比只看父子关系的错误判断：`6 < 15` 局部成立，但它继承了祖先根 `10` 的下界约束 `(10, ...)`，上下界递归正是通过传递区间把这种「祖先约束」一路带下来的。
