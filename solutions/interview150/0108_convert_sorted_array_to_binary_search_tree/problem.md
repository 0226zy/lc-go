# 108. 将有序数组转换为二叉搜索树 (Convert Sorted Array to Binary Search Tree)

## 题目描述

给你一个整数数组 `nums`，其中元素已经按**升序**排列，请你将其转换为一棵**高度平衡**的二叉搜索树。

**高度平衡**二叉树是指：每个节点的左右两个子树的高度差的绝对值不超过 1。

### 示例 1

```
输入: nums = [-10,-3,0,5,9]
输出: [0,-3,9,-10,null,5]
解释: [0,-10,5,null,-3,null,9] 也将被视为正确答案
```

### 示例 2

```
输入: nums = [1,3]
输出: [3,1]
解释: [1,null,3] 和 [3,1] 都是高度平衡的二叉搜索树
```

### 提示

- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` 按**严格递增**排列

## 题目解析

### 核心思路

这道题要同时满足两个性质：**二叉搜索树**（中序遍历升序）和**高度平衡**（左右子树高度差 ≤ 1）。

关键观察：数组本身就是“中序遍历”的结果。二叉搜索树的中序遍历是升序的，所以只要数组里**左边的元素都在左子树、右边的元素都在右子树**，BST 性质天然成立。

剩下的问题只是“怎么分才能让树平衡”。答案很直觉：**每次取当前区间的正中间元素当根节点**。

- 中点左边的元素全部进左子树，右边的全部进右子树 → 左子树节点数最多比右子树多 1；
- 左右子树分别递归用同样方式构造 → 每一层都满足“两边节点数相差不超过 1”；
- 节点数差不超过 1，高度自然差不超过 1 → 整棵树高度平衡。

这类“取中点递归”的做法是**分治**的典型模板，和二分查找的“折半”思想一脉相承。

### 算法步骤

1. 定义递归函数 `build(low, high)`，表示用 `nums[low..high]` 构造平衡 BST：
   - 若 `low > high`，区间为空，返回 `nil`；
   - 取 `mid = low + (high - low + 1) / 2`（上中点），以 `nums[mid]` 为根；
   - 左子树 = `build(low, mid-1)`，右子树 = `build(mid+1, high)`。
2. 从 `build(0, len(nums)-1)` 开始递归即可。

### 复杂度分析

- **时间复杂度**: O(n)，每个元素恰好作为一次根/子节点被访问
- **空间复杂度**: O(log n)，递归栈深度，树是平衡的所以是 log n

## 代码实现

```go
func SortedArrayToBST(nums []int) *datastructures.TreeNode {
    return build(nums, 0, len(nums)-1)
}

func build(nums []int, low, high int) *datastructures.TreeNode {
    if low > high {
        return nil
    }
    mid := low + (high-low)/2
    return &datastructures.TreeNode{
        Val:   nums[mid],
        Left:  build(nums, low, mid-1),   // 左半区间 → 左子树
        Right: build(nums, mid+1, high),  // 右半区间 → 右子树
    }
}
```

**执行过程示例**（`nums = [-10,-3,0,5,9]`）：

```
build(0,4): mid=2，根=0
├─ build(0,1): mid=1，根=-3
│   ├─ build(0,0): 根=-10
│   └─ build(2,1): 空
└─ build(3,4): mid=3，根=5
    ├─ build(3,2): 空
    └─ build(4,4): 根=9

结果树:
        0
      /   \
    -3     5
   /        \
 -10         9
```
