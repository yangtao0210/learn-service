package examination

func LastValidateCharPosition(s, l string) int {
	res := -1
	//动态规划解法
	//dp[i][j]:表示s[0:i-1]与l[0:j-1]最长公共子序列长度
	dp := make([][]int, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]int, len(l)+1)
	}
	//递推公式，分两种情况
	//1）s[i-1] == l[j-1]:dp[i][j] = dp[i-1][j-1]+1
	//2) s[i-1] != l[j-1]:dp[i][j] = max(dp[i-1][j],dp[i][j-1])
	//遍历顺序：自上而下，自左而右
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(l); j++ {
			if s[i-1] == l[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}

			if dp[i][j] == len(s) {
				res = j - 1
				break
			}
		}
	}
	return res
}
