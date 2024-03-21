package interview

// n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
// 你需要按照以下要求，给这些孩子分发糖果：
// 每个孩子至少分配到 1 个糖果。
// 相邻两个孩子评分更高的孩子会获得更多的糖果。
// 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
func AssignCandies(ratings []int) int {
	res := len(ratings)
	candies := make([]int, len(ratings))
	//双向贪心
	//第一次贪心策略:从左到右遍历：若右孩子分数较大，则其糖果数量为左边孩子糖果数+1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	//第二次贪心策略：从右到左遍历：若左孩子分数较大，则其糖果数为max(当前糖果数，右孩子糖果数+1)==》防止覆盖前一次的值
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}
	//求最小需要的糖果总数
	for i := range candies {
		res += candies[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
