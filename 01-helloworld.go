package main

import "fmt"

func main() {
	fmt.Print("hello world!")
}

// 1、Goland环境搭建时将GOROOT和GOPATH设置路径相同会报warning：GOPATH set to GOROOT has no effect
// 2、没有将包名设置为package main会报runnerw.exe: CreateProcess failed with error 216 (no message available)
