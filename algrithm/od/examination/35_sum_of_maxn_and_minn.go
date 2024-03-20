package examination

import (
	"fmt"
	"sort"
)

func SumOfMaxAndMinN() {
	//1.处理输入
	var m, n int
	fmt.Scan(&m)
	nums := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Scan(&n)

	//边界条件
	if m < 2*n {
		fmt.Println(-1)
	}
	//2.求解：排序 + 去重
	sort.Ints(nums)
	unRepeatedNums := make([]int, 0)
	unRepeatedNums = append(unRepeatedNums, nums[0])
	for i := 1; i < m; i++ {
		if nums[i] != nums[i-1] {
			unRepeatedNums = append(unRepeatedNums, nums[i])
		}
	}

	fmt.Println(unRepeatedNums)

	//基于去重后的数组打印结果
	if len(unRepeatedNums) < 2*n {
		fmt.Println(-1)
	} else {
		fmt.Println(unRepeatedNums[0] + unRepeatedNums[1] +
			unRepeatedNums[len(unRepeatedNums)-1] + unRepeatedNums[len(unRepeatedNums)-2])
	}

}
