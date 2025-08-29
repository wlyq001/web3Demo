package main

import (
	"fmt"
	"unsafe"
)

var p1 *int
var p2 *string

func main() {
	var i int = 4
	var str string = "hello"

	p1 = &i
	p2 = &str
	p3 := &p2

	var p4 *int

	fmt.Println(p1, p2, p3, p4, &p4)

	px := uintptr(unsafe.Pointer(p1))
	px -= 49104
	var p5 = (*string)(unsafe.Pointer(px))
	fmt.Println(*p5)
}
