package main

import (
	"fmt"
	"sync"
)

/*
并发：同一时段内，交替执行多个任务
并行：同一时刻执行多个任务

goroutine
channel
*/
var wg sync.WaitGroup

func hello() {
	fmt.Println("hello 李文周")
	wg.Done()
}

func main() { // 主goroutine执行main函数
	wg.Add(1)
	go hello() // 开启一个goroutine 执行say函数
	fmt.Println("hello 古力娜扎")
	wg.Wait() // 阻塞，等所有goroutine执行完再实现主程序退出
}
