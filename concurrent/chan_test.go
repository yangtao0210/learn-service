package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	ch := make(chan int)
	// 应该先执行读取（阻塞），再执行channel写入,否则会因为没有读取线程产生死锁
	go read(ch)
	ch <- 10
	time.Sleep(time.Second * 2)
}

func TestUnCacheChannel(t *testing.T) {
	ch := make(chan int)
	go write(ch)
	go read(ch)
	time.Sleep(time.Second * 2)
	fmt.Println("run end")
}

func TestReplaceNumber(t *testing.T) {
	var s string
	//golang处理用户输入
	fmt.Scanln(&s)
	replace := "number"
	replaceBytes := []byte(replace)
	bytes := []byte(s)
	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= '0' && bytes[i] <= '9' {
			bytes = append(bytes[:i], append(replaceBytes, bytes[i+1:]...)...)
			i += len(replaceBytes) - 1
		}
	}
	fmt.Println(string(bytes))
}
