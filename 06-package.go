package main

import (
	"geometry/rectangle"
	"fmt"
)

func main() {
	packageExample2()
}

func packageExample1() {
	// 在项目中建立一个src目录
	// 在该目录中添加一个geometry并添加一个geometry包
	// 添加geometry将其声明为main包
	// 使用 go install geometry
	// 会在src同级目录下生成一个bin包，并在其中生成一个geometry.exe文件
}

func packageExample2() {
	// 在 src 目录下新建rectangle目录并新建一个rectangle包
	// 导出该包调用其方法
	// ps: 在goland中将GOPATH中的project path 修改为src的父目录
	area := rectangle.Area(1.0, 2.0)
	fmt.Println("rectangle area is :", area)
}

func packageExample3() {
	// init() 方法一般用于初始化任务，也可以用于在开始执行之前验证程序的正确性
	// 包的初始化顺序：
	// 1、首先初始化包级别的变量
	// 2、调用init方法
}

func packageExample4() {
	// 在GO中导入了包却不使用这样是会报错的，
	// 同理变量在GO中声明了却不使用同样也是会报错的
}
