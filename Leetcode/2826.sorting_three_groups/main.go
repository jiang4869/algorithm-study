package main

func minimumOperations(nums []int) int {
	ans := 200
	n := len(nums)
	a, b, c := make([]int, len(nums)+1), make([]int, len(nums)+1), make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		a[i] = a[i-1]
		b[i] = b[i-1]
		c[i] = c[i-1]
		if nums[i-1] == 1 {
			a[i] = a[i-1] + 1
		} else if nums[i-1] == 2 {
			b[i] = b[i-1] + 1
		} else {
			c[i] = c[i-1] + 1
		}
	}
	ans = min(min(min(ans, n-a[n]), n-b[n]), n-c[n])

	// [1,a] [a+1,b],[b+1,n]
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			num1 := a[i]
			num2 := b[j] - b[i]
			num3 := c[n] - c[j]
			num := (i - num1) + (j - i - num2) + (n - j - num3)
			ans = min(num, ans)
		}
	}

	for i := 1; i < n; i++ {
		num1 := i
		num2 := n - i
		// [1,i] [i+1,n]
		// 1 2
		ans = min(ans, (num1-a[i])+(num2-b[n]+b[i]))
		// 1 3
		ans = min(ans, (num1+num2)-1)
		ans = min(ans, (num1-a[i])+(num2-c[n]+c[i]))
		// 2 3
		ans = min(ans, (num1-b[i])+(num2-c[n]+c[i]))
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
