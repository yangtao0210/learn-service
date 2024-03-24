package tanxin

import (
	"fmt"
	"sort"
)

//问题描述
// 某公司,每天早上都有很多人去坐电梯,每个人都可能到不同的楼层.同时电梯还有一个容量限制.电梯最多只能带K个人.
// 电梯从第a层到第b层,会花费|a-b|的时间.现在有N个人,以及知道每个人想要去的地方,
// 请问如何坐电梯,才能使每个人到达到他们对应的楼层,且所花费时间最少.电梯最后要会到第1层

func MinTime() int {
	//1.处理输入
	var N, K int
	fmt.Scan(&N, &K)
	floors := make([]int, N)
	for i := range floors {
		fmt.Scan(&floors[i])
	}
	sort.Ints(floors)

	//解法1：贪心
	time := 0
	if K >= N {
		return 2 * (floors[len(floors)-1] - 1)
	}
	for i := N - 1; i >= 0; i -= K {
		index := Max(0, i)
		fmt.Println("下标：", index)
		time += 2 * (floors[index] - 1)
		fmt.Println("时间", time)
	}
	return time
	// 解法2：动态规划
	// 1. dp[i][j]:表示前i个人都已经到达目标层，且电梯当前在第j层时花费的最短时间
	//dp := make([][]int, N+1)
	//for i := range dp {
	//	dp[i] = make([]int, N+1)
	//	for j := range dp[i] {
	//		dp[i][j] = math.MaxInt32
	//	}
	//}
	//dp[0][0] = 0
	//for i := 1; i <= N; i++ {
	//	for j := 0; j < i; j++ {
	//		for k := 0; k <= j; k++ {
	//			dp[i][j] = min(dp[i][j], dp[j][k]+abs(floors[i-1]-floors[k]))
	//		}
	//		if i < N {
	//			dp[i+1][i] = min(dp[i+1][i], dp[i][j]+abs(floors[i-1]-floors[i]))
	//		}
	//	}
	//}
	//
	//minVal := math.MaxInt32
	//for i := 0; i < N; i++ {
	//	minVal = min(minVal, dp[N][i]+floors[i]-1)
	//}
	//
	//return minVal
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
