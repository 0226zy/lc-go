package designsearchautocompletesystem

import (
	"sort"
	"strings"
)

// AutocompleteSystem 设计搜索自动补全系统
// 维护句子热度表与当前输入前缀，每输入一个字符返回以该前缀开头、热度最高的至多 3 个句子；
// 输入 '#' 时把当前句子记入历史并清空输入。
// 时间复杂度: input('#') 为 O(1)，其余为 O(n·L + k·log k)
// 空间复杂度: O(n·L)
type AutocompleteSystem struct {
	hot map[string]int // 句子 -> 热度
	cur string         // 当前已输入的前缀
}

// Constructor 用历史句子及其热度初始化系统
func Constructor(sentences []string, times []int) AutocompleteSystem {
	hot := make(map[string]int, len(sentences))
	for i, s := range sentences {
		hot[s] += times[i]
	}
	return AutocompleteSystem{hot: hot}
}

// Input 输入一个字符，返回自动补全结果；c 为 '#' 时结束当前输入并保存句子
func (as *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		as.hot[as.cur]++
		as.cur = ""
		return []string{}
	}
	as.cur += string(c)

	// 收集所有以当前输入为前缀的句子
	type candidate struct {
		sentence string
		times    int
	}
	cands := make([]candidate, 0)
	for s, t := range as.hot {
		if strings.HasPrefix(s, as.cur) {
			cands = append(cands, candidate{s, t})
		}
	}

	// 热度降序；热度相同按 ASCII 字典序升序
	sort.Slice(cands, func(i, j int) bool {
		if cands[i].times != cands[j].times {
			return cands[i].times > cands[j].times
		}
		return cands[i].sentence < cands[j].sentence
	})

	res := make([]string, 0, 3)
	for i := 0; i < len(cands) && i < 3; i++ {
		res = append(res, cands[i].sentence)
	}
	return res
}
