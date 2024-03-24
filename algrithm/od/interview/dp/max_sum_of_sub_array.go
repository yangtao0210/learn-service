package dp

import (
	"learn_code/algrithm/od/interview/tanxin"
)

//题目描述
//给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//子数组是数组中的一个连续部分。

func GetMaxSumOfSubArray(nums []int) int {
	maxSum := nums[0]
	//动态规划
	//dp[i]:表示前i个元素的连续子数组最大和,结果为计算所有位置的dp[i],取全局最大值
	// 递推公式:dp[i]的值可以由dp[i-1]和当前位置的元素值推出来,即dp[i] = max(dp[i],dp[i]+dp[i-1])
	//1）接受dp[i-1]:表示前i-1个元素的子数组最大和对于当前连续子数组和有正向左右，将dp[i-1]加入
	//2)不接受dp[i-1]：表示前i-1个元素的子数组和对当前位置有反向作用，不予接纳（加了还不如不加）
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = nums[i]
	}
	for i := 1; i < len(nums); i++ {
		dp[i] = tanxin.Max(dp[i], dp[i]+dp[i-1])
		maxSum = tanxin.Max(dp[i], maxSum)
	}
	return maxSum
}
