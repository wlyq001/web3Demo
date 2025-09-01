package main

import "fmt"

func task1(intP *int) {
	*intP += 10
}

func task2(intSliceP []int) {
	for i, v := range intSliceP {
		intSliceP[i] = v * 2
	}
}
func main() {
	x := 2
	task1(&x)
	fmt.Println(x)

	intSlice := []int{1, 2, 3, 8}
	task2(intSlice)
	fmt.Println(intSlice)
}
