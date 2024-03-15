package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//1.建立与服务端的连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("connection failed,err:", err)
		return
	}
	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//2.进行数据收发
		//读取用户输入
		inputStr := strings.Trim(scanner.Text(), "\r\n")
		//情况1：用户输入Q/q时退出
		if strings.EqualFold(strings.ToUpper(inputStr), "Q") {
			fmt.Println("用户退出！")
			return
		}
		//情况2：用户输入非q，发送数据给服务器
		_, err := conn.Write([]byte(inputStr))
		if err != nil {
			fmt.Println("数据发送失败")
			return
		}

		//接收服务端的数据
		buf := make([]byte, 512)
		number, err := conn.Read(buf)
		if err != nil {
			fmt.Println("recv failed,err:", err)
			return
		}
		fmt.Println(string(buf[:number]))
	}
}
