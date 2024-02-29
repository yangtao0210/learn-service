package od

import (
	"fmt"
	"strconv"
)

func TaxiBilling() {
	var n int
	fmt.Scan(&n)
	//初始化实际费用
	cost := 0
	//遍历读取数字的每一位
	for _, c := range strconv.Itoa(n) {
		digit := c - '0'
		//如果当前位大于4，则--
		if digit > 4 {
			digit--
		}
		//计算实际费用
		cost = cost*9 + int(digit)
	}
	fmt.Println(cost)
}
