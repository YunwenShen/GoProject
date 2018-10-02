package main

import "fmt"

func main() {
	ifElseExample(5)
}

func ifElseExample(number int) {
	if number % 2 == 0 {
		fmt.Printf("number %d 是一个偶数", number)
	} else {
		fmt.Printf("number %d 是一个奇数", number)
	}
}
