# 169. 多数元素 (Majority Element)

## 题目描述

给定一个大小为 `n` 的数组 `nums`，返回其中的多数元素。多数元素是指在数组中出现次数**严格大于** `⌊n / 2⌋` 的元素。

你可以假设数组是非空的，并且**给定的数组总是存在多数元素**。

### 示例 1

```
输入: nums = [3,2,3]
输出: 3
```

### 示例 2

```
输入: nums = [2,2,1,1,1,2,2]
输出: 2
```

### 提示

- `n == nums.length`
- `1 <= n <= 5 * 10^4`
- `-10^9 <= nums[i] <= 10^9`

## 题目解析

### 核心思路

题目的关键信息是"多数元素出现次数**超过一半**"。这意味着：**把所有等于多数元素的票记 +1，不等于它的记 -1，全部加起来总和一定为正**。

由此得到 **Boyer-Moore 投票算法**：

- 维护一个"候选人"和它当前的"净票数" `count`。
- 遍历数组：净票数为 0 时换个候选人（前面的元素互相抵消光了，多数元素在剩余部分仍然占多数，所以换人重新开始不影响答案）。
- 遍历结束后留下的候选人就是多数元素。

为什么是对的？因为多数元素的数量超过其他所有元素之和，无论它们怎么穿插，多数元素都"抵消不完"，最后一定剩下来。整个过程 O(n) 时间、O(1) 空间。

直观但不省空间的替代方案是**哈希计数**：统计每个元素次数，超过 `n/2` 即返回，O(n) 时间、O(n) 空间。代码里也保留了它作对照。

### 算法步骤

1. 初始化 `candidate = 任意值`、`count = 0`。
2. 遍历数组每个元素 `v`：
   - 若 `count == 0`，令 `candidate = v`（更换候选人）；
   - 若 `v == candidate`，`count++`，否则 `count--`。
3. 返回 `candidate`。

### 复杂度分析

- **投票法 —— 时间复杂度**: O(n)，**空间复杂度**: O(1)
- **哈希计数法 —— 时间复杂度**: O(n)，**空间复杂度**: O(n)

## 代码实现

### 主解：Boyer-Moore 投票

```go
func MajorityElement(nums []int) int {
    candidate := 0
    count := 0 // 候选人与已遍历元素之间的“净票数”
    for _, v := range nums {
        if count == 0 {
            candidate = v
        }
        if v == candidate {
            count++
        } else {
            count--
        }
    }
    return candidate
}
```

**执行过程示例**（`nums = [2,2,1,1,1,2,2]`，多数元素 2 出现 4 次 > 7/2）：

```
v=2: count=0 → 候选人换为2; v==2 → count=1
v=2: count=2
v=1: count=1
v=1: count=0
v=1: count=0 → 候选人换为1; v==1 → count=1
v=2: count=0
v=2: count=0 → 候选人换为2; v==2 → count=1
结果: 返回 2
```

### 对照：哈希计数

```go
func MajorityElementHash(nums []int) int {
    half := len(nums) / 2
    counts := make(map[int]int)
    for _, v := range nums {
        counts[v]++
        if counts[v] > half {
            return v
        }
    }
    return 0
}
```
