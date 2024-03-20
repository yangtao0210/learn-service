package examination

import (
	"fmt"
	"sort"
	"strings"
)

func AbstractOfString() {
	//处理输入
	var str string
	fmt.Scan(&str)
	strLower := strings.ToLower(str)

	//存储去除纯字母的字符串
	strBytes := make([]byte, 0)
	//统计每个字符的个数
	countChar := make([]int, 128)
	for i := 0; i < len(strLower); i++ {
		if strLower[i] >= 'a' && strLower[i] <= 'z' {
			countChar[strLower[i]]++
			strBytes = append(strBytes, strLower[i])
		}
	}

	//基于规则2.3处理字符串
	//存储结果
	type ByteCount struct {
		Char  byte
		Count int
	}
	res := make([]ByteCount, 0)
	//字符数组结尾加空格，保证最后一个有效字符的处理
	strBytes = append(strBytes, ' ')
	pre, repeat := strBytes[0], 1
	countChar[pre]--
	for i := 1; i < len(strBytes); i++ {
		cur := strBytes[i]
		countChar[cur]--
		if cur == pre {
			repeat++
		} else {
			//判断前一个字符是否为连续字符
			bc := ByteCount{Char: pre}
			if repeat > 1 {
				//出现连续字符
				bc.Count = repeat
			} else {
				//非连续字符
				bc.Count = countChar[pre]
			}
			res = append(res, bc)
			pre = cur
			repeat = 1
		}
	}

	//对结果集进行排序：大数字在前，数字相同时按照字母升序排列
	sort.Slice(res, func(i, j int) bool {
		if res[i].Count > res[j].Count {
			return true
		} else if res[i].Count == res[j].Count {
			return res[i].Char < res[j].Char
		}
		return false
	})
	//打印结果
	for _, elem := range res {
		fmt.Print(string(elem.Char), elem.Count)
	}
}
