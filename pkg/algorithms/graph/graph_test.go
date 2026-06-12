package graph

import (
	"math"
	"testing"
)

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestNewGraph(t *testing.T) {
	g := NewGraph(5)
	if g.Size() != 5 {
		t.Errorf("Size() = %d, want 5", g.Size())
	}
}

func TestBFS(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][3]int
		start    int
		expected []int
	}{
		{
			name:     "基本有向图",
			n:        6,
			edges:    [][3]int{{0, 1, 1}, {0, 2, 1}, {1, 3, 1}, {1, 4, 1}, {2, 4, 1}, {3, 5, 1}, {4, 5, 1}},
			start:    0,
			expected: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "无向图",
			n:        5,
			edges:    [][3]int{{0, 1, 1}, {0, 2, 1}, {1, 3, 1}, {2, 4, 1}},
			start:    0,
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "存在不可达节点",
			n:        4,
			edges:    [][3]int{{0, 1, 1}, {2, 3, 1}},
			start:    0,
			expected: []int{0, 1},
		},
		{
			name:     "单节点图",
			n:        1,
			edges:    nil,
			start:    0,
			expected: []int{0},
		},
		{
			name:     "非法起点",
			n:        3,
			edges:    [][3]int{{0, 1, 1}},
			start:    -1,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.n)
			for _, e := range tt.edges {
				g.AddEdge(e[0], e[1], e[2])
			}
			got := g.BFS(tt.start)
			if !intSliceEqual(got, tt.expected) {
				t.Errorf("BFS() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDFS(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][3]int
		start    int
		expected []int
	}{
		{
			name:     "基本有向图",
			n:        6,
			edges:    [][3]int{{0, 1, 1}, {0, 2, 1}, {1, 3, 1}, {1, 4, 1}, {2, 4, 1}, {3, 5, 1}, {4, 5, 1}},
			start:    0,
			expected: []int{0, 1, 3, 5, 4, 2},
		},
		{
			name:     "无向图",
			n:        5,
			edges:    [][3]int{{0, 1, 1}, {0, 2, 1}, {1, 3, 1}, {2, 4, 1}},
			start:    0,
			expected: []int{0, 1, 3, 2, 4},
		},
		{
			name:     "存在不可达节点",
			n:        4,
			edges:    [][3]int{{0, 1, 1}, {2, 3, 1}},
			start:    0,
			expected: []int{0, 1},
		},
		{
			name:     "单节点图",
			n:        1,
			edges:    nil,
			start:    0,
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.n)
			for _, e := range tt.edges {
				g.AddEdge(e[0], e[1], e[2])
			}
			got := g.DFS(tt.start)
			if !intSliceEqual(got, tt.expected) {
				t.Errorf("DFS() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestBellmanFord(t *testing.T) {
	tests := []struct {
		name       string
		n          int
		edges      [][3]int
		start      int
		expected   []int
		hasNegCycle bool
	}{
		{
			name:       "基本最短路",
			n:          5,
			edges:      [][3]int{{0, 1, -1}, {0, 2, 4}, {1, 2, 3}, {1, 3, 2}, {1, 4, 2}, {3, 2, 5}, {3, 1, 1}, {4, 3, -3}},
			start:      0,
			expected:   []int{0, -1, 2, -2, 1},
			hasNegCycle: false,
		},
		{
			name:       "不可达节点",
			n:          4,
			edges:      [][3]int{{0, 1, 5}, {2, 3, 1}},
			start:      0,
			expected:   []int{0, 5, math.MaxInt64, math.MaxInt64},
			hasNegCycle: false,
		},
		{
			name:       "负权环",
			n:          3,
			edges:      [][3]int{{0, 1, 1}, {1, 2, -3}, {2, 0, -1}},
			start:      0,
			expected:   nil,
			hasNegCycle: true,
		},
		{
			name:       "单节点",
			n:          1,
			edges:      nil,
			start:      0,
			expected:   []int{0},
			hasNegCycle: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.n)
			for _, e := range tt.edges {
				g.AddEdge(e[0], e[1], e[2])
			}
			dist, hasNegCycle := g.BellmanFord(tt.start)
			if hasNegCycle != tt.hasNegCycle {
				t.Errorf("BellmanFord() hasNegCycle = %v, want %v", hasNegCycle, tt.hasNegCycle)
			}
			if tt.expected != nil && !intSliceEqual(dist, tt.expected) {
				t.Errorf("BellmanFord() dist = %v, want %v", dist, tt.expected)
			}
		})
	}
}

func TestDijkstra(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][3]int
		start    int
		expected []int
	}{
		{
			name:     "基本最短路",
			n:        5,
			edges:    [][3]int{{0, 1, 10}, {0, 4, 5}, {1, 4, 2}, {4, 1, 3}, {1, 2, 1}, {4, 2, 9}, {4, 3, 2}, {3, 2, 6}, {2, 3, 4}},
			start:    0,
			expected: []int{0, 8, 9, 7, 5},
		},
		{
			name:     "无向图最短路",
			n:        4,
			edges:    [][3]int{{0, 1, 1}, {1, 2, 2}, {2, 3, 3}, {0, 3, 10}},
			start:    0,
			expected: []int{0, 1, 3, 6},
		},
		{
			name:     "不可达节点",
			n:        4,
			edges:    [][3]int{{0, 1, 1}, {2, 3, 1}},
			start:    0,
			expected: []int{0, 1, math.MaxInt64, math.MaxInt64},
		},
		{
			name:     "单节点图",
			n:        1,
			edges:    nil,
			start:    0,
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGraph(tt.n)
			for _, e := range tt.edges {
				g.AddEdge(e[0], e[1], e[2])
			}
			got := g.Dijkstra(tt.start)
			if !intSliceEqual(got, tt.expected) {
				t.Errorf("Dijkstra() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkBFS(b *testing.B) {
	g := NewGraph(100)
	for i := 0; i < 99; i++ {
		g.AddEdge(i, i+1, 1)
	}
	for i := 0; i < b.N; i++ {
		_ = g.BFS(0)
	}
}

func BenchmarkDFS(b *testing.B) {
	g := NewGraph(100)
	for i := 0; i < 99; i++ {
		g.AddEdge(i, i+1, 1)
	}
	for i := 0; i < b.N; i++ {
		_ = g.DFS(0)
	}
}

func BenchmarkBellmanFord(b *testing.B) {
	g := NewGraph(20)
	for i := 0; i < 19; i++ {
		g.AddEdge(i, i+1, i+1)
	}
	for i := 0; i < b.N; i++ {
		_, _ = g.BellmanFord(0)
	}
}

func BenchmarkDijkstra(b *testing.B) {
	g := NewGraph(100)
	for i := 0; i < 99; i++ {
		g.AddEdge(i, i+1, i+1)
	}
	for i := 0; i < b.N; i++ {
		_ = g.Dijkstra(0)
	}
}
