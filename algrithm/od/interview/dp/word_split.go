package dp

func WordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	//完全背包+排列问题（先背包，后物品）
	dp := make([]bool, len(s)+1)
	//dp[i]:表示长度为i的子串是否可由字典单词组成
	dp[0] = true                   //空字符串可由字典中的空串组成，即便不存在
	for i := 1; i <= len(s); i++ { //遍历背包
		for j := 0; j < i; j++ { //遍历物品
			//s[j:i]存在于字典且前面长度为j的子串可由字典单词组成
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
