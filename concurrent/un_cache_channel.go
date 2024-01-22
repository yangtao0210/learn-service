package main

import (
	"fmt"
)

/**
1.channel是引用类型,声明时需要指定元素类型;
2.类似于队列 - 先进先出
3.无缓冲的通道(阻塞通道)，必须有接收者接收通道中的数据
4.有缓冲的通道
*/

func recv(c chan int) {
	res := <-c
	fmt.Println("接收成功:", res)
}

func read(ch chan int) {
	for {
		val, ok := <-ch
		fmt.Println(val)
		if !ok {
			break
		}
	}
}

func write(ch chan int) {
	defer close(ch)
	for i := 0; i < 4; i++ {
		ch <- i
	}
}
