# 46 题当前实现分析

## 当前实现

```go
func Permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	n := len(nums)
	visited := make(map[int]bool, n)
	for _, v := range nums {
		visited[v] = false
	}
	ret := [][]int{}
	var backtrack func(int, *[]int)
	backtrack = func(idx int, path *[]int) {
		if idx == n {
			tmp := make([]int, n)
			copy(tmp, *path)
			ret = append(ret, tmp)
			return
		}
		for _, v := range nums {
			if visited[v] {
				continue
			}
			*path = append(*path, v)
			visited[v] = true
			backtrack(idx+1, path)
			visited[v] = false
			*path = (*path)[:len(*path)-1]
		}
	}
	path := []int{}
	backtrack(0, &path)
	return ret
}
```

## 运行测试结果

执行 `go test` 后**全部通过**。

```
PASS
ok  	github.com/0226zy/lc-go/solutions/hot100/0046_permutations
```

---

## 结论：功能正确，无逻辑 Bug

当前实现使用经典的**回溯法（深度优先搜索）**生成全排列，核心逻辑正确：

1. 使用 `visited` 记录每个数字是否已被当前路径使用
2. 递归到 `idx == n` 时，复制当前路径并加入结果
3. 回溯时恢复 `visited` 和 `path` 状态

测试结果也验证了所有示例、边界和去重场景均正确。

---

## 可优化点

虽然功能正确，但代码在写法上有几个可以改进的地方：

### 1. `visited` 使用 map，可替换为 bool 切片

当前代码：

```go
visited := make(map[int]bool, n)
for _, v := range nums {
    visited[v] = false
}
```

问题：
- map 的查询/写入时间复杂度虽然是 O(1)，但常数较大
- `visited[v] = false` 的初始化是**冗余**的，因为 map 中不存在的键默认就是 `false`
- 题目约束 `nums.length <= 6`，使用 map 有些"杀鸡用牛刀"

建议改用索引访问的 bool 切片：

```go
visited := make([]bool, n)
```

循环中使用索引 `i` 而不是值 `v`：

```go
for i, v := range nums {
    if visited[i] {
        continue
    }
    *path = append(*path, v)
    visited[i] = true
    backtrack(idx+1, path)
    visited[i] = false
    *path = (*path)[:len(*path)-1]
}
```

### 2. `*[]int` 指针传递可简化为切片值传递

当前代码：

```go
var backtrack func(int, *[]int)
backtrack = func(idx int, path *[]int) {
    *path = append(*path, v)
    backtrack(idx+1, path)
    *path = (*path)[:len(*path)-1]
}
path := []int{}
backtrack(0, &path)
```

Go 的切片本身就是引用类型，函数内修改切片的底层数组会影响外部。使用 `*[]int` 指针虽然更安全（能修改切片头），但写起来冗余。可以简化为：

```go
var backtrack func(int, []int)
backtrack = func(idx int, path []int) {
    if idx == n {
        tmp := make([]int, n)
        copy(tmp, path)
        ret = append(ret, tmp)
        return
    }
    for i, v := range nums {
        if visited[i] {
            continue
        }
        backtrack(idx+1, append(path, v))
    }
}
backtrack(0, []int{})
```

这种方式无需手动回溯 `path`，但每次递归都会创建新切片，空间开销略大。两种方式都可以接受。

### 3. 复杂度注释需要更新

当前注释仍是 TODO 状态：

```go
// 时间复杂度: O(?)  空间复杂度: O(?)
```

应补充为：

```go
// 时间复杂度: O(n * n!)  空间复杂度: O(n)
```

- 时间复杂度：共有 n! 个排列，每个排列需要 O(n) 时间复制，所以是 O(n * n!)
- 空间复杂度：递归栈深度为 O(n)，path 长度也为 O(n)

### 4. 空输入处理可更一致

当前代码：

```go
if len(nums) == 0 {
    return nil
}
```

虽然题目约束 `nums.length >= 1`，但如果从健壮性角度，返回空二维切片 `[][]int{}` 可能比 `nil` 更一致，因为调用者可能期望一个可 `range` 的值。不过这是一个风格问题。

---

## 优化后的实现参考

```go
func Permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	n := len(nums)
	visited := make([]bool, n)
	var ret [][]int

	var backtrack func(idx int, path []int)
	backtrack = func(idx int, path []int) {
		if idx == n {
			tmp := make([]int, n)
			copy(tmp, path)
			ret = append(ret, tmp)
			return
		}
		for i, v := range nums {
			if visited[i] {
				continue
			}
			visited[i] = true
			backtrack(idx+1, append(path, v))
			visited[i] = false
		}
	}

	backtrack(0, []int{})
	return ret
}
```

---

## 总结

| 维度 | 评价 |
|------|------|
| 功能正确性 | ✅ 正确，测试全部通过 |
| 算法选择 | ✅ 回溯法适合此题 |
| 代码风格 | ⚠️ 指针传递略显冗余，map 可以用切片替代 |
| 性能 | ⚠️ map 常数较大，但在 n <= 6 时无实际影响 |
| 复杂度注释 | ❌ 仍是 TODO，需要补充 |
| 空输入处理 | ⚠️ 返回 nil，可改为空切片 |

**核心结论：无 Bug，主要是代码简化和注释完善的空间。**
