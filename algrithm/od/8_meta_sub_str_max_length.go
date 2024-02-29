package od

import (
	"fmt"
)

func TestMetaSubstrMaxLength() {
	var flaw int
	var str string
	fmt.Scan(&flaw)
	fmt.Scan(&str)
	fmt.Println(MetaSubstrMaxLength(str, flaw))
}

func MetaSubstrMaxLength(str string, gap int) int {
	res := 0
	metaSet := make(map[int32]struct{})
	metaStr := "aoeiuAOEIU"
	for _, meta := range metaStr {
		metaSet[meta] = struct{}{}
	}
	metaIndexs := make([]int, 0)
	for i, c := range str {
		if _, ok := metaSet[c]; ok {
			metaIndexs = append(metaIndexs, i)
		}
	}
	fmt.Println(metaIndexs)
	//双指针解法
	left, right := 0, 0
	for right < len(metaIndexs) {
		diff := metaIndexs[right] - metaIndexs[left] - (right - left)
		if diff > gap {
			left++
		} else {
			if diff == gap {
				res = max(res, metaIndexs[right]-metaIndexs[left]+1)
			}
			right++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
