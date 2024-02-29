package od

import (
	"fmt"
	"sort"
)

func TestMinDiffOfScore() {
	var n int
	fmt.Scan(&n)
	scores := make([][2]int, n)
	for i, _ := range scores {
		var index, score int
		fmt.Scan(&index, &score)
		scores[i][0] = index
		scores[i][1] = score
	}
	res := MinDiffOfScore(scores)
	for _, r := range res {
		fmt.Printf("%v %v\n", r[0], r[1])
	}
}

func MinDiffOfScore(scores [][2]int) [][2]int {
	res := make([][2]int, 0)
	//按照分数升序排序
	sort.Slice(scores, func(i, j int) bool {
		if scores[i][1] < scores[j][1] {
			return true
		}
		return false
	})
	//求结果
	min := scores[1][1] - scores[0][1]
	if scores[0][0] < scores[1][0] {
		res = append(res, [2]int{scores[0][0], scores[1][0]})
	} else {
		res = append(res, [2]int{scores[1][0], scores[0][0]})
	}
	for i := 2; i < len(scores); i++ {
		diff := scores[i][1] - scores[i-1][1]
		if diff == min {
			if scores[i-1][0] < scores[i][0] {
				res = append(res, [2]int{scores[i-1][0], scores[i][0]})
			} else {
				res = append(res, [2]int{scores[i][0], scores[i-1][0]})
			}
		} else if diff < min {
			min = diff
			res = res[:0]
			if scores[i-1][0] < scores[i][0] {
				res = append(res, [2]int{scores[i-1][0], scores[i][0]})
			} else {
				res = append(res, [2]int{scores[i][0], scores[i-1][0]})
			}
		}
	}
	//对结果按p1升序排序
	sort.Slice(res, func(i, j int) bool {
		if res[i][0] < res[j][0] {
			return true
		}
		return false
	})
	return res
}
