package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPalindrome(90))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	length := 0
	y := x
	for y > 0 {
		y /= 10
		length++
	}

	for i := 1; i <= length/2; i++ {
		left := x / int(math.Pow(10, float64(length-i))) % 10
		right := x / int(math.Pow(10, float64(i-1))) % 10
		if left != right {
			return false
		}
	}
	return true

}
