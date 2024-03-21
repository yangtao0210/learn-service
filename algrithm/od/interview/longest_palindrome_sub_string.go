package interview

// 题目描述：
// 给你一个字符串 s，找到 s 中最长的回文子串
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

func LongestPalindrome(s string) string {
	maxLength := 0
	res := ""
	//解法：动态规划
	//1) dp[i][j]: s[i:j]是否为回文串
	//2) 确定递推公式：分为两种情况
	//2.1:s[i] != s[j] ==> dp[i][j] == false
	//2.2:s[i] == s[j] ==>
	// 若i,j相邻/相等,dp[i][j] = true
	// 否则:可由dp[i+1][j-1]推出
	//3.遍历顺序：从下至上，从左至右
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j <= len(s)-1; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] && j-i+1 > maxLength {
					res = s[i : j+1]
					maxLength = j - i + 1
				}
			}
		}
	}
	return res
}
