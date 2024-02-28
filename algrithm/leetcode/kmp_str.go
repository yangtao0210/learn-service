package leetcode

func getNext(s string) []int {
	//1.初始化
	j, next := 0, make([]int, len(s))
	bytes := []byte(s)
	for i := 1; i < len(s); i++ {
		//2.i,j位置字符不相等:更新j为next[j-1]，直到两者相等，或者j==0
		for j > 0 && bytes[i] != bytes[j] {
			j = next[j-1] //找前一位元素的最长前缀位置
		}
		//3.i,j位置字符相等
		if bytes[i] == bytes[j] {
			j++
		}
		next[i] = j
	}
	return next
}

func strStr(haystack string, needle string) int {
	j, next := 0, getNext(needle)
	//遍历主串，比较主串和模式串的值
	for i := 0; i < len(haystack); i++ {
		//若某位置上主串和模式串不匹配，则更新模式串的新开始匹配位置，即j = next[j-1]，直到两者相等或者已取到next数组的第一个元素值
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}
