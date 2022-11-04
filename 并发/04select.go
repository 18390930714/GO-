package main

import "fmt"

/*
select 多路复用
*/

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: //如果能从ch通道取值
			fmt.Println(x)
		case ch <- i: //如果能发送值到ch通道中
		default:
			fmt.Println("啥都不干")
		}
	}
}
