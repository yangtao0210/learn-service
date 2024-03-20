package examination

import (
	"fmt"
	"sort"
	"strings"
)

func ExchangeToMinStr() {
	var origin string
	fmt.Scan(&origin)

	//解题思路：
	originBytes, sortedBytes := []byte(origin), []byte(origin)
	sort.Slice(sortedBytes, func(i, j int) bool {
		return sortedBytes[i] < sortedBytes[j]
	})
	//若排序后的数组与原数组相同，则直接输出
	if strings.EqualFold(origin, string(sortedBytes)) {
		fmt.Println(origin)
	} else {
		for i := 0; i < len(origin); i++ {
			if originBytes[i] != sortedBytes[i] {
				//1.寻找strBytes[i]在原字符串的位置
				index := -1
				for j := i + 1; j < len(origin); j++ {
					if originBytes[j] == sortedBytes[i] {
						index = j
					}
				}
				//2.交换origin中i与index位置的值
				if index != -1 {
					originBytes[i], originBytes[index] = originBytes[index], originBytes[i]
					break
				}
			}
		}
	}

	fmt.Println(string(originBytes))
}
