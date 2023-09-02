package algorithm

import "fmt"

var Total = 0

// Tower 汉罗塔问题
func Tower(n int, a, b, c string) {
	if n == 1 {
		Total = Total + 1
		fmt.Println(a, "->", c)
		return
	}
	// 前n-1个移动到b
	Tower(n-1, a, c, b)
	Total = Total + 1
	fmt.Println(a, "->", c)
	Tower(n-1, b, a, c)
}
