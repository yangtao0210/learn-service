package od

import "fmt"

func OpenResourceHotSort() {
	var n int
	//输入n
	fmt.Scan(&n)
	//输入权重
	weights := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&weights[i])
	}
	//输入项目数据
	datas := make([][6]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&datas[i][0], &datas[i][1], &datas[i][2], &datas[i][3], &datas[i][4], &datas[i][5])
	}
}
