package constructquadtree

// Node 四叉树节点
// 每个节点有四个子节点：TopLeft（左上）、TopRight（右上）、
// BottomLeft（左下）、BottomRight（右下）。
// IsLeaf 为 true 表示该节点是叶子节点，此时 Val 就是该区域的统一值。
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

// Construct 建立四叉树
// 给定一个 n×n 的矩阵 grid，矩阵中每个格子的值要么是 0 要么是 1。
// 请用四叉树表示这个矩阵：
//   - 如果一个区域内所有格子的值都相同，那么该区域对应一个叶子节点；
//   - 否则将该区域均分为四个子区域，递归构造四个子节点。
//
// 核心思路：递归判断当前区域值是否一致（一次遍历即可确认），
// 一致就返回叶子节点，不一致就“十字切两刀”分成四块继续递归。
// 每个格子在最深的叶子处被检查一次，中间节点会重复检查子区域，
// 所有一致性检查的总代价为 O(n^2)。
// 时间复杂度: O(n^2)  空间复杂度: O(log n) 递归栈深度
func Construct(grid [][]int) *Node {
	return build(grid, 0, 0, len(grid))
}

// build 递归构造 grid 中以 (r, c) 为左上角、边长为 size  的区域对应的四叉树节点
func build(grid [][]int, r, c, size int) *Node {
	// 先检查该区域所有值是否一致
	first := grid[r][c]
	uniform := true
	for i := r; i < r+size && uniform; i++ {
		for j := c; j < c+size; j++ {
			if grid[i][j] != first {
				uniform = false
				break
			}
		}
	}
	val := first == 1
	if uniform {
		return &Node{Val: val, IsLeaf: true}
	}
	half := size / 2
	return &Node{
		Val:         false, // 非叶子节点的 Val 官方定义为 false
		IsLeaf:      false,
		TopLeft:     build(grid, r, c, half),
		TopRight:    build(grid, r, c+half, half),
		BottomLeft:  build(grid, r+half, c, half),
		BottomRight: build(grid, r+half, c+half, half),
	}
}
