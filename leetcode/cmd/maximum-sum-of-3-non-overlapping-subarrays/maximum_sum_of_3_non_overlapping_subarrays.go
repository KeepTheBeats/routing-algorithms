package main

import (
	"fmt"

	"r2tdsdn-configuration/leetcode/pkg/maximumsumof3nonoverlappingsubarrays"
)

func main() {
	fmt.Println(maximumsumof3nonoverlappingsubarrays.MaxSumOfThreeSubarrays([]int{1, 2, 1, 2, 6, 7, 5, 1}, 2))
	fmt.Println(maximumsumof3nonoverlappingsubarrays.MaxSumOfThreeSubarrays([]int{1, 2, 1, 2, 1, 2, 1, 2, 1}, 2))
}
