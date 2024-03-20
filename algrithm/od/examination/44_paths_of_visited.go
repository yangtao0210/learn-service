package examination

import "fmt"

func PathsOfVisited() {
	var l, w int
	fmt.Scan(&l, &w)
	datas := make([][]int, l)
	for i := 0; i < l; i++ {
		datas[i] = make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Scan(&datas[i][j])
		}
	}

	fmt.Println(getPathsCount(datas, l, w))
}

func getPathsCount(datas [][]int, l, w int) int {
	//动态规划解决 dp[i][j]表示从[0,0]到i,j的路径数量
	dp := make([][]int, l)
	//初始化：递推公式
	//datas[i][j] == 1:dp[i][j] = 0
	//datas[i][j] != 1:dp[i][j] = dp[i][j-1] + dp[i-1][j]
	for i := 0; i < l; i++ {
		dp[i] = make([]int, w)
		if datas[i][0] != 1 {
			dp[i][0] = 1
		}
		for j := 0; j < w; j++ {
			if datas[0][j] != 1 {
				dp[0][j] = 1
			}
		}
	}
	//dp过程处理
	for i := 1; i < l; i++ {
		for j := 1; j < w; j++ {
			if datas[i][j] != 1 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[l-1][w-1]
}
