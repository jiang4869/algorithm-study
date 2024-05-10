package main

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	ans := 0
	i := 0
	j := len(plants) - 1
	curCapacityA := capacityA
	curCapacityB := capacityB
	for i <= j {
		if i == j {
			if curCapacityA < plants[i] && curCapacityB < plants[j] {
				ans++
			}
			break
		}
		if curCapacityA >= plants[i] {
			curCapacityA -= plants[i]
		} else {
			curCapacityA = capacityA - plants[i]
			ans++
		}
		if curCapacityB >= plants[j] {
			curCapacityB -= plants[j]
		} else {
			curCapacityB = capacityA - plants[j]
			ans++
		}

		i++
		j--
	}
	return ans
}
