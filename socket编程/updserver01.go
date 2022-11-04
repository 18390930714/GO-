package main

import (
	"fmt"
	"net"
)

/*GO语言实现UPD协议通信*/
func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{ // 服务端监听端口
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
	for {
		var buf [1024]byte                         // 定义变量buf 字节切片
		n, addr, err := listen.ReadFromUDP(buf[:]) //从udp通信收消息
		if err != nil {
			fmt.Printf("read from udp failed, err:%v\n", err)
			return
		}
		fmt.Println("接收到数据:", string(buf[:]))
		_, err = listen.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Printf("write to %v failed, err:%v\n", addr, err)
			return
		}
	}
}
