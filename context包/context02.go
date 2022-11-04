package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // context.Background()：Context上下文创建方法；
	massege := make(chan int)                               //无缓存通道
	go son(ctx, massege)
	for i := 0; i < 10; i++ {
		massege <- i
	}
	cancel() //取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。
	time.Sleep(time.Second)
	fmt.Println("主进程结束")
}

func son(ctx context.Context, msg chan int) {
	t := time.Tick(time.Second)
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("接收到的值：%d\n", m)
		case <-ctx.Done(): // Done方法需要返回一个channel，这个channel会在当前工作完成或者上下文被取消后关闭，多次调用Done会返回同一个channel
			fmt.Println("son 协程结束运行")
			return
		}
	}
}
