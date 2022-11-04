package main

import (
	"fmt"
	"sync"
)

// 原生map不支持并发安全
// sync.Map 支持并发安全
var (
	wg sync.WaitGroup
)

var m = make(map[int]int)

func get(key int) int {
	return m[key]
}

func set(key int, value int) {
	m[key] = value
}

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1) // wg.Add(int) add函数表示起了多少个协程
		go func(i int) {
			set(i, i+100)
			fmt.Printf("key:%v value:%v", i, get(i))
			wg.Done() //Done 方法用于将 WaitGroup 的计数值减一，可以理解为完成一个子任务
		}(i)
	}
	wg.Wait() // main主协程使用wait方法等待子协程执行完毕：直到 WaitGroup 的计数值为0
}
