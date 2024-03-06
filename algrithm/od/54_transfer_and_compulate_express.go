package od

import (
	"container/list"
	"fmt"
	"strconv"
)

type OperationType string

const (
	t1 OperationType = "#"
	t2 OperationType = "$"
)

func TransferAndCompulateExpress() {
	var str string
	fmt.Scan(&str)

	//使用list结构，实现栈操作
	stack := list.New()
	index := 0
	for index < len(str) {
		c := str[index]
		//如果当前字符是数字
		fmt.Println(c, '#', '$')
		if c != '#' || c != '$' {
			start := index
			for index < len(str) && c != '#' || c != '$' {
				index++
			}
			fmt.Println(start, index, str[start:index])
			num, _ := strconv.Atoi(str[start:index])
			stack.PushBack(num)
		} else {
			fmt.Println("运算符")
			if c == '$' {
				elem := stack.Back()
				firstOperate := elem.Value.(int)
				stack.Remove(elem)
				//索引后移，获取下一个操作数
				index++
				start := index
				for index < len(str) && c >= '0' && c <= '9' {
					index++
				}
				secondOperate, _ := strconv.Atoi(str[start:index])
				//计算结果入栈
				stack.PushBack(computeResByType(firstOperate, secondOperate, t2))
			} else if c == '#' {
				index++
			}
		}
	}

	//此时栈只有数字，通过#运算获取最后结果
	elem := stack.Front()
	operateRes := elem.Value.(int)
	stack.Remove(elem)
	for stack.Len() != 0 {
		elem := stack.Front()
		operate := elem.Value.(int)
		operateRes = computeResByType(operateRes, operate, t1)
		stack.Remove(elem)
	}
	fmt.Println(operateRes)
}

func computeResByType(a, b int, ot OperationType) int {
	switch ot {
	case t1:
		return 2*a + 3*b + 4
	case t2:
		return 3*a + b + 4
	default:
		return 0
	}
}
