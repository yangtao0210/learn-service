package interview

import (
	"strconv"
	"strings"
)

func DigitSum(s string, k int) string {
	if len(s) <= k {
		return s
	}
	for len(s) > k {
		tmp := make([]string, 0)
		n := len(s)
		//遍历字符串s,以k个字符分组，计算每个分组的和，并添加到tmp中
		for i := 0; i < n; i += k {
			sum := getSum(s, i, min(i+k, n))
			tmp = append(tmp, strconv.Itoa(sum))
		}
		//用tmp替换s，判断s是否符合要求
		s = strings.Join(tmp, "")
	}

	return s
}

func getSum(s string, start, end int) int {
	val := 0
	//计算每k个字符的和:start,end
	for i := start; i < end; i++ {
		val += int(s[i] - '0')
	}
	return val
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
