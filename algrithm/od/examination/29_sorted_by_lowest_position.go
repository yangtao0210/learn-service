package examination

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SortedByLowestPosition() {
	var s string
	fmt.Scan(&s)

	strs := strings.Split(s, ",")
	nums := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		num, _ := strconv.Atoi(strs[i])
		nums[i] = num
	}
	sort.Slice(nums, func(i, j int) bool {
		return getKey(nums[i]) < getKey(nums[j])
	})
	fmt.Println(nums)
}

func getKey(num int) int {
	res := 0
	if num < 0 {
		res = (-1 * num) % 10
	} else {
		res = num % 10
	}
	return res
}
