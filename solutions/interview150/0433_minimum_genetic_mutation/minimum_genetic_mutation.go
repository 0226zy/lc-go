package minimumgeneticmutation

// MinMutation 最小基因变化
// 基因由 8 个字符组成，每个字符是 A/C/G/T 之一。每次变化只能改变一个字符，
// 且变化结果必须在基因库 bank 中。求从 startGene 到 endGene 的最少变化次数，无法到达返回 -1。
// 时间复杂度: O(n * 8 * 3) 每个基因尝试 8 个位置各 3 种替换  空间复杂度: O(n) 队列与 visited 集合
func MinMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	bankSet := make(map[string]bool, len(bank))
	for _, g := range bank {
		bankSet[g] = true
	}
	if !bankSet[endGene] {
		return -1
	}

	const genes = "ACGT"
	visited := map[string]bool{startGene: true}
	queue := []string{startGene}
	steps := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur == endGene {
				return steps
			}
			// 逐位尝试替换为其他 3 种字符
			for j := 0; j < len(cur); j++ {
				for _, ch := range genes {
					if byte(ch) == cur[j] {
						continue
					}
					next := cur[:j] + string(ch) + cur[j+1:]
					if bankSet[next] && !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
			}
		}
		steps++
	}
	return -1
}
