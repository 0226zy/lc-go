# 105 题当前实现问题分析

## 当前实现

```go
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	n := len(inorder)
	rootIdx := n / 2
	leftNum := rootIdx - 1
	root.Left = BuildTree(preorder[1:leftNum+1], inorder[0:rootIdx-1])
	root.Right = BuildTree(preorder[leftNum+1:], inorder[rootIdx+1:])
	return root
}
```

## 运行测试结果

执行 `go test` 后直接触发 **panic**：

```
runtime error: slice bounds out of range [1:0]
```

发生在 `construct_binary_tree_from_preorder_and_inorder_traversal.go:19`。

---

## 问题 1：根节点在中序序列中的位置定位错误

### 错误代码

```go
rootIdx := n / 2
```

### 分析

当前实现将根节点在中序数组中的索引简单取为数组长度的一半，这是**完全错误**的。

105 题的核心性质是：

- **前序遍历**的第一个元素 `preorder[0]` 就是当前子树的根节点
- **中序遍历**中，根节点左侧的所有元素构成左子树，右侧的所有元素构成右子树

因此，正确的做法是在 `inorder` 数组中**查找 `preorder[0]` 的位置**：

```go
rootIdx := 0
for rootIdx < n && inorder[rootIdx] != preorder[0] {
    rootIdx++
}
```

使用 `n / 2` 只有在根节点恰好位于中序数组正中间时才成立，对一般输入完全不适用。

---

## 问题 2：左子树节点数计算错误

### 错误代码

```go
leftNum := rootIdx - 1
```

### 分析

`leftNum` 被定义为左子树的节点数量，但由于 `rootIdx` 本身计算错误，加上这里混淆了"索引"与"数量"的概念，导致后续切片全部出错。

正确关系：
- 中序数组中，根节点索引为 `rootIdx`
- 左子树节点数就是 `rootIdx`（即 `inorder[0:rootIdx]` 的长度）
- 右子树节点数为 `n - rootIdx - 1`

因此：

```go
leftNum := rootIdx  // 左子树共有 rootIdx 个节点
```

---

## 问题 3：递归切片边界错误

### 错误代码

```go
root.Left = BuildTree(preorder[1:leftNum+1], inorder[0:rootIdx-1])
root.Right = BuildTree(preorder[leftNum+1:], inorder[rootIdx+1:])
```

### 分析

#### 左子树切片

- **前序部分**：左子树有 `leftNum` 个节点，前序中从索引 `1` 开始，应取 `preorder[1 : 1+leftNum]`
- **中序部分**：左子树为 `inorder[0:rootIdx]`，而不是 `inorder[0:rootIdx-1]`

当前代码写成 `preorder[1:leftNum+1]` 在 `leftNum = 0` 时等价于 `preorder[1:1]`，此时会触发 panic（当只有一个元素时，`leftNum = 0`，`preorder[1:1]` 没问题，但如果 `rootIdx = 0`，`leftNum = -1`，则变成 `preorder[1:0]`，导致 `[1:0]` 越界 panic）。

#### 右子树切片

- **前序部分**：右子树应取 `preorder[1+leftNum:]`
- **中序部分**：右子树应取 `inorder[rootIdx+1:]`，这部分当前是正确的

### 正确写法

```go
leftNum := rootIdx
root.Left = BuildTree(preorder[1:1+leftNum], inorder[0:rootIdx])
root.Right = BuildTree(preorder[1+leftNum:], inorder[rootIdx+1:])
```

---

## 问题 4：未对 `preorder` 和 `inorder` 长度不一致做防御

虽然题目约束保证 `len(preorder) == len(inorder)`，但健壮性实现应至少保证两者长度相等，否则切片划分可能错位。

```go
if len(preorder) != len(inorder) {
    return nil
}
```

---

## 问题 5：未查找根节点值是否存在

当前代码没有校验 `preorder[0]` 是否真实存在于 `inorder` 中。如果输入异常（违反题目约束），会导致 `rootIdx` 越界或产生错误切片。

虽然题目已保证两数组由相同唯一值组成，但添加校验能避免潜在 panic：

```go
rootIdx := -1
for i, v := range inorder {
    if v == preorder[0] {
        rootIdx = i
        break
    }
}
if rootIdx == -1 {
    return nil
}
```

---

## 修正后的实现思路

```go
func BuildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }

    rootVal := preorder[0]
    rootIdx := 0
    for rootIdx < len(inorder) && inorder[rootIdx] != rootVal {
        rootIdx++
    }

    leftNum := rootIdx
    root := &TreeNode{Val: rootVal}
    root.Left = BuildTree(preorder[1:1+leftNum], inorder[0:rootIdx])
    root.Right = BuildTree(preorder[1+leftNum:], inorder[rootIdx+1:])
    return root
}
```

### 复杂度补充

- **时间复杂度**：最坏情况下每次递归都要线性查找根节点位置，为 O(n²)；使用哈希表预处理中序索引可优化到 O(n)
- **空间复杂度**：递归栈深度 O(h)，h 为树高；最坏 O(n)

---

## 总结

当前实现的核心问题是**没有正确利用前序与中序遍历的关系来定位根节点**，导致：

1. `rootIdx` 错误地取为中序数组中点
2. 左子树节点数计算错误
3. 递归切片边界越界，直接触发 panic

必须先在前序数组中取第一个元素作为根，再在中序数组中查找该根的位置，才能正确划分左右子树。
