package od

import (
	"fmt"
)

func CoverMaxArea() {
	//输入h,n
	var m, n int
	fmt.Scan(&m, &n)
	//输入m*n矩阵
	masks := make([][]int, m)
	for i := 0; i < m; i++ {
		masks[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&masks[i][j])
		}
	}
	//使用两个map记录每个标识覆盖范围的最大，最小行列下标
	minIndexMap := make(map[int][2]int)
	maxIndexMap := make(map[int][2]int)
	//遍历矩阵，更新最大最小下标
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			flag := masks[i][j]
			if flag != 0 {
				if _, ok := minIndexMap[flag]; ok {
					//如果标识已经存在
					min, max := minIndexMap[flag], maxIndexMap[flag]
					minIndexMap[flag] = [2]int{minVal(min[0], i), minVal(min[1], j)}
					maxIndexMap[flag] = [2]int{maxVal(max[0], i), maxVal(max[1], j)}
				} else {
					//标识不存在
					minIndexMap[flag] = [2]int{i, j}
					maxIndexMap[flag] = [2]int{i, j}
				}
			}
		}
	}
	//遍历矩阵计算最大覆盖面积
	maxArea := 0
	fmt.Println(minIndexMap, maxIndexMap)
	for k := range minIndexMap {
		minPos, maxPos := minIndexMap[k], maxIndexMap[k]
		maxArea = maxVal(maxArea, (maxPos[0]-minPos[0]+1)*(maxPos[1]-minPos[1]+1))
	}
	fmt.Println(maxArea)
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxVal(a, b int) int {
	if a < b {
		return b
	}
	return a
}
