# 11. 盛最多水的容器 (Container With Most Water)

## 题目描述

给定一个长度为 `n` 的整数数组 `height`。有 `n` 条竖线，第 `i` 条线的两个端点是 `(i, 0)` 和 `(i, height[i])`。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

**说明**：你不能倾斜容器。

### 示例 1

```
输入: height = [1,8,6,2,5,4,8,3,7]
输出: 49
解释: 选择第 2 条线（高 8）和第 9 条线（高 7），容器容量 = min(8,7) * (8-1) = 49，为最大值。
```

### 示例 2

```
输入: height = [1,1]
输出: 1
```

## 提示

- `n == height.length`
- `2 <= n <= 10^5`
- `0 <= height[i] <= 10^4`

## 题目解析

### 核心思路

暴力解法是枚举所有两条线的组合，计算 `min(height[i], height[j]) * (j - i)` 取最大，时间复杂度 O(n²)。在 `n` 最大为 `10^5` 的约束下会超时（10¹⁰ 次运算），必须优化。

观察：**容器面积 = 底宽 × 较矮一侧的高度**（短板效应）。使用**双指针**：

- 左指针 `left` 指向数组开头，右指针 `right` 指向数组末尾，此时底宽最大。
- 每轮计算当前面积并更新答案，然后**移动较矮一侧的指针**，向中间收缩，直到两指针相遇。

### 为什么移动较矮一侧不会错过最优解？

设当前两指针处的容器高度为短板 `h = min(height[left], height[right])`，宽度为 `w = right - left`：

- **若移动较高一侧的指针**：新宽度必然变小，而新短板 `≤` 原短板 `h`（短板由较矮侧决定，较矮侧没有动，新高度最多等于原短板），所以新面积一定**严格小于**当前面积 `h * w`——这个方向没有任何希望。
- **若移动较矮一侧的指针**：新宽度同样变小，但有可能遇到更高的线，把短板抬高，面积可能变大。这是唯一可能产生更优解的方向。

更严谨地说：假设 `height[left] < height[right]`，那么对于任何在 `left` 右侧的指针 `j`（`j < right`），组合 `(left, j)` 的面积 `= min(height[left], height[j]) * (j - left) ≤ height[left] * (right - left)`，即**不可能超过当前面积**。也就是说，`left` 这条线作为短板的所有配对都已被当前状态“封顶”，可以放心把它排除掉（`left++`），不会错过最优解。

因此每轮排除一条不可能成为最优解的线，n 条线最多扫描 n 次，时间复杂度降为 O(n)。

### 算法步骤

1. 初始化 `left = 0`、`right = n - 1`、`ans = 0`。
2. 当 `left < right` 时循环：
   - 计算当前面积 `area = min(height[left], height[right]) * (right - left)`，用 `area` 更新 `ans`。
   - 若 `height[left] < height[right]`，则 `left++`（排除较矮的左线）；否则 `right--`（排除较矮或等高的右线）。
3. 返回 `ans`。

### 复杂度分析

- **时间复杂度**: O(n)，双指针每轮至少收缩一格，总共最多移动 n 次。
- **空间复杂度**: O(1)，只使用常数个变量。

## 代码实现

```go
func MaxArea(height []int) int {
    left, right := 0, len(height)-1
    ans := 0

    for left < right {
        // 容器面积 = 底宽 * 较矮一侧的高度（短板效应）
        width := right - left
        h := height[left]
        if height[right] < h {
            h = height[right]
        }
        if area := width * h; area > ans {
            ans = area
        }

        // 移动较矮一侧的指针才有可能找到更高的短板
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return ans
}
```

**执行过程示例**（`height = [1,8,6,2,5,4,8,3,7]`）：

```
left=0(h=1), right=8(h=7): 面积=min(1,7)*8=8   左矮 → left++
left=1(h=8), right=8(h=7): 面积=min(8,7)*7=49  右矮 → right--   ← 最大
left=1(h=8), right=7(h=3): 面积=min(8,3)*6=18  右矮 → right--
left=1(h=8), right=6(h=8): 面积=min(8,8)*5=40  等高 → right--
left=1(h=8), right=5(h=4): 面积=min(8,4)*4=16  右矮 → right--
left=1(h=8), right=4(h=5): 面积=min(8,5)*3=15  右矮 → right--
left=1(h=8), right=3(h=2): 面积=min(8,2)*2=4   右矮 → right--
left=1(h=8), right=2(h=6): 面积=min(8,6)*1=6   右矮 → right--
left==right，结束，答案 = 49
```
