package examination

import (
	"fmt"
)

func GetLuckyNumber() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	introducts := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&introducts[i])
	}
	maxPosition, position := 0, 0
	for _, intro := range introducts {
		position += intro
		if intro == m {
			if intro < 0 {
				position += -1
			} else {
				position += 1
			}
		}
		maxPosition = max(maxPosition, position)
	}
	fmt.Println(maxPosition)
}
