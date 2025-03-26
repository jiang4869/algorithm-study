package main

func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	pCnt := [26]int{}
	sCnt := [26]int{}
	for _, c := range p {
		pCnt[c-'a']++
	}
	pLen := len(p)
	l, r := 0, 0
	for l < len(s) && r < len(s) {
		sCnt[s[r]-'a']++
		r++
		if (r - l) < pLen {
			continue
		}
		if pCnt == sCnt {
			ans = append(ans, l)
		}
		sCnt[s[l]-'a']--
		l++
	}
	return ans
}
