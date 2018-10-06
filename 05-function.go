package main

import (
	"math"
	"fmt"
)

func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

// 函数示例
func function(param string) (int, int) {
	// func为函数声明关键字
	// function为函数名称
	// param为函数入参
	// int为函数返回类型，如果不存在返回值则不需要写返回类型, 多个返回值将其写在括号内
	return 0, 0
}

func myAbs(param int) (int, float64) {
	return param, math.Abs(float64(param))
}

func test() {
	// _空白符表示不接受该值
	a, _ := myAbs(1)
	fmt.Println(a)
}

func find(num int, nums ...int) {
	// 可变参数函数
	fmt.Print(num, len(nums))
}

func change(s ...string) {
	s[0] = "Go"
	// append 会创建一个新的内存地址来存放append之后的值
	a := append(s, "playground")
	fmt.Println(a)
}
