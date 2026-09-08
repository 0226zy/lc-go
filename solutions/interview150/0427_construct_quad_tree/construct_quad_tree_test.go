package constructquadtree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConstruct(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want string // 层序序列化（官方格式：[isLeaf,val] 或 null），内部节点的 val 恒为 0
	}{
		{
			// LeetCode 官方示例 1
			name: "示例1: 2x2交错",
			grid: [][]int{{0, 1}, {1, 0}},
			want: "[[0,0],[1,0],[1,1],[1,1],[1,0]]",
		},
		{
			// LeetCode 官方示例 2：只有右上象限需要继续划分
			name: "示例2: 8x8",
			grid: [][]int{
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
			},
			want: "[[0,0],[1,1],[0,0],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]",
		},
		{
			// LeetCode 官方示例 3
			name: "示例3: 1x1全1",
			grid: [][]int{{1}},
			want: "[[1,1]]",
		},

		// 边界：1x1 全 0
		{name: "1x1全0", grid: [][]int{{0}}, want: "[[1,0]]"},

		// 边界：2x2 全相同
		{name: "2x2全1", grid: [][]int{{1, 1}, {1, 1}}, want: "[[1,1]]"},
		{name: "2x2全0", grid: [][]int{{0, 0}, {0, 0}}, want: "[[1,0]]"},

		// 边界：每个格子都需细分到 1x1 才统一
		{
			name: "4x4棋盘交错",
			grid: [][]int{{0, 1, 0, 1}, {1, 0, 1, 0}, {0, 1, 0, 1}, {1, 0, 1, 0}},
			want: "[[0,0],[0,0],[0,0],[0,0],[0,0]," +
				"[1,0],[1,1],[1,1],[1,0]," +
				"[1,0],[1,1],[1,1],[1,0]," +
				"[1,0],[1,1],[1,1],[1,0]," +
				"[1,0],[1,1],[1,1],[1,0]]",
		},
		// 边界：左半全0右半全1
		{
			name: "4x4左右对半",
			grid: [][]int{{0, 0, 1, 1}, {0, 0, 1, 1}, {0, 0, 1, 1}, {0, 0, 1, 1}},
			want: "[[0,0],[1,0],[1,1],[1,0],[1,1]]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := Construct(tt.grid)
			got := serialize(root)
			if got != tt.want {
				t.Errorf("Construct 层序序列化 = %v, want %v", got, tt.want)
			}
			// 额外校验：把四叉树还原成矩阵，必须与输入一致
			if !reflect.DeepEqual(decode(root, len(tt.grid)), tt.grid) {
				t.Errorf("Construct 结果还原的矩阵与输入不一致")
			}
		})
	}
}

// serialize 层序序列化四叉树（官方输出格式：[isLeaf,val] 表示节点，null 表示路径终止符）
// 注意：内部节点的 val 官方允许任意值，本实现统一为 0
func serialize(root *Node) string {
	if root == nil {
		return "[]"
	}
	var result []string
	queue := []*Node{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, "null")
			continue
		}
		v := 0
		if node.Val {
			v = 1
		}
		l := 0
		if node.IsLeaf {
			l = 1
		}
		result = append(result, fmt.Sprintf("[%d,%d]", l, v))
		// 每个节点都占 4 个子节点槽位：叶子节点的子节点输出 null
		queue = append(queue, node.TopLeft, node.TopRight, node.BottomLeft, node.BottomRight)
	}
	// 去掉末尾的 null
	for len(result) > 0 && result[len(result)-1] == "null" {
		result = result[:len(result)-1]
	}
	out := "["
	for i, s := range result {
		if i > 0 {
			out += ","
		}
		out += s
	}
	return out + "]"
}

// decode 把四叉树还原为 n×n 矩阵，用于校验建树正确性
func decode(root *Node, n int) [][]int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	fill(root, 0, 0, n, grid)
	return grid
}

// fill 把 node 代表的值填入 grid[r..r+size) × grid[c..c+size)
func fill(node *Node, r, c, size int, grid [][]int) {
	if node.IsLeaf {
		v := 0
		if node.Val {
			v = 1
		}
		for i := r; i < r+size; i++ {
			for j := c; j < c+size; j++ {
				grid[i][j] = v
			}
		}
		return
	}
	half := size / 2
	fill(node.TopLeft, r, c, half, grid)
	fill(node.TopRight, r, c+half, half, grid)
	fill(node.BottomLeft, r+half, c, half, grid)
	fill(node.BottomRight, r+half, c+half, half, grid)
}

func BenchmarkConstruct(b *testing.B) {
	benchmarks := []struct {
		name string
		grid [][]int
	}{
		{"n=8", generateGrid(8)},
		{"n=64", generateGrid(64)},
		{"n=256", generateGrid(256)},
		{"n=512", generateGrid(512)},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Construct(bm.grid)
			}
		})
	}
}

// generateGrid 生成 n×n 伪随机 01 矩阵
func generateGrid(n int) [][]int {
	grid := make([][]int, n)
	seed := 1
	for i := range grid {
		grid[i] = make([]int, n)
		for j := range grid[i] {
			seed = (seed*1103515245 + 12345) % (1 << 31)
			grid[i][j] = seed % 2
		}
	}
	return grid
}
