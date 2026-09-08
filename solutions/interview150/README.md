# 面试经典 150 题（Top Interview 150）

LeetCode 官方题单「面试经典 150 题」的 Go 题解合集，共 **150 题 / 23 个分类**。题目归属与顺序以官方题单为准（见 [LIST.md](LIST.md)）。

每道题一个独立目录，包含三个文件：

- `problem.md` —— 通俗易懂的题解：核心思路（为什么这么做）、算法步骤、复杂度分析、带注释的 Go 代码、执行过程示例
- `<name>.go` —— 题解实现（导出函数 + 复杂度注释）
- `<name>_test.go` —— 中文表驱动单元测试 + Benchmark

刷题建议：先看每类下面的「思路与模板」，再按顺序做题；同类题共用同一套模板，吃透模板比背单题重要。

## 分类总览

| # | 分类 | 题数 | 一句话认出它 |
| --- | --- | --- | --- |
| 1 | 数组 / 字符串 | 24 | 原地修改数组、字符串模拟 |
| 2 | 双指针 | 5 | 有序数组 / 两端收敛找组合 |
| 3 | 滑动窗口 | 4 | 连续子串 / 子数组的约束最值 |
| 4 | 矩阵 | 5 | 二维网格上的遍历与变换 |
| 5 | 哈希表 | 9 | “有没有 / 出现过没有 / 分组” |
| 6 | 区间 | 4 | 区间重叠、合并、插入 |
| 7 | 栈 | 5 | 括号匹配、表达式求值、路径简化 |
| 8 | 链表 | 11 | 指针穿针引线、哑节点 |
| 9 | 二叉树（一般） | 14 | 递归三部曲、后序贡献值 |
| 10 | 二叉树 BFS | 4 | 层序遍历、按层处理 |
| 11 | 二叉搜索树 | 3 | 中序遍历 = 有序序列 |
| 12 | 图（一般） | 6 | 网格 DFS、拓扑排序、建图 |
| 13 | 图 BFS | 3 | 求最少步数 / 最短变换 |
| 14 | 字典树 | 3 | 前缀匹配、批量单词查询 |
| 15 | 回溯 | 7 | 暴力枚举所有方案 + 剪枝 |
| 16 | 分治 | 4 | 对半拆开、分别解决、合并结果 |
| 17 | Kadane 算法 | 2 | 最大子数组和 |
| 18 | 二分查找 | 7 | 有序 / 部分有序、求边界 |
| 19 | 堆 | 4 | Top K、动态中位数、多路归并 |
| 20 | 位运算 | 6 | 异或、n&(n-1)、位计数 |
| 21 | 数学 | 6 | 模拟竖式、快速幂、数论技巧 |
| 22 | 一维 DP | 5 | 一维状态的递推 |
| 23 | 多维 DP | 9 | 网格 / 字符串 / 股票状态机 |

## 各类思路与模板

### 1. 数组 / 字符串

**核心模型**：`O(1)` 空间原地修改数组，几乎都靠「读写双指针」。

```go
// 读写指针模板：slow 指向下一个写入位置，fast 负责扫描
slow := 0
for fast := 0; fast < len(nums); fast++ {
    if 应该保留(nums[fast]) {
        nums[slow] = nums[fast]
        slow++
    }
}
return slow // 新长度
```

- **保留非 X 元素 / 去重**：直接套上面的模板；去重 II（允许两个）把条件改成 `nums[fast] != nums[slow-2]`（027、026、080、088）。
- **轮转数组**：`k %= n`，整体翻转 → 翻转前 k 个 → 翻转剩下的（189）。
- **前后缀分解**：答案 = 每个位置「左边所有数 × 右边所有数」，正反两趟扫描（238）。
- **贪心**：跳跃游戏维护「最远可达」；加油站从累积亏损点之后重新开始；股票一次交易 = 维护最低买入价，多次交易 = 把所有上坡都吃满（55、45、134、121、122）。
- **左右两次遍历**：分发糖果、接雨水——每个位置的值同时依赖两侧信息时，就分别从左、从右各扫一遍取约束（135、42）。

### 2. 双指针

**前提**：数组有序（先排序也行）。两端各一个指针，根据大小关系向内收缩。

```go
// 对撞指针模板（以 Two Sum II 为例）
l, r := 0, len(nums)-1
for l < r {
    sum := nums[l] + nums[r]
    if sum == target { /* 命中 */ break }
    if sum < target { l++ } else { r-- }
}
```

