package main

import (
	"bufio"
	"fmt"
	"net"
)

var answers map[string]string

//tcp服务端

func main() {
	answers = make(map[string]string)
	answers["你好,服务器"] = "Hello,May I help you?"
	answers["我想要一些关于BD的keyword"] = "OK! I has found words as 'BingDu,BaiDu,BuDong' and so on,from web of google!"
	answers["好的，谢谢！"] = "You`re welcome!"
	answers["Bye"] = "Bye"

	//1.监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed,err:", err)
	}
	for {
		//2.接收客户端请求建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}
		//3.启动goroutine处理连接
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("connection closed failed,err:", err)
		}
	}(conn)

	for {
		reader := bufio.NewReader(conn)
		//buf缓存数据
		buf := make([]byte, 128)
		number, err := reader.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("客户端关闭连接")
			} else {
				fmt.Println("read bytes from client failed,err:", err)
			}
			break
		}
		recv := string(buf[:number])
		fmt.Println("收到客户端的数据：", recv)
		//发送数据给客户端
		conn.Write([]byte(answers[recv]))
	}
}
