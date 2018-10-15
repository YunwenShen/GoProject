package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	wg.Done()
	fmt.Printf("Goroutine %d ended \n", i)
}

func waitGroupExample1() {
	size := 3
	var wg sync.WaitGroup
	for i := 0; i < size; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("all goroutines finished executing")
}

func main() {
	waitGroupExample1()
}
