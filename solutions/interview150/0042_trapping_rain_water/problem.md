# 42. 接雨水 (Trapping Rain Water)

## 题目描述

给定 `n` 个非负整数，表示柱状图中每个柱子的高度，每个柱子紧挨着放在一起，宽度为 1。计算下雨之后这个柱状图能接多少雨水。

### 示例 1

```
输入: height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
解释: 下雨后接到的雨水如下图所示，总共可以接 6 个单位的雨水。
```

### 示例 2

```
输入: height = [4,2,0,3,2,5]
输出: 9
```

## 提示

- `1 <= n <= 2 * 10^4`
- `0 <= height[i] <= 10^5`

## 题目解析

### 核心思路

关键在于想清楚：**位置 `i` 处能接多少水由什么决定？**

答案：`i` 处的水量 = `min(左边最高柱子, 右边最高柱子) - height[i]`（若结果为正）。因为水位到达两侧最高柱中较矮的那一根就会被「溢出」，而柱子本身占掉了一部分体积。

由此可以演化出三种解法：

1. **动态规划**：预处理出每个位置的 `leftMax[i]` 和 `rightMax[i]`，直接套公式。O(n) 时间、O(n) 空间。
2. **单调栈**：栈里存高度单调递增的下标，遇到更高的柱子时，以栈顶为「谷底」弹出计算一段横向的积水。O(n) 时间、O(n) 空间。
3. **双指针（最优解）**：仔细观察 DP 公式——每个位置只取决于 `min(leftMax, rightMax)`。用两个指针从两端往中间夹：**哪边的柱子矮，就先处理哪边**。因为矮的那一侧决定了水位上限，它自己的 `leftMax`（或 `rightMax`）就是限制值，此时另一侧的最高柱是多少根本不影响结果。这样只需两个指针和两个变量，O(n) 时间、**O(1) 空间**。

本题最推荐的解法是**双指针**，属于「相向双指针」模型。

### 算法步骤（双指针）

1. 初始化指针 `left = 0`，`right = n - 1`，以及 `leftMax = rightMax = 0`。
2. 当 `left < right` 时循环：
   - 若 `height[left] < height[right]`（左边矮）：
     - 若 `height[left] >= leftMax`，更新 `leftMax = height[left]`（这根柱子比左侧所有柱都高，接不到水）；
     - 否则它能接 `leftMax - height[left]` 的水，累加进答案。
     - `left++`
   - 否则（右边矮或等高）：对称地处理右边，能接 `rightMax - height[right]` 的水，`right--`。
3. 循环结束返回累加值。

**为什么这样是对的？** 当 `height[left] < height[right]` 时，对 `left` 位置而言，`rightMax` 至少和 `height[right]` 一样大，因此 `min(leftMax, rightMax) = leftMax` 一定成立——直接用 `leftMax` 计算即可，无需知道 `rightMax` 的精确值。右侧同理。每个位置的水量计算与 DP 公式完全一致，所以结果正确。

### 复杂度分析

- **时间复杂度**: O(n)，每个指针各走一遍
- **空间复杂度**: O(1)，双指针最优解；DP 与单调栈解法需要 O(n) 空间

## 代码实现

### 最优解：双指针

```go
func Trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	total := 0
	for left < right {
		// 矮的一侧决定水位；移动矮侧指针，若当前柱比该侧最高矮则能接水
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				total += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				total += rightMax - height[right]
			}
			right--
		}
	}
	return total
}
```

### 对照解法：动态规划

```go
func TrapDP(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	total := 0
	for i := 1; i < n-1; i++ {
		total += min(leftMax[i], rightMax[i]) - height[i]
	}
	return total
}
```

### 对照解法：单调栈

```go
func TrapStack(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	stack := make([]int, 0, n)
	total := 0
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			bottom := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			leftIdx := stack[len(stack)-1]
			width := i - leftIdx - 1
			boundedHeight := min(height[leftIdx], height[i]) - bottom
			total += width * boundedHeight
		}
		stack = append(stack, i)
	}
	return total
}
```

**执行过程示例**（`height = [0,1,0,2,1,0,1,3,2,1,2,1]`，双指针）：

```
初始化: left=0, right=11, leftMax=0, rightMax=0

left 侧移动（height[left] < height[right]）:
  left=0: h=0 >= leftMax=0 -> leftMax=0, left=1
  left=1: h=1 >= 0 -> leftMax=1, left=2
  left=2: h=0 < 1 -> 接 1-0=1, total=1, left=3
  left=3: h=2 >= 1 -> leftMax=2, left=4
  left=4: h=1 < 2 -> 接 2-1=1, total=2, left=5
  left=5: h=0 < 2 -> 接 2-0=2, total=4, left=6
  left=6: h=1 < 2 -> 接 2-1=1, total=5, left=7

right 侧移动（h[left]=3 > h[right]=1，处理右侧）:
  right=11: h=1 >= rightMax=0 -> rightMax=1, right=10
  right=10: h=2 >= 1 -> rightMax=2, right=9
  right=9: h=1 < 2 -> 接 2-1=1, total=6, right=8
  right=8: h=2 >= 2 -> rightMax=2, right=7
  left == right，循环结束

结果: 6
```
