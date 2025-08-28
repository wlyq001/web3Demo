package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3}))
}

func singleNumber(nums []int) int {
	result := make(map[int]int)
	for _, x := range nums {
		_, exist := result[x]
		if exist {
			delete(result, x)
		} else {
			result[x] = 1
		}
	}

	var x int
	for k := range result {
		x = k
	}
	return x
}
