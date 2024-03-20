package examination

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Elem struct {
	Num   int
	Count int
	Index int
}

func ClearRepeatElem() {
	var s string
	fmt.Scan(&s)

	//统计每个字符出现的次数
	strs := strings.Split(s, ",")
	numsMap := make(map[int]*Elem, len(strs))
	for i := 0; i < len(strs); i++ {
		num, _ := strconv.Atoi(strs[i])
		if _, ok := numsMap[num]; ok {
			numsMap[num].Count += 1
		} else {
			numsMap[num] = &Elem{
				Num:   num,
				Count: 1,
				Index: i,
			}
		}
	}

	//遍历map将val存储到结果集中
	elems := make([]Elem, 0)
	for _, v := range numsMap {
		elems = append(elems, *v)
	}

	//对结果集按照规则排序：次数降序，次数相等时，下标升序
	sort.SliceStable(elems, func(i, j int) bool {
		//次数降序
		if elems[i].Count > elems[j].Count {
			return true
		} else if elems[i].Count == elems[j].Count {
			//次数相等时，下标升序
			return elems[i].Index < elems[j].Index
		}
		return false
	})

	//输出结果
	for _, elem := range elems {
		fmt.Print(elem.Num)
	}

}
