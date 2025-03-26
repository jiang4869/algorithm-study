package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func differenceOfDistinctValues(grid [][]int) [][]int {
	n := len(grid)
	m := len(grid[0])
	ans := make([][]int, n)

	mmid := make(map[int]struct{})

	for i := 0; i < n; i++ {
		ans[i] = make([]int, m)
		for j := 0; j < m; j++ {
			clear(mmid)
			for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
				mmid[grid[x][y]] = struct{}{}
			}
			topLeft := len(mmid)
			clear(mmid)
			for x, y := i+1, j+1; x < n && y < m; x, y = x+1, y+1 {
				mmid[grid[x][y]] = struct{}{}
			}
			bottomRight := len(mmid)
			ans[i][j] = abs(topLeft - bottomRight)
		}
	}

	return ans
}
