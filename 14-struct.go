package main

import (
	"fmt"
	"structs/computer"
)

// 结构体的定义：结构体是用户定义的类型，表示若干个字段的集合。类似于Java中的类
// 创建命名结构体
type Employee struct {
	firstName string
	lastName  string
	age       int8
}

// 匿名字段结构体
// 不能包含相同类型的匿名字段
type Anonymity struct {
	string
	int
}

type Address struct {
	city  string
	state string
}

type Person struct {
	firstName string
	lastName  string
	age       int8
	Address
}

func structExample1() {
	// 创建匿名结构体
	employee := struct {
		firstName string
		lastName  string
		age       int8
	}{
		firstName: "yunwen",
		lastName:  "shen",
		age:       24,
	}
	fmt.Println(employee)
}

func structExample2() {
	// 结构体的零值
	var employee Employee
	// 定义了employee却没有初始化，firstName和lastName默认为"",而age默认为0
	fmt.Println(employee)
}

func structExample3() {
	// 访问结构体中的字段
	var employee Employee = Employee{
		firstName: "yunwen",
		lastName:  "shen",
		age:       24,
	}
	fmt.Println(employee.lastName)
	fmt.Println(employee.firstName)
	fmt.Println(employee.age)
}

func structExample4() {
	// 提升字段
	// Person 结构体有一个匿名字段 Address，而 Address 是一个结构体。
	// 现在结构体 Address 有 city 和 state 两个字段，
	// 访问这两个字段就像在 Person 里直接声明的一样，因此我们称之为提升字段。
	var person Person
	person.age = 24
	person.firstName = "guilan"
	person.lastName = "chan"
	person.city = "anHui"
	person.state = "huaiNing"
}

func structExample5() {
	var spec computer.Spec
	spec.Maker = "yigao"
	spec.Price = 2000
	// 在src下的structs目录下的computer包中的spec没有将model声明为大写
	// 在此处便不能访问其model属性
}

func main() {
	structExample3()
}
