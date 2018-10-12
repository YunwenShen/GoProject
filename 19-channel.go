package main

import "fmt"

// 信道可以想象成GO协程之间通信的管道。如同管道中的水从一端流到另一端

func channelExample1() {
	// 信道的声明
	var c chan int
	if c == nil {
		fmt.Println()
		c = make(chan int)
		fmt.Printf("Type of c is %T", c)
	}
}

func channelExample2() {
	// 通过信道接受和发送，信道的发送和接受默认是阻塞，
	// 阻塞的意味着接受不到发送端的值就会一直等待
	// 信道的实现的功能就是等同于Java多线程中wait和notify
	done := make(chan bool)
	go channelExample3(done)
	// 从信道中将值取出
	fmt.Println(<-done, <-done)
	fmt.Println("main function")
}

func channelExample3(done chan bool) {
	fmt.Println("hello world goroutine")
	// 将值存入到信道中
	done <- true
	done <- false
}

func channelExample4() {
	// 当GO协程使用信道发送数据时，按理来说会有其他协程来接受该数据。
	// 如果没有其他协程来接受那么就会造成死锁。
	var ch = make(chan int)
	ch <- 5
}

func channelExample5(ch chan<- int) {
	// 信道可以分为单向信道和双向信道
	// 默认的信道是双向信道
}

func channelExample6(intch chan int) {
	// 关闭信道和遍历信道内容
	for i := 1; i < 10; i++ {
		intch <- i
	}
	close(intch)
}

func channelExample7() {
	intch := make(chan int)
	go channelExample6(intch)
	for {
		value, ok := <-intch
		if ok == false {
			break
		}
		fmt.Println(value)
	}
}

func main() {
	channelExample7()
}
