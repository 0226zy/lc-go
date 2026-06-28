package longestvalidparentheses

// LongestValidParentheses 最长有效括号
// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效括号子串的长度。
// 时间复杂度: O(n)  空间复杂度: O(n)
func LongestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	dp := make([]int, len(s))
	ret := 0

	for i := 1; i < len(s); i++ {
		if s[i] != ')' {
			continue // 只处理 ')'
		}

		if s[i-1] == '(' {
			// 情况1：...()
			if i >= 2 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}
		} else if dp[i-1] > 0 {
			// 情况2：...))
			//   x x x   ((  xxx  ))
			//      j-1  j         i
			j := i - dp[i-1] - 1
			if j >= 0 && s[j] == '(' {
				dp[i] = dp[i-1] + 2
				if j >= 1 {
					dp[i] += dp[j-1]
				}
			}
		}

		if dp[i] > ret {
			ret = dp[i]
		}
	}

	return ret
}
