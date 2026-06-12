package permutations

// Permute 全排列
// 给定一个不含重复数字的数组 nums，返回其所有可能的全排列。
// 你可以按任意顺序返回答案。
// 时间复杂度: O(?)  空间复杂度: O(?)
func Permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	n := len(nums)
	visited := make([]bool, n)
	ret := [][]int{}
	var backtrack func(int, *[]int)
	backtrack = func(idx int, path *[]int) {
		if idx == n {
			tmp := make([]int, n)
			copy(tmp, *path)
			ret = append(ret, tmp)
			return
		}
		for i, v := range nums {
			if visited[i] {
				continue
			}
			*path = append(*path, v)
			visited[i] = true
			backtrack(idx+1, path)
			visited[i] = false
			*path = (*path)[:len(*path)-1]
		}
	}
	path := []int{}
	backtrack(0, &path)

	return ret
}
