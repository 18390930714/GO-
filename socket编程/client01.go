package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20001") // Dial net包建立连接的函数
	if err != nil {
		fmt.Printf("dial failed. err:%v\n", err)
		return
	}
	defer conn.Close()                 //关闭连接
	input := bufio.NewReader(os.Stdin) // bufio包的NewReader()函数 读取输入的函数; os.Stdin 标准输入;读取用户输入
	for {
		s, _ := input.ReadString('\n') //delim:
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}

		_, err := conn.Write([]byte(s)) //Write() 发送数据函数
		if err != nil {
			fmt.Printf("send faild, err:%v\n", err)
			return
		}

		var buf [1024]byte
		n, err := conn.Read(buf[:]) //Read()接收数据函数
		if err != nil {
			fmt.Printf("read faild, err:%v\n", err)
			return
		}
		fmt.Println("收到服务端回复：", string(buf[:n]))
	}

}
