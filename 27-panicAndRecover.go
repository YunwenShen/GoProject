package main

import "fmt"

// 当程序发生异常时，应该使用panic来终止异常，而不是忽略该错误继续执行下去
func panicExample1(firstName *string, lastName *string) {
	defer recoverExample1()
	if firstName == nil {
		panic("runtimeError: firstName is nil")
	}
	if lastName == nil {
		panic("runtimeError: lastName is nil")
	}
	fmt.Println(firstName, lastName)
}

// 当程序发生异常时可以使用recover来捕获异常
func recoverExample1() {
	if r := recover(); r != nil {
		fmt.Println("recover from ", r)
	}
}

func main() {
	panicExample1(nil, nil)
}
