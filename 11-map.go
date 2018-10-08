package main

import "fmt"

func main() {
	// 创建map
	dict := make(map[string]string)
	// 给map添加元素
	dict["a"] = "b"
	// 获取map中的元素
	fmt.Println(dict["a"])
	// 判断map是否存在key
	value, exist := dict["a"]
	fmt.Println(value, exist)
	// 删除map中的元素
	delete(dict, "a")
	// 获取map的长度
	fmt.Println(len(dict))
}