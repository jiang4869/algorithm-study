package main

import "strings"

// TODO 用 Rabin-Karp search 自己实现一遍
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
