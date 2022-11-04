package main

import (
	"fmt"
	"sync"
	"time"
)

/*
读写互斥锁
读写互斥锁: 一个goroutine进入写锁时，其他goroutine不能读写，当一个goroutine进入读锁时，其他goroutine可以读不能写
sync包
*/

var (
	x      int64
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)

func read() {
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg.Done()
}

func write() {
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 10)
	rwlock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go write()
	}
	wg.Wait() //阻塞直到计数器变为0
	fmt.Println(time.Now().Sub(start))
}
