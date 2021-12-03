package maximizesumofarrayafterknegations

func LargestSumAfterKNegations(nums []int, k int) int {
	return largestSumAfterKNegations(nums, k)
}

// 每次都将最小的数取负，通过小根堆找最小的数
func largestSumAfterKNegations(nums []int, k int) int {
	buildHeap(nums)
	for i := 0; i < k; i++ {
		nums[0] = -nums[0]
		minHeap(nums, 0, len(nums)-1) // 取负之后进行小根堆的调整，保持最小的数在数组第1个
	}
	return sum(nums)
}

func minHeap(nums []int, start, end int) {
	var dad, son int
	dad = start
	son = 2*dad + 1
	for son <= end { // 若子节点指标在范围内才做比较
		if son+1 <= end && nums[son+1] < nums[son] {
			// 先比较两个子节点大小，选择最大的
			son++
		}
		if nums[dad] <= nums[son] { // 如果父节点大于子节点代表调整完毕，直接跳出函数
			return
		}
		// 否则交换父子内容再继续子节点和孙节点比较
		nums[dad], nums[son] = nums[son], nums[dad]
		dad = son
		son = 2*dad + 1
	}
	return
}

func buildHeap(nums []int) {
	// 初始化，i从最后一个父节点开始调整
	for i := len(nums)/2 - 1; i >= 0; i-- {
		minHeap(nums, i, len(nums)-1)
	}
}

func sum(nums []int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}
