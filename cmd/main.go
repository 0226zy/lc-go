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
	maximumsubarray "github.com/0226zy/lc-go/solutions/hot100/0053_maximum_subarray"
	mergeintervals "github.com/0226zy/lc-go/solutions/hot100/0056_merge_intervals"
	destroyingasteroids "github.com/0226zy/lc-go/solutions/daily/2126_destroying_asteroids"
	movezeroes "github.com/0226zy/lc-go/solutions/hot100/0283_move_zeroes"
	rotatearray "github.com/0226zy/lc-go/solutions/hot100/0189_rotate_array"
	productofarrayexceptself "github.com/0226zy/lc-go/solutions/hot100/0238_product_of_array_except_self"
	setmatrixzeroes "github.com/0226zy/lc-go/solutions/hot100/0073_set_matrix_zeroes"
	spiralmatrix "github.com/0226zy/lc-go/solutions/hot100/0054_spiral_matrix"
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
		{
			name:        "max-subarray",
			alias:       "53",
			description: "最大子数组和 (Maximum Subarray)",
			fn: func() {
				nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
				result := maximumsubarray.MaxSubArray(nums)
				fmt.Printf("MaxSubArray(%v) = %d\n", nums, result)
			},
		},
		{
			name:        "merge-intervals",
			alias:       "56",
			description: "合并区间 (Merge Intervals)",
			fn: func() {
				intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
				result := mergeintervals.Merge(intervals)
				fmt.Printf("Merge(%v) = %v\n", intervals, result)
			},
		},
		{
			name:        "destroying-asteroids",
			alias:       "2126",
			description: "摧毁小行星 (Destroying Asteroids)",
			fn: func() {
				mass := 10
				asteroids := []int{3, 9, 19, 5, 21}
				result := destroyingasteroids.AsteroidsDestroyed(mass, asteroids)
				fmt.Printf("AsteroidsDestroyed(%d, %v) = %v\n", mass, asteroids, result)
			},
		},
		{
			name:        "move-zeroes",
			alias:       "283",
			description: "移动零 (Move Zeroes)",
			fn: func() {
				nums := []int{0, 1, 0, 3, 12}
				movezeroes.MoveZeroes(nums)
				fmt.Printf("MoveZeroes([0,1,0,3,12]) = %v\n", nums)
			},
		},
		{
			name:        "rotate-array",
			alias:       "189",
			description: "旋转数组 (Rotate Array)",
			fn: func() {
				nums := []int{1, 2, 3, 4, 5, 6, 7}
				rotatearray.Rotate(nums, 3)
				fmt.Printf("Rotate([1,2,3,4,5,6,7], 3) = %v\n", nums)
			},
		},
		{
			name:        "product-except-self",
			alias:       "238",
			description: "除自身以外数组的乘积 (Product of Array Except Self)",
			fn: func() {
				nums := []int{1, 2, 3, 4}
				result := productofarrayexceptself.ProductExceptSelf(nums)
				fmt.Printf("ProductExceptSelf(%v) = %v\n", nums, result)
			},
		},
		{
			name:        "set-matrix-zeroes",
			alias:       "73",
			description: "矩阵置零 (Set Matrix Zeroes)",
			fn: func() {
				matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
				setmatrixzeroes.SetZeroes(matrix)
				fmt.Printf("SetZeroes([[1,1,1],[1,0,1],[1,1,1]]) = %v\n", matrix)
			},
		},
		{
			name:        "spiral-matrix",
			alias:       "54",
			description: "螺旋矩阵 (Spiral Matrix)",
			fn: func() {
				matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
				result := spiralmatrix.SpiralOrder(matrix)
				fmt.Printf("SpiralOrder([[1,2,3],[4,5,6],[7,8,9]]) = %v\n", result)
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
