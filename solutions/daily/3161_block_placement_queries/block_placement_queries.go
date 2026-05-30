package blockplacementqueries

import "sort"

// fenwick 树状数组，维护前缀最大值
type fenwick struct {
	n    int
	tree []int
}

func newFenwick(n int) *fenwick {
	return &fenwick{
		n:    n,
		tree: make([]int, n+2),
	}
}

func (f *fenwick) maximize(i, val int) {
	for i <= f.n {
		if val > f.tree[i] {
			f.tree[i] = val
		}
		i += i & -i
	}
}

func (f *fenwick) get(i int) int {
	res := 0
	for i > 0 {
		if f.tree[i] > res {
			res = f.tree[i]
		}
		i -= i & -i
	}
	return res
}

// GetResults 物块放置查询
// 给定一个二维数组 queries，包含两种操作：
// 操作类型 1：queries[i] = [1, x]，在距离原点 x 处建一个障碍物。
// 操作类型 2：queries[i] = [2, x, sz]，判断在数轴范围 [0, x] 内是否可以放置一个长度为 sz 的物块。
// 返回一个布尔数组，对每个类型 2 的操作给出结果。
//
// 核心思路：离线逆序处理 + 并查集维护前驱 + 树状数组维护前缀最大间隔。
// 先正序收集所有障碍物并初始化树状数组；然后逆序遍历，将“添加障碍物”转为“删除障碍物”，
// 删除时只需更新右侧邻居的间隔，查询时取 [0, 前驱] 的最大间隔与 x-前驱 的最大值。
//
// 时间复杂度: O(n log n)  空间复杂度: O(n)
func GetResults(queries [][]int) []bool {
	n := len(queries) * 3
	if n > 50000 {
		n = 50000
	}

	obsSet := make(map[int]struct{})
	for _, q := range queries {
		if q[0] == 1 {
			obsSet[q[1]] = struct{}{}
		}
	}

	coords := make([]int, 0, len(obsSet)+2)
	coords = append(coords, 0)
	coords = append(coords, n)
	for x := range obsSet {
		coords = append(coords, x)
	}
	sort.Ints(coords)

	// 去重
	uniq := coords[:1]
	for i := 1; i < len(coords); i++ {
		if coords[i] != coords[i-1] {
			uniq = append(uniq, coords[i])
		}
	}
	coords = uniq
	m := len(coords)

	posMap := make(map[int]int, m)
	for i, v := range coords {
		posMap[v] = i
	}

	prev := make([]int, m)
	next := make([]int, m)
	for i := 0; i < m; i++ {
		prev[i] = i - 1
		next[i] = i + 1
	}
	next[m-1] = -1

	fa := make([]int, m)
	for i := 0; i < m; i++ {
		fa[i] = i
	}

	find := func(x int) int {
		root := x
		for fa[root] != root {
			root = fa[root]
		}
		for x != root {
			fa[x], x = root, fa[x]
		}
		return root
	}

	tree := newFenwick(m)

	// 初始化所有相邻障碍物之间的间隔，在右端点位置更新最大值
	for i := 1; i < m; i++ {
		tree.maximize(i+1, coords[i]-coords[i-1])
	}

	res := make([]bool, 0, len(queries))
	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		if q[0] == 1 {
			x := q[1]
			idx := posMap[x]
			p := prev[idx]
			ne := next[idx]
			if ne != -1 {
				next[p] = ne
				prev[ne] = p
				fa[idx] = p
				tree.maximize(ne+1, coords[ne]-coords[p])
			} else {
				next[p] = -1
				fa[idx] = p
			}
		} else {
			x, sz := q[1], q[2]
			pos := sort.Search(m, func(i int) bool { return coords[i] > x }) - 1
			predIdx := find(pos)
			maxGap := tree.get(predIdx + 1)
			if gap := x - coords[predIdx]; gap > maxGap {
				maxGap = gap
			}
			res = append(res, maxGap >= sz)
		}
	}

	// 逆序恢复正序结果
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
