package examination

import (
	"fmt"
	"math"
	"sort"
)

func MaxWeightRemainSilver() {
	//输入n
	var n int
	fmt.Scan(&n)
	//输入n个元素
	weights := make([]int, 0)
	for i := 0; i < n; i++ {
		var elem int
		fmt.Scan(&elem)
		weights = append(weights, elem)
	}
	sort.Ints(weights)
	for len(weights) >= 3 {
		length := len(weights)
		x, y, z := weights[length-3], weights[length-2], weights[length-1]
		weights = weights[:length-3]
		remain := 0
		//按照规则融化银饰
		if x == y && y == z {
			continue
		} else if x == y && y != z {
			remain = z - y
		} else if x != y && y == z {
			remain = y - x
		} else {
			remain = int(math.Abs(float64(z - y - y + x)))
		}
		if remain != 0 {
			weights = append(weights, remain)
			sort.Ints(weights)
		}
	}
	if len(weights) == 2 {
		fmt.Println(weights[1])
	} else if len(weights) == 1 {
		fmt.Println(weights[0])
	} else {
		fmt.Println(0)
	}
}
