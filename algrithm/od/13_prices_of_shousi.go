package od

import (
	"container/list"
)

func PricesOfShousi(prices []int) []int {
	res, length := make([]int, len(prices)), len(prices)
	//遍历第一遍时，用栈存储未更新价格的下标
	stack := list.New()
	for i := 0; i < 2*length-1; i++ {
		index := i % length
		//遍历时判断：栈顶下标对应元素（index的前一个元素） 是否大于 当前下标元素
		for stack.Len() != 0 && prices[stack.Back().Value.(int)] > prices[index] {
			top := stack.Back()
			res[top.Value.(int)] = prices[top.Value.(int)] + prices[index]
			stack.Remove(top)
		}
		//遍历第一遍时，存储未更新价格的下标
		if i < length {
			stack.PushBack(index)
		}
	}
	for stack.Len() != 0 {
		pop := stack.Back()
		res[pop.Value.(int)] = prices[pop.Value.(int)]
		stack.Remove(pop)
	}
	return res
}
