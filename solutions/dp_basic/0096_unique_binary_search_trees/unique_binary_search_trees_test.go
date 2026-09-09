package uniquebinarysearchtrees

import "testing"

var numTreesCases = []struct {
	name string
	n    int
	want int
}{
	{name: "1个节点", n: 1, want: 1},
	{name: "2个节点", n: 2, want: 2},
	{name: "3个节点（示例）", n: 3, want: 5},
	{name: "4个节点", n: 4, want: 14},
	{name: "5个节点", n: 5, want: 42},
	{name: "10个节点", n: 10, want: 16796},
	{name: "19个节点（约束上限）", n: 19, want: 1767263190},
}

func TestNumTrees(t *testing.T) {
	for _, tt := range numTreesCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumTrees(tt.n); got != tt.want {
				t.Errorf("NumTrees(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestNumTreesAlternative(t *testing.T) {
	for _, tt := range numTreesCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumTreesAlternative(tt.n); got != tt.want {
				t.Errorf("NumTreesAlternative(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkNumTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumTrees(19)
	}
}

func BenchmarkNumTreesAlternative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumTreesAlternative(19)
	}
}
