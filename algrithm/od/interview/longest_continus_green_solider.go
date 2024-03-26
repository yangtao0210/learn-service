package interview

func LongestContinuousGreenSolider(s string) int {
	//解法：滑动窗口
	longest := 0
	left, right, zeroCount := 0, 0, 0
	for right < len(s) {
		if s[right] == '0' {
			//[left,right]窗口0的个数+1
			zeroCount++
			if zeroCount <= 1 {
				//当前窗口有一个0：表示连续1的个数为窗口长度-1（将left右移一个即可）
				left++
			} else {
				//当前窗口有2个0：不符合连续1的要求，则以当前位置的下一个位置为左边界，重新寻找符合条件的右边界
				longest = max(longest, right-left)
				left = right + 1
				zeroCount = 0
			}
		}
		right++
	}
	longest = max(longest, right-left)
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
