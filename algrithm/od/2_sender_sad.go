package od

import (
	"container/list"
	"fmt"
	"math"
)

func SenderSad() {
	var n, m int
	fmt.Scan(&n, &m)
	//定义客户之间的距离矩阵&初始化，大小为n+1 * n+1
	distances := make([][]int, n+1)
	for i := range distances {
		distances[i] = make([]int, n+1)
		for j := range distances[i] {
			distances[i][j] = math.MaxInt
		}
	}
	//客户id到矩阵下标的映射
	customerIdMap := make(map[int]int)
	//根据输入数据更新快递站到各个客户之间的距离
	for i := 0; i < n; i++ {
		var cId, distance int
		fmt.Scan(&cId, &distance)
		customerIdMap[cId] = i + 1
		distances[0][i+1], distances[i+1][0] = distance, distance
	}
	//根据快递员收集客户数据，更新距离矩阵
	for i := 0; i < m; i++ {
		var cId_1, cId_2, distance int
		fmt.Scan(&cId_1, &cId_2, &distance)
		cId1Index, cId2Index := customerIdMap[cId_1], customerIdMap[cId_2]
		distances[cId1Index][cId2Index], distances[cId2Index][cId1Index] = distance, distance
	}

	//dp记录每个状态的最短距离
	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = 0

	//广度优先搜索
	queue := list.New()
	queue.PushBack([]int{0, 0, 0})
	for queue.Len() != 0 {
		cur := queue.Front().Value.([]int)
		queue.Remove(queue.Front())
		crow, ccol := cur[0], cur[1]
		dis := dp[crow][ccol]
		for j := 0; j <= n; j++ {
			if j != ccol && distances[ccol][j] != math.MaxInt {
				var i int
				if j != 0 {
					i = crow | 1<<(j-1)
				} else {
					i = crow
				}
				if dp[i][j] > dis+distances[ccol][j] {
					dp[i][j] = dis + distances[ccol][j]
					queue.PushBack([]int{i, j, 0})
				}
			}
		}
	}

	minDistance := dp[(1<<n)-1][0]
	if minDistance != math.MaxInt {
		fmt.Println(minDistance)
	} else {
		fmt.Println(-1)
	}
}
