package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func ListenAndServe(address string) {
	// 绑定和监听地址
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Sprintf("listen err: %v", err))
	}
	defer listener.Close()
	log.Println(fmt.Sprintf("bind: %s, start listening...", address))

	for {
		// Accept 阻塞直至建立链接
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(fmt.Sprintf("accept err : %v", err))
		}
		// 开启新的 goroutine 处理
		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	// bufio 标准库缓冲区功能
	reader := bufio.NewReader(conn)
	for {
		// ReadString 会一直阻塞，直到遇到分隔符 \n
		msg, err := reader.ReadString('\n')
		if err != nil {
			// 链接终端或被关闭, 用io.EOF
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}
			return
		}
		// 将收到的信息发送给客户端
		b := []byte(msg)
		conn.Write(b)
	}
}

func main() {
	ListenAndServe(":8080")
}