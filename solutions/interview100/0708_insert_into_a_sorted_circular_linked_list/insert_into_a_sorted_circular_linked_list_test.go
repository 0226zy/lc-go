package insertintoasortedcircularlinkedlist

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
)

// newCircularList 从切片构造循环链表（尾节点指回头节点），空切片返回 nil
func newCircularList(vals []int) *datastructures.ListNode {
	if len(vals) == 0 {
		return nil
	}
	// pos = 0 表示尾节点指回头节点，形成完整闭环
	return datastructures.NewCycleLinkedList(vals, 0)
}

// circularToSlice 从 head 开始读取恰好 n 个节点的值（用于校验循环链表）
func circularToSlice(head *datastructures.ListNode, n int) []int {
	res := make([]int, 0, n)
	cur := head
	for i := 0; i < n; i++ {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name      string
		vals      []int // 循环链表的节点值（升序，从某个最小值开始描述）
		insertVal int
		want      []int // 插入后从返回节点开始读 len(vals)+1 个节点
	}{
		// LeetCode 官方示例
		{"示例1: 常规插入", []int{3, 4, 1}, 2, []int{3, 4, 1, 2}},
		{"示例2: 空链表", nil, 1, []int{1}},
		{"示例3: 单节点", []int{1}, 0, []int{1, 0}},

		// 边界：插入新的最小值/最大值，应插在断点处
		{"插入新的最小值", []int{3, 4, 1}, 0, []int{3, 4, 1, 0}},
		{"插入新的最大值", []int{3, 4, 1}, 5, []int{3, 4, 5, 1}},

		// 边界：所有值相等，转满一圈后任意插入
		{"所有值相等", []int{1, 1, 1}, 2, []int{1, 2, 1, 1}},

		// 边界：插入与已有值相同的值
		{"插入重复值", []int{1, 3, 5}, 3, []int{1, 3, 3, 5}},

		// 负数场景
		{"含负数", []int{-3, 0, -5}, -4, []int{-3, 0, -5, -4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := newCircularList(tt.vals)
			got := Insert(head, tt.insertVal)
			if got == nil {
				t.Fatalf("Insert(%v, %d) 返回 nil", tt.vals, tt.insertVal)
			}
			if res := circularToSlice(got, len(tt.vals)+1); !utils.EqualIntSlice(res, tt.want) {
				t.Errorf("Insert(%v, %d) 后循环链表为 %v, want %v", tt.vals, tt.insertVal, res, tt.want)
			}
		})
	}
}

func BenchmarkInsert(b *testing.B) {
	// 构造 5*10^4 节点的循环有序链表
	vals := make([]int, 50000)
	for i := range vals {
		vals[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 每次重新构造链表，保证插入规模一致
		head := newCircularList(vals)
		// 插入一个比所有值都大的数，触发最坏情况（遍历到断点）
		Insert(head, len(vals))
	}
}
