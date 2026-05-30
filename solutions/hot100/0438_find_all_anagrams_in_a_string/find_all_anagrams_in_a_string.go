package findallanagramsinastring

func FindAnagrams(s string, p string) []int {
	if len(s) < len(p) || len(p) == 0 {
		return []int{}
	}

	result := []int{}
	pCount := make([]int, 26)
	sCount := make([]int, 26)

	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}

	if matches(pCount, sCount) {
		result = append(result, 0)
	}

	for i := len(p); i < len(s); i++ {
		sCount[s[i]-'a']++
		sCount[s[i-len(p)]-'a']--

		if matches(pCount, sCount) {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}

func matches(pCount, sCount []int) bool {
	for i := 0; i < 26; i++ {
		if pCount[i] != sCount[i] {
			return false
		}
	}
	return true
}
