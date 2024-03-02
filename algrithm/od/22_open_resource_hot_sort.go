package od

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type OpenHot struct {
	name string
	hot  int
}

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
	//计算每个开源项目的热度
	openHots := make([]OpenHot, n)
	for i := 0; i < len(datas); i++ {
		openHots[i].name = datas[i][0]
		for j := 1; j < len(datas[i]); j++ {
			hot, _ := strconv.Atoi(datas[i][j])
			openHots[i].hot += hot * weights[j-1]
		}
	}
	//按照规则排序
	sort.Slice(openHots, func(i, j int) bool {
		//热度逆序
		if openHots[i].hot > openHots[j].hot {
			return true
		} else if openHots[i].hot == openHots[j].hot {
			//热度相等时，按照字典顺序正序
			return strings.ToLower(openHots[i].name) < strings.ToLower(openHots[j].name)
		}
		return false
	})

	//输出结果
	for _, openHot := range openHots {
		fmt.Println(openHot.name)
	}
}
