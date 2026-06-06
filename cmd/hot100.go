package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/0226zy/lc-go/pkg/datastructures"
	"github.com/0226zy/lc-go/pkg/utils"
	addtwonumbers "github.com/0226zy/lc-go/solutions/hot100/0002_add_two_numbers"
	longestsubstring "github.com/0226zy/lc-go/solutions/hot100/0003_longest_substring_without_repeating_characters"
	medianoftwosortedarrays "github.com/0226zy/lc-go/solutions/hot100/0004_median_of_two_sorted_arrays"
	longestpalindromicsubstring "github.com/0226zy/lc-go/solutions/hot100/0005_longest_palindromic_substring"
	containerwithmostwater "github.com/0226zy/lc-go/solutions/hot100/0011_container_with_most_water"
	removenthnodefromendoflist "github.com/0226zy/lc-go/solutions/hot100/0019_remove_nth_node_from_end_of_list"
	groupanagrams "github.com/0226zy/lc-go/solutions/hot100/0049_group_anagrams"
	maximumsubarray "github.com/0226zy/lc-go/solutions/hot100/0053_maximum_subarray"
	spiralmatrix "github.com/0226zy/lc-go/solutions/hot100/0054_spiral_matrix"
	mergeintervals "github.com/0226zy/lc-go/solutions/hot100/0056_merge_intervals"
	setmatrixzeroes "github.com/0226zy/lc-go/solutions/hot100/0073_set_matrix_zeroes"
	rotatearray "github.com/0226zy/lc-go/solutions/hot100/0189_rotate_array"
	productofarrayexceptself "github.com/0226zy/lc-go/solutions/hot100/0238_product_of_array_except_self"
	movezeroes "github.com/0226zy/lc-go/solutions/hot100/0283_move_zeroes"
	longestincreasingsubsequence "github.com/0226zy/lc-go/solutions/hot100/0300_longest_increasing_subsequence"
	findallanagramsinastring "github.com/0226zy/lc-go/solutions/hot100/0438_find_all_anagrams_in_a_string"
	subarraysumequalk "github.com/0226zy/lc-go/solutions/hot100/0560_subarray_sum_equals_k"
	intersectionoftwolinkedlists "github.com/0226zy/lc-go/solutions/hot100/0160_intersection_of_two_linked_lists"
	reverselinkedlist "github.com/0226zy/lc-go/solutions/hot100/0206_reverse_linked_list"
	palindromelinkedlist "github.com/0226zy/lc-go/solutions/hot100/0234_palindrome_linked_list"
	linkedlistcycle "github.com/0226zy/lc-go/solutions/hot100/0141_link_ed_list_cycle"
	linkedlistcycleii "github.com/0226zy/lc-go/solutions/hot100/0142_link_ed_list_cycle_ii"
	mergetwosortedlists "github.com/0226zy/lc-go/solutions/hot100/0021_merge_two_sorted_lists"
)

func init() {
	rootCmd.AddCommand(
		addTwoNumbersCmd,
		longestSubstringCmd,
		medianSortedArraysCmd,
		longestPalindromeCmd,
		containerWithMostWaterCmd,
		removeNthNodeFromEndOfListCmd,
		groupAnagramsCmd,
		maxSubArrayCmd,
		spiralMatrixCmd,
		mergeIntervalsCmd,
		setMatrixZeroesCmd,
		rotateArrayCmd,
		productExceptSelfCmd,
		moveZeroesCmd,
		longestIncreasingSubsequenceCmd,
		findAllAnagramsCmd,
		subarraySumEqualsKCmd,
		getIntersectionNodeCmd,
		reverseListCmd,
		isPalindromeCmd,
		hasCycleCmd,
		detectCycleCmd,
		mergeTwoListsCmd,
	)
}

var addTwoNumbersCmd = &cobra.Command{
	Use:     "add-two-numbers",
	Short:   "两数相加 (Add Two Numbers)",
	Aliases: []string{"2"},
	Run: func(cmd *cobra.Command, args []string) {
		l1 := datastructures.NewLinkedList([]int{2, 4, 3})
		l2 := datastructures.NewLinkedList([]int{5, 6, 4})
		result := addtwonumbers.AddTwoNumbers(l1, l2)
		fmt.Printf("AddTwoNumbers(%v, %v) = %v\n", l1.ToSlice(), l2.ToSlice(), result.ToSlice())
	},
}

