package od

import (
	"fmt"
)

var m, n, k int
var visited [][]bool

func SearchGold() {
	fmt.Scan(&m, &n, &k)
	visited = make([][]bool, m)
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}
	fmt.Println(dfsVisited(0, 0))
}

func dfsVisited(x, y int) int {
	//递归出口：x|y到达边界 或者 x,y的数位之和大于k
	if x < 0 || y < 0 || x > m-1 || y > n-1 || visited[x][y] || getDigistSum(x)+getDigistSum(y) > k {
		return 0
	}
	visited[x][y] = true
	return 1 + dfsVisited(x-1, y) + dfsVisited(x+1, y) + dfsVisited(x, y-1) + dfsVisited(x, y+1)
}

func getDigistSum(num int) int {
	sum := 0
	for num != 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