- **两数之和**：和太小左移，太大右移（167）。
- **三数之和**：排序后固定一个数，剩下变成有序两数之和；**去重是核心细节**——命中后两个指针都要跳过相同值，固定的数也要跳过（15）。
- **盛最多水的容器**：短板决定高度，每次移动较短的那一边（11）。

### 3. 滑动窗口

**识别特征**：求「连续子串 / 子数组」满足某条件的最长 / 最短 / 个数。

```go
// 可变窗口模板（最短覆盖窗口）
left := 0
for right := 0; right < len(s); right++ {
    加入 s[right]              // 扩大窗口
    for 窗口已满足条件 {
        更新答案
        移除 s[left]          // 收缩窗口
        left++
    }
}
```

- **最短**：满足条件就收缩（209、76）；**最长**：不满足条件才收缩（3）。
- 字符窗口用 `map[byte]int` 计数；固定长度单词窗口按「单词长度」为单位滑动（30）。

### 4. 矩阵

- **螺旋遍历**：维护上下左右四个边界，每走一条边就收缩对应边界，注意单行 / 单列的边界重叠（54）。
- **原地旋转 90°**：先沿主对角线转置，再水平翻转每一行；映射公式 `(r,c) → (c, n-1-r)`（48）。
- **原地标记**：矩阵置零用首行首列当标记位；生命游戏用一个 cell 里编码两种状态（如 `2` = 活→死，`-1` = 死→活），最后统一还原（73、289）。
- **有效性检查**：数独按行 / 列 / 宫三个维度查重，宫编号 `r/3*3 + c/3`（36）。

### 5. 哈希表

- **计数**：字符统计后比较（383 赎金信、242 异位词）。
- **双向映射**：同构字符串必须同时维护 `a→b` 和 `b→a` 两张表（205、290）。
- **分组**：异位词分组以「排序后的字符串」或「计数签名」为 key（49）。
- **下标索引**：两数之和存 `值 → 下标`；存在重复元素 II 存 `值 → 最近一次下标`（1、219）。
- **序列检测**：最长连续序列先全部入集合，只从「序列起点」（`num-1` 不存在）开始向右扩张，保证 O(n)（128）。
- **环检测**：快乐数用快慢指针找环，或哈希记录出现过的数（202）。

### 6. 区间

```go
// 区间通用模板：按起点（或终点）排序 + 一次扫描
sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
for _, iv := range intervals {
    if iv[0] <= 当前合并区间.end { // 有重叠
        当前.end = max(当前.end, iv[1])
    } else { // 无重叠
        输出当前; 开启新区间
    }
}
```

- 合并区间（56）、插入区间（与新区间分类讨论：在左边 / 在右边 / 重叠，57）。
- **按右端点排序的贪心**：射箭问题，每次取当前最早的结束点放一支箭，跳过所有被覆盖的起点（452）。
- 汇总区间：扫描相邻差是否为 1（228）。

### 7. 栈

- **匹配消去**：遇左括号入栈，遇右括号看栈顶是否配对（20）。
- **表达式求值**：操作数入栈，遇运算符弹两个算完再入栈（150）；带括号与负号时用「当前符号位 + 结果栈」通用模板（224）。
- **辅助栈**：最小栈额外维护一个同步最小值栈（155）。
- **路径简化**：按 `/` 切分，`.` 忽略、`..` 弹栈、正常段入栈，最后拼接（71）。

### 8. 链表

**三板斧**：

1. **哑节点**（dummy head）：删除 / 插入时统一头节点的边界处理。
2. **快慢指针**：找中点、倒数第 N 个（先走 N 步再同步走）、判环（Floyd）。
3. **局部反转**：记住 `prev` 穿针引线，K 个一组翻转 = 每 K 个先断开、反转、再接回（25、92）。

- 随机指针复制：哈希表 `旧节点 → 新节点` 两趟，或原地交错插入（138）。
- LRU 缓存：哈希表定位 + 双向链表维护使用顺序，`Get/Put` 都先摘除再头插（146）。

### 9. 二叉树（一般）

**递归三要素**：定义函数语义 → 边界（nil）→ 递推（左、右子树结果如何组合）。104、100、226、101 都是直接套用。

