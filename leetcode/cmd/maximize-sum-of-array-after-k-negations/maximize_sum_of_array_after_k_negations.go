package main

import (
	"fmt"

	"r2tdsdn-configuration/leetcode/pkg/maximizesumofarrayafterknegations"
)

func main() {
	fmt.Println(maximizesumofarrayafterknegations.LargestSumAfterKNegations([]int{4, 2, 3}, 1))
	fmt.Println(maximizesumofarrayafterknegations.LargestSumAfterKNegations([]int{3, -1, 0, 2}, 3))
	fmt.Println(maximizesumofarrayafterknegations.LargestSumAfterKNegations([]int{3, -1, 0, 2}, 3))
	fmt.Println(maximizesumofarrayafterknegations.LargestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
}
