package leetcode

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	fmt.Println(strings.TrimSpace(s))
	strs := strings.Split(strings.TrimSpace(s), " ")

	// 双指针移除空值
	slow, fast := 0, 0
	for fast < len(strs) {
		if !strings.EqualFold(strs[fast], "") {
			if strings.EqualFold(strs[slow], "") {
				strs[fast], strs[slow] = strs[slow], strs[fast]
			}
			fast++
			slow++
		} else {
			fast++
		}
	}
	fmt.Println(strs)
	//翻转字符串
	start, end := 0, slow-1
	for start < end {
		strs[start], strs[end] = strs[end], strs[start]
		start++
		end--
	}
	fmt.Println(strings.TrimSpace(strings.Join(strs, " ")))
	return strings.TrimSpace(strings.Join(strs, " "))
}
