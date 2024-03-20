package examination

import (
	"fmt"
	"math"
	"sort"
)

type HeightDiff struct {
	Height int
	Diff   int
}

func DiffAbsOfHeight() {
	//输入h,n
	var h, n int
	fmt.Scan(&h, &n)
	heights := make([]int, n)
	heightDiffs := make([]HeightDiff, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&heights[i])
		heightDiffs[i].Height = heights[i]
		heightDiffs[i].Diff = int(math.Abs(float64(heights[i] - h)))
	}
	sort.Slice(heightDiffs, func(i, j int) bool {
		//身高差升序
		if heightDiffs[i].Diff < heightDiffs[j].Diff {
			return true
		} else if heightDiffs[i].Diff == heightDiffs[j].Diff {
			//身高差相等，身高升序
			return heightDiffs[i].Height < heightDiffs[j].Height
		}
		return false
	})

	for i, diff := range heightDiffs {
		if i != len(heightDiffs)-1 {
			fmt.Printf("%v ", diff.Height)
		} else {
			fmt.Printf("%v\n", diff.Height)
		}
	}
}
