package serializeanddeserializenarytree

import (
	"math"
	"testing"
)

// nilV 表示层序切片中每组子节点之间的分隔符（null）
const nilV = math.MinInt32

// newNAryTree 按 LeetCode 层序表示构造 N 叉树，nilV 分隔每个节点的孩子组
func newNAryTree(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}
	root := &Node{Val: vals[0]}
	queue := []*Node{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		i++ // 跳过孩子组前的分隔符
		node := queue[0]
		queue = queue[1:]
		for i < len(vals) && vals[i] != nilV {
			child := &Node{Val: vals[i]}
			node.Children = append(node.Children, child)
			queue = append(queue, child)
			i++
		}
	}
	return root
}

// equalNAryTree 递归比较两棵 N 叉树的结构与节点值是否完全一致
func equalNAryTree(a, b *Node) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.Val != b.Val || len(a.Children) != len(b.Children) {
		return false
	}
	for i := range a.Children {
		if !equalNAryTree(a.Children[i], b.Children[i]) {
			return false
		}
	}
	return true
}

func TestCodecSerialize(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want string
	}{
		// 官方示例1对应的先序编码：值 + 孩子数
		{"示例1的树", []int{1, nilV, 3, 2, 4, nilV, 5, 6}, "1 3 3 2 5 0 6 0 2 0 4 0"},
		{"空树", nil, ""},
		{"单节点", []int{1}, "1 0"},
		{"长链", []int{1, nilV, 2, nilV, 3, nilV, 4}, "1 1 2 1 3 1 4 0"},
	}

	codec := Constructor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := newNAryTree(tt.vals)
			if got := codec.serialize(root); got != tt.want {
				t.Errorf("serialize() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCodecDeserialize(t *testing.T) {
	tests := []struct {
		name string
		data string
		vals []int
	}{
		{"示例1字符串", "1 3 3 2 5 0 6 0 2 0 4 0", []int{1, nilV, 3, 2, 4, nilV, 5, 6}},
		{"空字符串", "", nil},
		{"单节点字符串", "1 0", []int{1}},
	}

	codec := Constructor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := codec.deserialize(tt.data)
			want := newNAryTree(tt.vals)
			if !equalNAryTree(got, want) {
				t.Errorf("deserialize(%q) 结构不符，重新序列化得 %q，期望 %q",
					tt.data, codec.serialize(got), codec.serialize(want))
			}
		})
	}
}

func TestCodecRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		vals []int
	}{
		// LeetCode 官方示例
		{"示例1: [1,null,3,2,4,null,5,6]", []int{1, nilV, 3, 2, 4, nilV, 5, 6}},
		{"示例2: 深层多叉树", []int{1, nilV, 2, 3, 4, 5, nilV, nilV, 6, 7, nilV, 8, nilV, 9, 10, nilV, nilV, 11, nilV, 12, nilV, 13, nilV, nilV, 14}},

		// 边界情况
		{"空树", nil},
		{"单节点", []int{42}},
		{"只有一层多孩子", []int{1, nilV, 2, 3, 4, 5, 6, 7}},
		{"退化为链", []int{1, nilV, 2, nilV, 3, nilV, 4, nilV, 5}},
		{"极大节点值", []int{10000, nilV, 0, 9999}},
	}

	codec := Constructor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := newNAryTree(tt.vals)
			restored := codec.deserialize(codec.serialize(root))
			if !equalNAryTree(restored, root) {
				t.Errorf("往返后结构不一致：原树编码 %q，还原树编码 %q",
					codec.serialize(root), codec.serialize(restored))
			}
		})
	}
}

func BenchmarkCodec(b *testing.B) {
	// 构造一棵 5 叉、深度 5 的满树（1+5+25+125+625 = 781 个节点）
	var build func(depth int, val *int) *Node
	build = func(depth int, val *int) *Node {
		node := &Node{Val: *val}
		*val++
		if depth > 0 {
			for i := 0; i < 5; i++ {
				node.Children = append(node.Children, build(depth-1, val))
			}
		}
		return node
	}

	benchmarks := []struct {
		name  string
		depth int
	}{
		{"depth=2", 2},
		{"depth=4", 4},
		{"depth=5", 5},
	}

	codec := Constructor()
	for _, bm := range benchmarks {
		v := 1
		root := build(bm.depth, &v)
		data := codec.serialize(root)
		b.Run("serialize/"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				codec.serialize(root)
			}
		})
		b.Run("deserialize/"+bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				codec.deserialize(data)
			}
		})
	}
}
