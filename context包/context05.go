package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TraceCode string

var wg sync.WaitGroup

// WithValue

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace dode
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Second * 5)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done() //执行完协程数减一
}

func main() {
	// 设置一个5秒的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// 在系统的入口中设置trace code 传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "19992937123")
	wg.Add(1) // 协程数
	go worker(ctx)
	time.Sleep(time.Second * 10)
	cancel()  //通知子goroutine结束
	wg.Wait() //等待协程执行完
	fmt.Println("over")
}
