package examination

import (
	"container/list"
	"fmt"
	"strings"
)

func TestSubstrLength() {
	var s string
	for {
		fmt.Scan(&s)
		if strings.EqualFold(s, "exit") {
			break
		}
		fmt.Println(CompulateSubstrLength(s))
	}
}

func CompulateSubstrLength(str string) int {
	res, hasLetter := -1, false
	//使用队列存储字母的索引位置：判断字母数量 & 获取下次判断子串的开始索引
	queue := list.New()
	l, r := 0, 0
	for r < len(str) {
		c := str[r]
		//如果当前字符是字母
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			//设置是否有字母标志位
			hasLetter = true
			//当前字母下标放入队列
			queue.PushBack(r)
			//如果字母数量大于1个，则移除队列头元素，并将左指针移动到该索引的下一个位置，继续遍历
			if queue.Len() > 1 {
				head := queue.Front()
				l = head.Value.(int) + 1
				queue.Remove(head)
			}
			//左右指针指向同一个位置，则移动右指针，并进行下一次遍历
			if l == r {
				r++
				continue
			}
		}
		res = max(res, r-l+1)
		r++
	}
	if !hasLetter {
		return -1
	}
	return res
}
