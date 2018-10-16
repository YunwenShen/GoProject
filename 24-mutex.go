package main

import (
	"fmt"
	"sync"
)

// 临界区：当程序并发运行时，多个协程不应该同时访问那些修改共享资源的代码。这些修改共享资源的代码被称为临界区
// Mutex用于提供一种加锁的机制，可确保在某时刻只有一个协程在临界区运行，防止出现竞态的情况

// 共享资源
var a = 1

// 使用Mutex处理竞态
func increaseByMutex(wg *sync.WaitGroup, lock *sync.Mutex) {
	lock.Lock()
	a++
	lock.Unlock()
	wg.Done()
}

// 使用缓冲信道来处理竞态
func increaseByChannel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	a++
	<-ch
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	//var lock sync.Mutex
	var ch = make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		//go increaseByMutex(&wg, &lock)
		go increaseByChannel(&wg, ch)
	}
	wg.Wait()
	fmt.Println("the result of a is ", a)
}
