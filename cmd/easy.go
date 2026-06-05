package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/0226zy/lc-go/pkg/utils"
	twosum "github.com/0226zy/lc-go/solutions/easy/0001_two_sum"
)

func init() {
	rootCmd.AddCommand(twoSumCmd)
}

var twoSumCmd = &cobra.Command{
	Use:     "two-sum",
	Short:   "两数之和 (Two Sum)",
	Aliases: []string{"1"},
	Run: func(cmd *cobra.Command, args []string) {
		nums := []int{2, 7, 11, 15}
		target := 9
		result := twosum.TwoSum(nums, target)
		fmt.Printf("TwoSum(%v, %d) = ", nums, target)
		utils.PrintIntSlice(result)
	},
}
