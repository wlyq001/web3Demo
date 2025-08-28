package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}
	slice := ints[0:1]

	fmt.Println(slice)

	result := append(ints[:1], ints[1:]...)
	fmt.Println(result)

}
