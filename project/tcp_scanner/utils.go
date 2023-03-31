package tcp_scanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// isOpen 判断端口是否可访问
func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	// 设置连接超时时间
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}

// PortScanner 端口扫描器,返回可访问的端口列表
func PortScanner(host string) []int {
	ports := []int{}
	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	timeout := time.Millisecond * 200
	for port := 1; port < 100; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(host, port, timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}
			wg.Done()
		}(port)
	}
	wg.Wait()
	return ports
}
