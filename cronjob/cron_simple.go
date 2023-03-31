package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	c := cron.New(cron.WithLogger(
		cron.VerbosePrintfLogger(log.New(os.Stdout, "cron:", log.LstdFlags))))
	c.AddFunc("@every 1s", func() {
		var m sync.Mutex
		m.Lock()
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}
		m.Unlock()
	})
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Sprintf("cost : %v", elapsed)
	}()
	c.Start()
	time.Sleep(time.Second * 5)
}
