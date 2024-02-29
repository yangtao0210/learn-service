package od

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TestAmountCount() {
	for {
		reader := bufio.NewReader(os.Stdin)
		lineBytes, _, _ := reader.ReadLine()
		strs := strings.Split(string(lineBytes), " ")
		if len(strs) <= 2 {
			break
		}
		var nums []int
		for _, str := range strs {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}
		fmt.Println(AmountCount(nums))
	}
}

func AmountCount(nums []int) int {
	res := 0
	pre, cur := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		cur = nums[i+1] - nums[i]
		fmt.Println(pre, cur)
		//转折点处更新pre
		if (pre >= 0 && cur < 0) || (pre <= 0 && cur > 0) {
			//正峰值时更新res
			if pre >= 0 && cur < 0 {
				res++
			}
			pre = cur
		}
		//判断最后一个元素是否为正峰值
		if i == len(nums)-2 && cur > 0 {
			res++
		}

	}
	return res
}
