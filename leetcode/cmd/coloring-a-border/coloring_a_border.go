package main

import (
	"fmt"

	"r2tdsdn-configuration/leetcode/pkg/coloringaborder"
)

func main() {
	fmt.Println(coloringaborder.ColorBorder([][]int{{1, 1}, {1, 2}}, 0, 0, 3))
	fmt.Println(coloringaborder.ColorBorder([][]int{{1, 2, 2}, {2, 3, 2}}, 0, 1, 3))
	fmt.Println(coloringaborder.ColorBorder([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 1, 1, 2))
}
