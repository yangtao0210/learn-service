package examination

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func PrintWordMatchedWithInputPrefix() {
	var str, prefix string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		str = scanner.Text()
	}
	fmt.Scan(&prefix)

	//构造正则表达式:[[:punct:]]匹配标点符号;匹配除字符和空格外的所有字符
	//re := regexp.MustCompile(`[^a-zA-Z\s]`)
	re := regexp.MustCompile(`[[:punct:]]`)
	//替换原字符串中正则匹配到的字符
	dstStr := re.ReplaceAllString(str, " ")
	words := strings.Split(dstStr, " ")
	//构造word字典
	wordMap := make(map[string]struct{})
	for _, word := range words {
		if _, ok := wordMap[word]; !ok && !strings.EqualFold(word, "") {
			wordMap[word] = struct{}{}
		}
	}
	//基于pre前缀匹配字典key
	matchedRes := MatchedWithPrefix(wordMap, prefix)
	for _, w := range matchedRes {
		fmt.Printf("%v ", w)
	}
}

func MatchedWithPrefix(wordMap map[string]struct{}, pre string) []string {
	result := make([]string, 0)
	for k := range wordMap {
		//判断pre是否为k的前缀
		if strings.HasPrefix(k, pre) {
			result = append(result, k)
		}
	}

	if len(result) == 0 {
		return []string{pre}
	}
	//按字典排序
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}
