package examination

import (
	"fmt"
	"sort"
	"strings"
)

func PositionSmallestKAscii() {
	var str string
	fmt.Scan(&str)
	var k int
	fmt.Scan(&k)

	chars := []byte(str)
	//排序
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	if k >= len(chars) {
		fmt.Println(strings.Index(str, string(chars[len(chars)-1])))
	} else {
		fmt.Println(strings.Index(str, string(chars[k-1])))
	}
}
