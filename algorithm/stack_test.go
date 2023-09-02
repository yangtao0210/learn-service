package algorithm

import (
	"fmt"
	"testing"
	"time"
)

func TestStack(t *testing.T) {
	arrayStack := new(ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("rabbit")
	arrayStack.Push("hem")
	fmt.Println("栈大小:", arrayStack.Size())
	fmt.Println("出栈:", arrayStack.Pop())
	fmt.Println("出栈:", arrayStack.Pop())
	fmt.Println("栈顶元素:", arrayStack.Peek())
	fmt.Println("栈大小:", arrayStack.Size())
	fmt.Println("是否为空:", arrayStack.IsEmpty())
	arrayStack.Clear()
	fmt.Println("是否为空:", arrayStack.IsEmpty())
}

func TestMultiGoRoutine(t *testing.T) {
	arrayStack := new(ArrayStack)
	go func() {
		arrayStack.Push("cat")
		arrayStack.Push("dog")
		fmt.Println("栈大小:", arrayStack.Size())
		fmt.Println("出栈:", arrayStack.Pop())
	}()

	go func() {
		arrayStack.Push("rabbit")
		arrayStack.Push("hem")
		fmt.Println("栈大小:", arrayStack.Size())
		fmt.Println("出栈:", arrayStack.Pop())
	}()
	time.Sleep(3 * time.Second)
	fmt.Println(arrayStack.array)
}
