package main

import (
	"fmt"
	"time"
)

func main() {
	// task1:
	//ch := make(chan int)
	//go func() {
	//	for i := 1; i <= 10; i++ {
	//		ch <- i
	//	}
	//}()
	//
	//go func() {
	//	for i := 1; i <= 10; i++ {
	//		fmt.Println(<-ch)
	//	}
	//}()
	//
	//time.Sleep(1 * time.Second)

	// task2:
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(<-ch)
		}
	}()
	time.Sleep(3 * time.Second)
}
