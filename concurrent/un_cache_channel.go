package main

import (
	"fmt"
	"time"
)

/**
1.channel是引用类型,声明时需要指定元素类型;
2.类似于队列 - 先进先出
3.无缓冲的通道(阻塞通道)，必须有接收者接收通道中的数据
4.有缓冲的通道
*/

func send(c chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("发送数据...")
	c <- 10
}

func main() {
	// 新建通道
	ch := make(chan int)
	//启用另一个协程接收通道元素
	go send(ch)
	fmt.Println("取数据...", <-ch)
}
