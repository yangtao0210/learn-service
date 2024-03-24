package dfs

import "fmt"

var res []string
var path []byte
var used []bool

func ArrangementString() {
	var str string
	fmt.Scan(&str)
	res = make([]string, 0)
	path = make([]byte, 0)
	used = make([]bool, len(str))
	tracebacking(str)
	fmt.Println(res)
}

func tracebacking(str string) {
	//递归出口
	if len(path) == len(str) {
		tmp := make([]byte, len(path))
		copy(tmp, path)
		res = append(res, string(tmp))
		return
	}
	//递归式
	for i := 0; i < len(str); i++ {
		//同层剪枝:去重同层重复元素
		if i > 0 && str[i] == str[i-1] && !used[i-1] {
			continue
		}
		//同枝剪枝：同一个元素不能在某个树枝重复使用
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, str[i])
		tracebacking(str)
		path = path[:len(path)-1]
		used[i] = false
	}
}
