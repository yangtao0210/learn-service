package od

import (
	"fmt"
	"sort"
)

func AssignCpuFromTwoGroups() {
	var l1, l2 int
	fmt.Scan(&l1, &l2)

	//输入A组数据并计算总算力
	cpuOfGroupA := make([]int, l1)
	totalServerGroupA := 0
	for i := 0; i < l1; i++ {
		fmt.Scan(&cpuOfGroupA[i])
		totalServerGroupA += cpuOfGroupA[i]
	}

	//输入A组数据并计算总算力
	cpuOfGroupB := make(map[int]int)
	totalServerGroupB := 0
	for i := 0; i < l2; i++ {
		var data int
		fmt.Scan(&data)
		totalServerGroupB += data
		cpuOfGroupB[data]++
	}

	sort.Ints(cpuOfGroupA)
	//计算总算力的一半：(a-b)/2 = x - y ==>遍历x,求对应的y
	diffCpu := (totalServerGroupA - totalServerGroupB) / 2
	//遍历A的算力:x
	for _, pA := range cpuOfGroupA {
		pB := pA - diffCpu
		//判断B中是否存在pB的算力值
		if val, ok := cpuOfGroupB[pB]; ok && val > 0 {
			fmt.Printf("%v %v", pA, pB)
			break
		}
	}
}
