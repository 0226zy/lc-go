package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
	twosum "github.com/0226zy/lc-go/solutions/easy/0001_two_sum"
	addtwonumbers "github.com/0226zy/lc-go/solutions/hot100/0002_add_two_numbers"
	longestsubstring "github.com/0226zy/lc-go/solutions/hot100/0003_longest_substring_without_repeating_characters"
	medianoftwosortedarrays "github.com/0226zy/lc-go/solutions/hot100/0004_median_of_two_sorted_arrays"
	longestpalindromicsubstring "github.com/0226zy/lc-go/solutions/hot100/0005_longest_palindromic_substring"
	containerwithmostwater "github.com/0226zy/lc-go/solutions/hot100/0011_container_with_most_water"
	groupanagrams "github.com/0226zy/lc-go/solutions/hot100/0049_group_anagrams"
	longestincreasingsubsequence "github.com/0226zy/lc-go/solutions/hot100/0300_longest_increasing_subsequence"
	findallanagramsinastring "github.com/0226zy/lc-go/solutions/hot100/0438_find_all_anagrams_in_a_string"
	subarraysumequalk "github.com/0226zy/lc-go/solutions/hot100/0560_subarray_sum_equals_k"
)

type command struct {
	name        string
	alias       string
	description string
	fn          func()
}

func main() {
	commands := []command{
		{
			name:        "two-sum",
			alias:       "1",
			description: "两数之和 (Two Sum)",
			fn: func() {
				nums := []int{2, 7, 11, 15}
				target := 9
				result := twosum.TwoSum(nums, target)
				fmt.Printf("TwoSum(%v, %d) = ", nums, target)
				utils.PrintIntSlice(result)
			},
		},
		{
			name:        "add-two-numbers",
			alias:       "2",
			description: "两数相加 (Add Two Numbers)",
			fn: func() {
				l1 := datastructures.NewLinkedList([]int{2, 4, 3})
				l2 := datastructures.NewLinkedList([]int{5, 6, 4})
				result := addtwonumbers.AddTwoNumbers(l1, l2)
				fmt.Printf("AddTwoNumbers(%v, %v) = %v\n", l1.ToSlice(), l2.ToSlice(), result.ToSlice())
			},
		},
		{
			name:        "longest-substring",
			alias:       "3",
			description: "无重复字符的最长子串 (Longest Substring Without Repeating Characters)",
			fn: func() {
				s := "abcabcbb"
				result := longestsubstring.LengthOfLongestSubstring(s)
				fmt.Printf("LengthOfLongestSubstring(%q) = %d\n", s, result)
			},
		},
		{
			name:        "median-sorted-arrays",
			alias:       "4",
			description: "寻找两个正序数组的中位数 (Median of Two Sorted Arrays)",
			fn: func() {
				nums1 := []int{1, 3}
				nums2 := []int{2}
				result := medianoftwosortedarrays.FindMedianSortedArrays(nums1, nums2)
				fmt.Printf("FindMedianSortedArrays(%v, %v) = %.1f\n", nums1, nums2, result)
			},
		},
		{
			name:        "longest-palindrome",
			alias:       "5",
			description: "最长回文子串 (Longest Palindromic Substring)",
			fn: func() {
				s := "babad"
				result := longestpalindromicsubstring.LongestPalindrome(s)
				fmt.Printf("LongestPalindrome(%q) = %q\n", s, result)
			},
		},
		{
			name:        "group-anagrams",
			alias:       "49",
			description: "字母异位词分组 (Group Anagrams)",
			fn: func() {
				strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
				result := groupanagrams.GroupAnagrams(strs)
				fmt.Printf("GroupAnagrams(%v) = %v\n", strs, result)
			},
		},
		{
			name:        "container-with-most-water",
			alias:       "11",
			description: "盛最多水的容器 (Container With Most Water)",
			fn: func() {
				height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
				result := containerwithmostwater.MaxArea(height)
				fmt.Printf("MaxArea(%v) = %d\n", height, result)
			},
		},
		{
			name:        "longest-increasing-subsequence",
			alias:       "300",
			description: "最长递增子序列 (Longest Increasing Subsequence)",
			fn: func() {
				nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
				result := longestincreasingsubsequence.LengthOfLIS(nums)
				fmt.Printf("LengthOfLIS(%v) = %d\n", nums, result)
			},
		},
		{
			name:        "find-all-anagrams",
			alias:       "438",
			description: "找到字符串中所有字母异位词 (Find All Anagrams in a String)",
			fn: func() {
				s := "cbaebabacd"
				p := "abc"
				result := findallanagramsinastring.FindAnagrams(s, p)
				fmt.Printf("FindAnagrams(%q, %q) = ", s, p)
				utils.PrintIntSlice(result)
			},
		},
		{
			name:        "subarray-sum-equals-k",
			alias:       "560",
			description: "和为 K 的子数组 (Subarray Sum Equals K)",
			fn: func() {
				nums := []int{1, 1, 1}
				k := 2
				result := subarraysumequalk.SubarraySum(nums, k)
				fmt.Printf("SubarraySum(%v, %d) = %d\n", nums, k, result)
			},
		},
	}

	if len(os.Args) < 2 {
		printHelp(commands)
		os.Exit(0)
	}

	arg := strings.ToLower(os.Args[1])
	for _, cmd := range commands {
		if arg == cmd.name || arg == cmd.alias {
			cmd.fn()
			return
		}
	}

	fmt.Fprintf(os.Stderr, "未知命令: %s\n\n", os.Args[1])
	printHelp(commands)
	os.Exit(1)
}

func printHelp(commands []command) {
	fmt.Println("LeetCode Go 示例运行器")
	fmt.Println()
	fmt.Println("用法: go run ./cmd/main.go <命令>")
	fmt.Println()
	fmt.Println("可用命令:")
	for _, cmd := range commands {
		fmt.Printf("  %-22s (别名: %s)  %s\n", cmd.name, cmd.alias, cmd.description)
	}
}
