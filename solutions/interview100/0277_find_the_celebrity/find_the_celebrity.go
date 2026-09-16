package findthecelebrity

// KnowsFunc 交互接口：判断 a 是否认识 b
// 由平台提供实现，解法只通过该接口获取认识关系。
type KnowsFunc func(a int, b int) bool

// FindCelebrity 搜寻名人
// 在 n 个人的聚会中找出名人：其他所有人都认识名人，而名人不认识任何人。
// 若不存在名人，返回 -1。
// 思路：先用 n-1 次询问淘汰出唯一候选人，再用 2(n-1) 次询问验证。
// 时间复杂度: O(n)  空间复杂度: O(1)
func FindCelebrity(n int, knows KnowsFunc) int {
	// 第一轮：淘汰赛。knows(cand, i) 为真则 cand 出局（名人不认识任何人），
	// 为假则 i 出局（名人都被所有人认识），每轮恰淘汰一人。
	cand := 0
	for i := 1; i < n; i++ {
		if knows(cand, i) {
			cand = i
		}
	}
	// 第二轮：验证候选人。名人必须不认识任何人、且被所有人认识。
	for i := 0; i < n; i++ {
		if i == cand {
			continue
		}
		if knows(cand, i) || !knows(i, cand) {
			return -1
		}
	}
	return cand
}
