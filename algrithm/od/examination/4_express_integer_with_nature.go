package examination

import (
	"fmt"
	"sort"
	"strconv"
)

func TestExpressIntegerWithNature() {
	for {
		var n int
		fmt.Scan(&n)
		if n == -1 {
			break
		}
		ExpressIntegerWithNature(n)
	}
}

func ExpressIntegerWithNature(target int) {
	if target <= 2 {
		fmt.Println(strconv.Itoa(target) + "=" + strconv.Itoa(target))
		fmt.Println("Result:1")
		return
	}
	res := make([]string, 0)
	res = append(res, strconv.Itoa(target)+"="+strconv.Itoa(target))
	for i := 1; i < target; i++ {
		sum := 0
		express := ""
		for j := i; j < target; j++ {
			sum += j
			express += strconv.Itoa(j) + "+"
			if sum == target {
				res = append(res, strconv.Itoa(target)+"="+express[:len(express)-1])
				break
			}
		}
	}
	//按照表达式长度递增排序
	sort.Slice(res, func(i, j int) bool {
		if len(res[i]) < len(res[j]) {
			return true
		}
		return false
	})

	for _, exp := range res {
		fmt.Println(exp)
	}
	fmt.Println("Result:" + strconv.Itoa(len(res)))
}
