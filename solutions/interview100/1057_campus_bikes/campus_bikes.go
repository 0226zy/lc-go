package campusbikes

import "sort"

// AssignBikes 校园自行车分配
// n 名工人、m 辆自行车（n <= m），每轮从「未分配的工人 × 未占用的自行车」中选出
// 曼哈顿距离最短的一对分配；距离相同时优先选工人下标小、再选自行车下标小的组合。
// 返回长度为 n 的数组，answer[i] 表示第 i 名工人分配到的自行车下标。
// 时间复杂度: O(n*m*log(n*m))  空间复杂度: O(n*m)
func AssignBikes(workers [][]int, bikes [][]int) []int {
	n, m := len(workers), len(bikes)

	// 预计算所有 (工人, 自行车) 配对的曼哈顿距离
	type pair struct {
		dist, worker, bike int
	}
	pairs := make([]pair, 0, n*m)
	for i, w := range workers {
		for j, b := range bikes {
			d := abs(w[0]-b[0]) + abs(w[1]-b[1])
			pairs = append(pairs, pair{d, i, j})
		}
	}

	// 按 (距离, 工人下标, 自行车下标) 三级关键字升序排序
	sort.Slice(pairs, func(a, b int) bool {
		if pairs[a].dist != pairs[b].dist {
			return pairs[a].dist < pairs[b].dist
		}
		if pairs[a].worker != pairs[b].worker {
			return pairs[a].worker < pairs[b].worker
		}
		return pairs[a].bike < pairs[b].bike
	})

	answer := make([]int, n)
	usedBike := make([]bool, m)
	// answer[i] = -1 表示工人 i 尚未分配
	for i := range answer {
		answer[i] = -1
	}

	// 按贪心顺序依次配对，所有工人都分到自行车后提前结束
	assigned := 0
	for _, p := range pairs {
		if assigned == n {
			break
		}
		if answer[p.worker] != -1 || usedBike[p.bike] {
			continue
		}
		answer[p.worker] = p.bike
		usedBike[p.bike] = true
		assigned++
	}
	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
