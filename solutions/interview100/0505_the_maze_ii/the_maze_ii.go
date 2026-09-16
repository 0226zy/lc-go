package themazeii

import "container/heap"

// state 堆中的一个状态：球停在 (x, y)，从起点滚到这里的总距离为 d
type state struct {
	x, y int
	d    int
}

// minHeap 以距离 d 为键的小顶堆
type minHeap []state

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].d < h[j].d }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) { *h = append(*h, x.(state)) }

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	s := old[n-1]
	*h = old[:n-1]
	return s
}

// ShortestDistance 迷宫 II
// 迷宫由 m×n 矩阵表示，0 为空地、1 为墙。球从 start 出发，每次选一个方向一直滚直到撞墙才停下，
// 求滚到 destination 并恰好停下的最短距离（经过的空格数），无法到达返回 -1。
// 时间复杂度: O(m*n*max(m,n))  空间复杂度: O(m*n)
func ShortestDistance(maze [][]int, start []int, destination []int) int {
	m, n := len(maze), len(maze[0])
	const inf = int(^uint(0) >> 1)
	// dist[x][y] 表示球停在 (x, y) 的最短距离
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = inf
		}
	}
	dist[start[0]][start[1]] = 0

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	h := &minHeap{{start[0], start[1], 0}}
	heap.Init(h)
	for h.Len() > 0 {
		cur := heap.Pop(h).(state)
		// 过期状态：已有更短距离，跳过
		if cur.d > dist[cur.x][cur.y] {
			continue
		}
		for _, d := range dirs {
			// 沿该方向一直滚到撞墙（或出界）前的最后一个空格
			x, y, step := cur.x, cur.y, 0
			for x+d[0] >= 0 && x+d[0] < m && y+d[1] >= 0 && y+d[1] < n && maze[x+d[0]][y+d[1]] == 0 {
				x += d[0]
				y += d[1]
				step++
			}
			if cur.d+step < dist[x][y] {
				dist[x][y] = cur.d + step
				heap.Push(h, state{x, y, cur.d + step})
			}
		}
	}
	if dist[destination[0]][destination[1]] == inf {
		return -1
	}
	return dist[destination[0]][destination[1]]
}
