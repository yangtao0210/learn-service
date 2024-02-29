package od

import "fmt"

var used []bool
var path []byte
var res int

func NumberOfSpecialStr(str string, n int) int {
	res = 0
	used = make([]bool, len(str))
	backtracking(str, n)
	return res
}

func backtracking(str string, n int) {
	if len(path) == n {
		res++
		return
	}
	for i := 0; i < len(str); i++ {
		//同树层去重
		if i > 0 && used[i] == false && str[i] == str[i-1] {
			continue
		}
		//路径的相邻元素不能重复
		if len(path) > 0 && path[len(path)-1] == str[i] {
			continue
		}
		path = append(path, str[i])
		used[i] = true
		fmt.Println(path, used)
		//递归调用
		backtracking(str, n)
		//回溯
		used[i] = false
		path = path[:len(path)-1]
	}
}
