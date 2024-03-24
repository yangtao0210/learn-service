package interview

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)
	if n == 0 || nums[0] > 0 || nums[n-1] < 0 {
		return res
	}
	//排序 + 三指针解法
	for i := 0; i < n; i++ {
		//对第一个元素去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				//添加结果时统一对2，3元素去重
				//第二个元素去重
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				right--
				//第三个元素去重
				for right > left && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return res
}
