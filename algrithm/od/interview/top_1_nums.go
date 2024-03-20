package interview

func Top1Nums(nums []int) int {
	countMap := make(map[int]int)
	maxCount, maxNum := 0, -1<<31
	for _, num := range nums {
		countMap[num]++
		if countMap[num] > maxCount {
			maxCount = countMap[num]
			maxNum = num
		} else if countMap[num] == maxCount && num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}
