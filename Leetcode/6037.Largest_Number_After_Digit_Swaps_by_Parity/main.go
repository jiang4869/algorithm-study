package main

func largestInteger(num int) int {
	arr := make([]int, 0)
	for true {
		arr = append(arr, num%10)
		num /= 10
		if num == 0 {
			break
		}
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]%2 == arr[j]%2 && arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	num = 0
	for i := len(arr) - 1; i >= 0; i-- {
		num = num*10 + arr[i]
	}

	return num
}
