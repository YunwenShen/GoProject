package main

import "fmt"

func main() {
	pointerExample4()
}

func pointerExample1() {
	// 变量 b 的值为 156，而 b 的内存地址为 0x1040a124。
	// 变量 a 存储了 b 的地址。我们就称 a 指向了 b。
	b := 156
	var a *int = &b
	fmt.Printf("type of a is %T\n", a)
	fmt.Println("the address of b is ", a)
}

func pointerExample2() {
	// 指针的零值
	a := 100
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after initialize is", b)
	}
}

func pointerExample3() {
	// 指针的解引用: 可以获取指针所指向的变量值
	a := 188
	var b = &a
	fmt.Println("address of b is", b)
	fmt.Println("value of b is", *b)
}

func pointerExample4() {
	// 通过指针修改a的值
	a := 1
	b := &a
	*b++
	fmt.Println(a)
}

func pointerExample5() {
	// go中不支持指针运算
	var a int = 1
	var b *int = &a
	b++
}