var longestSubstringCmd = &cobra.Command{
	Use:     "longest-substring",
	Short:   "无重复字符的最长子串 (Longest Substring Without Repeating Characters)",
	Aliases: []string{"3"},
	Run: func(cmd *cobra.Command, args []string) {
		s := "abcabcbb"
		result := longestsubstring.LengthOfLongestSubstring(s)
		fmt.Printf("LengthOfLongestSubstring(%q) = %d\n", s, result)
	},
}

var medianSortedArraysCmd = &cobra.Command{
	Use:     "median-sorted-arrays",
	Short:   "寻找两个正序数组的中位数 (Median of Two Sorted Arrays)",
	Aliases: []string{"4"},
	Run: func(cmd *cobra.Command, args []string) {
		nums1 := []int{1, 3}
		nums2 := []int{2}
		result := medianoftwosortedarrays.FindMedianSortedArrays(nums1, nums2)
		fmt.Printf("FindMedianSortedArrays(%v, %v) = %.1f\n", nums1, nums2, result)
	},
}

var longestPalindromeCmd = &cobra.Command{
	Use:     "longest-palindrome",
	Short:   "最长回文子串 (Longest Palindromic Substring)",
	Aliases: []string{"5"},
	Run: func(cmd *cobra.Command, args []string) {
		s := "babad"
		result := longestpalindromicsubstring.LongestPalindrome(s)
		fmt.Printf("LongestPalindrome(%q) = %q\n", s, result)
	},
}

var containerWithMostWaterCmd = &cobra.Command{
	Use:     "container-with-most-water",
	Short:   "盛最多水的容器 (Container With Most Water)",
	Aliases: []string{"11"},
	Run: func(cmd *cobra.Command, args []string) {
		height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		result := containerwithmostwater.MaxArea(height)
		fmt.Printf("MaxArea(%v) = %d\n", height, result)
	},
}

var removeNthNodeFromEndOfListCmd = &cobra.Command{
	Use:     "remove-nth-node-from-end-of-list",
	Short:   "删除链表的倒数第 N 个结点 (Remove Nth Node From End of List)",
	Aliases: []string{"19"},
	Run: func(cmd *cobra.Command, args []string) {
		head := datastructures.NewLinkedList([]int{1, 2, 3, 4, 5})
		n := 2
		result := removenthnodefromendoflist.RemoveNthNodeFromEndOfList(head, n)
		fmt.Printf("RemoveNthNodeFromEndOfList([1,2,3,4,5], %d) = %v\n", n, result.ToSlice())
	},
}

var groupAnagramsCmd = &cobra.Command{
	Use:     "group-anagrams",
	Short:   "字母异位词分组 (Group Anagrams)",
	Aliases: []string{"49"},
	Run: func(cmd *cobra.Command, args []string) {
		strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		result := groupanagrams.GroupAnagrams(strs)
		fmt.Printf("GroupAnagrams(%v) = %v\n", strs, result)
	},
}

var maxSubArrayCmd = &cobra.Command{
	Use:     "max-subarray",
	Short:   "最大子数组和 (Maximum Subarray)",
	Aliases: []string{"53"},
	Run: func(cmd *cobra.Command, args []string) {
		nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		result := maximumsubarray.MaxSubArray(nums)
		fmt.Printf("MaxSubArray(%v) = %d\n", nums, result)
	},
}

var spiralMatrixCmd = &cobra.Command{
	Use:     "spiral-matrix",
	Short:   "螺旋矩阵 (Spiral Matrix)",
	Aliases: []string{"54"},
	Run: func(cmd *cobra.Command, args []string) {
		matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		result := spiralmatrix.SpiralOrder(matrix)
		fmt.Printf("SpiralOrder([[1,2,3],[4,5,6],[7,8,9]]) = %v\n", result)
	},
}

var mergeIntervalsCmd = &cobra.Command{
	Use:     "merge-intervals",
	Short:   "合并区间 (Merge Intervals)",
	Aliases: []string{"56"},
	Run: func(cmd *cobra.Command, args []string) {
		intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
		result := mergeintervals.Merge(intervals)
		fmt.Printf("Merge(%v) = %v\n", intervals, result)
	},
}

var setMatrixZeroesCmd = &cobra.Command{
	Use:     "set-matrix-zeroes",
	Short:   "矩阵置零 (Set Matrix Zeroes)",
	Aliases: []string{"73"},
	Run: func(cmd *cobra.Command, args []string) {
		matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
		setmatrixzeroes.SetZeroes(matrix)
		fmt.Printf("SetZeroes([[1,1,1],[1,0,1],[1,1,1]]) = %v\n", matrix)
	},
}

