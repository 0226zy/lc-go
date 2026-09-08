# 15. 三数之和 (3Sum)

## 题目描述

给你一个整数数组 `nums` ，判断是否存在三元组 `[nums[i], nums[j], nums[k]]` 满足 `i != j`、`i != k` 且 `j != k` ，同时还满足 `nums[i] + nums[j] + nums[k] == 0` 。请你返回所有和为 `0` 且**不重复**的三元组。

注意：答案中不可以包含重复的三元组。

### 示例 1

```
输入: nums = [-1,0,1,2,-1,-4]
输出: [[-1,-1,2],[-1,0,1]]
解释:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
```

### 示例 2

```
输入: nums = [0,1,1]
输出: []
解释: 唯一可能的三元组和不为 0 。
```

### 示例 3

```
输入: nums = [0,0,0]
输出: [[0,0,0]]
解释: 唯一可能的三元组和为 0 。
```

## 提示

- `3 <= nums.length <= 3000`
- `-10^5 <= nums[i] <= 10^5`

## 题目解析

### 核心思路

暴力三重循环是 O(n³)，n 最大 3000 完全不可行。经典做法是**排序 + 双指针**，把问题转化为「两数之和」：

1. 先把数组**升序排序**。排序后相同元素相邻，天然方便去重；同时可以利用有序性做单调剪枝。
2. 枚举第一个数 `nums[i]`，问题退化为在 `nums[i+1..n-1]` 中找两个数之和等于 `-nums[i]`。有序数组上找两数之和可以用**对撞双指针**：`left` 从 `i+1` 开始，`right` 从 `n-1` 开始，根据当前三数之和移动指针。
3. **去重是本题的关键**：
   - `i` 去重：若 `nums[i] == nums[i-1]`，说明以该值为首元素的所有三元组在上一轮已找过，直接跳过。
   - 找到一个三元组后，`left` 右移跳过所有与 `nums[left]` 相同的元素，`right` 左移跳过所有与 `nums[right]` 相同的元素。
4. **剪枝**：排序后若 `nums[i] > 0`，其后所有数都非负，三数之和不可能为 0，直接结束。

排序保证了每个三元组内部天然按升序输出，且每组解只会被发现一次。

### 算法步骤

1. 对 `nums` 升序排序。
2. 遍历 `i` 从 `0` 到 `n-1`：
   - 若 `nums[i] > 0`，提前结束（剪枝）。
   - 若 `i > 0` 且 `nums[i] == nums[i-1]`，跳过（`i` 去重）。
   - 令 `left = i+1`，`right = n-1`，当 `left < right` 时：
     - 计算 `sum = nums[i] + nums[left] + nums[right]`。
     - `sum < 0`：`left++`（需要更大的数）。
     - `sum > 0`：`right--`（需要更小的数）。
     - `sum == 0`：记录三元组，然后 `left` 右移、`right` 左移并各自跳过重复值。
3. 返回收集到的所有三元组。

### 复杂度分析

- **时间复杂度**: O(n²)，排序 O(n log n)，外层枚举 O(n)、内层双指针总共 O(n)，整体 O(n²)。
- **空间复杂度**: O(log n)，排序所需的栈空间（不计输出结果）。

## 代码实现

```go
func ThreeSum(nums []int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    result := make([][]int, 0)

    for i := 0; i < n-2; i++ {
        // 剪枝：排序后 nums[i] > 0，后面不可能凑出和为 0
        if nums[i] > 0 {
            break
        }
        // i 去重：与上一个首元素相同，解已找过
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        left, right := i+1, n-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            switch {
            case sum < 0:
                left++
            case sum > 0:
                right--
            default:
                result = append(result, []int{nums[i], nums[left], nums[right]})
                // 左右指针去重：跳过所有与当前值相同的元素
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
                left++
                right--
            }
        }
    }
    return result
}
```

**执行过程示例**（`nums = [-1,0,1,2,-1,-4]`，排序后为 `[-4,-1,-1,0,1,2]`）：

```
i=1 (nums[i]=-1): left=2(-1), right=5(2) → sum=0 → 记录 [-1,-1,2]
                  去重后 left=3(0), right=4(1) → sum=0 → 记录 [-1,0,1]
i=2: nums[2] == nums[1]，跳过（i 去重）
i=3 (nums[i]=0): left=4(1), right=5(2) → sum=3 > 0 → right--，left==right 结束
结果: [[-1,-1,2], [-1,0,1]]
```
