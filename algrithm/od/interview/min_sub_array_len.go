package interview

import "math"

func minSubArrayLen(target int, nums []int) int {
	left, right, sum, minLen := 0, 0, 0, math.MaxInt
	//滑动窗口解法
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			//情况1:当前窗口和大于等于目标值：右边界不变，左边界收缩 & 更新结果 & 更新sum
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
		//当前窗口和小于目标值：左边界不变，右边界伸展
		right++
	}
	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}
