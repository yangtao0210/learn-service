package interview

import "sort"

// GroupAnagrams 字母异位词分组：给定一个字符串数组 strs ，将 变位词 组合在一起。 可以按任意顺序返回结果列表
func GroupAnagrams(strs []string) [][]string {
	groups := make([][]string, 0)
	if len(strs) == 0 {
		return groups
	}
	keyMap := make(map[string][]string)
	//1.以排序后的单词作为map 的 key，其值为异位词的分组
	for _, str := range strs {
		keyMap[sortString(str)] = append(keyMap[sortString(str)], str)
	}
	//2.遍历map将值加入结果集中
	for _, val := range keyMap {
		groups = append(groups, val)
	}
	return groups
}

func sortString(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}
