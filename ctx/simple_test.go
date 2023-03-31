package ctx

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var wg sync.WaitGroup
var exit bool

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	// 如何接收外部命令实现退出
	wg.Done()
}

// 全局变量实现接收外部命令
func globalWorker() {
	//全局变量方式存在问题：1.跨包调用时不易统一；2.如果worker中再启动goroutine，不好控制
	for {
		fmt.Println("global_worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	wg.Done()
}

// 通道传递外部命令
func channelWorker(exitChan chan struct{}) {
LOOP:
	for {
		fmt.Println("channel_worker")
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			break LOOP
		default:
		}
	}
	wg.Done()
}

// 官方处理
func officeWork(ctx context.Context) {
LOOP:
	for {
		fmt.Println("office_worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	wg.Done()
}
func TestWorker(t *testing.T) {
	wg.Add(1)
	go worker()
	// 如何优雅的实现结束子goroutine
	wg.Wait()
	fmt.Println("over")
}

func TestGlobalWork(t *testing.T) {
	wg.Add(1)
	go globalWorker()
	// 如何优雅的实现结束子goroutine
	time.Sleep(time.Second * 3)
	exit = true
	wg.Wait()
	fmt.Println("over")
}

func TestChannelWork(t *testing.T) {
	exitChan := make(chan struct{})
	wg.Add(1)
	go channelWorker(exitChan)
	// 如何优雅的实现结束子goroutine
	time.Sleep(time.Second * 3)
	exitChan <- struct{}{}
	close(exitChan)
	wg.Wait()
	fmt.Println("over")
}

func TestOfficeWork(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go officeWork(ctx)
	// 如何优雅的实现结束子goroutine
	time.Sleep(time.Second * 3)
	cancel() // 通知goroutine结束
	wg.Wait()
	fmt.Println("over")
}
