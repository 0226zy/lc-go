# 383. 赎金信 (Ransom Note)

## 题目描述

给你两个字符串：`ransomNote` 和 `magazine`，判断 `ransomNote` 能不能由 `magazine` 里面的字符构成。

如果可以，返回 `true`；否则返回 `false`。

`magazine` 中的每个字符只能在 `ransomNote` 中使用一次。

**注意：** 你可以认为两个字符串均只包含小写字母。

### 示例 1

```
输入: ransomNote = "a", magazine = "b"
输出: false
```

### 示例 2

```
输入: ransomNote = "aa", magazine = "ab"
输出: false
```

### 示例 3

```
输入: ransomNote = "aa", magazine = "aab"
输出: true
```

### 提示

- `1 <= ransomNote.length, magazine.length <= 10^5`
- `ransomNote` 和 `magazine` 由小写英文字母组成

## 题目解析

### 核心思路

这是一道典型的 **哈希计数** 题，和 242「有效的字母异位词」是同一类模型：问“一个字符串的字符能不能凑出另一个字符串”。

怎么想？把 `magazine` 看成你的“材料仓库”，先盘点仓库里每个字母有多少库存；再拿着 `ransomNote` 逐字去仓库里取货。取某个字母时发现库存为 0（取成负数），就说明材料不够，直接返回 `false`；全部取完都没有透支，返回 `true`。

为什么这个解法是对的？因为题目只要求 `ransomNote` 的每个字符都能在 `magazine` 中找到**足够的数量**，而对字符的顺序、位置没有任何要求。因此问题完全等价于：`magazine` 中每种字母的计数 ≥ `ransomNote` 中对应字母的计数。用计数数组比较即可。

由于字符集只有 26 个小写字母，不需要真的用哈希表，一个长度为 26 的数组就够了（本质上仍是哈希的思想）。

### 算法步骤

1. 创建计数数组 `count[26]`。
2. 遍历 `magazine`，对其中每个字符 `count[ch-'a']++`，统计库存。
3. 遍历 `ransomNote`，对每个字符 `count[ch-'a']--`。
4. 一旦某个计数变成负数，说明该字符不够用，返回 `false`。
5. 遍历结束没有返回 `false`，返回 `true`。

### 复杂度分析

- **时间复杂度**: O(m + n)，m、n 分别为两个字符串的长度，各遍历一次
- **空间复杂度**: O(1)，固定 26 个计数位置

## 代码实现

```go
func CanConstruct(ransomNote string, magazine string) bool {
    count := [26]int{}
    for i := 0; i < len(magazine); i++ {
        count[magazine[i]-'a']++ // 统计 magazine 中每个字母的库存
    }
    for i := 0; i < len(ransomNote); i++ {
        count[ransomNote[i]-'a']-- // 消耗 ransomNote 中的字符
        if count[ransomNote[i]-'a'] < 0 {
            return false // 库存透支，材料不够
        }
    }
    return true
}
```

**执行过程示例**（`ransomNote = "aa"`, `magazine = "ab"`）：

```
统计 magazine:  a:1, b:1
消耗 'a' (第1个):  a:0
消耗 'a' (第2个):  a:-1 → 负数，返回 false
```
