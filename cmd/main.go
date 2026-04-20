package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
	"github.com/0226zy/lc-go/solutions/easy"
	"github.com/0226zy/lc-go/solutions/hot100/0002_add_two_numbers"
	"github.com/0226zy/lc-go/solutions/hot100/0003_longest_substring_without_repeating_characters"
	"github.com/0226zy/lc-go/solutions/hot100/0004_median_of_two_sorted_arrays"
	"github.com/0226zy/lc-go/solutions/hot100/0005_longest_palindromic_substring"
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
				result := easy.TwoSum(nums, target)
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
