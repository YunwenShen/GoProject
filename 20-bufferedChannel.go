package main

import (
	"fmt"
	"time"
)

func bufferedChannelExample1() {
	// 定义缓冲信道, 缓冲信道的发送和接受是非阻塞的
	// 缓冲信道会存储数据, 无缓冲信道永远都不会存储数据，只负责数据的流通
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func bufferedChannelExample2() {
	// 会发生信道阻塞的情况
	// 1、缓冲区已经满了时，向缓冲区put是阻塞的
	// 2、当缓冲区是空的，向缓冲区get是阻塞的
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}

func bufferedChannelExample3() {
	// 缓冲信道的容量是指信道可以存储的值的数量
	// 缓冲信道的长度指的是当前排队的元素的个数
	buffered := make(chan int, 2)
	unbuffered := make(chan int)
	fmt.Println("buffered channel capicity is ", cap(buffered), " size is ", len(buffered))
	fmt.Println("unbuffered channel capicity is ", cap(unbuffered), " size is ", len(unbuffered))
}

func main() {
	bufferedChannelExample3()
}
