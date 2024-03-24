package lands

//题目描述：
//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//此外，你可以假设该网格的四条边均被水包围。

type dir struct {
	x, y int
}

var dirs = []dir{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var visited [][]bool

func GetNumOfLands(grid [][]byte) int {
	count := 0
	//1.初始化visited
	visited = make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
	//2.遍历grid,调用DFS进行处理
	for i, g := range grid {
		for j, val := range g {
			if val == '1' && !visited[i][j] {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

// 深度优先遍历:将当前元素可达岛屿进行标记
func dfs(grid [][]byte, x, y int) {
	//递归出口：当前元素为‘0’ 或是已被访问的‘1’
	if grid[x][y] == '0' || visited[x][y] {
		return
	}
	//递归式
	//1）更新当前元素访问标识
	visited[x][y] = true
	//2) 处理当前元素的上下左右四个方向
	for _, dir := range dirs {
		nextX, nextY := x+dir.x, y+dir.y
		//剪枝：若当前元素为边界元素，则无需遍历某些方向
		if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
			continue
		}
		//递归处理
		dfs(grid, nextX, nextY)
	}
}
