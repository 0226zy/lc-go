# 112. 路径总和 (Path Sum)

## 题目描述

给你二叉树的根节点 `root` 和一个表示目标和的整数 `targetSum`。判断该树中是否存在 **根节点到叶节点** 的路径，这条路径上所有节点值相加的和等于 `targetSum`。

**叶节点** 是指没有子节点的节点。

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
解释: 树中存在两条根节点到叶节点的路径：
1 -> 2 和为 3，1 -> 3 和为 4，都不等于 5
```

### 示例 3

```
输入: root = [], targetSum = 0
输出: false
解释: 树为空，不存在任何根节点到叶节点的路径
```

## 提示

- 树中节点的数目在范围 `[0, 5000]` 内
- `-1000 <= Node.val <= 1000`
- `-1000 <= targetSum <= 1000`

## 题目解析

### 核心思路

本题是树上 DFS 的「**自上而下携带状态**」模板，模板要点如下：

1. **状态随递归向下传递**：从根出发，每经过一个节点就把目标和减去该节点的值，原问题就缩小为「以当前节点为根的子树中，是否存在到叶节点的路径和等于剩余值」的子问题。状态只需一次减法运算即可更新，无需维护全局路径。
2. **终止条件明确且必须落在叶节点**：
   - 走到空节点：路径不存在，返回 `false`；
   - 走到叶节点：路径结束，比较剩余值是否恰好等于叶节点的值。
   之所以判定必须落在叶节点，是因为题目要求路径**必须到达叶子**，中途经过内部节点时即使剩余值恰好为 0 也不能提前结算（例如 `root = [1,2]`、`targetSum = 1`，根节点值已经为 1，但 1 不是叶节点，答案仍为 `false`）。
3. **无需标记/还原**：本题只向下传递一个数值（剩余目标和），它是按值传递的，天然在每次递归调用中独立，不存在回溯题中「修改全局状态、返回前还原」的问题，因此代码没有撤销操作。

### 算法步骤

1. 若 `root` 为空，返回 `false`（空树没有任何路径）。
2. 若 `root` 是叶节点（左右孩子都为空），返回 `root.Val == targetSum`。
3. 否则计算剩余目标和 `remain = targetSum - root.Val`，递归判断
   `HasPathSum(root.Left, remain) || HasPathSum(root.Right, remain)`，
   利用逻辑或的短路特性：左子树一旦找到，右子树不再访问。

### 复杂度分析

- **时间复杂度**: O(n)，n 为节点数，最坏情况下需要遍历所有节点
- **空间复杂度**: O(h)，h 为树高，即递归栈的深度；平衡树为 O(log n)，最坏（链状树）为 O(n)

## 代码实现

```go
func HasPathSum(root *datastructures.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 到达叶节点：判断剩余目标和是否恰好等于当前节点值
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	// 内部节点：扣除当前节点值后，转化为左右子树的子问题，任一边成立即可
	remain := targetSum - root.Val
	return HasPathSum(root.Left, remain) || HasPathSum(root.Right, remain)
}
```

**执行过程示例**（示例 1，`root = [5,4,8,11,null,13,4,7,2,null,null,null,1]`，`targetSum = 22`）：

```
HasPathSum(5, 22)     → 5 不是叶节点，remain = 17
 ├─ HasPathSum(4, 17)    → remain = 13
 │   └─ HasPathSum(11, 13)   → remain = 2
 │       ├─ HasPathSum(7, 2)     → 叶节点，7 ≠ 2，返回 false
 │       └─ HasPathSum(2, 2)     → 叶节点，2 = 2，返回 true ✓
 │       向上逐层返回 true，左子树已命中
 └─ 右子树（8 一侧）因逻辑或短路，不再访问
最终结果: true（路径 5 -> 4 -> 11 -> 2，和恰为 22）
```
