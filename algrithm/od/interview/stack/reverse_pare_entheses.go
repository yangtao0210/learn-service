package stack

import "container/list"

//题目描述：
//给出一个字符串 s（仅含有小写英文字母和括号）。
//请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
//注意，您的结果中 不应 包含任何括号

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
