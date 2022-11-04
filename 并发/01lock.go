package main

import (
	"fmt"
	"sync"
)

/*
互斥锁 : 保证同一时间有且只有一个goroutine 进入临界区，其他goroutine在等待
*/

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex //互斥锁
)

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() //加锁
		x = x + 1
		lock.Unlock() //解锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
