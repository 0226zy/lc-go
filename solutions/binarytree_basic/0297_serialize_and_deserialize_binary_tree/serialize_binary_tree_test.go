package serializebinarytree

import (
	"math"
	"reflect"
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

const nilNode = math.MinInt32

func TestSerializeDeserialize(t *testing.T) {
	tests := []struct {
		name string
		vals []int
	}{
		// LeetCode 官方示例
		{"示例1: [1,2,3,null,null,4,5]", []int{1, 2, 3, nilNode, nilNode, 4, 5}},
		{"示例2: 空树", nil},

		// 边界
		{"单节点", []int{1}},
		{"只有左子树", []int{1, 2, nilNode, 3}},
		{"只有右子树", []int{1, nilNode, 2, nilNode, 3}},
		{"满二叉树", []int{1, 2, 3, 4, 5, 6, 7}},
		{"含负数与零", []int{0, -1, 2, nilNode, -3, nilNode, 4}},
		{"右斜链", []int{1, nilNode, 2, nilNode, 3, nilNode, 4}},
		{"左斜链", []int{4, 3, nilNode, 2, nilNode, 1}},
	}

	codec := Constructor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := datastructures.NewTreeFromSlice(tt.vals)
			data := codec.Serialize(root)
			got := codec.Deserialize(data)

			wantLevel := root.LevelOrder()
			gotLevel := got.LevelOrder()
			if !reflect.DeepEqual(gotLevel, wantLevel) {
				t.Errorf("往返后层序 = %v, want %v（序列化串 %q）", gotLevel, wantLevel, data)
			}
		})
	}
}

func TestSerializeEmpty(t *testing.T) {
	codec := Constructor()
	data := codec.Serialize(nil)
	if data != "N" {
		t.Errorf("空树序列化 = %q, want %q", data, "N")
	}
	if got := codec.Deserialize(data); got != nil {
		t.Errorf("空树反序列化应返回 nil, got %v", got.LevelOrder())
	}
}

func TestSerializeFormat(t *testing.T) {
	// 固定一棵小树，校验前序 + N 标记格式（便于锁定实现约定）
	root := datastructures.NewTreeFromSlice([]int{1, 2, 3})
	codec := Constructor()
	got := codec.Serialize(root)
	want := "1,2,N,N,3,N,N"
	if got != want {
		t.Errorf("Serialize([1,2,3]) = %q, want %q", got, want)
	}
}

func BenchmarkSerializeDeserialize(b *testing.B) {
	vals := make([]int, 0, 2047)
	for size := 1; len(vals)+size <= 2047; size *= 2 {
		for i := 0; i < size; i++ {
			vals = append(vals, i+1)
		}
	}
	root := datastructures.NewTreeFromSlice(vals)
	codec := Constructor()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := codec.Serialize(root)
		_ = codec.Deserialize(data)
	}
}
