package examination

import (
	"fmt"
	"sort"
)

func MaxDispatchTeamsCount() {
	var n, minAblity int
	fmt.Scan(&n)
	ablities := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ablities[i])
	}
	fmt.Scan(&minAblity)

	//排序
	sort.Ints(ablities)
	teamCount, l, r := 0, 0, n-1
	for l < r {
		if ablities[r] >= minAblity {
			teamCount++
			r--
		} else if ablities[l]+ablities[r] >= minAblity {
			teamCount++
			r--
			l++
		} else {
			l++
		}
	}
	if l == r && ablities[l] >= minAblity {
		teamCount++
	}
	fmt.Println(teamCount)
}