var rotateArrayCmd = &cobra.Command{
	Use:     "rotate-array",
	Short:   "旋转数组 (Rotate Array)",
	Aliases: []string{"189"},
	Run: func(cmd *cobra.Command, args []string) {
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		rotatearray.Rotate(nums, 3)
		fmt.Printf("Rotate([1,2,3,4,5,6,7], 3) = %v\n", nums)
	},
}

var productExceptSelfCmd = &cobra.Command{
	Use:     "product-except-self",
	Short:   "除自身以外数组的乘积 (Product of Array Except Self)",
	Aliases: []string{"238"},
	Run: func(cmd *cobra.Command, args []string) {
		nums := []int{1, 2, 3, 4}
		result := productofarrayexceptself.ProductExceptSelf(nums)
		fmt.Printf("ProductExceptSelf(%v) = %v\n", nums, result)
	},
}

var moveZeroesCmd = &cobra.Command{
	Use:     "move-zeroes",
	Short:   "移动零 (Move Zeroes)",
	Aliases: []string{"283"},
	Run: func(cmd *cobra.Command, args []string) {
		nums := []int{0, 1, 0, 3, 12}
		movezeroes.MoveZeroes(nums)
		fmt.Printf("MoveZeroes([0,1,0,3,12]) = %v\n", nums)
	},
}

var longestIncreasingSubsequenceCmd = &cobra.Command{
	Use:     "longest-increasing-subsequence",
	Short:   "最长递增子序列 (Longest Increasing Subsequence)",
	Aliases: []string{"300"},
	Run: func(cmd *cobra.Command, args []string) {
		nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
		result := longestincreasingsubsequence.LengthOfLIS(nums)
		fmt.Printf("LengthOfLIS(%v) = %d\n", nums, result)
	},
}

var findAllAnagramsCmd = &cobra.Command{
	Use:     "find-all-anagrams",
	Short:   "找到字符串中所有字母异位词 (Find All Anagrams in a String)",
	Aliases: []string{"438"},
	Run: func(cmd *cobra.Command, args []string) {
		s := "cbaebabacd"
		p := "abc"
		result := findallanagramsinastring.FindAnagrams(s, p)
		fmt.Printf("FindAnagrams(%q, %q) = ", s, p)
		utils.PrintIntSlice(result)
	},
}

var subarraySumEqualsKCmd = &cobra.Command{
	Use:     "subarray-sum-equals-k",
	Short:   "和为 K 的子数组 (Subarray Sum Equals K)",
	Aliases: []string{"560"},
	Run: func(cmd *cobra.Command, args []string) {
		nums := []int{1, 1, 1}
		k := 2
		result := subarraysumequalk.SubarraySum(nums, k)
		fmt.Printf("SubarraySum(%v, %d) = %d\n", nums, k, result)
	},
}

var getIntersectionNodeCmd = &cobra.Command{
	Use:     "get-intersection-node",
	Short:   "相交链表 (Intersection of Two Linked Lists)",
	Aliases: []string{"160"},
	Run: func(cmd *cobra.Command, args []string) {
		// 创建相交链表: listA = [4,1,8,4,5], listB = [5,6,1,8,4,5]
		common := &datastructures.ListNode{Val: 8}
		common.Next = &datastructures.ListNode{Val: 4}
		common.Next.Next = &datastructures.ListNode{Val: 5}

		headA := &datastructures.ListNode{Val: 4}
		headA.Next = &datastructures.ListNode{Val: 1}
		headA.Next.Next = common

		headB := &datastructures.ListNode{Val: 5}
		headB.Next = &datastructures.ListNode{Val: 6}
		headB.Next.Next = &datastructures.ListNode{Val: 1}
		headB.Next.Next.Next = common

		result := intersectionoftwolinkedlists.GetIntersectionNode(headA, headB)
		if result != nil {
			fmt.Printf("GetIntersectionNode([4,1,8,4,5], [5,6,1,8,4,5]) = %d\n", result.Val)
		} else {
			fmt.Println("GetIntersectionNode() = nil")
		}
	},
}

var reverseListCmd = &cobra.Command{
	Use:     "reverse-list",
	Short:   "反转链表 (Reverse Linked List)",
	Aliases: []string{"206"},
	Run: func(cmd *cobra.Command, args []string) {
		head := datastructures.NewLinkedList([]int{1, 2, 3, 4, 5})
		fmt.Printf("原始链表: %s\n", head)
		
		result := reverselinkedlist.ReverseList(head)
		fmt.Printf("反转后: %s\n", result)
	},
}

