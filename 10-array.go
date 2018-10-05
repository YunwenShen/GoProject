package main

import "fmt"

func main() {
	arrayExample7()
}

func arrayExample1() {
	// 声明数组指定长度的数组
	// go中的变量声明真的好奇怪啊
	var array [3]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	fmt.Println(array)
}

func arrayExample2() {
	// 声明不固定长度的数组, 初始化数组
	var array = [...]string{"a", "b", "c"}
	fmt.Println(array)
}

func arrayExample3() {
	// 在go中数组是值类型而不是引用类型，
	// 所以array1和array2并不是引用同一地址
	var array1 = [...]int{1, 2, 3}
	var array2 = array1
	array1[1] = 0
	fmt.Println(array1)
	fmt.Println(array2)
}

func arrayExample4() {
	// 查看数组长度, 这个len方法和Python中的len似乎是一样的
	var array1 = [...]int{1, 2, 3, 4}
	fmt.Println(len(array1))
}

func arrayExample5() {
	// 循环数组
	var array1 = [...]int{1, 2, 3}
	for i := 0; i < len(array1); i++ {
		fmt.Printf("index is %d, value is %d \n", i, array1[i])
	}
	// 使用range方式循环数组
	for i, v := range array1 {
		fmt.Printf("index is %d, value is %d \n", i, v)
	}
}

func arrayExample6() {
	// 多维数组 3行2列
	var array = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	// 循环多维数组
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Println(array[i][j])
		}
	}
}

func arrayExample7() {
	// 切片其本身不拥有任何数组，只是对现有数组的引用
	var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 切片的容量是从创建切片开始的底层数组中的元素数
	// arraySlice的切片为array[1:5],容量为10-1=9
	var arraySlice = array[4:5]
	for i := 0; i < len(arraySlice); i++ {
		fmt.Println(arraySlice[i])
	}
}

func arrayExample8() {
	// 使用make创建切片
	var array = make([]int, 5, 5)
	fmt.Println(array)
}

func arrayExample9() {
	// 在go中数组长度是固定，它的长度不能添加
	// 但是在go中切片的长度是动态的, 通过append将原来的切片增加到原来的2倍
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is 6
}

func arrayExample10() {
	var names []string
	// nil 表示一个切片的长度和容量都为0
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}
}

// 在GO中的数组值得注意的点：GO中的数组值类型，而GO中的切片是引用类型
// 将GO中的数组传递到函数中对其做操作不会改变其数组
// 而对GO中的切片传递到函数中对其做操作会改变其切片