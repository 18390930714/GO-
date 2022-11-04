package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
TCP服务端程序处理流程
1.监听端口哦
2.接受客户端请求简历链接
3.创建goroutine处理链接
*/

//tcp server demo

func process(conn net.Conn) {
	defer conn.Close() //处理完之后要关闭这个连接
	// 针对当前的连接做数据的发送和接收操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}
		recv := string(buf[:n])
		fmt.Println("接收到的数据：", recv)
		conn.Write([]byte("ok"))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20001") // 开启socket监听
	if err != nil {
		fmt.Println("listen faild, err: %v\n", err) //判断报错
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept faild, err: ", err)
			continue
		}

		go process(conn) // 启动一个goroutine处理连接
	}
}
