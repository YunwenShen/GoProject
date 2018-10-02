package main

import "fmt"

func main() {
	finger := 5
	switch finger {
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("default")
	}
}
