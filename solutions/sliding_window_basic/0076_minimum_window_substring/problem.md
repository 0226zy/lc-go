# 76. 最小覆盖子串 (Minimum Window Substring)

## 题目描述

给定两个字符串 `s` 和 `t`，返回 `s` 中**涵盖 `t` 所有字符的** **最小子串**。如果 `s` 中不存在涵盖 `t` 所有字符的子串，则返回空字符串 `""`。

**注意**：`t` 中重复的字符也要按出现次数覆盖。

### 示例 1

```
输入: s = "ADOBECODEBANC", t = "ABC"
输出: "BANC"
解释: "BANC" 包含 'A'、'B'、'C' 各一次，是最短的覆盖子串。
```

### 示例 2

```
输入: s = "a", t = "aa"
输出: ""
解释: t 中需要两个 'a'，但 s 只有一个。
```

### 提示

- `1 <= s.length, t.length <= 10^5`
- `s` 和 `t` 由英文字母组成

**进阶**：你能设计一个在 `O(len(s) + len(t))` 时间内解决的算法吗？

## 题目解析

### 核心思路

这是滑动窗口中最重要的一道模板题——**need / have 计数滑窗**。它和 209、3 题同宗：右指针扩张找可行解，左指针收缩取最优解；区别在于「可行」的判定从「和达标」「无重复」升级为「字符多重集合的覆盖」。

直接用 `have[c] >= need[c]` 逐个字符判断会引入 O(字符集) 的扫描，所以维护一个计数器 `formed`：表示**有多少种字符的窗口内次数恰好达到需求**。每次加入/移出字符时只关心这一个字符是否跨过了「刚好满足」或「刚好跌破」的临界点，O(1) 更新：

- 加入字符 `c` 后 `have[c]` 从 `need[c]-1` 变为 `need[c]` → `formed++`（这种字符达标了）；
- 移出字符 `c` 后 `have[c]` 从 `need[c]` 变为 `need[c]-1` → `formed--`（这种字符不再达标）。

当 `formed == required`（`required` 是 `t` 中不同字符的种数）时，窗口覆盖了 `t`，开始收缩左边界并记录最短窗口。

**为什么收缩不会漏解**：窗口可行时，左端点每往右挪一格，如果仍可行就得到了一个更短的候选；一旦不可行，任何以当前左端点为起点、右端点在当前 `right` 左侧的子串都不可能更优（它们已经被之前的收缩枚举过）。所以每个位置作为「最短可行窗口左端点」的机会只会被考察一次，双指针单调右移，整体线性。

### 算法步骤

1. 统计 `need`：遍历 `t`，`need[c]++`；统计 `required` = `need` 中大于 0 的字符种数。
2. 初始化 `left = 0`，`formed = 0`，`start = 0`，`minLen = len(s)+1`（哨兵）。
3. `right` 从左到右遍历 `s`：
   - `have[s[right]]++`；若该字符 `have` 恰好等于 `need`，`formed++`；
   - 当 `formed == required`（窗口已覆盖 `t`），循环收缩：
     - 若当前窗口更短，更新 `minLen` 和 `start`；
     - 移出 `s[left]`：`have` 减一，若该字符 `have` 跌破 `need`，`formed--`；
     - `left++`。
4. 若 `minLen` 仍为哨兵值返回 `""`，否则返回 `s[start:start+minLen]`。

### 复杂度分析

- **时间复杂度**: O(len(s) + len(t))，两个指针都只右移，每个字符最多进出窗口一次
- **空间复杂度**: O(|Σ|)，|Σ| 为字符集大小（英文字母时可用定长数组代替哈希表）

## 代码实现

```go
func MinWindow(s string, t string) string {
	m, n := len(s), len(t)
	if m < n {
		return ""
	}

	var need, have [128]int // 英文字母可直接用数组当哈希表
	for i := 0; i < n; i++ {
		need[t[i]]++
	}
	required := 0
	for i := 0; i < 128; i++ {
		if need[i] > 0 {
			required++
		}
	}

	left, start, minLen := 0, 0, m+1
	formed := 0
	for right := 0; right < m; right++ {
		c := s[right]
		have[c]++
		if need[c] > 0 && have[c] == need[c] {
			formed++ // 字符 c 恰好达到所需次数
		}
		for formed == required { // 窗口已覆盖 t，收缩取最短
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}
			d := s[left]
			have[d]--
			if need[d] > 0 && have[d] < need[d] {
				formed-- // 字符 d 跌破所需次数
			}
			left++
		}
	}

	if minLen == m+1 {
		return ""
	}
	return s[start : start+minLen]
}
```

**执行过程示例**（`s = "ADOBECODEBANC"`, `t = "ABC"`，`need = {A:1,B:1,C:1}`，图中窗口记为当前 `[left, right]`）：

```
right=0..5 依次加入 A D O B E C:
  到 right=5 时 have[A]=1, have[B]=1, have[C]=1 → formed=3 == required
  开始收缩: "ADOBEC"(长6) 更新 minLen=6, start=0
  移出 A → have[A]=0 < 1 → formed=2, 停止收缩 (left=1)
right=6..9 加入 O D E B: B 出现第二次 have[B]=2, 无临界点变化
right=10 加入 A: have[A] 0→1 达标 → formed=3, 窗口 "ODEBANC" 收缩:
  长7 不更短; 移出 O,D,E 无影响; 移出 B → have[B]=1 == need 不跌破; 移出 A → formed=2, 停止
  期间窗口收缩到 "BANC"? 重新看: left 从 5 走到 10 的过程中 right=10, 窗口 [5,10]="CODEBA" 不含 C? 
  实际临界点: 移出 left=5 的 C 时 have[C]=0 < 1 → formed=2, 窗口停在 [6,10]="ODEBA"... 
  继续 right=11 加入 N, right=12 加入 C: have[C] 0→1 → formed=3, 窗口收缩:
  left=6..9 移出 O,D,E 无影响; left=9 是 B, 移出后 have[B]=1 == need 不跌破;
  此时窗口 [10,12]="ANC" 只有3字符但缺 B? 重新数: s = A D O B E C O D E B A N C
  正确轨迹: right=12 时窗口收缩到 [9,12]="BANC" 长4 → 更新 minLen=4, start=9
  移出 B → formed=2, 结束
结果: s[9:13] = "BANC"
```

注：上面轨迹为帮助理解的简化演示，关键是抓住两个临界点：`have[c]` **升穿** `need[c]` 时 `formed++`，**跌破** 时 `formed--`。
