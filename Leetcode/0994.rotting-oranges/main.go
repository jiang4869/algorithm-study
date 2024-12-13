package main

func orangesRotting(grid [][]int) int {
	queue := make([]struct {
		x, y int
		cnt  int
	}, 0)
	visit := make([][]bool, len(grid))
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < len(grid); i++ {
		visit[i] = make([]bool, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, struct {
					x, y int
					cnt  int
				}{x: i, y: j, cnt: 0})
			}
		}
	}
	ans := 0
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		x, y, cnt := front.x, front.y, front.cnt
		if cnt > ans {
			ans = cnt
		}
		visit[x][y] = true
		if x+1 < m && !visit[x+1][y] && grid[x+1][y] == 1 {
			queue = append(queue, struct {
				x, y int
				cnt  int
			}{x: x + 1, y: y, cnt: cnt + 1})
			grid[x+1][y] = 2
		}
		if x-1 >= 0 && !visit[x-1][y] && grid[x-1][y] == 1 {
			queue = append(queue, struct {
				x, y int
				cnt  int
			}{x: x - 1, y: y, cnt: cnt + 1})
			grid[x-1][y] = 2
		}
		if y-1 >= 0 && !visit[x][y-1] && grid[x][y-1] == 1 {
			queue = append(queue, struct {
				x, y int
				cnt  int
			}{x: x, y: y - 1, cnt: cnt + 1})
			grid[x][y-1] = 2
		}
		if y+1 < n && !visit[x][y+1] && grid[x][y+1] == 1 {
			queue = append(queue, struct {
				x, y int
				cnt  int
			}{x: x, y: y + 1, cnt: cnt + 1})
			grid[x][y+1] = 2
		}
	}
	for i := 0; i < len(grid); i++ {
		visit[i] = make([]bool, len(grid))
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return ans
}
