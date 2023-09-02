package algorithm

import (
	"sync"
)

// ArrayStack 数组实现栈及其操作
type ArrayStack struct {
	array []string   // 底层切片
	size  int        //元素数量
	lock  sync.Mutex // 锁
}

// Push 入栈操作
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	// 后进元素放在最后面，取的时候从后往前取
	stack.array = append(stack.array, v)
	stack.size += 1
}

// Pop 出栈
func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.size == 0 {
		panic("stack is empty")
	}
	val := stack.array[stack.size-1]
	//更新栈属性值
	newArr := make([]string, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArr[i] = stack.array[i]
	}
	stack.array = newArr
	stack.size -= 1
	return val
}

// Peek 获取栈顶元素
func (stack *ArrayStack) Peek() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.size == 0 {
		panic("stack is empty")
	}
	return stack.array[stack.size-1]
}

// Size 获取栈大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

// IsEmpty 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}

// Clear 清空栈元素
func (stack *ArrayStack) Clear() {
	stack.array = []string{}
	stack.size = 0
}
