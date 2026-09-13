# 257. 二叉树的所有路径 (Binary Tree Paths)

## 题目描述

给你一个二叉树的根节点 `root`，按 **任意顺序**，返回所有从根节点到叶子节点的路径。

**叶子节点** 是指没有子节点的节点。路径用 `"->"` 连接节点值，例如 `"1->2->5"`。

### 示例 1

```
输入: root = [1,2,3,null,5]
输出: ["1->2->5","1->3"]

      1
    /   \
   2     3
    \
     5
```

### 示例 2

```
输入: root = [1]
输出: ["1"]
```

## 提示

- 树中节点的数目在范围 `[1, 100]` 内
- `-100 <= Node.val <= 100`

## 题目解析

### 核心思路

典型的 **二叉树先序 DFS + 路径字符串拼接**。沿根向下走时维护一个字符串缓冲区，到达叶节点时把当前缓冲区内容收集进答案。

DFS 模板要点：

- **递归边界**：`node == nil` 直接返回，不产生任何路径。
- **收集条件**：`node.Left == nil && node.Right == nil` 时当前节点是叶子，把缓冲区里的完整路径拷贝进结果。
- **标记/还原策略**：本题不修改树本身，"标记"的是路径缓冲区。进入节点前用 `mark := path.Len()` 记下缓冲区长度，写入 `"->"`（非根节点）和节点值；递归返回后用 `path.Truncate(mark)` 截断回进入时的长度，等价于回溯中的「撤销选择」。这样左右子树共用同一个缓冲区，避免每层都新建字符串。
- 注意必须先判断叶节点再递归左右孩子，否则叶节点会先递归两个 nil 孩子再返回，虽然结果相同，但显式收集更清晰。

### 算法步骤

1. 初始化空结果集 `result` 和字符串缓冲区 `path`。
2. 从根节点开始 DFS：
   - 节点为空，直接返回；
   - 记录 `mark = path.Len()`；若 `mark > 0`（非根）先写入 `"->"`，再写入节点值；
   - 若是叶节点，把 `path.String()` 加入结果；否则递归左右子树；
   - `path.Truncate(mark)` 撤销本次写入。
3. 返回 `result`。

### 复杂度分析

- **时间复杂度**: O(n·h)，n 为节点数，h 为树高。每个节点访问一次，叶节点处拼接路径开销 O(h)；平衡树时 h = O(log n)，退化为链时 h = O(n)，最坏 O(n²)。
- **空间复杂度**: O(h)，递归栈深度与路径缓冲区长度（不计输出）；最坏 O(n)。

## 代码实现

```go
func BinaryTreePaths(root *datastructures.TreeNode) []string {
    var result []string
    var path bytes.Buffer

    var dfs func(node *datastructures.TreeNode)
    dfs = func(node *datastructures.TreeNode) {
        if node == nil {
            return
        }
        // 记录进入前的缓冲区长度，退出时还原，实现路径的撤销
        mark := path.Len()
        if mark > 0 {
            path.WriteString("->")
        }
        path.WriteString(strconv.Itoa(node.Val))

        if node.Left == nil && node.Right == nil {
            // 到达叶节点：收集当前完整路径
            result = append(result, path.String())
        } else {
            dfs(node.Left)
            dfs(node.Right)
        }
        // 回溯：把缓冲区截断回进入时的长度，撤销本节点的写入
        path.Truncate(mark)
    }
    dfs(root)
    return result
}
```

**执行过程示例**（`root = [1,2,3,null,5]`）：

```
dfs(1):  path="1"，非叶
  dfs(2): path="1->2"，非叶
    dfs(nil): 返回
    dfs(5): path="1->2->5"，是叶 → 收集 "1->2->5"
          还原 path="1->2"
    还原 path="1->2" 后回到 dfs(2) 收尾，还原 path="1"
  dfs(3): path="1->3"，是叶 → 收集 "1->3"
        还原 path="1"
结果: ["1->2->5", "1->3"]
```