- **构造类**：前序 + 中序（或中序 + 后序）构造，用哈希表记录中序下标，递归划分区间（105、106）。
- **展开 / 拼接**：114 用「右子树挂到左子树最右下」或倒序前序指针法。
- **后序「贡献值」**：124 最大路径和——函数返回「单边向下最大贡献」，全局答案在递归中顺带更新，负贡献直接舍弃。
- **最近公共祖先**：后序递归，左右各找到目标则在当前节点汇合（236）。
- **完全二叉树节点数**：有右子树高度等于整树高度则为满树，直接 `2^h - 1`（222）。

### 10. 二叉树 BFS

```go
// 层序遍历模板
queue := []*TreeNode{root}
for len(queue) > 0 {
    level := queue; queue = nil
    for _, node := range level { /* 处理本层，子节点入新 queue */ }
}
```

右视图 = 每层最后一个（199）；锯齿形 = 奇数层倒序或双端切换方向（103）；层平均值直接按层求和（637）。

### 11. 二叉搜索树

**中序遍历 = 递增序列**，这一性质覆盖本类全部题目：

- 最小绝对差：中序时比较相邻差值（530）。
- 第 K 小：中序计数到 K 即停（230）。
- 验证 BST：递归时携带合法区间 `(min, max)`，不能只比较父子节点（98）。

### 12. 图（一般）

- **网格 DFS**：岛屿数量，遍历到 `1` 就计数 + 淹岛（改写成 `0`）（200）；被围绕的区域从边界上的 `O` 反向 DFS，能到达的标记为安全（130）。
- **克隆类**：DFS/BFS + `visited` 哈希表，第一次访问就创建副本（133）。
- **建图 + 路径搜索**：除法求值把 `a/b=k` 建成 `a→b(k)` 和 `b→a(1/k)` 两条边，查询就是求路径权值积（399）。
- **拓扑排序**：课程表 = 检测有向环。BFS 入度法（入度 0 入队）或 DFS 三色标记（207、210）。

### 13. 图 BFS

**求最少步数 / 最短变换 → BFS**，一层一层扩展：

```go
// BFS 模板
queue := []起点; visited := map[起点]bool{}
for step := 1; len(queue) > 0; step++ {
    for 本层每个节点 {
        for 每个邻居 {
            if visited { continue }
            if 是目标 { return step }
            visited[邻居] = true
            queue = append(queue, 邻居)
        }
    }
}
```

蛇梯棋注意格子编号方向（909）；基因变化每次尝试改一位（433）；单词接龙可以预建「通配符中间态」图加速（127）。

### 14. 字典树

```go
type TrieNode struct {
    children [26]*TrieNode
    isEnd    bool
}
```

- 实现 Trie：insert / search / startsWith（208）。
- 通配符搜索：遇到 `.` 对 26 个分支递归（211）。
- 单词搜索 II：所有单词入 Trie，棋盘 DFS 时按 Trie 剪枝，走过的格子标记访问后还原（212）。

### 15. 回溯

```go
// 回溯通用框架
var path []元素
var result [][]元素
func backtrack(选择列表, start int) {
    if 满足结束条件 { result = append(result, 拷贝(path)); return }
    for i := start; i < len(选择列表); i++ {
        path = append(path, 选择列表[i])   // 做选择
        backtrack(选择列表, i + 1)          // 组合/子集传 i+1（不可重复选）
        path = path[:len(path)-1]          // 撤销选择
    }
}
```

- **组合 / 子集**：`start` 保证不重不漏（77）；元素可重复选时传 `i` 而不是 `i+1`（39）。
- **排列**：用 `used` 数组标记已选元素（46）。
- **括号生成**：左括号随时加，右括号数量不超过左括号才加（22）。
- **N 皇后**：按行放，用列 / 主对角线 / 副对角线三个集合判冲突（52）。
- **网格回溯**：单词搜索，DFS 四个方向，标记访问、递归后还原（79）。

### 16. 分治

- 将有序数组转 BST：每次取中点作根，左右半区递归（108）。
- 排序链表：快慢指针找中点断开，归并两个有序链表，递归或自底向上（148）。
- 合并 K 个链表：两两归并（23）。
- 四叉树：区域值全同则建叶子，否则四分递归（427）。

### 17. Kadane 算法

