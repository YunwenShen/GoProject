package main

func main() {

}

func example1() {
	// go中使用const关键字表示常量
	const a = 5
}

func example2() {
	// 常量的值必须在编译是确定，不能使用函数返回值因为调用函数发生在运行时
	// const b = math.Sqrt(4)
}

func example3() {
	// 常量是无类型的，常量可以赋值给 “合适的” 类型，而不需要类型转换。
	// f1编译通过，f2编译不通过
	const c = 0
	var d = 0
/*	var f1 float32 = c
	var f2 float32 = d*/
}

func example4() {
	// 创建带类型的常量
	const a string = "123"
}
