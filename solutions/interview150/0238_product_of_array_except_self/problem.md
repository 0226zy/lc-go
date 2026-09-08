# 238. 除了自身以外数组的乘积 (Product of Array Except Self)

## 题目描述

给你一个整数数组 `nums`，返回数组 `answer`，其中 `answer[i]` 等于 `nums` 中除 `nums[i]` 之外其余各元素的乘积。

**注意：** 题目数据保证数组中任意前缀或后缀的乘积都 **32 位整数** 范围内。

**要求：** 请 **不要使用除法**，且需要在 O(n) 时间复杂度内完成。

### 示例 1

```
输入: nums = [1,2,3,4]
输出: [24,12,8,6]
解释: 除自身以外的乘积：
      answer[0] = 2*3*4 = 24
      answer[1] = 1*3*4 = 12
      answer[2] = 1*2*4 = 8
      answer[3] = 1*2*3 = 6
```

### 示例 2

```
输入: nums = [-1,1,0,-3,3]
输出: [0,0,9,0,0]
```

## 提示

- `2 <= nums.length <= 10^5`
- `-30 <= nums[i] <= 30`
- 保证任意前缀或后缀的乘积都在 32 位整数范围内
- 进阶：你可以在 O(1) 的额外空间复杂度内完成（返回数组不计入额外空间）

## 题目解析

### 核心思路

最直观的想法是「先算全部乘积，再除以自己」，但题目明确**禁止除法**（而且遇到 0 时除法本身就不可行）。

换个角度想：`answer[i]` = 左边所有数的乘积 × 右边所有数的乘积，也就是 `nums[0..i-1]` 的乘积（前缀积）乘以 `nums[i+1..n-1]` 的乘积（后缀积）。这属于典型的**前后缀分解**模型。

关键在于如何把空间压下来：先把结果数组 `answer` 当作「前缀积数组」，再用一个滚动的 `suffix` 变量从右往左把后缀积乘进去。这样除了返回数组外，只用一个变量。

### 算法步骤

1. 创建结果数组 `answer`，长度与 `nums` 相同。
2. **第一遍（从左往右）**：令 `answer[0] = 1`，然后 `answer[i] = answer[i-1] * nums[i-1]`，循环结束后 `answer[i]` 就是 `nums[i]` 左边所有元素的乘积。
3. **第二遍（从右往左）**：用一个变量 `suffix`（初值为 1）记录右边所有元素的乘积；每到一个位置，先 `answer[i] *= suffix`，再 `suffix *= nums[i]`。
4. 返回 `answer`。

### 复杂度分析

- **时间复杂度**: O(n)，前后各遍历一次
- **空间复杂度**: O(1)，除返回数组外只用一个变量（不计入额外空间）

## 代码实现

```go
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// 第一遍：answer[i] 先存 nums[0..i-1] 的前缀乘积
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	// 第二遍：用后缀乘积从右往左累乘到 answer 中
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}
	return answer
}
```

**执行过程示例**（`nums = [1,2,3,4]`）：

```
第一遍（前缀积）:
  answer[0] = 1
  answer[1] = answer[0]*nums[0] = 1*1 = 1
  answer[2] = answer[1]*nums[1] = 1*2 = 2
  answer[3] = answer[2]*nums[2] = 2*3 = 6
  得到 answer = [1, 1, 2, 6]（每个位置左边元素的乘积）

第二遍（乘上后缀积）:
  i=3: answer[3] *= suffix(1) -> 6, suffix *= 4 -> 4
  i=2: answer[2] *= suffix(4) -> 8, suffix *= 3 -> 12
  i=1: answer[1] *= suffix(12) -> 12, suffix *= 2 -> 24
  i=0: answer[0] *= suffix(24) -> 24, suffix *= 1 -> 24
  得到 answer = [24, 12, 8, 6]
```