var isPalindromeCmd = &cobra.Command{
	Use:     "is-palindrome",
	Short:   "回文链表 (Palindrome Linked List)",
	Aliases: []string{"234"},
	Run: func(cmd *cobra.Command, args []string) {
		// 测试示例1: [1,2,2,1]
		head1 := datastructures.NewLinkedList([]int{1, 2, 2, 1})
		result1 := palindromelinkedlist.IsPalindrome(head1)
		fmt.Printf("IsPalindrome([1,2,2,1]) = %v\n", result1)

		// 测试示例2: [1,2]
		head2 := datastructures.NewLinkedList([]int{1, 2})
		result2 := palindromelinkedlist.IsPalindrome(head2)
		fmt.Printf("IsPalindrome([1,2]) = %v\n", result2)
	},
}

var hasCycleCmd = &cobra.Command{
	Use:     "has-cycle",
	Short:   "环形链表 (Linked List Cycle)",
	Aliases: []string{"141"},
	Run: func(cmd *cobra.Command, args []string) {
		// 测试示例1: [3,2,0,-4] 有环，pos=1
		head1 := datastructures.NewCycleLinkedList([]int{3, 2, 0, -4}, 1)
		result1 := linkedlistcycle.HasCycleTwoPointers(head1)
		fmt.Printf("HasCycle([3,2,0,-4], pos=1) = %v\n", result1)

		// 测试示例2: [1,2] 有环，pos=0
		head2 := datastructures.NewCycleLinkedList([]int{1, 2}, 0)
		result2 := linkedlistcycle.HasCycleTwoPointers(head2)
		fmt.Printf("HasCycle([1,2], pos=0) = %v\n", result2)

		// 测试示例3: [1] 无环，pos=-1
		head3 := datastructures.NewCycleLinkedList([]int{1}, -1)
		result3 := linkedlistcycle.HasCycleTwoPointers(head3)
		fmt.Printf("HasCycle([1], pos=-1) = %v\n", result3)
	},
}

	var detectCycleCmd = &cobra.Command{
		Use:     "detect-cycle",
		Short:   "环形链表 II (Linked List Cycle II)",
		Aliases: []string{"142"},
		Run: func(cmd *cobra.Command, args []string) {
			// 测试示例1: [3,2,0,-4] 有环，pos=1
			head1 := datastructures.NewCycleLinkedList([]int{3, 2, 0, -4}, 1)
			result1 := linkedlistcycleii.DetectCycleTwoPointers(head1)
			if result1 != nil {
				fmt.Printf("DetectCycle([3,2,0,-4], pos=1) = node with val %d\n", result1.Val)
			} else {
				fmt.Println("DetectCycle([3,2,0,-4], pos=1) = nil")
			}

			// 测试示例2: [1,2] 有环，pos=0
			head2 := datastructures.NewCycleLinkedList([]int{1, 2}, 0)
			result2 := linkedlistcycleii.DetectCycleTwoPointers(head2)
			if result2 != nil {
				fmt.Printf("DetectCycle([1,2], pos=0) = node with val %d\n", result2.Val)
			} else {
				fmt.Println("DetectCycle([1,2], pos=0) = nil")
			}

			// 测试示例3: [1] 无环，pos=-1
			head3 := datastructures.NewCycleLinkedList([]int{1}, -1)
			result3 := linkedlistcycleii.DetectCycleTwoPointers(head3)
			if result3 != nil {
				fmt.Printf("DetectCycle([1], pos=-1) = node with val %d\n", result3.Val)
			} else {
				fmt.Println("DetectCycle([1], pos=-1) = nil")
			}
		},
	}

var mergeTwoListsCmd = &cobra.Command{
	Use:     "merge-two-lists",
	Short:   "合并两个有序链表 (Merge Two Sorted Lists)",
	Aliases: []string{"21"},
	Run: func(cmd *cobra.Command, args []string) {
		// 测试示例1: [1,2,4] 和 [1,3,4]
		list1 := datastructures.NewLinkedList([]int{1, 2, 4})
		list2 := datastructures.NewLinkedList([]int{1, 3, 4})
		result := mergetwosortedlists.MergeTwoListsIterative(list1, list2)
		fmt.Printf("MergeTwoListsIterative([1,2,4], [1,3,4]) = %v\n", result.ToSlice())

		// 测试递归法
		list3 := datastructures.NewLinkedList([]int{1, 2, 4})
		list4 := datastructures.NewLinkedList([]int{1, 3, 4})
		result2 := mergetwosortedlists.MergeTwoListsRecursive(list3, list4)
		fmt.Printf("MergeTwoListsRecursive([1,2,4], [1,3,4]) = %v\n", result2.ToSlice())
	},
}
