package od

func MaxLengthContainsO(str string) int {
	//思路：因为首位相连，故只需统计字符串中o出现的次数，分为以下两种情况
	//奇数个O:删除一个o即为结果
	//偶数个O：返回字符串长度
	oCount := 0
	for _, c := range str {
		if c == 'o' {
			oCount++
		}
	}
	if oCount%2 == 1 {
		return len(str) - 1
	}
	return len(str)
}
