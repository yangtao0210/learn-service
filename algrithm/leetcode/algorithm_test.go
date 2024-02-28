package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	reverseWords("the sky is blue")
}

func TestKMPStr(t *testing.T) {
	next := getNext("aabaaf")
	fmt.Println(next)
}

func TestRepeatedSubString(t *testing.T) {
	fmt.Println(repeatedSubstringPattern("ababab"))
}

func TestCharByte(t *testing.T) {
	isValid("(){}[]")
}

func isValid(s string) bool {
	//思路：使用栈 + map 完成字符串匹配
	matchMap := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	sBytes := []byte(s)
	stack := make([]byte, 0)
	for i := 1; i < len(sBytes); i++ {
		elem := sBytes[i]
		if len(stack) == 0 {
			stack = append(stack, elem)
			continue
		}
		//获取栈顶元素
		top := stack[len(stack)-1]
		if val, ok := matchMap[top]; ok {
			if val == elem {
				//出栈
				stack = stack[:len(stack)-1]
				fmt.Println("出栈：", stack)
			} else {
				//入栈
				stack = append(stack, elem)
				fmt.Println("入栈：", stack)
			}
		} else {
			return false
		}
	}
	return len(stack) == 0
}
