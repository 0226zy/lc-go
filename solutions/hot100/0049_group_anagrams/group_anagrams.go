package groupanagrams

// GroupAnagrams 字母异位词分组
// 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
// 时间复杂度:   空间复杂度:
func GroupAnagrams(strs []string) [][]string {
	// 字母异位词，意味这字母相同。
	// 可以构建一个数组 [26]int，基于str 在数组中计数
	// 如果字符出现，则在对应的数组累加出现的次数
	// 如果两个str 是字母异位词，则对应的 [26]int 相同
	// 最后用 [26]int 做为map 的key 进行分类

	// key: [26]int{};value: []string{}
	dict := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		dict[cnt] = append(dict[cnt], str)
	}
	res := [][]string{}
	for _, strs := range dict {
		res = append(res, strs)
	}
	return res
}
