package interview

import "container/list"

func ReverseParentheses(s string) string {
	//当前字符为'(': 将str入栈，并将str置空
	//当前字符为')': 反转str,并将该str添加到栈顶字符串末尾，结果赋给str，栈顶元素出栈
	//为字符：将字符添加到str末尾
	stack := list.New()
	strBytes := make([]byte, 0)
	for i := range s {
		if s[i] == '(' {
			stack.PushBack(strBytes)
			strBytes = []byte{}
		} else if s[i] == ')' {
			//反转str
			reverse(strBytes)
			//将str追加到栈顶元素末尾,并赋值给str
			top := stack.Back()
			strBytes = append(top.Value.([]byte), strBytes...)
			//栈顶元素出栈
			stack.Remove(top)
		} else {
			strBytes = append(strBytes, s[i])
		}
	}
	return string(strBytes)
}

func reverse(bytes []byte) {
	left, right := 0, len(bytes)-1
	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
}
