# 4. 寻找两个正序数组的中位数 (Median of Two Sorted Arrays)

## 题目描述

给定两个大小分别为 `m` 和 `n` 的正序（从小到大）数组 `nums1` 和 `nums2`。请你找出并返回这两个正序数组的 **中位数** 。

算法的时间复杂度应该为 `O(log (m+n))`。

### 示例 1

```
输入: nums1 = [1,3], nums2 = [2]
输出: 2.00000
解释: 合并数组 = [1,2,3] ，中位数 2
```

### 示例 2

```
输入: nums1 = [1,2], nums2 = [3,4]
输出: 2.50000
解释: 合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
```

### 提示

- `nums1.length == m`
- `nums2.length == n`
- `0 <= m <= 1000`
- `0 <= n <= 1000`
- `1 <= m + n <= 2000`
- `-10^6 <= nums1[i], nums2[i] <= 10^6`

## 题目解析

### 核心思路

中位数的本质是：**把所有元素分成左右两半，左半的最大值 ≤ 右半的最小值，且两半元素个数相等（或左半多一个）**。偶数总长度时答案 = (左半最大值 + 右半最小值) / 2，奇数时答案 = 左半最大值。

如果直接把两个数组合并排序再取中位数，复杂度是 `O(m+n)`，不满足要求。要拿到 `O(log (m+n))`，就要用 **二分查找**——但二分什么？

**关键转化：在两个数组中各找一个“切割位置”**，使左边总元素个数为 `(m+n+1)/2`（左半可以多一个，正好覆盖奇偶两种情况）。设 `nums1` 在 `i` 处切割（`i` 表示左半在 nums1 中取前 `i` 个），`nums2` 在 `j` 处切割，那么 `i + j = (m+n+1)/2`，所以只要确定了 `i`，`j` 就跟着确定了。

**`i` 满足什么条件才算切对了？** 切割点左边的最大值 ≤ 切割点右边的最小值，即：

```
nums1[i-1] <= nums2[j]  且  nums2[j-1] <= nums1[i]
```

对照下图理解（竖线是切割位置）：

```
nums1:  1   3  |  4   6      （i = 2）
nums2:  2   5  |  7   8      （j = 2）

左半: 1 3 | 2 5    → 最大值 = max(3, 5) = 5
右半: 4 6 | 7 8    → 最小值 = min(4, 7) = 4
```

若 `nums1[i-1] > nums2[j]`：说明 `nums1` 的左半段切得太靠右了，`i` 要减小；若 `nums2[j-1] > nums1[i]`：`i` 要增大。这正是单调的，可以对 `i` 做二分。

还有两个细节保证实现干净：

1. **始终让 `nums1` 是较短的那个数组**（长度不够就交换），这样二分区间是 `[0, m]`（`m ≤ n`），且切割位置 `j = half - i` 永远不会越界。
2. **边界处理**：`i == 0` 时 nums1 左半为空，左半最大值只取 `nums2[j-1]`；`i == m` 时 nums1 右半为空，右半最小值只取 `nums2[j]`。`j` 为 0 或 `n` 同理。这覆盖了“某个数组为空”“中位数完全来自另一个数组”的情况。

### 算法步骤

1. 若 `len(nums1) > len(nums2)`，交换两者，保证 `m ≤ n`。
2. 计算 `half = (m + n + 1) / 2`（左半应含的元素个数），在 `i ∈ [0, m]` 上二分：
   - `j = half - i`。
   - 若 `i < m` 且 `nums2[j-1] > nums1[i]`：`i` 太小，`left = i + 1`。
   - 若 `i > 0` 且 `nums1[i-1] > nums2[j]`：`i` 太大，`right = i - 1`。
   - 否则切割正确，退出循环。
3. 计算 `maxLeft`（左半最大值，注意 `i==0`/`j==0` 的边界）：
   - 总长度为奇数，直接返回 `maxLeft`；
   - 否则再计算 `minRight`（右半最小值，注意 `i==m`/`j==n` 的边界），返回 `(maxLeft + minRight) / 2.0`。

### 复杂度分析

- **时间复杂度**: O(log(min(m, n)))，只对较短的数组做二分
- **空间复杂度**: O(1)，只使用常数额外空间

## 代码实现

```go
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // 保证 nums1 是较短的数组，使二分区间 [0, m] 有意义
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }
    m, n := len(nums1), len(nums2)
    half := (m + n + 1) / 2 // 左半部分应包含的元素个数（奇数时左半多一个）

    left, right := 0, m
    for left <= right {
        i := left + (right-left)/2 // 防溢出写法
        j := half - i              // j 由 i 唯一确定，且保证 0 <= j <= n
        if i < m && nums2[j-1] > nums1[i] {
            left = i + 1 // i 太小，左半最大值越界了
        } else if i > 0 && nums1[i-1] > nums2[j] {
            right = i - 1 // i 太大
        } else {
            // 切割正确：求左半最大值 maxLeft
            maxLeft := 0
            if i == 0 {
                maxLeft = nums2[j-1]
            } else if j == 0 {
                maxLeft = nums1[i-1]
            } else if nums1[i-1] > nums2[j-1] {
                maxLeft = nums1[i-1]
            } else {
                maxLeft = nums2[j-1]
            }
            if (m+n)%2 == 1 {
                return float64(maxLeft) // 奇数：左半多一个元素，maxLeft 就是中位数
            }
            // 偶数：再求右半最小值 minRight
            minRight := 0
            if i == m {
                minRight = nums2[j]
            } else if j == n {
                minRight = nums1[i]
            } else if nums1[i] < nums2[j] {
                minRight = nums1[i]
            } else {
                minRight = nums2[j]
            }
            return float64(maxLeft+minRight) / 2.0
        }
    }
    return 0.0 // 题目保证 m+n >= 1，不会走到这里
}
```

**执行过程示例**（`nums1 = [1,3]`, `nums2 = [2,4,5]`）：

```
m=2, n=3, half=(2+3+1)/2=3，对 i ∈ [0,2] 二分

i=1, j=2: nums1[0]=1 <= nums2[2]=5 ✓；nums2[1]=4 > nums1[1]=3 ✗ → i 太小, left=2
i=2, j=1: nums1[1]=3 <= nums2[1]=4 ✓；nums2[0]=2 <= nums1[2]? i==m 无 nums1[2]，跳过 ✓ → 切割正确

切割: [1,3 | ] 与 [2 | 4,5]
maxLeft = max(3, 2) = 3；m+n=5 为奇数 → 返回 3.0
合并数组 [1,2,3,4,5] 的中位数正是 3 ✓
```
