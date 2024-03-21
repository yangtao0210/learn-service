package interview

// 题目：岛屿周长
type pair struct {
	x int
	y int
}

var pairs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func LandPerimeter(grid [][]int) int {
	res := 0
	n, m := len(grid), len(grid[0])
	//解法1:迭代法
	for i, row := range grid {
		for j, val := range row {
			if val == 1 {
				//如果是陆地,判断四个边的状态：是否为边界，是否为陆地
				for _, p := range pairs {
					x, y := i+p.x, j+p.y
					if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
						res++
					}
				}
			}
		}
	}
	return res
}
