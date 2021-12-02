package relativeranks

import (
	"fmt"
	"sort"
)

func findRelativeRanks(score []int) []string {
	n := len(score)
	var ans []string = make([]string, n)

	var scoreToIndex map[int]int = make(map[int]int)
	for i := 0; i < n; i++ {
		scoreToIndex[score[i]] = i // 记录分数和下标的对应关系
	}
	sort.Sort(sort.Reverse(sort.IntSlice(score))) // 分数降序排列

	for i := 0; i < n; i++ {
		var rank string
		switch i {
		case 0:
			rank = "Gold Medal"
		case 1:
			rank = "Silver Medal"
		case 2:
			rank = "Bronze Medal"
		default:
			rank = fmt.Sprintf("%d", i+1)
		}
		ans[scoreToIndex[score[i]]] = rank // 根据降序排列的结果和分数与下标的对应关系输出结果
	}

	return ans
}

func FindRelativeRanks(score []int) []string {
	return findRelativeRanks(score)
}
