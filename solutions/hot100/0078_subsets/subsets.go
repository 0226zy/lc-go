package subsets

// Subsets 子集
// 给定互不相同的整数数组 nums，返回所有可能的子集（幂集）。
// 时间复杂度: O(n × 2^n)  空间复杂度: O(n)
func Subsets(nums []int) [][]int {
	ret := [][]int{}
	set := []int{}
	var dfs func(curr int)
	dfs = func(curr int) {
		if curr == len(nums) {
			ret = append(ret, append([]int{}, set...))
			return
		}
		set = append(set, nums[curr])
		dfs(curr + 1)
		set = set[:len(set)-1]
		dfs(curr + 1)
	}
	dfs(0)
	return ret

}

// SubsetsBitmask 子集（位掩码迭代）
// 利用二进制位掩码枚举所有子集。
// 时间复杂度: O(n × 2^n)  空间复杂度: O(n × 2^n)
func SubsetsBitmask(nums []int) [][]int {
	panic("not implemented")
}
