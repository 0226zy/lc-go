package meetingroomsii

import (
	"container/heap"
	"sort"
)

// intHeap 最小堆，存各会议室的结束时间
type intHeap []int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// MinMeetingRooms 会议室 II
// 给定会议时间区间数组 intervals，返回容纳所有会议所需的最小会议室数量。
// 时间复杂度: O(n log n) 排序 + 每次堆操作 O(log n)  空间复杂度: O(n) 堆
func MinMeetingRooms(intervals [][]int) int {
	// 按开始时间升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 堆顶是最早空闲的会议室的结束时间
	rooms := &intHeap{}
	heap.Init(rooms)
	for _, meeting := range intervals {
		start, end := meeting[0], meeting[1]
		// 最早空闲的会议室在当前会议开始前已空出，可复用
		if rooms.Len() > 0 && (*rooms)[0] <= start {
			heap.Pop(rooms)
		}
		heap.Push(rooms, end)
	}
	return rooms.Len()
}
