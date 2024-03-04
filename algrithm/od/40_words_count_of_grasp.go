package od

import (
	"fmt"
	"maps"
)

func WordCountOfGrasp() {
	var n int
	fmt.Scan(&n)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&words[i])
	}
	var chars string
	fmt.Scan(&chars)

	charMap := make(map[int32]int)
	for _, c := range chars {
		if _, ok := charMap[c]; ok {
			charMap[c] += 1
		} else {
			charMap[c] = 1
		}
	}

	count := 0
	for _, word := range words {
		if isComposed(charMap, word) {
			count++
		}
	}
	fmt.Println(count)
}

func isComposed(charMap map[int32]int, word string) bool {
	chars := make(map[int32]int)

	//copymap数据到临时map，防止原始数组被修改
	maps.Copy(chars, charMap)

	composed := true
	universalChar := '?'
	for _, wc := range word {
		if val, ok := chars[wc]; ok && val > 0 {
			//字符存在
			chars[wc]--
		} else {
			if val, ok := chars[universalChar]; ok && val > 0 {
				//字符不存在,但存在万能字符
				chars[universalChar]--
			} else {
				composed = false
			}
		}
	}
	return composed
}
