package main

import (
	"fmt"
	"time"
)

// 协程的定义： 协程是与其他函数或方法一起并发运行的函数或方法

// 协程可以看做是轻量级的线程。与创建线程相比，创建一个GO协程的成本很小

// 协程相比于线程的优势：
// 1、相比于线程而言，GO协程的成本极低。堆栈大小只有若干kb
// 2、GO协程会复用数量更少的OS线程
// 3、GO协程使用信道来进行通信。信道用于防止多个协程访问共享内存时发生竞态条件。

// 使用协程
func hello() {
	fmt.Println("hello world goroutine")
}

func numbers() {
	for i := 1; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Print(i)
	}
}

func alphabets() {
	for i := 'a'; i < 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}

func goroutineExample1() {
	// 此处只会输出 main function
	// go hello 启动一个新的协程，正常理解的是 hello 方法 和 goroutine 方法会并发执行
	// 然而 goroutine 运行在主协程上，当主协程运行结束后，而主协程终止，程序也运行终止了
	// 如果 想让程序执行 hello 方法 需要在主协程运行结束后等待其他协程运行结束
	// goroutineExample2 等待了其他线程结束
	go hello()
	fmt.Println("main function")
}

func goroutineExample2() {
	// 其实可以使用线程的思想来理解协程
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}

func goroutineExample3() {
	// 解析请见goroutine.png
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}

func main() {
	goroutineExample3()
}
