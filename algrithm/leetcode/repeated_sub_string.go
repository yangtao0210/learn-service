package leetcode

import "strings"

func repeatedSubstringPattern(s string) bool {
	ns := []byte(s + s)
	return strings.Contains(string(ns[1:len(ns)-2]), s)
}
