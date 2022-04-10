package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	arr := make([][]uint8, numRows)
	numCols := len(s)
	for i := range arr {
		arr[i] = make([]uint8, numCols)
		for j := range arr[i] {
			arr[i][j] = 32
		}
	}
	i, j := 0, 0
	cur := 0
	lens := len(s)
	flag := false
	maxi, maxj := 0, 0
	for j < numCols && cur < lens {
		if flag {
			for i >= 0 && cur < lens {
				arr[i][j] = s[cur]
				maxi = max(maxi, i)
				maxj = max(maxj, j)
				j++
				cur++
				i--
			}
			i += 2
			j--
			flag = false
		} else {
			for i < numRows && cur < lens {
				arr[i][j] = s[cur]
				maxi = max(maxi, i)
				maxj = max(maxj, j)
				cur++
				i++
			}
			j++
			i -= 2
			flag = true
		}
	}
	var ans string
	for h := 0; h <= maxi; h++ {
		for k := 0; k <= maxj; k++ {
			if (arr[h][k] >= 'a' && arr[h][k] <= 'z') || (arr[h][k] >= 'A' && arr[h][k] <= 'Z') {
				ans += string(arr[h][k])

			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
