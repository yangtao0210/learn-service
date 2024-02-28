package od

import (
	"container/list"
	"math"
)

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
