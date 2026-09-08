# 112. 路径总和 (Path Sum)

## 题目描述

给你二叉树的根节点 `root` 和一个表示目标和的整数 `targetSum`。判断该树中是否存在**根节点到叶节点**的路径，这条路径上所有节点值相加的和等于 `targetSum`。

**叶节点**是指没有子节点的节点。

### 示例 1

```
输入: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出: true
解释: 等于目标和的根节点到叶节点路径为 5 -> 4 -> 11 -> 2
```

### 示例 2

```
输入: root = [1,2,3], targetSum = 5
输出: false
解释: 树中存在两条根节点到叶节点的路径 1 -> 2 和 1 -> 3，和分别为 3、4
```

### 示例 3

```
输入: root = [], targetSum = 0
输出: false
解释: 树为空，不存在根节点到叶节点的路径
```

### 提示

- 树中节点数的范围是 `[0, 5000]`
- `-1000 <= Node.val <= 1000`
- `-1000 <= targetSum <= 1000`

## 题目解析

### 核心思路

这是**「自上而下传递状态」的 DFS 模板**：从根出发递归向下走，每到一个节点就把目标和减去当前节点值，问题就变成「以该节点为根的子树中，是否存在到叶节点的路径和等于剩余值」的子问题。

子问题的终止条件非常清晰：

- 走到空节点：路径不存在，返回 `false`；
- 走到**叶节点**：路径到此为止，直接比较剩余值是否等于叶节点的值。

之所以必须在叶节点处判定，是因为题目要求路径**必须到达叶节点**——不能到某个内部节点就提前结算。

### 算法步骤

1. 若 `root` 为空，返回 `false`（空树没有任何路径）。
2. 若 `root` 是叶节点（左右孩子都为空），返回 `root.Val == targetSum`。
3. 否则令 `remain = targetSum - root.Val`，递归判断：
   `HasPathSum(root.Left, remain) || HasPathSum(root.Right, remain)`，任一边成立即可。

### 复杂度分析

- **时间复杂度**: O(n)，最坏情况下需要遍历所有节点
- **空间复杂度**: O(h)，h 为树高，即递归栈深度；最坏（链状树）为 O(n)

## 代码实现

```go
func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 到达叶节点：检查剩余目标和是否恰好等于叶节点的值
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	// 内部节点：转化为左右子树的子问题
	remain := targetSum - root.Val
	return HasPathSum(root.Left, remain) || HasPathSum(root.Right, remain)
}
```

**执行过程示例**（示例 1，`targetSum = 22`）：

```
HasPathSum(5, 22)   → 5 不是叶节点，remain = 17
 ├─ HasPathSum(4, 17)  → remain = 13
 │   └─ HasPathSum(11, 13) → remain = 2
 │       ├─ HasPathSum(7, 2)  → 叶节点，7 ≠ 2，false
 │       └─ HasPathSum(2, 2)  → 叶节点，2 = 2，true  ✓
 └─ （右侧子树不再访问，短路返回）
```
