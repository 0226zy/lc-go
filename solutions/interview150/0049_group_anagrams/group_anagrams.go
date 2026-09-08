package groupanagrams

import "sort"

// GroupAnagrams 字母异位词分组
// 将字符串数组 strs 中的字母异位词（字符重新排列可互相得到）分到同一组，返回全部分组。
// 时间复杂度: O(n * k log k)，n 为字符串个数，k 为字符串平均长度  空间复杂度: O(n * k)
func GroupAnagrams(strs []string) [][]string {
	// 关键观察：互为异位词的字符串，其排序后的结果完全相同
	// 用排序后的字符串作为分组 key，一次遍历完成分组
	groups := make(map[string][]string, len(strs))
	for _, s := range strs {
		key := sortedString(s)
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

// sortedString 返回字符串排序后的副本，作为异位词的统一 key
func sortedString(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	return string(bytes)
}
