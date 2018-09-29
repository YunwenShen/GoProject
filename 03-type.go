package main

import "fmt"

func main() {
	BoolExample()
}

// 布尔类型
func BoolExample() {
	var t1 = true
	var t2 = false
	fmt.Println(t1 && t2)
	fmt.Println(t1 || t2)
	fmt.Println("t1 is ", t1)
}

// 有符号整型
func intExample() {
	var t1 int8  // int8的长度为8位，相当于java中的byte
	var t2 int16 // int16的长度为16位，相当于java中的short
	var t3 int32 // int32的长度为32位，相当于java中的int,又名rune
	var t4 int64 // int64的长度为64位，相当于java中的long
	var t5 int   // int的长度根据底层平台不同决定的，32位系统中为32位，64位系统中为64位
}

// 无符号整型和有符号整型的区别在于取值范围不能为负数
func uintExample() {
	var t1 uint8 // 在go中的别名为byte
	var t2 uint16
	var t3 uint32
	var t4 uint64
}

// 浮点型
func floatExample() {
	var t1 float32 // 32位浮点型，单精度
	var t2 float64 // 64位浮点型，双精度
}

// 复数类型
func complexExample() {
	var t1 complex64  // 实部和虚部都是float32类型的复数
	var t2 complex128 // 实部和虚部都是float64类型的复数
}

// 字符串类型
func StringExample() {
	var t1 string = "123"
}

// go中没有类型自动提升和类型转换，必须是相同类型才能进行操作
