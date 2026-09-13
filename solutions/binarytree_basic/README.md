# 二叉树基础题合集

二叉树入门与常考基础题的 Go 题解合集，共 **30 题 / 4 个分类**。

每道题一个独立目录，包含：

- `problem.md` —— 题目描述、核心思路、复杂度分析
- `<name>.go` —— 题解实现（导出函数 + 复杂度注释）
- `<name>_test.go` —— 中文表驱动单元测试 + Benchmark

刷题建议：先看下面的「思路与模板」，再按分类做题；同类题共用同一套骨架，吃透模板比背单题重要。

## 分类总览

| # | 分类 | 题数 | 一句话认出它 |
| --- | --- | --- | --- |
| 1 | 遍历 | 7 | DFS 三序 + BFS 层序 |
| 2 | 属性 / 递归 | 13 | 递归三要素、后序贡献值 |
| 3 | 构造 | 4 | 划分子树区间建树 |
| 4 | 二叉搜索树 | 6 | 中序有序、区间约束 |

## 思路与模板

### 1. 递归三要素

**定义函数语义 → 边界（`nil`）→ 组合左右子树结果。**

```go
func dfs(node *TreeNode) T {
    if node == nil { return 边界值 }
    l, r := dfs(node.Left), dfs(node.Right)
    return 组合(node, l, r)
}
```

- 深度 / 相同 / 对称 / 翻转 / 平衡：直接套用（104、100、101、226、110）。
- 最小深度注意：单侧为空时不能算叶子（111）。
- LCA：后序汇合，左右都找到则当前为祖先（236）。

### 2. 层序 BFS

按层处理：先取出本层全部节点，再扩展下一层。

```go
queue := []*TreeNode{root}
for len(queue) > 0 {
    level := queue; queue = nil
    for _, node := range level {
        // 处理本层；子节点入新 queue
    }
}
```

- 层序 / 锯齿形 / 右视图 / 层平均值：同一模板（102、103、199、637）。
- 右视图 = 每层最后一个；锯齿形 = 奇数层倒序。

### 3. 构造

用中序确定左右子树范围，前序/后序取根；哈希表预存中序下标，递归划分区间。

```go
// 前序+中序：根 = pre[0]，中序找根下标 i，左长 = i
// 后序+中序：根 = post[末]，同理划分
root := &TreeNode{Val: 根值}
root.Left  = build(左区间)
root.Right = build(右区间)
```

- 105 / 106 按上述划区间；654 每次取区间最大值作根；108 有序数组取中点作根（平衡 BST）。

### 4. BST：中序 = 有序

```go
// 验证：携带合法区间 (lo, hi)
func valid(node *TreeNode, lo, hi *int) bool {
    if node == nil { return true }
    if lo != nil && node.Val <= *lo { return false }
    if hi != nil && node.Val >= *hi { return false }
    return valid(node.Left, lo, &node.Val) && valid(node.Right, &node.Val, hi)
}
```

- 搜索 / 插入：一路比较往左或往右（700、701）。
- 删除：叶子直接删；单子树接上来；双子树用后继（右子树最左）替换（450）。
- 第 K 小 / 最小差：中序计数或比相邻差值（230、530）。

### 5. 贡献值（后序）

函数返回「单边向下的最大贡献」，路径答案在递归中全局更新。

```go
var ans int // 或 math.MinInt
func gain(node *TreeNode) int {
    if node == nil { return 0 }
    l := max(0, gain(node.Left))   // 负贡献舍弃
    r := max(0, gain(node.Right))
    ans = max(ans, node.Val+l+r)   // 经过当前节点的路径
    return node.Val + max(l, r)    // 只能选一边向上
}
```

- 直径 = 左右深度之和的最大值（543）；最大路径和同模板（124）。

### 6. 序列化

层序或前序 + 空节点占位（如 `"null"`），反序列化按同样约定还原。

```go
// 层序序列化示意：BFS，空孩子也写入占位符
// 反序列化：按队列顺序消费字符串，挂左右孩子
```

- 297 任选一种约定即可，关键是编解码一致；114 展开为链表可看作「右指针串前序」。

## 题目清单

### 1. 遍历（7）

| 题号 | 题目 |
| --- | --- |
| 94 | [二叉树的中序遍历](0094_binary_tree_inorder_traversal/problem.md) |
| 144 | [二叉树的前序遍历](0144_binary_tree_preorder_traversal/problem.md) |
| 145 | [二叉树的后序遍历](0145_binary_tree_postorder_traversal/problem.md) |
| 102 | [二叉树的层序遍历](0102_binary_tree_level_order_traversal/problem.md) |
| 103 | [二叉树的锯齿形层序遍历](0103_binary_tree_zigzag_level_order_traversal/problem.md) |
| 199 | [二叉树的右视图](0199_binary_tree_right_side_view/problem.md) |
| 637 | [二叉树的层平均值](0637_average_of_levels_in_binary_tree/problem.md) |

### 2. 属性 / 递归（13）

| 题号 | 题目 |
| --- | --- |
| 100 | [相同的树](0100_same_tree/problem.md) |
| 101 | [对称二叉树](0101_symmetric_tree/problem.md) |
| 104 | [二叉树的最大深度](0104_maximum_depth_of_binary_tree/problem.md) |
| 110 | [平衡二叉树](0110_balanced_binary_tree/problem.md) |
| 111 | [二叉树的最小深度](0111_minimum_depth_of_binary_tree/problem.md) |
| 226 | [翻转二叉树](0226_invert_binary_tree/problem.md) |
| 404 | [左叶子之和](0404_sum_of_left_leaves/problem.md) |
| 543 | [二叉树的直径](0543_diameter_of_binary_tree/problem.md) |
| 124 | [二叉树中的最大路径和](0124_binary_tree_maximum_path_sum/problem.md) |
| 222 | [完全二叉树的节点个数](0222_count_complete_tree_nodes/problem.md) |
| 114 | [二叉树展开为链表](0114_flatten_binary_tree_to_linked_list/problem.md) |
| 236 | [二叉树的最近公共祖先](0236_lowest_common_ancestor_of_a_binary_tree/problem.md) |
| 297 | [二叉树的序列化与反序列化](0297_serialize_and_deserialize_binary_tree/problem.md) |

### 3. 构造（4）

| 题号 | 题目 |
| --- | --- |
| 105 | [从前序与中序遍历序列构造二叉树](0105_construct_binary_tree_from_preorder_and_inorder_traversal/problem.md) |
| 106 | [从中序与后序遍历序列构造二叉树](0106_construct_binary_tree_from_inorder_and_postorder_traversal/problem.md) |
| 654 | [最大二叉树](0654_maximum_binary_tree/problem.md) |
| 108 | [将有序数组转换为二叉搜索树](0108_convert_sorted_array_to_binary_search_tree/problem.md) |

### 4. 二叉搜索树（6）

| 题号 | 题目 |
| --- | --- |
| 98 | [验证二叉搜索树](0098_validate_binary_search_tree/problem.md) |
| 700 | [二叉搜索树中的搜索](0700_search_in_a_binary_search_tree/problem.md) |
| 701 | [二叉搜索树中的插入操作](0701_insert_into_a_binary_search_tree/problem.md) |
| 450 | [删除二叉搜索树中的节点](0450_delete_node_in_a_bst/problem.md) |
| 230 | [二叉搜索树中第 K 小的元素](0230_kth_smallest_element_in_a_bst/problem.md) |
| 530 | [二叉搜索树的最小绝对差](0530_minimum_absolute_difference_in_bst/problem.md) |