```go
// 最大子数组和：以 i 结尾的最大和 = max(nums[i], 以 i-1 结尾 + nums[i])
cur := nums[0]; ans := nums[0]
for i := 1; i < len(nums); i++ {
    cur = max(nums[i], cur+nums[i])
    ans = max(ans, cur)
}
```

环形子数组：答案 = max(Kadane, 总和 - 最小子数组和)，但要特判全负数情况（918）。

### 18. 二分查找

```go
// lower_bound 模板：第一个 >= target 的位置
lo, hi := 0, len(nums) // 注意 hi 开区间
for lo < hi {
    mid := (lo + hi) / 2
    if nums[mid] >= target { hi = mid } else { lo = mid + 1 }
}
return lo
```

- 搜索插入位置 = lower_bound（35）；首尾位置 = 两次 lower_bound（34）。
- **旋转数组**：二分后必有一半有序，判断 target 是否落在有序半区（33、153）。
- **二分爬坡**：峰值满足 `nums[mid] < nums[mid+1]` 往右走，否则往左（162）。
- 二维矩阵二分：把下标映射成一维（74）；两数组中位数 = 找第 k 小（4）。

### 19. 堆

- **Top K**：维护大小为 K 的小顶堆，堆顶就是第 K 大（215）。
- **双堆中位数**：大顶堆存较小一半，小顶堆存较大一半，保持平衡（295）。
- **多路归并**：K 对最小和 = 每行一个指针入最小堆，弹出后同行后移再入堆（373）。
- **双堆贪心**：IPO 按资本解锁项目入最大堆，每次取最大利润（502）。

### 20. 位运算

- `n & (n-1)`：消去最低位的 1（数 1 的个数，191）。
- 异或：相同为 0，只出现一次的数直接全员异或（136）；出现三次则按位统计 1 的个数再对 3 取模（137）。
- 区间按位与：结果就是公共前缀，`n & (n-1)` 循环收缩右边界（201）。
- 二进制加法 / 位反转：逐位处理或用分治交换（67、190）。

### 21. 数学

- 回文数：只反转一半数字比较，提前防溢出（9）。
- 快速幂：`x^n = (x^2)^(n/2)`，负数指数取倒数，`n` 取最小 int 时先转 int64（50）。
- 阶乘后的零：2 因子远多于 5，数 5 的倍数个数（172）。
- 直线上最多点：以每个点为基准统计斜率（gcd 约分，垂直线单独算），取最大（149）。

### 22. 一维 DP

- **选 / 不选**：打家劫舍 `dp[i] = max(dp[i-1], dp[i-2] + nums[i])`（198）。
- **完全背包最值**：零钱兑换 `dp[x] = min(dp[x-coin] + 1)`（322）。
- **布尔判定**：单词拆分 `dp[i] = 任意 dp[j] && s[j:i] 在字典中`（139）。
- **最值 + 二选一**：爬楼梯（70）；LIS 有 O(n²) 递推与 O(n log n) 贪心二分两种（300）。

### 23. 多维 DP

- **网格路径**：`dp[i][j]` 从左上到 (i,j) 的最值 / 方案数，取上、左两个来源（64、63、120）。
- **字符串编辑**：编辑距离，增 / 删 / 改三种转移（72）；交错字符串判定 s3 前缀由谁拼成（97）。
- **回文**：最长回文子串——中心扩展或 `dp[i][j]` 区间 DP（5）。
- **股票状态机**：最多 k 笔交易用 `buy[k]` / `sell[k]` 两组状态滚动，k 很大时退化为不限次数（122 的贪心）（123、188）。
- **最大正方形**：`dp[i][j]` 表示以 (i,j) 为右下角的正方形边长，取左上三格最小值 + 1（221）。

## 题目清单

