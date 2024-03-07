package od

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const side int = 18

func Airs() {
	reader := bufio.NewScanner(os.Stdin)
	var blacks, whites []int
	row := 0
	for row < 2 && reader.Scan() {
		strs := strings.Split(reader.Text(), " ")
		nums := make([]int, len(strs))
		for i, s := range strs {
			nums[i], _ = strconv.Atoi(s)
		}
		if row == 0 {
			blacks = nums
		} else {
			whites = nums
		}
		row++
	}
	fmt.Println(getAirsOfFirst(blacks, whites), " ", getAirsOfFirst(whites, blacks))
}

func getAirsOfFirst(first, second []int) int {
	set := make(map[string]struct{})
	for i := 0; i < len(first); i += 2 {
		x, y := first[i], first[i+1]
		//记录当前坐标：方便处理当前坐标可能是另一个坐标的气点的情况
		set[fmt.Sprintf("%v_%v", x, y)] = struct{}{}
		if x > 0 {
			set[fmt.Sprintf("%v_%v", x-1, y)] = struct{}{}
		}
		if x < side {
			set[fmt.Sprintf("%v_%v", x+1, y)] = struct{}{}
		}
		if y > 0 {
			set[fmt.Sprintf("%v_%v", x, y-1)] = struct{}{}
		}
		if y < side {
			set[fmt.Sprintf("%v_%v", x, y+1)] = struct{}{}
		}
	}
	//去除包含在second坐标列表中的气点
	for i := 0; i < len(second); i += 2 {
		x, y := second[i], second[i+1]
		k := fmt.Sprintf("%v_%v", x, y)
		if _, ok := set[k]; ok {
			delete(set, k)
		}
	}
	////去除包含本身坐标列表中的气点：存在当前坐标可能是另一个坐标的气点的情况
	//for i := 0; i < len(first); i += 2 {
	//	x, y := first[i], first[i+1]
	//	k := fmt.Sprintf("%v_%v", x, y)
	//	if _, ok := set[k]; ok {
	//		delete(set, k)
	//	}
	//}
	return len(set) - len(first)/2
}
