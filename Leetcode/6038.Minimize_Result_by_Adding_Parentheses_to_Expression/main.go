package main

import (
	"math"
	"strconv"
	"strings"
)

func minimizeResult(expression string) string {
	splits := strings.Split(expression, "+")
	var minNum int32
	minNum = math.MaxInt32
	res := ""
	for i := 0; i <= len(splits[0]); i++ {
		for j := 0; j <= len(splits[1]); j++ {
			num1 := splits[0][0:i]
			num2 := splits[0][i:]
			num3 := splits[1][0:j]
			num4 := splits[1][j:]
			if len(num3) == 0 {
				continue
			}
			if len(num2) == 0 {
				continue
			}
			var tmp int32
			tmp = 1
			if len(num1) > 0 {
				tmp *= toNum(num1)
			}
			if len(num4) > 0 {
				tmp *= toNum(num4)
			}
			n2 := toNum(num2)
			n3 := toNum(num3)
			tmp *= n2 + n3
			if tmp < minNum {
				minNum = tmp
				res = num1 + "(" + num2 + "+" + num3 + ")" + num4
			}
		}
	}

	return res
}

func toNum(str string) int32 {
	atoi1, _ := strconv.Atoi(str)
	return int32(atoi1)
}
