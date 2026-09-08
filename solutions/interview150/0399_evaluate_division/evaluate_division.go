package evaluatedivision

// CalcEquation 除法求值
// 给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，
// 其中 equations[i] = [Ai, Bi] 表示 Ai / Bi = values[i]。
// 请根据已知条件求出 queries[j] = [Cj, Dj] 的结果，即 Cj / Dj 的值。
// 如果结果不存在，返回 -1.0。题目保证不会有重复或矛盾的方程。
// 时间复杂度: O((E+Q)*V) E 为方程数，Q 为查询数，V 为变量数（每次查询一次 BFS）
// 空间复杂度: O(V+E) 建图的邻接表
func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 建图：边 a->b 权值为 values[i]，边 b->a 权值为 1/values[i]
	type edge struct {
		to string
		w  float64
	}
	graph := make(map[string][]edge)
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		graph[a] = append(graph[a], edge{to: b, w: values[i]})
		graph[b] = append(graph[b], edge{to: a, w: 1.0 / values[i]})
	}

	results := make([]float64, len(queries))
	for i, q := range queries {
		from, to := q[0], q[1]
		// 起点或终点不在图中，或起点即终点
		if _, ok := graph[from]; !ok || len(graph[from]) == 0 {
			results[i] = -1.0
			continue
		}
		if _, ok := graph[to]; !ok || len(graph[to]) == 0 {
			results[i] = -1.0
			continue
		}
		if from == to {
			results[i] = 1.0
			continue
		}

		// BFS：搜索 from 到 to 的路径，累乘边权
		visited := map[string]bool{from: true}
		type state struct {
			node string
			prod float64
		}
		queue := []state{{node: from, prod: 1.0}}
		found := -1.0
		for len(queue) > 0 && found < 0 {
			curr := queue[0]
			queue = queue[1:]
			for _, e := range graph[curr.node] {
				if visited[e.to] {
					continue
				}
				prod := curr.prod * e.w
				if e.to == to {
					found = prod
					break
				}
				visited[e.to] = true
				queue = append(queue, state{node: e.to, prod: prod})
			}
		}
		results[i] = found
	}
	return results
}
