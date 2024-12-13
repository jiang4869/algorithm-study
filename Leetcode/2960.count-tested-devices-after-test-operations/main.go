package main

func countTestedDevices(batteryPercentages []int) int {
	ans := 0
	cnt := 0
	for _, val := range batteryPercentages {
		if val > cnt {
			ans++
			cnt++
		}
	}
	return ans
}
