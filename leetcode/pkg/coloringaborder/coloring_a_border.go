package coloringaborder

func ColorBorder(grid [][]int, row int, col int, color int) [][]int {
	return colorBorder(grid, row, col, color)
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	var needColor [][2]int
	var visited [][]bool = make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	var try func(int, int, int) int = func(i, j int, color int) int {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) { // 出界
			return -1
		}
		if grid[i][j] != color { // 颜色不同
			return -1
		}
		if visited[i][j] { // 访问过了
			return 0
		}
		return 1
	}

	dirs := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 官解的dirs，上下左右

	var dfs func(int, int)
	dfs = func(i, j int) {
		var isBorder bool
		visited[i][j] = true

		for _, dir := range dirs {
			ni, nj := i+dir.x, j+dir.y
			switch try(ni, nj, grid[i][j]) {
			case -1:
				isBorder = true
			case 1:
				dfs(ni, nj)
			}
		}

		if isBorder {
			needColor = append(needColor, [2]int{i, j})
		}
	}
	dfs(row, col)

	for i := 0; i < len(needColor); i++ {
		grid[needColor[i][0]][needColor[i][1]] = color
	}
	return grid
}
