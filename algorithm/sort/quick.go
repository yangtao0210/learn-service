package sort

func QuickSort(nums []int, begin, end int) {
	if begin < end {
		// 切分
		loc := partition(nums, begin, end)
		// 递归左排序
		QuickSort(nums, begin, loc-1)
		// 递归右排序
		QuickSort(nums, loc+1, end)
	}
}

// 切分元素：返回目标元素下标
func partition(nums []int, begin int, end int) int {
	// 选择第一个元素作为基准
	pivot := nums[begin]
	i := begin + 1
	j := end

	// 遍历比较：交换符合条件的元素
	for i < j {
		if nums[i] > pivot {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	// 遍历一遍后，比较最后遍历位置的元素与基准元素的大小，若大于等于，则需要前移一位，保证位于基准后的都大于等于其
	if nums[i] >= pivot {
		i--
	}
	nums[i], nums[begin] = nums[begin], nums[i]
	return i
}
