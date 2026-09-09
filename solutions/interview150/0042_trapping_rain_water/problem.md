# 0042. 接雨水 (Trapping Rain Water)

## 题目描述

给定 `n` 个非负整数表示每个宽度为 `1` 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

### 示例

```
输入: height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
解释: 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，
     在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
```

```
输入: height = [4,2,0,3,2,5]
输出: 9
```

### 提示

- `n == height.length`
- `1 <= n <= 2 * 10^4`
- `0 <= height[i] <= 10^5`

## 四步拆解

关键观察：**位置 `i` 能接的水量，只取决于它左边最高柱子和右边最高柱子中较矮的那个**——水会被「短板」拦住（木桶效应）。即：

```
water[i] = min(左边最高, 右边最高) - height[i]   （若结果为正）
```

### 1. 状态定义

- `leftMax[i]` 表示**下标 `i` 及其左侧所有柱子的最大高度**；
- `rightMax[i]` 表示**下标 `i` 及其右侧所有柱子的最大高度**。
- 最终答案：`sum(min(leftMax[i], rightMax[i]) - height[i])`，对所有 `i` 求和。

为什么这样定义？因为「`i` 左边最高柱」和「`i` 右边最高柱」都天然满足递推：第 `i` 根柱子的左/右最高值，等于前一个位置的对应值与 `height[i]` 取 max，正是标准的 DP 子结构。

### 2. 边界条件 + 遍历顺序

- base case：
  - `leftMax[0] = height[0]`（最左端左边没有别的柱子）；
  - `rightMax[n-1] = height[n-1]`（最右端同理）。
- 遍历顺序：
  - `leftMax` **从左到右**正序填，`leftMax[i]` 依赖 `leftMax[i-1]`；
  - `rightMax` **从右到左**倒序填，`rightMax[i]` 依赖 `rightMax[i+1]`。
  - 两个方向的依赖都在已计算的一侧，顺序天然保证子问题先求解。

### 3. 状态转移方程

```
leftMax[i]  = max(leftMax[i-1], height[i])
rightMax[i] = max(rightMax[i+1], height[i])
water[i]    = min(leftMax[i], rightMax[i]) - height[i]
```

推导：考虑第 `i` 根柱子，它上方的水位由左右两侧最高的「围墙」中较矮的那一侧决定，水位减去自身高度就是接水量。由于 `leftMax[i] >= height[i]` 且 `rightMax[i] >= height[i]`，结果一定非负，无需特判。两端柱子至少一侧的最大值等于自身，接水量为 0，自然被公式覆盖。

### 4. 样例验证 + 代码实现

手动推导 `height = [4,2,0,3,2,5]`：

| i | height | leftMax | rightMax | min(L,R) - height |
| --- | --- | --- | --- | --- |
| 0 | 4 | 4 | 5 | 4-4=0 |
| 1 | 2 | 4 | 5 | 4-2=2 |
| 2 | 0 | 4 | 5 | 4-0=4 |
| 3 | 3 | 4 | 5 | 4-3=1 |
| 4 | 2 | 4 | 5 | 4-2=2 |
| 5 | 5 | 5 | 5 | 5-5=0 |

求和：`0+2+4+1+2+0 = 9` ✅（与示例一致）

代码与上面的推导一一对应：

```go
// Trap 接雨水（标准 DP 数组版）
func Trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}
```

## 双指针版（不使用 DP 数组）

重新理解 `water[i] = min(leftMax[i], rightMax[i]) - height[i]`：真正起决定作用的是**左右两侧最大值中较小的那一侧**。用两个指针 `left`、`right` 从两端向中间走，同时维护两个变量 `lMax`、`rMax`：

- 若 `lMax < rMax`，说明位置 `left` 的接水量由左侧短板决定（此时右侧最大值是多少根本不影响 `min`），即 `lMax - height[left]`，然后 `left++`；
- 否则由右侧短板决定，即 `rMax - height[right]`，然后 `right--`。

每一步都确定一个位置的接水量，与 DP 公式完全等价，但全程只用 O(1) 额外空间。本题属于「相向双指针」模型。需要说明的是，interview150 官方题单把本题归在「数组/字符串」章节，推荐解法是双指针；这里仍按 DP 框架以 DP 数组版为主版，双指针作为 Alternative 版本。

```go
// TrapAlternative 接雨水（双指针版）
func TrapAlternative(height []int) int {
	left, right := 0, len(height)-1
	lMax, rMax := 0, 0
	ans := 0
	for left < right {
		lMax = max(lMax, height[left])
		rMax = max(rMax, height[right])
		if lMax < rMax { // 左侧是短板，结算 left 位置
			ans += lMax - height[left]
			left++
		} else { // 右侧是短板，结算 right 位置
			ans += rMax - height[right]
			right--
		}
	}
	return ans
}
```

**执行过程示例**（`height = [0,1,0,2,1,0,1,3,2,1,2,1]`，双指针）：

```
初始化: left=0, right=11, lMax=0, rMax=0

left 侧移动（lMax <= rMax）:
  left=0: h=0, lMax=0, 接 0, left=1
  left=1: h=1, lMax=1, 接 0, left=2
  left=2: h=0 < 1 -> 接 1-0=1, total=1, left=3
  left=3: h=2, lMax=2, 接 0, left=4
  left=4: h=1 < 2 -> 接 2-1=1, total=2, left=5
  left=5: h=0 < 2 -> 接 2-0=2, total=4, left=6
  left=6: h=1 < 2 -> 接 2-1=1, total=5, left=7

right 侧移动（h[left]=3 使 lMax=3 > rMax=1，处理右侧）:
  right=11: h=1, rMax=1, 接 0, right=10
  right=10: h=2, rMax=2, 接 0, right=9
  right=9: h=1 < 2 -> 接 2-1=1, total=6, right=8
  right=8: h=2, rMax=2, 接 0, right=7
  left == right，循环结束

结果: 6
```

## 复杂度分析

| 版本 | 时间复杂度 | 空间复杂度 |
| --- | --- | --- |
| DP 数组版 | O(n) | O(n) |
| 双指针版 | O(n) | O(1) |

**进阶**：本题还可以用**单调栈**求解——栈中保存高度单调递减的下标，遇到更高的柱子时以栈顶为「谷底」弹出，横向计算一段积水面积，时间 O(n)、空间 O(n)。单调栈视角适合理解「按层横向接水」，面试中可作为扩展思路了解。
