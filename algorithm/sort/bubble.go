package sort

func BubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 1; i > 0; i-- {
		// 优化：添加交换标志：若一轮过后交换标志仍为false，则表示数组本身有序
		swap := false
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}