<!-- CHECKLIST-START -->
### 1. 数组 / 字符串 (Array / String)（24/24）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 88 | [合并两个有序数组](0088_merge_sorted_array/problem.md) | 简单 |
| [ ] | 27 | [移除元素](0027_remove_element/problem.md) | 简单 |
| [ ] | 26 | [删除有序数组中的重复项](0026_remove_duplicates_from_sorted_array/problem.md) | 简单 |
| [ ] | 80 | [删除有序数组中的重复项 II](0080_remove_duplicates_from_sorted_array_ii/problem.md) | 中等 |
| [ ] | 169 | [多数元素](0169_majority_element/problem.md) | 简单 |
| [ ] | 189 | [轮转数组](0189_rotate_array/problem.md) | 中等 |
| [ ] | 121 | [买卖股票的最佳时机](0121_best_time_to_buy_and_sell_stock/problem.md) | 简单 |
| [ ] | 122 | [买卖股票的最佳时机 II](0122_best_time_to_buy_and_sell_stock_ii/problem.md) | 中等 |
| [ ] | 55 | [跳跃游戏](0055_jump_game/problem.md) | 中等 |
| [ ] | 45 | [跳跃游戏 II](0045_jump_game_ii/problem.md) | 中等 |
| [ ] | 274 | [H 指数](0274_h_index/problem.md) | 中等 |
| [ ] | 380 | [O(1) 时间插入、删除和获取随机元素](0380_insert_delete_getrandom_o1/problem.md) | 中等 |
| [ ] | 238 | [除了自身以外数组的乘积](0238_product_of_array_except_self/problem.md) | 中等 |
| [ ] | 134 | [加油站](0134_gas_station/problem.md) | 中等 |
| [ ] | 135 | [分发糖果](0135_candy/problem.md) | 困难 |
| [ ] | 42 | [接雨水](0042_trapping_rain_water/problem.md) | 困难 |
| [ ] | 13 | [罗马数字转整数](0013_roman_to_integer/problem.md) | 简单 |
| [ ] | 12 | [整数转罗马数字](0012_integer_to_roman/problem.md) | 中等 |
| [ ] | 58 | [最后一个单词的长度](0058_length_of_last_word/problem.md) | 简单 |
| [ ] | 14 | [最长公共前缀](0014_longest_common_prefix/problem.md) | 简单 |
| [ ] | 151 | [反转字符串中的单词](0151_reverse_words_in_a_string/problem.md) | 中等 |
| [ ] | 6 | [Z 字形变换](0006_zigzag_conversion/problem.md) | 中等 |
| [ ] | 28 | [找出字符串中第一个匹配项的下标](0028_find_the_index_of_the_first_occurrence_in_a_string/problem.md) | 简单 |
| [ ] | 68 | [文本左右对齐](0068_text_justification/problem.md) | 困难 |

### 2. 双指针 (Two Pointers)（5/5）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 125 | [验证回文串](0125_valid_palindrome/problem.md) | 简单 |
| [ ] | 392 | [判断子序列](0392_is_subsequence/problem.md) | 简单 |
| [ ] | 167 | [两数之和 II - 输入有序数组](0167_two_sum_ii_input_array_is_sorted/problem.md) | 中等 |
| [ ] | 11 | [盛最多水的容器](0011_container_with_most_water/problem.md) | 中等 |
| [ ] | 15 | [三数之和](0015_3sum/problem.md) | 中等 |

### 3. 滑动窗口 (Sliding Window)（4/4）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 209 | [长度最小的子数组](0209_minimum_size_subarray_sum/problem.md) | 中等 |
| [ ] | 3 | [无重复字符的最长子串](0003_longest_substring_without_repeating_characters/problem.md) | 中等 |
| [ ] | 30 | [串联所有单词的子串](0030_substring_with_concatenation_of_all_words/problem.md) | 困难 |
| [ ] | 76 | [最小覆盖子串](0076_minimum_window_substring/problem.md) | 困难 |

### 4. 矩阵 (Matrix)（5/5）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 36 | [有效的数独](0036_valid_sudoku/problem.md) | 中等 |
| [ ] | 54 | [螺旋矩阵](0054_spiral_matrix/problem.md) | 中等 |
| [ ] | 48 | [旋转图像](0048_rotate_image/problem.md) | 中等 |
| [ ] | 73 | [矩阵置零](0073_set_matrix_zeroes/problem.md) | 中等 |
| [ ] | 289 | [生命游戏](0289_game_of_life/problem.md) | 中等 |

### 5. 哈希表 (Hashmap)（9/9）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 383 | [赎金信](0383_ransom_note/problem.md) | 简单 |
| [ ] | 205 | [同构字符串](0205_isomorphic_strings/problem.md) | 简单 |
| [ ] | 290 | [单词规律](0290_word_pattern/problem.md) | 简单 |
| [ ] | 242 | [有效的字母异位词](0242_valid_anagram/problem.md) | 简单 |
| [ ] | 49 | [字母异位词分组](0049_group_anagrams/problem.md) | 中等 |
| [ ] | 1 | [两数之和](0001_two_sum/problem.md) | 简单 |
| [ ] | 202 | [快乐数](0202_happy_number/problem.md) | 简单 |
| [ ] | 219 | [存在重复元素 II](0219_contains_duplicate_ii/problem.md) | 简单 |
| [ ] | 128 | [最长连续序列](0128_longest_consecutive_sequence/problem.md) | 中等 |

