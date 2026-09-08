# 33. 搜索旋转排序数组 (Search in Rotated Sorted Array)

## 题目描述

整数数组 `nums` 按升序排列，数组中的值 **互不相同**。

在传递给函数之前，`nums` 在预先未知的某个下标 `k`（`0 <= k < nums.length`）上进行了 **旋转**，使数组变为 `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]`（下标从 0 开始计数）。例如，`[0,1,2,4,5,6,7]` 在下标 3 处旋转后变为 `[4,5,6,7,0,1,2]`。

给定旋转后的数组 `nums` 和一个整数 `target`，如果 `nums` 中存在这个目标值 `target`，则返回它的下标，否则返回 `-1`。

你必须设计一个时间复杂度为 `O(log n)` 的算法解决此问题。

### 示例 1

```
输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
```

### 示例 2

```
输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
```

### 示例 3

```
输入: nums = [1], target = 0
输出: -1
```

### 提示

- `1 <= nums.length <= 5000`
- `-10^4 <= nums[i] <= 10^4`
- `nums` 中的每个值都 **独一无二**
- `nums` 是一个旋转过的升序数组
- `-10^4 <= target <= 10^4`

## 题目解析

### 核心思路

这是一道经典的 **“旋转排序数组 + 二分查找”** 模型题。数组本身是有序的，只是被从中间“掰弯”了，但关键在于：**旋转后的数组从中间切一刀，必然有一半是完全有序的**。

以 `[4,5,6,7,0,1,2]` 为例，取中点 `mid = 3`（值为 7）：
- 左半 `[4,5,6,7]` 是严格递增的（因为 `nums[0] < nums[mid]`）；
- 右半 `[0,1,2]` 虽短，但下一轮二分它也会成为“有序的一半”。

于是每轮二分的判断逻辑是：

1. 判断哪一半有序（比较 `nums[0]` 与 `nums[mid]`，或比较 `nums[mid]` 与 `nums[right]`）。
2. 如果 `target` 落在这有序的一半内（闭区间判断），就丢弃另一半；否则丢弃这一半。
3. 每轮都能把搜索范围缩小一半，最终要么命中 `target`，要么区间为空返回 `-1`。

由于每次都砍掉一半元素，总轮数是 `O(log n)`，满足题目要求。

**为什么正确？** 旋转排序数组只有一个“下降点”（旋转点）。中点要么在下降点左边（左半有序），要么在右边（右半有序）。有序的那半段内的元素大小关系完全正常，可以直接用区间包含判断 `target` 是否在里面，从而放心地丢弃另一半。

### 算法步骤

1. 初始化 `left = 0`，`right = len(nums) - 1`。
2. 当 `left <= right` 时循环：
   - 计算中点 `mid = left + (right - left) / 2`（避免溢出）。
   - 若 `nums[mid] == target`，直接返回 `mid`。
   - 若 `nums[left] <= nums[mid]`：左半 `[left, mid]` 有序。
     - 若 `nums[left] <= target < nums[mid]`，说明 `target` 在左半，`right = mid - 1`；
     - 否则 `left = mid + 1`。
   - 否则：右半 `(mid, right]` 有序。
     - 若 `nums[mid] < target <= nums[right]`，说明 `target` 在右半，`left = mid + 1`；
     - 否则 `right = mid - 1`。
3. 循环结束仍未命中，返回 `-1`。

### 复杂度分析

- **时间复杂度**: O(log n)，每轮二分搜索范围缩小一半
- **空间复杂度**: O(1)，只使用常数额外空间

## 代码实现

```go
func Search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2 // 防溢出
        if nums[mid] == target {
            return mid
        }
        if nums[left] <= nums[mid] { // 左半有序
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else { // 右半有序
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    return -1
}
```

**执行过程示例**（`nums = [4,5,6,7,0,1,2]`, `target = 0`）：

```
left=0, right=6, mid=3: nums[3]=7 ≠ 0
  左半 [4,5,6,7] 有序，0 < 4 不在其中 → left = 4
left=4, right=6, mid=5: nums[5]=1 ≠ 0
  nums[4]=0 <= nums[5]=1，左半 [0,1] 有序，0 <= 0 < 1 在其中 → right = 4
left=4, right=4, mid=4: nums[4]=0 == 0 → 返回 4
```
