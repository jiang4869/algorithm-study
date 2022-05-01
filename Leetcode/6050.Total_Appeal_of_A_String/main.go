package main

func appealSum(s string) int64 {
	res := int64(0)
	mmid := make(map[int32][]int)
	for idx, val := range s {
		mmid[val] = append(mmid[val], idx)
	}
	length := len(s)
	for _, v := range mmid {
		if len(v) > 0 {
			res += int64((v[0] + 1) * (length - v[0]))
			lenv := len(v)
			for i := 1; i < lenv; i++ {
				res += int64((v[i] - v[i-1]) * (length - v[i]))
			}
		}
	}
	return res
}