### 6. 区间 (Intervals)（4/4）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 228 | [汇总区间](0228_summary_ranges/problem.md) | 简单 |
| [ ] | 56 | [合并区间](0056_merge_intervals/problem.md) | 中等 |
| [ ] | 57 | [插入区间](0057_insert_interval/problem.md) | 中等 |
| [ ] | 452 | [用最少数量的箭引爆气球](0452_minimum_number_of_arrows_to_burst_balloons/problem.md) | 中等 |

### 7. 栈 (Stack)（5/5）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 20 | [有效的括号](0020_valid_parentheses/problem.md) | 简单 |
| [ ] | 71 | [简化路径](0071_simplify_path/problem.md) | 中等 |
| [ ] | 155 | [最小栈](0155_min_stack/problem.md) | 中等 |
| [ ] | 150 | [逆波兰表达式求值](0150_evaluate_reverse_polish_notation/problem.md) | 中等 |
| [ ] | 224 | [基本计算器](0224_basic_calculator/problem.md) | 困难 |

### 8. 链表 (Linked List)（11/11）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 141 | [环形链表](0141_linked_list_cycle/problem.md) | 简单 |
| [ ] | 2 | [两数相加](0002_add_two_numbers/problem.md) | 中等 |
| [ ] | 21 | [合并两个有序链表](0021_merge_two_sorted_lists/problem.md) | 简单 |
| [ ] | 138 | [随机链表的复制](0138_copy_list_with_random_pointer/problem.md) | 中等 |
| [ ] | 92 | [反转链表 II](0092_reverse_linked_list_ii/problem.md) | 中等 |
| [ ] | 25 | [K 个一组翻转链表](0025_reverse_nodes_in_k_group/problem.md) | 困难 |
| [ ] | 19 | [删除链表的倒数第 N 个结点](0019_remove_nth_node_from_end_of_list/problem.md) | 中等 |
| [ ] | 82 | [删除排序链表中的重复元素 II](0082_remove_duplicates_from_sorted_list_ii/problem.md) | 中等 |
| [ ] | 61 | [旋转链表](0061_rotate_list/problem.md) | 中等 |
| [ ] | 86 | [分隔链表](0086_partition_list/problem.md) | 中等 |
| [ ] | 146 | [LRU 缓存](0146_lru_cache/problem.md) | 中等 |

### 9. 二叉树（一般）(Binary Tree General)（14/14）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 104 | [二叉树的最大深度](0104_maximum_depth_of_binary_tree/problem.md) | 简单 |
| [ ] | 100 | [相同的树](0100_same_tree/problem.md) | 简单 |
| [ ] | 226 | [翻转二叉树](0226_invert_binary_tree/problem.md) | 简单 |
| [ ] | 101 | [对称二叉树](0101_symmetric_tree/problem.md) | 简单 |
| [ ] | 105 | [从前序与中序遍历序列构造二叉树](0105_construct_binary_tree_from_preorder_and_inorder_traversal/problem.md) | 中等 |
| [ ] | 106 | [从中序与后序遍历序列构造二叉树](0106_construct_binary_tree_from_inorder_and_postorder_traversal/problem.md) | 中等 |
| [ ] | 117 | [填充每个节点的下一个右侧节点指针 II](0117_populating_next_right_pointers_in_each_node_ii/problem.md) | 中等 |
| [ ] | 114 | [二叉树展开为链表](0114_flatten_binary_tree_to_linked_list/problem.md) | 中等 |
| [ ] | 112 | [路径总和](0112_path_sum/problem.md) | 简单 |
| [ ] | 129 | [求根节点到叶节点数字之和](0129_sum_root_to_leaf_numbers/problem.md) | 中等 |
| [ ] | 124 | [二叉树中的最大路径和](0124_binary_tree_maximum_path_sum/problem.md) | 困难 |
| [ ] | 173 | [二叉搜索树迭代器](0173_binary_search_tree_iterator/problem.md) | 中等 |
| [ ] | 222 | [完全二叉树的节点个数](0222_count_complete_tree_nodes/problem.md) | 中等 |
| [ ] | 236 | [二叉树的最近公共祖先](0236_lowest_common_ancestor_of_a_binary_tree/problem.md) | 中等 |

