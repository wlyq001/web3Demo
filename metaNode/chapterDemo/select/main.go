package main

import "fmt"

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)
	//go
	func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			ch2 <- i
			ch3 <- i
		}
		//close(ch1)
		//close(ch2)
		//close(ch3)
	}()

	for i := 0; i < 1; i++ {
		select {
		//case i1, ok := <-ch1:
		case ch1 <- 11:
			//if !ok {
			//	fmt.Println("ch1 closed")
			//	return
			//}
			fmt.Println("ch1 set number:", 11)
		//case i2, ok := <-ch2:
		case ch2 <- 11:
			//if !ok {
			//	fmt.Println("ch1 closed")
			//	return
			//}
			fmt.Println("ch2 set number:", 11)
		//case i3, ok := <-ch3:
		case ch3 <- 11:
			//if !ok {
			//	fmt.Println("ch2 closed")
			//	return
			//}
			fmt.Println("ch3 set number:", 11)
			//default:
			//	fmt.Println("no channel ready")
		}
	}
}
