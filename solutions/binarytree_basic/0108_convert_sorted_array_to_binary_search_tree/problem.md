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

## 提示

- `1 <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`
- `nums` 按**严格递增**排列

## 题目解析

### 核心思路

题目要求同时满足两个性质：**二叉搜索树**（中序遍历结果为升序）和**高度平衡**（任意节点左右子树高度差 ≤ 1）。

关键观察：升序数组本身就是这棵 BST 的**中序遍历**结果。因此只要保证「数组中点左边的元素全部落在左子树、右边的元素全部落在右子树」，BST 性质就天然成立，无需任何额外判断。

于是问题只剩“怎么分才能让树平衡”，答案很直接：**每次取当前区间的中间元素作为根节点**。

- **递归定义**：`buildBST(low, high)` 表示用 `nums[low..high]` 构造一棵平衡 BST 并返回根节点；
- **状态携带**：只需携带当前区间的左右端点下标，不需要额外路径状态；
- **平衡性保证**：中点左右两侧元素个数最多相差 1，递归到每一层都成立，因此每个节点的左右子树节点数差 ≤ 1，高度差自然 ≤ 1；
- 这是典型的**分治**模板，与二分查找的“折半”思想一脉相承——二分查找是在中点处“选一边丢弃”，本题是在中点处“两边都递归”。

### 算法步骤

1. 定义递归函数 `buildBST(nums, low, high)`：
   - 若 `low > high`，区间为空，返回 `nil`（递归出口）；
   - 取下中点 `mid = low + (high - low) / 2`，以 `nums[mid]` 创建根节点；
   - 左子树 = `buildBST(nums, low, mid-1)`，右子树 = `buildBST(nums, mid+1, high)`；
   - 返回根节点。
2. 从 `buildBST(nums, 0, len(nums)-1)` 开始递归，返回结果即为答案。

> 取上中点 `low + (high - low + 1) / 2` 同样正确，只是偶数长度区间根的位置不同，两种写法都是 LeetCode 认可的合法答案。本文取下中点。

### 复杂度分析

- **时间复杂度**: O(n)，数组中每个元素恰好作为一个节点被创建一次。
- **空间复杂度**: O(log n)，递归调用栈的深度等于树高，平衡树高度为 log n（不计返回值占用的空间）。

## 代码实现

```go
package sortedarraytobst

import "github.com/0226zy/lc-go/pkg/datastructures"

// SortedArrayToBST 将有序数组转换为二叉搜索树
// 给定一个按升序排列的整数数组 nums，将其转换为一棵高度平衡的二叉搜索树。
// 时间复杂度: O(n) 每个元素恰好访问一次  空间复杂度: O(log n) 递归栈深度（平衡树）
func SortedArrayToBST(nums []int) *datastructures.TreeNode {
	return buildBST(nums, 0, len(nums)-1)
}

// buildBST 递归地用 nums[low..high] 构造高度平衡二叉搜索树
func buildBST(nums []int, low, high int) *datastructures.TreeNode {
	if low > high {
		return nil
	}
	mid := low + (high-low)/2
	return &datastructures.TreeNode{
		Val:   nums[mid],
		Left:  buildBST(nums, low, mid-1),
		Right: buildBST(nums, mid+1, high),
	}
}
```

**执行过程示例**（`nums = [-10,-3,0,5,9]`，取下中点）：

```
buildBST(0,4): mid=2，根=0
├─ buildBST(0,1): mid=0，根=-10
│   ├─ buildBST(0,-1): 空 → nil
│   └─ buildBST(1,1): mid=1，根=-3（左右均为空）
└─ buildBST(3,4): mid=3，根=5
    ├─ buildBST(3,2): 空 → nil
    └─ buildBST(4,4): mid=4，根=9（左右均为空）

结果树（层序 [0,-10,5,null,-3,null,9]，即官方示例认可的另一种答案）:
        0
      /   \
    -10    5
      \     \
      -3     9
```
