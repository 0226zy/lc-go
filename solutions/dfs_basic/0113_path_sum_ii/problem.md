# 113. 路径总和 II (Path Sum II)

## 题目描述

给你二叉树的根节点 `root` 和一个整数目标和 `targetSum`，找出所有 **从根节点到叶子节点** 路径总和等于给定目标和的路径。

**叶子节点** 是指没有子节点的节点。

### 示例 1

```
输入: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出: [[5,4,11,2],[5,8,4,5]]
```

### 示例 2

```
输入: root = [1,2,3], targetSum = 5
输出: []
```

### 示例 3

```
输入: root = [1,2], targetSum = 0
输出: []
```

## 提示

- 树中节点总数在范围 `[0, 5000]` 内
- `-1000 <= Node.val <= 1000`
- `-1000 <= targetSum <= 1000`

## 题目解析

### 核心思路

本题是 112 题「路径总和」的进阶版：不再是判断“是否存在”，而是要 **收集所有满足条件的路径**，因此采用 **前序 DFS + 回溯**。

DFS 模板要点：

- **前序遍历**：从根到叶的路径恰好对应前序遍历的递归栈顺序，进入节点时“做选择”（加入路径），离开时“撤销选择”（弹出路径）。
- **路径拷贝**：到达叶节点且路径和满足条件时，必须对当前路径做一次 **深拷贝** 后再加入结果——否则后续回溯弹出元素会篡改已收集的答案。
- **结束位置必须是叶节点**：路径和达标但节点不是叶节点（如示例 3 的 `[1,2], targetSum=0` 变形）不能收集，继续向下搜索。
- **不要提前剪枝**：节点值和 `targetSum` 都可能为负，路径和没有单调性，不能因为当前和超过 `targetSum` 就停止向下搜索。

回溯三要素：

- **路径**：`path` 当前从根到该节点的节点值序列，以及剩余目标和 `remain`。
- **选择列表**：当前节点的左、右孩子。
- **结束条件**：`node` 为叶节点且 `remain == 0`，收集路径拷贝；`node == nil` 直接返回。

### 算法步骤

1. 若树为空，直接返回空结果。
2. 从根节点开始 DFS，进入节点时：
   - 把 `node.Val` 追加到 `path`，`remain -= node.Val`；
   - 若为叶节点：检查 `remain == 0`，满足则把 `path` 的拷贝加入结果；
   - 否则递归搜索左右子树；
   - 回溯：弹出 `path` 末尾元素，恢复现场。
3. 返回 `result`。

### 复杂度分析

- **时间复杂度**: O(n·h)，n 为节点数，每个节点访问一次；收集答案时需要拷贝长度为 O(h) 的路径，h 为树高。
- **空间复杂度**: O(h)，递归栈深度与当前路径切片长度均为树高，最坏（链状树）为 O(n)；不计输出本身的空间。

## 代码实现

```go
func PathSum(root *datastructures.TreeNode, targetSum int) [][]int {
    var result [][]int
    var path []int

    var dfs func(node *datastructures.TreeNode, remain int)
    dfs = func(node *datastructures.TreeNode, remain int) {
        if node == nil {
            return
        }
        // 做选择：把当前节点加入路径，并更新剩余目标和
        path = append(path, node.Val)
        remain -= node.Val
        // 到达叶节点且剩余和恰好为 0：收集一条合法路径（必须拷贝，否则后续回溯会覆盖）
        if node.Left == nil && node.Right == nil && remain == 0 {
            p := make([]int, len(path))
            copy(p, path)
            result = append(result, p)
        } else {
            dfs(node.Left, remain)
            dfs(node.Right, remain)
        }
        // 撤销选择：弹出当前节点，恢复现场，继续探索其他分支
        path = path[:len(path)-1]
    }
    dfs(root, targetSum)
    return result
}
```

**执行过程示例**（`root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22`）：

```
进入 5: path=[5] remain=17，非叶
  进入 4: path=[5,4] remain=13，非叶
    进入 11: path=[5,4,11] remain=2，非叶
      进入 7: path=[5,4,11,7] remain=-5，叶但和不为 0，不收集，回溯弹出 7
      进入 2: path=[5,4,11,2] remain=0，叶且和为 0，收集 [5,4,11,2]，回溯弹出 2
    回溯弹出 11
  回溯弹出 4
  进入 8: path=[5,8] remain=9，非叶
    进入 13: path=[5,8,13] remain=-4，叶但和不为 0，不收集，回溯弹出 13
    进入 4: path=[5,8,4] remain=5，非叶
      进入 5: path=[5,8,4,5] remain=0，叶且和为 0，收集 [5,8,4,5]，回溯弹出 5
      进入 1: path=[5,8,4,1] remain=4，叶但和不为 0，不收集，回溯弹出 1
    回溯弹出 4
  回溯弹出 8
结果: [[5,4,11,2],[5,8,4,5]]
```
