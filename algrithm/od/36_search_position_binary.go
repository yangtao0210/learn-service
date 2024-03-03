package od

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SearchPositionBinary() {
	//1.处理输入数据
	var sIDs string
	fmt.Scan(&sIDs)
	var sID int
	fmt.Scan(&sID)

	ids := strings.Split(sIDs, ",")
	nums := make([]int, len(ids))
	for i, id := range ids {
		id_num, _ := strconv.Atoi(id)
		nums[i] = id_num
	}

	sort.Ints(nums)
	//2.二分查找插入位置:左闭右闭区间
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > sID {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	//队列编号从1开始
	fmt.Println(l + 1)
}
