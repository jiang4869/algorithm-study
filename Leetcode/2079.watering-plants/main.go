package main

func wateringPlants(plants []int, capacity int) int {
	ans := 0
	curCapacity := capacity
	for i := 0; i < len(plants); i++ {
		curCapacity -= plants[i]
		if i != len(plants)-1 && curCapacity < plants[i+1] {
			curCapacity = capacity
			ans += (i + 1) * 2
		}
		ans++
	}

	return ans
}
