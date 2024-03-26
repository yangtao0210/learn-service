package interview

import (
	"strconv"
	"strings"
)

func MaxContinuousSubString(s string) string {
	if len(s) == 0 {
		return ""
	}
	ss := strings.Split(s, ",")
	maxSubStr := ss[0]
	curSubStr := ss[0]
	for i := 1; i < len(ss); i++ {
		pre, cur := ss[i-1], ss[i]
		if isDigit(pre) && isDigit(cur) {
			//1.pre,cur均为数字
			preNum, _ := strconv.Atoi(pre)
			curNum, _ := strconv.Atoi(cur)
			if curNum-preNum == 1 {
				curSubStr += cur
			} else {
				maxSubStr = maxSubString(maxSubStr, curSubStr)
				curSubStr = cur
			}

		} else if isLetter(pre) && isLetter(cur) {
			//2.pre,cur均为字母
			if cur[0]-pre[0] == 1 {
				curSubStr += cur
			} else {
				maxSubStr = maxSubString(maxSubStr, curSubStr)
				curSubStr = cur
			}
		} else if isDigit(pre) && isLetter(cur) {
			//3.pre为数字，cur为字母:
			//1）若当前数字为10且下一个字母为A时，将当前元素添加到子串中
			if pre == "10" && cur == "A" {
				curSubStr += cur
			} else {
				maxSubStr = maxSubString(maxSubStr, curSubStr)
				curSubStr = cur
			}
		}
	}
	return maxSubString(maxSubStr, curSubStr)
}

func isDigit(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isLetter(s string) bool {
	if s >= "A" && s <= "Z" {
		return true
	}
	return false
}

func maxSubString(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}
