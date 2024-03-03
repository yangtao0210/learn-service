package od

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func HeightAndWeightOrder() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
	}
	seq := 1
	var heights, weights []string

	//输入两个序列
	for seq < 3 && scanner.Scan() {
		if seq == 1 {
			heights = strings.Split(scanner.Text(), " ")[:n]
		} else {
			weights = strings.Split(scanner.Text(), " ")[:n]
		}
		seq++
	}
	fmt.Println(getOrderByHeightAndWeight(heights, weights))
}

func getOrderByHeightAndWeight(heights, weights []string) []int {
	type Student struct {
		Height int
		Weight int
		Serial int
	}
	stus := make([]Student, len(heights))
	for i := 0; i < len(heights); i++ {
		stus[i].Height, _ = strconv.Atoi(heights[i])
		stus[i].Weight, _ = strconv.Atoi(weights[i])
		stus[i].Serial = i + 1
	}

	//排序：身高升序，相同身高体重升序，相同体重身高，序号升序
	sort.Slice(stus, func(i, j int) bool {
		if stus[i].Height < stus[j].Height {
			return true
		} else if stus[i].Height == stus[j].Height {
			if stus[i].Weight < stus[j].Weight {
				return true
			} else if stus[i].Weight == stus[j].Weight {
				return stus[i].Serial < stus[j].Serial
			}
		}
		return false
	})

	fmt.Println("排序后：", stus)
	// 组装结果
	res := make([]int, len(stus))
	for i, stu := range stus {
		res[i] = stu.Serial
	}
	return res
}
