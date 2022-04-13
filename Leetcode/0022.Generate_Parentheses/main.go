package main

var ans []string

func dfs(left, right, n int, data []byte) {
	if left > n || right > left {
		return
	}
	if right == n {
		ans = append(ans, string(data))
		return
	}
	if left < n {
		data[left+right] = '('
		dfs(left+1, right, n, data)
	}
	if right < left {
		data[left+right] = ')'
		dfs(left, right+1, n, data)

	}

}

func generateParenthesis(n int) []string {
	ans = make([]string, 0)
	data := make([]byte, 2*n)
	dfs(0, 0, n, data)
	return ans
}
