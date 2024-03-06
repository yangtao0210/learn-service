package od

import (
	"fmt"
	"strings"
)

var resSlice []string

func SplitAndTransferString() {
	var k int
	fmt.Scan(&k)
	var str string
	fmt.Scan(&str)

	strs := strings.Split(str, "-")
	resSlice = make([]string, 0)
	if len(strs) <= 1 {
		fmt.Println(str)
	} else {
		resSlice = append(resSlice, strs[0])
		for i := 1; i < len(strs); i++ {
			appendStrToResult(strs[i], k)
		}
	}
	fmt.Println(strings.Join(resSlice, "-"))
}

func appendStrToResult(str string, k int) {
	c, loop := 0, len(str)/k
	for c < loop {
		resSlice = append(resSlice, transferStr(str[c*k:(c+1)*k]))
		c++
	}
	if len(str)%k != 0 {
		resSlice = append(resSlice, transferStr(str[c*k:]))
	}
}

func transferStr(str string) string {
	lower, upper := 0, 0
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			lower++
		} else if c >= 'A' && c <= 'Z' {
			upper++
		}
	}
	if lower > upper {
		return strings.ToLower(str)
	} else if lower < upper {
		return strings.ToUpper(str)
	}
	return str
}
