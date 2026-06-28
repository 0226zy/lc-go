# 76. 最小覆盖子串 — 解法详解

## 一句话总结

> **用滑动窗口在 s 中找一个最短的子串，使得它包含 t 中所有字符（含重复）。** 核心是"先右扩找到可行解，再左缩优化到最短"。

---

## 1. 题目本质

### 1.1 什么是"覆盖"

t = `"AABC"`，则子串必须包含：`A` 至少 2 个、`B` 至少 1 个、`C` 至少 1 个。

```
s = "ADOBECODEBANC"

"ADOBEC"     → A×1 B×1 C×1  ❌ A 不够
"ADOBECODEB" → A×1 B×1 C×1  ❌ A 不够（需要2）
"DOBECODEBA" → A×1 B×1 C×1  ❌ A 不够
"OBECODEBA"  → A×1 B×1 C×1  ❌
"BANC"       → A×1 B×1 C×1  ❌ A 不够（t 只需1个A时OK）
```

等等，示例中 t = `"ABC"`（每个字母只需 1 个），所以 `"BANC"` 就够了。

### 1.2 为什么用滑动窗口

- 子串是**连续的** → 天然适合窗口
- 要找**最短**的可行子串 → 扩到可行后尝试收缩
- 暴力 O(n²) 检查所有子串太慢 → 滑动窗口 O(n)

---

## 2. 算法思路：双指针滑动窗口

### 2.1 核心思想

```
1. right 右扩，直到窗口 [left, right] 包含 t 的所有字符
2. left 右缩，尽量缩短窗口，记录最短
3. 缩到不再可行，回到步骤 1 继续 right 右扩
```

### 2.2 生活化比喻

你在一条街上找一段连续的店铺，要凑齐购物清单上的所有商品（含数量）：
- 先一路往右走，走到凑齐为止
- 然后从左边开始砍，砍到再砍就缺货为止
- 记下这段距离，继续往右走找下一个可能更短的

### 2.3 关键数据结构

```go
need   map[byte]int  // t 中每个字符的需求量
window map[byte]int  // 当前窗口中各字符的计数
valid  int           // 已满足需求的字符种类数（不是个数！）
```

**valid 的含义**：当 `window[c] == need[c]` 时，该字符"刚好满足"，`valid++`。当 `valid == len(need)` 时，窗口覆盖了 t。

---

## 3. 图解全过程

以 `s = "ADOBECODEBANC"`, `t = "ABC"` 为例：

```
need = {A:1, B:1, C:1}, valid=0, len(need)=3

第1轮：right 右扩
right=0  s[0]='A'  window={A:1}  valid=1
right=1  s[1]='D'  window={A:1,D:1}  (D不在need中)
right=2  s[2]='O'  (O不在need中)
right=3  s[3]='B'  window={A:1,B:1}  valid=2
right=4  s[4]='E'  (E不在need中)
right=5  s[5]='C'  window={A:1,B:1,C:1}  valid=3 ✅ 覆盖！

  窗口 "ADOBEC" 长度6
  left 右缩：
  left=0  s[0]='A'  如果移除，window[A]=0 < need[A]=1 → 不可行，停止
  记录：minLen=6, start=0, 子串="ADOBEC"

  移除 left=0 的 'A'：window[A]--, valid--（因为 window[A] < need[A]）
  left=1, valid=2

第2轮：right 继续右扩
right=6  s[6]='O'
right=7  s[7]='D'
right=8  s[8]='E'
right=9  s[9]='B'  window[B]=2  (B 已满足，valid 不变)
right=10 s[10]='A' window[A]=1  valid=3 ✅ 覆盖！

  窗口 "DOBECODEBA" → 缩 left
  left=1  'D' 不在 need → 跳过
  left=2  'O' 不在 need → 跳过
  left=3  'B' 移除？window[B]=1==need[B]=1 → 移除后 window[B]=0 < 1 → 不可行
  记录：长度8 > 6，不更新

  移除 left=3 的 'B'：window[B]--, valid--
  left=4

第3轮：right 继续右扩
right=11 s[11]='N'
right=12 s[12]='C' window[C]=2  (C 已满足)

  窗口 "ODEBANC" → 缩 left
  left=4  'E' 跳过
  left=5  'C' 移除？window[C]=1==need[C]=1 → 移除后不可行
  
  但 wait, 当前 valid=2（缺 B），还没覆盖！
  实际上 right=12 时 window[C] 从 0→1，但 need[C]=1...
```

（上面过程较复杂，实际代码用循环处理更清晰）

