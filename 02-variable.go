package main

import (
	"fmt"
)

func main() {
	main4()
}

func main0() {
	// 声明单个变量
	var age int
	// 赋值
	age = 20
	fmt.Print("my age is ", age)
}

func main1() {
	// 声明变量并初始化
	var age int = 20
	fmt.Print("my age is ", age)
}

func main2() {
	// 类型推断
	var age = 10
	fmt.Print("my age is ", age)
}

func main3() {
	// 声明多个变量
	var age, width int
	age = 10
	width = 123
	fmt.Print("age is", age, "width is", width)
}

func main4() {
	// 简短声明
	a, b := 1, "2"
	c := 1.2
	fmt.Print(a, b, c)
}

// go属于强类型语言
