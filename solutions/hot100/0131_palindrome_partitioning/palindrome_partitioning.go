package palindromepartitioning

// Partition 分割回文串
// 给定字符串 s，返回所有使每个子串都是回文串的分割方案。
// 时间复杂度: O(n × 2^n)  空间复杂度: O(n)
func Partition(s string) [][]string {
	panic("not implemented")
}

// PartitionWithTable 分割回文串（预处理回文表）
// 预先用 DP 计算 isPalin[i][j] 表，避免回溯中重复判断回文。
// 时间复杂度: O(n² + n × 2^n)  空间复杂度: O(n²)
func PartitionWithTable(s string) [][]string {
	panic("not implemented")
}
