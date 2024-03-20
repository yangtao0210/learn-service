package examination

import (
	"fmt"
	"math"
)

func DrawingMachine() {
	//处理输入
	var n, e int
	fmt.Scan(&n, &e)
	offsetY := make([]int, e)
	for i := 0; i < n; i++ {
		var x, offset int
		fmt.Scan(&x, &offset)
		offsetY[x] = offset
	}

	if e == 0 {
		fmt.Println(0)
	}

	//动态规划：dp[i]:表示x坐标为i时的y偏移量,长度为e
	dp := make([]int, e)
	dp[0] = offsetY[0]
	for i := 1; i < e; i++ {
		dp[i] = offsetY[i] + dp[i-1]
	}

	area := 0
	for _, d := range dp {
		area += int(math.Abs(float64(d)))
	}
	fmt.Println(area)

}
