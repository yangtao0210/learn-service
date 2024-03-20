package examination

import "fmt"

func MaxCountGems() {
	var n, v int
	fmt.Scan(&n)
	gems := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&gems[i])
	}
	fmt.Scan(&v)

	fmt.Println(getMaxCount(gems, v))

}

func getMaxCount(gems []int, val int) int {
	//滑动窗口实现
	maxCount, l, r, sum := 0, 0, 0, 0
	for r < len(gems) {
		sum += gems[r]
		for sum > val {
			sum -= gems[l]
			l++
		}
		maxCount = max(maxCount, r-l+1)
		r++
	}
	return maxCount
}
