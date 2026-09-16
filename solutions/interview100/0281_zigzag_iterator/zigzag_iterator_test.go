package zigzagiterator

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestZigzagIterator(t *testing.T) {
	tests := []struct {
		name string
		v1   []int
		v2   []int
		// ops 仅含 "next" / "hasnext"，want 为每次操作对应的返回值
		ops  []string
		want []int // hasnext 用 1 表示 true，0 表示 false
	}{
		{
			name: "示例1: v1较短",
			v1:   []int{1, 2},
			v2:   []int{3, 4, 5, 6},
			ops:  []string{"next", "next", "next", "next", "next", "next", "hasnext"},
			want: []int{1, 3, 2, 4, 5, 6, 0},
		},
		{
			name: "示例2: v2为空",
			v1:   []int{1},
			v2:   []int{},
			ops:  []string{"next", "hasnext"},
			want: []int{1, 0},
		},
		{
			name: "v1为空",
			v1:   []int{},
			v2:   []int{7, 8},
			ops:  []string{"hasnext", "next", "next", "hasnext"},
			want: []int{1, 7, 8, 0},
		},
		{
			name: "两个向量等长",
			v1:   []int{1, 3, 5},
			v2:   []int{2, 4, 6},
			ops:  []string{"next", "next", "next", "next", "next", "next", "hasnext"},
			want: []int{1, 2, 3, 4, 5, 6, 0},
		},
		{
			name: "v2较短",
			v1:   []int{1, 2, 3, 4},
			v2:   []int{5},
			ops:  []string{"next", "next", "next", "next", "hasnext"},
			want: []int{1, 5, 2, 3, 1},
		},
		{
			name: "连续多次hasnext结果稳定",
			v1:   []int{1},
			v2:   []int{2},
			ops:  []string{"hasnext", "hasnext", "next", "hasnext", "next", "hasnext", "hasnext"},
			want: []int{1, 1, 1, 1, 2, 0, 0},
		},
		{
			name: "单元素v1单元素v2",
			v1:   []int{-2147483648},
			v2:   []int{2147483647},
			ops:  []string{"next", "next", "hasnext"},
			want: []int{-2147483648, 2147483647, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := Constructor(tt.v1, tt.v2)
			var got []int
			for _, op := range tt.ops {
				switch op {
				case "next":
					got = append(got, it.next())
				case "hasnext":
					if it.hasnext() {
						got = append(got, 1)
					} else {
						got = append(got, 0)
					}
				}
			}
			if !utils.EqualIntSlice(got, tt.want) {
				t.Errorf("操作序列 %v 结果 = %v, want %v", tt.ops, got, tt.want)
			}
		})
	}
}

// TestZigzagIteratorExhaustAll 完整遍历两个向量，校验输出是合法的锯齿序列
func TestZigzagIteratorExhaustAll(t *testing.T) {
	v1 := []int{1, 2, 3}
	v2 := []int{4, 5, 6, 7, 8}
	it := Constructor(v1, v2)
	var got []int
	for it.hasnext() {
		got = append(got, it.next())
	}
	want := []int{1, 4, 2, 5, 3, 6, 7, 8}
	if !utils.EqualIntSlice(got, want) {
		t.Errorf("完整遍历结果 = %v, want %v", got, want)
	}
}

func BenchmarkZigzagIterator(b *testing.B) {
	v1 := make([]int, 1000)
	v2 := make([]int, 1000)
	for i := range v1 {
		v1[i] = i
		v2[i] = i + 1000
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		it := Constructor(v1, v2)
		for it.hasnext() {
			it.next()
		}
	}
}
