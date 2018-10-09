package main

import "fmt"

// 在面对对象领域中，将接口描述为接口定义了对象的行为
// 在go中的接口和面对对象领域中的说法很类似，接口指定了一个类型具有的方法，并由这个类型决定如何实现这个接口

// 定义洗衣机的接口
type WashingMachine interface {
	Cleaning()
	Drying()
}


// 海尔牌洗衣机
type Haier struct {
	quality string
	price   float64
}

// 美的洗衣机
type Media struct {
	quality string
	price   float64
}

// 海尔实现洗衣机接口
func (haier Haier) Cleaning() {
	fmt.Println("haier washing machine can clean clothes")
}
func (haier Haier) Drying() {
	fmt.Println("haier washing machine can dry clothes")
}

// 美的实现洗衣机接口
func (media Media) Cleaning() {
	fmt.Println("media washing machine can clean clothes")
}
func (media Media) Drying() {
	fmt.Println("media washing machine can dry clothes")

}

func interfaceExample1() {
	var haier = Haier{}
	var wm WashingMachine
	// 由于go是强类型语言，vm在经过vm=haier，之后wm仍然属于interface类型
	wm = haier
	wm.Cleaning()
}

func interfaceExample2(wm WashingMachine) {
	// 接口的内部结构就是一个元组（type,value）
	// type 指的是实现该接口实现类的类型
	// value指的是该实现类的值
	fmt.Printf("Interface type %T value %v\n", wm, wm)
}

func interfaceExample3() {
	var vm1 WashingMachine
	var vm2 WashingMachine
	haier := Haier{}
	media := Media{}
	vm1 = haier
	vm2 = media
	interfaceExample2(vm1)
	interfaceExample2(vm2)
}

func interfaceExample4(wm WashingMachine) {
	// 类型断言用于提取接口的底层值以及底层值的实际类型
	switch wm.(type) {
	case Haier:
		fmt.Println("the name of wm is haier value is", wm.(Haier))
	case Media:
		fmt.Println("the name of wm is meida value is", wm.(Media))
	}
}

func interfaceExample5() {
	var vm1 WashingMachine
	haier := Haier{}
	vm1 = haier
	interfaceExample2(vm1)
}

func interfaceExample6() {
	// 空接口 没有包含方法的接口称为空接口。
	// 空接口 表示为 interface {}
	// 由于空接口没有方法，所以可以理解为所有方法都实现了该接口
}

func interfaceExample7() {
	// 接口零值指的是接口的type为nil，value也为nil类型
}

func main() {
	interfaceExample5()
}
