package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func task1(parity int) {
	for i := 1; i <= 10; i++ {
		if (i % 2) == parity {
			fmt.Printf("thread %d :%d \n", parity, i)
			time.Sleep(3 * time.Millisecond)
		}
	}
}

func task2Work(workNo int, ch chan string) {
	begin := time.Now()
	fmt.Println("workNo : ", workNo)
	ch <- "work" + strconv.Itoa(workNo) + ":" + time.Now().Sub(begin).String()
	group.Done()
}

var group = sync.WaitGroup{}

func main() {
	// task1:
	//go task1(1)
	//go task1(0)
	//
	//time.Sleep(3 * time.Second)

	// task2:
	group.Add(10)
	ch := make(chan string, 10)

	for i := 0; i < 10; i++ {
		go task2Work(i, ch)
	}
	//time.Sleep(1 * time.Second)
	//close(ch)
	group.Wait()
	close(ch)

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

}
