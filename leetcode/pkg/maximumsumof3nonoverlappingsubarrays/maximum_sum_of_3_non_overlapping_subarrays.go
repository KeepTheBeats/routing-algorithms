package maximumsumof3nonoverlappingsubarrays

func MaxSumOfThreeSubarrays(nums []int, k int) []int {
	return maxSumOfThreeSubarrays(nums, k)
}

// 这个代码超时了。
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	var sums []int = make([]int, len(nums)-k+1)
	var slidingWindow int
	for i := 0; i < k; i++ {
		slidingWindow += nums[i] // 滑动窗口求i开头长度为k的子数组长度
	}
	sums[0] = slidingWindow

	for i := 1; k-1+i < len(nums); i++ {
		slidingWindow -= nums[i-1]
		slidingWindow += nums[k-1+i]
		sums[i] = slidingWindow
	}
	maxThreeSum := 0
	var result []int = make([]int, 3)
	threeSum := 0
	var tmpResult []int
	var findSum func(int, int)
	findSum = func(start int, picked int) {
		for i := start; i < len(sums); i++ {
			threeSum += sums[i]
			picked++
			tmpResult = append(tmpResult, i)
			if picked == 3 { // 已经取了3个子数组了
				if threeSum > maxThreeSum {
					maxThreeSum = threeSum
					copy(result, tmpResult)
				}
			} else { // 未取到3个子数组，继续取
				// 取了i之后，后面k-1个子数组也都不能用了，因为不能重叠
				findSum(i+k, picked)
			}
			tmpResult = tmpResult[:len(tmpResult)-1]
			picked--
			threeSum -= sums[i]
		}
	}
	findSum(0, 0)
	return result
}
