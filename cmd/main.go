package main

import (
	"fmt"

	"github.com/0226zy/lc-go/pkg/utils"
	"github.com/0226zy/lc-go/solutions/easy"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := easy.TwoSum(nums, target)
	fmt.Printf("TwoSum(%v, %d) = ", nums, target)
	utils.PrintIntSlice(result)
}
