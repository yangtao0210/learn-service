package stack

import "container/list"

// 有效的括号：
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号
func BracketsIsValiate(s string) bool {
	//解法：栈
	stack := list.New()
	bracketsMap := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for i := range s {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			//遍历到左括号，直接入栈
			stack.PushBack(c)
		} else {
			//遍历到右括号：判断当前元素和栈顶元素是否匹配（匹配时栈顶出栈）
			if stack.Len() > 0 && stack.Back().Value.(byte) == bracketsMap[c] {
				stack.Remove(stack.Back())
			} else {
				return false
			}
		}
	}
	//若括号对全部匹配，则栈为空，否则栈不为空
	return stack.Len() == 0
}
