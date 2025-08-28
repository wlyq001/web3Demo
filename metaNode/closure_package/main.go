package main

import "fmt"

var func1 func()

type A struct {
	a int
}

func (a *A) test() {
	a.a++
}
func main() {
	a := A{a: 1}
	func1 = a.test
	func1()
	fmt.Println(a.a)
}
