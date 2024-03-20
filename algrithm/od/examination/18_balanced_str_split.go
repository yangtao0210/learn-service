package examination

import (
	"fmt"
	"strings"
)

func TestBalancedStringSplit() {
	var s string
	for {
		fmt.Scan(&s)
		if strings.EqualFold(s, "exit") {
			break
		}
		fmt.Println(BalancedStringSplit(s))
	}
}
func BalancedStringSplit(s string) int {
	res, count := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			count++
		} else {
			count--
		}
		if count == 0 {
			res++
		}
	}
	return res
}
