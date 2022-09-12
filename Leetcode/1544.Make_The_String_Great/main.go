package main

func makeGood(s string) string {
	res := ""

	for i := 0; i < len(s); i++ {
		if len(res) == 0 {
			res = res + string(s[i])
			continue
		}
		if res[len(res)-1] >= 'a' && res[len(res)-1] <= 'z' {
			if res[len(res)-1] == s[i]+32 {
				res = res[:len(res)-1]
			} else {
				res = res + string(s[i])
			}
		} else if res[len(res)-1] >= 'A' && res[len(res)-1] <= 'Z' {
			if res[len(res)-1] == s[i]-32 {
				res = res[:len(res)-1]
			} else {
				res = res + string(s[i])
			}
		}
	}

	return res
}
