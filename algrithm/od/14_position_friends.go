package od

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func TestFriendPosition() {
	read := bufio.NewReader(os.Stdin)
	//输入n
	nBytes, _, _ := read.ReadLine()
	n, _ := strconv.Atoi(string(nBytes))
	//输入n个身高
	bytes, _, _ := read.ReadLine()
	nums := strings.Split(string(bytes), " ")
	heights := make([]int, n)
	for i, num := range nums {
		number, _ := strconv.Atoi(num)
		heights[i] = number
	}
	fmt.Println(FriendPosition(heights))
}

func FriendPosition(heights []int) []int {
	math.Ceil(float64(6 / 2))
	res := make([]int, len(heights))
	stack := list.New()
	stack.PushBack(0)
	for i := 1; i < len(heights); i++ {
		for stack.Len() != 0 && heights[stack.Back().Value.(int)] < heights[i] {
			pop := stack.Back()
			res[pop.Value.(int)] = i
			stack.Remove(pop)
		}
		stack.PushBack(i)
	}
	for stack.Len() != 0 {
		pop := stack.Back()
		res[pop.Value.(int)] = 0
		stack.Remove(pop)
	}
	return res
}
