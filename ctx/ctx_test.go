package ctx

import (
	"context"
	"testing"
	"time"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return //return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func TestGen(t *testing.T) {
	/*
		Background()：返回最顶层Context,主要用于main函数/初始化函数/测试代码...
		1.WithCancel(ctx):返回带有Done通道的父节点副本，可调用cancel函数或者直接关闭Done通道。
	*/
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel() // 当我们取完需要的整数后调用cancel
	//for n := range gen(ctx) {
	//	fmt.Println(n)
	//	if n == 10 {
	//		break
	//	}
	//}
	// 2.WithDeadline(ctx,d):返回父节点副本，并将deadline调整为不迟于d
	d := time.Now().Add(time.Millisecond * 50)
	ctx, cancel := context.WithDeadline(context.Background(), d)

}
