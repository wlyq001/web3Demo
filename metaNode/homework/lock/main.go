package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	// task1
	//mutex := sync.Mutex{}
	//
	//x := 0
	//
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		for j := 0; j < 1000; j++ {
	//			mutex.Lock()
	//			x++
	//			mutex.Unlock()
	//		}
	//	}()
	//}
	//time.Sleep(5 * time.Second)
	//fmt.Println(x)

	// task2:
	x := atomic.Int64{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				x.Add(1)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(x.Load())

}
