package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	atoi, _ := strconv.Atoi(str)
	itoa := strconv.Itoa(atoi)
	fmt.Println(str == itoa)

}
