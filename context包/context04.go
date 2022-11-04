package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting")
		time.Sleep(time.Second * 5) // 假设正常连接数据库耗时5秒
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("work done")
	wg.Done()
}

func main() {
	// 设置一个10秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
