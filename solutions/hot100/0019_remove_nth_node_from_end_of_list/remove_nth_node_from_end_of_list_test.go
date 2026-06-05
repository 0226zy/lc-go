package removenthnodefromendoflist

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

func TestRemoveNthNodeFromEndOfList(t *testing.T) {
	tests := []struct {
		name string
		head []int
		n    int
		want []int
	}{
		{
			name: "官方示例1",
			head: []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			name: "官方示例2",
			head: []int{1},
			n:    1,
			want: []int{},
		},
		{
			name: "官方示例3",
			head: []int{1, 2},
			n:    1,
			want: []int{1},
		},
		{
			name: "删除头结点",
			head: []int{1, 2},
			n:    2,
			want: []int{2},
		},
		{
			name: "空链表",
			head: []int{},
			n:    0,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := datastructures.NewLinkedList(tt.head)
			got := RemoveNthNodeFromEndOfList(input, tt.n)
			var gotSlice []int
			if got != nil {
				gotSlice = got.ToSlice()
			}
			if !utils.EqualIntSlice(gotSlice, tt.want) {
				t.Errorf("RemoveNthNodeFromEndOfList(%v, %d) = %v, want %v", tt.head, tt.n, gotSlice, tt.want)
			}
		})
	}
}
