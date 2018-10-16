package main

import (
	"fmt"
	"time"
)

// select语句用于在多个发送/接受操作中进行选择。
// select语句会一直阻塞，直到发送/接受操作准备就绪。
// 如果有多个信道准备就绪，select会随机选择其中一个进行操作

// select的应用可以是使用select方法向多个服务器进行请求，
// 返回较快响应的服务器，而忽略其他的响应

func server1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from server2"
}

func selectExample1() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	for {
		time.Sleep(1 * time.Second)
		select {
		case s1 := <-output1:
			fmt.Println(s1)
			return
		case s2 := <-output2:
			fmt.Println(s2)
			return
		default:
			fmt.Println("default")
		}
	}
}

func main() {
	selectExample1()
}
