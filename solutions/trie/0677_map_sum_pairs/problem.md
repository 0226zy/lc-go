# 677. 键值映射 (Map Sum Pairs)

## 题目描述
设计一个 map，完成以下功能：
- `insert(key, val)`：如果 key 不存在，插入键值对；如果 key 已存在，更新 key 对应的值。
- `sum(prefix)`：返回所有以 prefix 为前缀的键的值的总和。

## 示例
```
输入：
["MapSum", "insert", "sum", "insert", "sum"]
[[], ["apple",3], ["ap"], ["app",2], ["ap"]]
输出：[null, null, 3, null, 5]
解释：
MapSum mapSum = new MapSum();
mapSum.insert("apple", 3);
mapSum.sum("ap");           // 返回 3 (apple = 3)
mapSum.insert("app", 2);
mapSum.sum("ap");           // 返回 5 (apple + app = 3 + 2 = 5)
```

## 约束
- key 和 prefix 仅由小写英文字母组成
- 1 <= key.length, prefix.length <= 50
- 0 <= val <= 1000
- 最多调用 50 次 insert 和 sum

## 分类
前缀树 / 哈希表
