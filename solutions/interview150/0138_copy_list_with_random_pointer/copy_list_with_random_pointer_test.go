package copylistwithrandompointer

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
)

// randomPairs 从随机链表提取 [val, random指向的值] 对，用于断言结构完全一致
func randomPairs(head *datastructures.RandomListNode) [][2]interface{} {
	var pairs [][2]interface{}
	for curr := head; curr != nil; curr = curr.Next {
		var r interface{}
		if curr.Random != nil {
			r = curr.Random.Val
		}
		pairs = append(pairs, [2]interface{}{curr.Val, r})
	}
	return pairs
}

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		name          string
		vals          []int
		randomIndices []int
	}{
		// LeetCode 官方示例
		{"示例1: [[7,null],[13,0],[11,4],[10,2],[1,0]]", []int{7, 13, 11, 10, 1}, []int{-1, 0, 4, 2, 0}},
		{"示例2: [[1,1],[2,1]]", []int{1, 2}, []int{1, 1}},
		{"示例3: [[3,null],[3,0],[3,null]]", []int{3, 3, 3}, []int{-1, 0, -1}},

		// 边界：空链表
		{"空链表", []int{}, nil},

		// 边界：单节点，random 为 nil
		{"单节点 random 为 nil", []int{42}, []int{-1}},

		// 边界：单节点，random 指向自身
		{"单节点 random 指向自身", []int{5}, []int{0}},

		// 边界：全部 random 为 nil
		{"全部 random 为 nil", []int{1, 2, 3, 4}, []int{-1, -1, -1, -1}},

		// 边界：全部 random 指向自身
		{"全部 random 指向自身", []int{1, 2, 3}, []int{0, 1, 2}},

		// 边界：random 全部指向下一个节点
		{"random 指向下一个节点", []int{10, 20, 30}, []int{1, 2, -1}},

		// 边界：random 全部指向前一个节点
		{"random 指向前一个节点", []int{10, 20, 30}, []int{-1, 0, 1}},

		// 边界：最后一个节点 random 指向头节点
		{"尾节点 random 指向头", []int{1, 2, 3, 4, 5}, []int{-1, 0, 1, 2, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := datastructures.NewRandomLinkedList(tt.vals, tt.randomIndices)
			copyHead := CopyRandomList(head)

			// 值序列一致
			gotVals := datastructures.RandomListToSlice(copyHead)
			var want []int
			if len(tt.vals) > 0 {
				want = tt.vals
			}
			if len(gotVals) != len(want) {
				t.Fatalf("拷贝链表长度 = %d, want %d", len(gotVals), len(want))
			}
			for i := range want {
				if gotVals[i] != want[i] {
					t.Fatalf("拷贝链表值序列 = %v, want %v", gotVals, want)
				}
			}

			// random 指向关系一致
			wantPairs := randomPairs(head)
			gotPairs := randomPairs(copyHead)
			if len(gotPairs) != len(wantPairs) {
				t.Fatalf("random 对数量 = %d, want %d", len(gotPairs), len(wantPairs))
			}
			for i := range wantPairs {
				if gotPairs[i] != wantPairs[i] {
					t.Errorf("第 %d 个节点: got (val=%v, random=%v), want (val=%v, random=%v)",
						i, gotPairs[i][0], gotPairs[i][1], wantPairs[i][0], wantPairs[i][1])
				}
			}

			// 深拷贝验证：拷贝节点的地址必须与所有原节点不同
			if head != nil && copyHead == head {
				t.Errorf("拷贝链表头节点与原链表是同一个节点，不是深拷贝")
			}
		})
	}
}

func BenchmarkCopyRandomList(b *testing.B) {
	vals := make([]int, 1000)
	randomIndices := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		vals[i] = i
		randomIndices[i] = (i * 7) % 1000 // random 随机落在链表中
	}
	head := datastructures.NewRandomLinkedList(vals, randomIndices)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CopyRandomList(head)
	}
}
