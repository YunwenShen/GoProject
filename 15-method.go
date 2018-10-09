package main

import "fmt"

// 方法的定义：方法就是函数，在关键字func和方法名之间添加了一个特殊的接收器类型
// 接收器的类型可以是结构体类型的也可以是非结构体类型

type Student struct {
	name   string
	age    int
	salary float64
}

func (student Student) displaySalary() {
	fmt.Println(student.name, student.age, student.salary)
}

func (student *Student) changeName() {
	student.name = "syw"
}

func (student Student) changeAge() {
	student.age = 18
}

func methodExample1() {
	student := Student{
		name:   "ShenYunWen",
		age:    24,
		salary: 1000,
	}
	// 使用函数同样可以实现这个功能
	// go中引入方法的概念就是为了弥补go中没有类的概念
	// 使用方法可以将相同方法名定义在不同的值接受器上
	student.displaySalary()
}

func methodExample2() {
	// go中的方法与java/python有很大的区别
	// java/python 他们的方法都是值传递，而go中的方法函数既可以值传递也可以引用传递
	student := Student{
		name: "shenyunwen",
		age:  24,
	}
	fmt.Println("before change name", student.name)
	student.changeName()
	fmt.Println("after change name", student.name)
	// changeName是值传递，changeAge则是引用传递
	// 引用传递的使用场景：对于一个结构体拥有多个字段，使用拷贝是代价较高的情况下可以使用引用传递
	fmt.Println("before change age", student.age)
	student.changeAge()
	fmt.Println("after change name", student.age)
}

func main() {
	methodExample2()
}
