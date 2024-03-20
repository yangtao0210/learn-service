package examination

import (
	"fmt"
)

func SumOfContinuousInterval() {
	var n, x int
	fmt.Scan(&n, &x)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	//滑动窗口
	l, r, sum, count := 0, 0, 0, 0
	for r < n {
		sum += nums[r]
		for sum >= x {
			count += n - r
			sum -= nums[l]
			l++
		}
		r++
	}
	fmt.Println(count)
}