### 10. 二叉树 BFS (Binary Tree BFS)（4/4）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 199 | [二叉树的右视图](0199_binary_tree_right_side_view/problem.md) | 中等 |
| [ ] | 637 | [二叉树的层平均值](0637_average_of_levels_in_binary_tree/problem.md) | 简单 |
| [ ] | 102 | [二叉树的层序遍历](0102_binary_tree_level_order_traversal/problem.md) | 中等 |
| [ ] | 103 | [二叉树的锯齿形层序遍历](0103_binary_tree_zigzag_level_order_traversal/problem.md) | 中等 |

### 11. 二叉搜索树 (Binary Search Tree)（3/3）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 530 | [二叉搜索树的最小绝对差](0530_minimum_absolute_difference_in_bst/problem.md) | 简单 |
| [ ] | 230 | [二叉搜索树中第 K 小的元素](0230_kth_smallest_element_in_a_bst/problem.md) | 中等 |
| [ ] | 98 | [验证二叉搜索树](0098_validate_binary_search_tree/problem.md) | 中等 |

### 12. 图（一般）(Graph General)（6/6）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 200 | [岛屿数量](0200_number_of_islands/problem.md) | 中等 |
| [ ] | 130 | [被围绕的区域](0130_surrounded_regions/problem.md) | 中等 |
| [ ] | 133 | [克隆图](0133_clone_graph/problem.md) | 中等 |
| [ ] | 399 | [除法求值](0399_evaluate_division/problem.md) | 中等 |
| [ ] | 207 | [课程表](0207_course_schedule/problem.md) | 中等 |
| [ ] | 210 | [课程表 II](0210_course_schedule_ii/problem.md) | 中等 |

### 13. 图 BFS (Graph BFS)（3/3）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 909 | [蛇梯棋](0909_snakes_and_ladders/problem.md) | 中等 |
| [ ] | 433 | [最小基因变化](0433_minimum_genetic_mutation/problem.md) | 中等 |
| [ ] | 127 | [单词接龙](0127_word_ladder/problem.md) | 困难 |

### 14. 字典树 (Trie)（3/3）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 208 | [实现 Trie (前缀树)](0208_implement_trie_prefix_tree/problem.md) | 中等 |
| [ ] | 211 | [添加与搜索单词 - 数据结构设计](0211_design_add_and_search_words_data_structure/problem.md) | 中等 |
| [ ] | 212 | [单词搜索 II](0212_word_search_ii/problem.md) | 困难 |

### 15. 回溯 (Backtracking)（7/7）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 17 | [电话号码的字母组合](0017_letter_combinations_of_a_phone_number/problem.md) | 中等 |
| [ ] | 77 | [组合](0077_combinations/problem.md) | 中等 |
| [ ] | 46 | [全排列](0046_permutations/problem.md) | 中等 |
| [ ] | 39 | [组合总和](0039_combination_sum/problem.md) | 中等 |
| [ ] | 52 | [N 皇后 II](0052_n_queens_ii/problem.md) | 困难 |
| [ ] | 22 | [括号生成](0022_generate_parentheses/problem.md) | 中等 |
| [ ] | 79 | [单词搜索](0079_word_search/problem.md) | 中等 |

### 16. 分治 (Divide & Conquer)（4/4）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 108 | [将有序数组转换为二叉搜索树](0108_convert_sorted_array_to_binary_search_tree/problem.md) | 简单 |
| [ ] | 148 | [排序链表](0148_sort_list/problem.md) | 中等 |
| [ ] | 427 | [建立四叉树](0427_construct_quad_tree/problem.md) | 中等 |
| [ ] | 23 | [合并 K 个升序链表](0023_merge_k_sorted_lists/problem.md) | 困难 |

### 17. Kadane 算法 (Kadane's Algorithm)（2/2）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 53 | [最大子数组和](0053_maximum_subarray/problem.md) | 中等 |
| [ ] | 918 | [环形子数组的最大和](0918_maximum_sum_circular_subarray/problem.md) | 中等 |

