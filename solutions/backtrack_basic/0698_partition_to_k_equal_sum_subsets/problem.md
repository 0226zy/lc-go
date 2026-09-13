# 698. 划分为k个相等的子集 (Partition to K Equal Sum Subsets)

## 题目描述

给定一个整数数组 `nums` 和一个正整数 `k`，找出是否有可能把这个数组分成 `k` 个非空子集，其总和都相等。

### 示例 1

```
输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
输出： True
说明： 有可能将其分成 4 个子集 {5}, {1,4}, {2,3}, {2,3} 等于总和。
```

### 示例 2

```
输入: nums = [1,2,3,4], k = 3
输出: false
```

## 提示

- `1 <= k <= nums.length <= 16`
- `1 <= nums[i] <= 10^4`
- `nums[i]` 的频率限制在 `[1,4]` 范围内

## 题目解析

### 核心思路

本题是典型的**回溯 + 桶（子集）视角**的划分问题。把数组想象成要装进 `k` 个容量为 `target = sum / k` 的桶里，每个数必须放进恰好一个桶，问能否把所有数放完且每个桶都恰好装满。

回溯三要素：

- **路径**：`k` 个桶当前已装入的元素和 `buckets[i]`，以及当前正在放置的元素下标 `index`。
- **选择列表**：对当前元素 `nums[index]`，可选的动作是放入 `k` 个桶中的任意一个。
- **结束条件**：`index == len(nums)`，所有数都放完。由于总量恰为 `k * target` 且每个桶都不超过 `target`，放完即意味着每个桶恰好等于 `target`，直接返回成功。

三个关键优化（缺一不可，否则大数据会超时）：

1. **前置判断**：`sum % k != 0` 必然失败；最大元素大于 `target` 必然失败。
2. **降序排序**：先放大数。大数选择余地小，能更早触发"桶放不下"的剪枝，把失败分支砍掉在浅层。
3. **桶等量剪枝**：在遍历桶时，若 `buckets[i] == buckets[i-1]`，说明前一个桶和当前桶完全等价——当前数放进前一个桶失败后，再放进这个桶必然得到同样的子问题，直接跳过。

回溯三要素的标记/还原：每次尝试放入时做 `buckets[i] += nums[index]`，递归返回失败后做 `buckets[i] -= nums[index]` 还原现场。

### 算法步骤

1. 计算 `sum`，若 `sum % k != 0` 返回 `false`；令 `target = sum / k`。
2. 将 `nums` 降序排序；若 `nums[0] > target` 返回 `false`。
3. 建立长度为 `k` 的桶数组 `buckets`，初始全为 0。
4. 从 `index = 0` 开始回溯：
   - 若 `index == len(nums)`，返回 `true`；
   - 遍历每个桶 `i`：
     - 若 `buckets[i] + nums[index] > target`，跳过（容量剪枝）；
     - 若 `i > 0 && buckets[i] == buckets[i-1]`，跳过（等量桶剪枝）；
     - 放入桶 `i`，递归 `backtrack(index + 1)`，成功则返回 `true`；
     - 回溯：从桶 `i` 中取出。
5. 所有桶都试过仍失败，返回 `false`。

### 复杂度分析

- **时间复杂度**: O(k^n)，n 为数组长度。每个数有 k 个桶可选，最坏为指数级；降序排序与等量桶剪枝能大幅削减实际搜索量，足以通过 n <= 16 的数据范围。
- **空间复杂度**: O(n)，递归栈深度不超过 n 层；桶数组占用 O(k)（不计排序开销）。

## 代码实现

```go
func CanPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// 总和不能整除 k，直接不可能均分
	if k <= 0 || sum%k != 0 {
		return false
	}
	target := sum / k

	// 降序排列：大数先放，能更早触发"桶放不下"的剪枝
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	// 最大的数超过单个桶容量，必然无法均分
	if nums[0] > target {
		return false
	}

	buckets := make([]int, k) // buckets[i] 表示第 i 个桶当前已装入的元素和

	// backtrack 尝试把 nums[index] 放进某个桶，成功放完所有数则返回 true
	var backtrack func(index int) bool
	backtrack = func(index int) bool {
		// 所有数都放完，且每个桶都恰好为 target，划分成功
		if index == len(nums) {
			return true
		}
		for i := 0; i < k; i++ {
			// 剪枝1：当前桶放不下 nums[index]
			if buckets[i]+nums[index] > target {
				continue
			}
			// 剪枝2：与上一个等量的桶等价，重复尝试没有意义
			if i > 0 && buckets[i] == buckets[i-1] {
				continue
			}
			buckets[i] += nums[index]
			if backtrack(index + 1) {
				return true
			}
			buckets[i] -= nums[index]
		}
		return false
	}
	return backtrack(0)
}
```

**执行过程示例**（`nums = [4,3,2,3,5,2,1], k = 4`，`sum = 20`，`target = 5`）：

```
降序排序后 nums = [5,4,3,3,2,2,1]，4 个桶初始为 [0,0,0,0]

放 5：桶0 装得下 → buckets=[5,0,0,0]
放 4：桶0 放不下(5+4>5)；桶1 装得下 → buckets=[5,4,0,0]
放 3：桶0 放不下；桶1 放不下(4+3>5)；桶2 装得下 → buckets=[5,4,3,0]
放 3：桶0、1、2 都放不下；桶3 装得下 → buckets=[5,4,3,3]
放 2：桶0 放不下；桶1 装得下(4+2=6>5?) 放不下！
     → 桶1 4+2=6>5 跳过，桶2 3+2=5 装得下 → buckets=[5,4,5,3]
放 2：桶0、1、2 放不下；桶3 3+2=5 装得下 → buckets=[5,4,5,5]
放 1：桶0 放不下；桶1 4+1=5 装得下 → buckets=[5,5,5,5]
index == 7 == len(nums)，所有数放完，返回 true

结果：true（对应划分 {5}, {1,4}, {2,3}, {2,3}）
```
