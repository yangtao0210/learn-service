package examination

import (
	"fmt"
	"regexp"
	"sort"
)

func ContinuousCharLength() {
	var str string
	fmt.Scan(&str)

	var k int
	fmt.Scan(&k)

	//统计str的字符set
	set := make(map[byte]struct{})
	for _, c := range str {
		if _, ok := set[byte(c)]; !ok {
			set[byte(c)] = struct{}{}
		}
	}

	if len(set) < k {
		fmt.Println(-1)
	} else {
		charMap := make(map[byte]int)
		for k := range set {
			reg := regexp.MustCompile(string(k) + "+")
			regRes := reg.FindAllString(str, -1)
			for _, s := range regRes {
				if _, ok := charMap[k]; !ok {
					charMap[k] = len(s)
				} else {
					charMap[k] = max(len(s), charMap[k])
				}
			}
		}

		type CharCount struct {
			Char  byte
			Count int
		}
		//将map组装成结构体数组
		ccSlice := make([]CharCount, 0)
		for k, v := range charMap {
			ccSlice = append(ccSlice, CharCount{
				Char:  k,
				Count: v,
			})
		}
		//排序
		sort.Slice(ccSlice, func(i, j int) bool {
			return ccSlice[i].Count > ccSlice[j].Count
		})
		fmt.Println(ccSlice[k-1].Count)
	}
}