### 18. 二分查找 (Binary Search)（7/7）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 35 | [搜索插入位置](0035_search_insert_position/problem.md) | 简单 |
| [ ] | 74 | [搜索二维矩阵](0074_search_a_2d_matrix/problem.md) | 中等 |
| [ ] | 162 | [寻找峰值](0162_find_peak_element/problem.md) | 中等 |
| [ ] | 33 | [搜索旋转排序数组](0033_search_in_rotated_sorted_array/problem.md) | 中等 |
| [ ] | 34 | [在排序数组中查找元素的第一个和最后一个位置](0034_find_first_and_last_position_of_element_in_sorted_array/problem.md) | 中等 |
| [ ] | 153 | [寻找旋转排序数组中的最小值](0153_find_minimum_in_rotated_sorted_array/problem.md) | 中等 |
| [ ] | 4 | [寻找两个正序数组的中位数](0004_median_of_two_sorted_arrays/problem.md) | 困难 |

### 19. 堆 (Heap)（4/4）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 215 | [数组中的第K个最大元素](0215_kth_largest_element_in_an_array/problem.md) | 中等 |
| [ ] | 502 | [IPO](0502_ipo/problem.md) | 困难 |
| [ ] | 373 | [查找和最小的 K 对数字](0373_find_k_pairs_with_smallest_sums/problem.md) | 中等 |
| [ ] | 295 | [数据流的中位数](0295_find_median_from_data_stream/problem.md) | 困难 |

### 20. 位运算 (Bit Manipulation)（6/6）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 67 | [二进制求和](0067_add_binary/problem.md) | 简单 |
| [ ] | 190 | [颠倒二进制位](0190_reverse_bits/problem.md) | 简单 |
| [ ] | 191 | [位1的个数](0191_number_of_1_bits/problem.md) | 简单 |
| [ ] | 136 | [只出现一次的数字](0136_single_number/problem.md) | 简单 |
| [ ] | 137 | [只出现一次的数字 II](0137_single_number_ii/problem.md) | 中等 |
| [ ] | 201 | [数字范围按位与](0201_bitwise_and_of_numbers_range/problem.md) | 中等 |

### 21. 数学 (Math)（6/6）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 9 | [回文数](0009_palindrome_number/problem.md) | 简单 |
| [ ] | 66 | [加一](0066_plus_one/problem.md) | 简单 |
| [ ] | 172 | [阶乘后的零](0172_factorial_trailing_zeroes/problem.md) | 中等 |
| [ ] | 69 | [x 的平方根](0069_sqrtx/problem.md) | 简单 |
| [ ] | 50 | [Pow(x, n)](0050_powx_n/problem.md) | 中等 |
| [ ] | 149 | [直线上最多的点数](0149_max_points_on_a_line/problem.md) | 困难 |

### 22. 一维 DP (1D DP)（5/5）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 70 | [爬楼梯](0070_climbing_stairs/problem.md) | 简单 |
| [ ] | 198 | [打家劫舍](0198_house_robber/problem.md) | 中等 |
| [ ] | 139 | [单词拆分](0139_word_break/problem.md) | 中等 |
| [ ] | 322 | [零钱兑换](0322_coin_change/problem.md) | 中等 |
| [ ] | 300 | [最长递增子序列](0300_longest_increasing_subsequence/problem.md) | 中等 |

### 23. 多维 DP (Multidimensional DP)（9/9）

| 状态 | 题号 | 题目 | 难度 |
| --- | --- | --- | --- |
| [ ] | 120 | [三角形最小路径和](0120_triangle/problem.md) | 中等 |
| [ ] | 64 | [最小路径和](0064_minimum_path_sum/problem.md) | 中等 |
| [ ] | 63 | [不同路径 II](0063_unique_paths_ii/problem.md) | 中等 |
| [ ] | 5 | [最长回文子串](0005_longest_palindromic_substring/problem.md) | 中等 |
| [ ] | 97 | [交错字符串](0097_interleaving_string/problem.md) | 中等 |
| [ ] | 72 | [编辑距离](0072_edit_distance/problem.md) | 中等 |
| [ ] | 123 | [买卖股票的最佳时机 III](0123_best_time_to_buy_and_sell_stock_iii/problem.md) | 困难 |
| [ ] | 188 | [买卖股票的最佳时机 IV](0188_best_time_to_buy_and_sell_stock_iv/problem.md) | 困难 |
| [ ] | 221 | [最大正方形](0221_maximal_square/problem.md) | 中等 |

<!-- CHECKLIST-END -->
