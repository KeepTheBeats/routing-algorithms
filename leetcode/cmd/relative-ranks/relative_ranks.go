package main

import (
	"fmt"

	"r2tdsdn-configuration/leetcode/pkg/relativeranks"
)

func main() {
	fmt.Println(relativeranks.FindRelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Println(relativeranks.FindRelativeRanks([]int{10, 3, 8, 9, 4}))
}
