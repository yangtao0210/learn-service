package od

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var totalSum int
var res int

func TestGetMinDiff() {
	read := bufio.NewReader(os.Stdin)
	//输入10个元素
	bytes, _, _ := read.ReadLine()
	nums := strings.Split(string(bytes), " ")
	scores := make([]int, 10)
	for i, num := range nums {
		number, _ := strconv.Atoi(num)
		if i < 10 {
			scores[i] = number
		}
	}
	fmt.Println(GetMinDiff(scores))
}

func GetMinDiff(scores []int) int {
	res = math.MaxInt
	for _, score := range scores {
		totalSum += score
	}
	dfs(scores, 0, 0, 0)
	return res
}

// 递归方法求解:sum表示当前分组的和，count表示当前分组中的元素个数，index表示已遍历元素个数
func dfs(nums []int, index, count, sum int) {
	//递归出口
	if count == 5 || index == 10 {
		//说明已经分好组
		if count == 5 {
			otherSum := totalSum - sum
			diff := abs(otherSum, sum)
			res = min(diff, res)
		}
		return
	}
	//递归式：分两种情况
	// 1）当前分组选当前元素
	dfs(nums, index+1, count+1, sum+nums[index])
	// 2)当前分组不选当前元素
	dfs(nums, index+1, count, sum)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a, b int) int {
	if a-b < 0 {
		return -1 * (a - b)
	}
	return a - b
}
