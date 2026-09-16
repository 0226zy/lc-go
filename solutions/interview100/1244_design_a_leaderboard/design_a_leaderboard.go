package designaleaderboard

import "sort"

// Leaderboard 力扣排行榜
// 维护参赛者到总分的映射，支持加分、重置和查询前 K 名分数之和。
type Leaderboard struct {
	scores map[int]int // playerId -> 总分
}

// Constructor 创建空的排行榜
func Constructor() Leaderboard {
	return Leaderboard{scores: make(map[int]int)}
}

// AddScore 将参赛者的分数增加 score；不在排行榜中则以 score 作为初始分加入
// 时间复杂度: O(1)
func (l *Leaderboard) AddScore(playerId int, score int) {
	l.scores[playerId] += score
}

// Top 返回当前排行榜上前 K 名参赛者的分数之和
// 时间复杂度: O(n log n)，n 为参赛者人数  空间复杂度: O(n)
func (l *Leaderboard) Top(K int) int {
	all := make([]int, 0, len(l.scores))
	for _, s := range l.scores {
		all = append(all, s)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(all)))
	sum := 0
	for i := 0; i < K && i < len(all); i++ {
		sum += all[i]
	}
	return sum
}

// Reset 重置参赛者分数为 0 并将其移出排行榜
// 时间复杂度: O(1)
func (l *Leaderboard) Reset(playerId int) {
	delete(l.scores, playerId)
}
