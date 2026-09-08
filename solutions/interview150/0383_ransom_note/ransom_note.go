package ransomnote

// CanConstruct 赎金信
// 判断 ransomNote 是否能由 magazine 中的字符构成，magazine 中每个字符只能用一次。
// 时间复杂度: O(m + n)  空间复杂度: O(1)（固定 26 个小写字母的计数数组）
func CanConstruct(ransomNote string, magazine string) bool {
	// 统计 magazine 中每个小写字母出现的次数
	count := [26]int{}
	for i := 0; i < len(magazine); i++ {
		count[magazine[i]-'a']++
	}
	// 消耗 ransomNote 中的字符，出现负值说明该字符不够用
	for i := 0; i < len(ransomNote); i++ {
		count[ransomNote[i]-'a']--
		if count[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