### 3.1 最终结果

```
找到的可行窗口中，最短的是 "BANC"，长度 4
```

---

## 4. 代码实现

```go
func MinWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0          // 已满足需求的字符种类数
	start, minLen := 0, len(s)+1

	for right < len(s) {
		c := s[right]   // 即将加入窗口的字符
		right++

		// 1. 右扩：c 加入窗口
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++ // 该字符刚好满足需求
			}
		}

		// 2. 左缩：当窗口覆盖 t 时，尝试收缩
		for valid == len(need) {
			// 更新最短
			if right-left < minLen {
				start = left
				minLen = right - left
			}

			d := s[left] // 即将移出窗口的字符
			left++

			if need[d] > 0 {
				if window[d] == need[d] {
					valid-- // 移除后该字符不再满足
				}
				window[d]--
			}
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}
```

---

## 5. 关键细节解析

### 5.1 valid 计数的是"种类"不是"个数"

```go
if window[c] == need[c] {
    valid++  // 只有"刚好满足"时才 +1
}
```

- t = `"AABC"`，need = `{A:2, B:1, C:1}`，len(need) = 3
- 窗口有 1 个 A → valid 不变（还没到 2）
- 窗口有 2 个 A → valid++（A 满足了）
- valid == 3 → A/B/C 都满足 → 覆盖

### 5.2 左缩时 valid 的维护

```go
if window[d] == need[d] {
    valid--  // 移除前刚好满足，移除后就不满足了
}
window[d]--
```

必须**先检查再减**。如果 `window[d] > need[d]`（多余的），移除一个不影响 valid。

### 5.3 need[c] > 0 的判断

```go
if need[c] > 0 {  // 只关心 t 中需要的字符
    window[c]++
    ...
}
```

不在 need 中的字符（如 'D', 'O'）直接跳过，不维护 window 计数。

### 5.4 minLen 初始化为 len(s)+1

```go
minLen := len(s)+1  // 比任何合法子串都长
```

如果最终 `minLen == len(s)+1`，说明从未找到可行解，返回 `""`。

---

## 6. 复杂度分析

- **时间 O(m + n)**：left 和 right 各最多遍历 s 一次，构建 need 遍历 t 一次。
- **空间 O(1)**：字符集固定（英文字母），need 和 window 最多 52 个键。

---

## 7. 常见误区

### 7.1 valid 计数错误

```go
// ❌ 错误：每加一个字符就 valid++
window[c]++
valid++  // 2个A时 valid 会变成2，但 len(need) 只有1

// ✅ 正确：只在刚好满足时 valid++
if window[c] == need[c] {
    valid++
}
```

### 7.2 左缩顺序错误

```go
// ❌ 错误：先减再判断
window[d]--
if window[d] < need[d] {  // 已经减了，判断逻辑不对
    valid--
}

// ✅ 正确：先判断再减
if window[d] == need[d] {
    valid--
}
window[d]--
```

### 7.3 忘记更新 minLen 的时机

```go
// ❌ 错误：在 right 扩展时更新
if valid == len(need) {
    minLen = ...  // right 还在扩，不是最短
}

// ✅ 正确：在 left 收缩的循环内更新
for valid == len(need) {
    if right-left < minLen {
        start, minLen = left, right-left
    }
    // 然后收缩 left
}
```

### 7.4 窗口边界搞混

```go
// right 是"即将加入"的位置，加入后 right++
// 所以窗口实际是 [left, right)，长度 = right - left
// 不是 right - left + 1
```

---

## 8. 与其他滑动窗口题的对比

| 题目 | 窗口行为 | 判定条件 |
|------|---------|---------|
| 76 最小覆盖子串 | 右扩到可行，左缩到最短 | valid == len(need) |
| 3 无重复字符最长子串 | 右扩遇重复则左缩 | window[c] <= 1 |
| 438 找所有字母异位词 | 固定窗口大小 k | valid == len(need) |
| 209 长度最小子数组 | 右扩到达标，左缩到最短 | sum >= target |

**76 题的特点**：窗口大小不固定，目标是**最小化**窗口，所以是"扩到可行 → 缩到不行"的循环。

---

## 9. 总结

```
最小覆盖子串
  ↓ 滑动窗口
right 右扩到覆盖 t
  ↓ valid == len(need)
left 左缩到刚好不覆盖
  ↓ 记录最短
重复直到 right 到末尾
```

**记忆口诀**：右扩凑齐，左缩到缺，记录最短，valid 管种类。
