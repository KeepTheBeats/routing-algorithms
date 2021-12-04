package main

import (
	"fmt"

	"r2tdsdn-configuration/leetcode/pkg/ransomnote"
)

func main() {
	fmt.Println(ransomnote.CanConstruct("a", "b"))
	fmt.Println(ransomnote.CanConstruct("aa", "ab"))
	fmt.Println(ransomnote.CanConstruct("aa", "aab"))
}
